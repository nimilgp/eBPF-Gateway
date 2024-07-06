package main

import (
	"crypto/rand"
	"ebpf-firewall/dbLayer"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (app *application) generateBearerToken(w http.ResponseWriter, acc dbLayer.Account) error {
	b := make([]byte, 128)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("<ERROR>\t\t[(gen bearer-token)failed to get random bytes]\n%s\n\n", err)
		return err
	}
	tokenString := base64.URLEncoding.EncodeToString(b)[:128]
	curTime := time.Now()
	validTill := curTime.Add(time.Minute * 30)

	timeStamp := pgtype.Timestamp{
		Time:  validTill,
		Valid: true,
	}
	arg := dbLayer.CreateBearerTokenParams{
		Tokenstring: tokenString,
		Validtill:   timeStamp,
		Username:    acc.Username,
	}
	if err := app.queries.CreateBearerToken(app.ctx, arg); err != nil {
		log.Printf("<ERROR>\t\t[(gen bearer-token)failed to create bearer token]\n%s\n\n", err)
		return err
	} else {
		if err := json.NewEncoder(w).Encode(arg); err != nil {
			log.Printf("<ERROR>\t\t[(gen bearer-token)failed to send bearer token]\n%s\n\n", err)
			return err
		}
		log.Printf("<INFO>\t\t[(gen bearer-token)succesfully generate bearer token]\ntoken sting :%s\n\n", tokenString)
		return nil
	}
}

func (app *application) verifyAndUpdateBearerToken(tokenString string) bool {
	bearerToken, err := app.queries.RetrieveBearerToken(app.ctx, tokenString)
	if err != nil {
		log.Printf("<WARNING>\t\t[(verify & update bearer token)failed to retrieve bearer token]\n%s\n\n", err)
		return false
	}
	curTime := time.Now()
	if curTime.Before(bearerToken.Validtill.Time) {
		validTill := curTime.Add(time.Minute * 30)
		timeStamp := pgtype.Timestamp{
			Time:  validTill,
			Valid: true,
		}
		arg := dbLayer.UpdateBearerTokenExpirationParams{
			Tokenstring: tokenString,
			Validtill:   timeStamp,
		}
		if err := app.queries.UpdateBearerTokenExpiration(app.ctx, arg); err != nil {
			log.Printf("<ERROR>\t\t[(verify & update bearer token)failed to update bearer token]\n%s\n\n", err)
			return false
		}
		log.Printf("<INFO>\t\t[(verify & update bearer token)updated bearer token expiration]\ntoken sting :%s\n\n", tokenString)
		return true
	} else {
		log.Printf("<INFO>\t\t[(verify & update bearer token)bearer token has expired]\ntoken sting :%s\n\n", tokenString)
		if err := app.queries.DeleteBearerToken(app.ctx, bearerToken.Username); err != nil {
			log.Printf("<ERROR>\t\t[(verify & update bearer token)failed to delete bearer token]\n%s\n\n", err)
		}
		return false
	}
}
