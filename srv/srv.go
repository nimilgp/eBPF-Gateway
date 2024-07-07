package main

import (
	"context"
	"ebpf-firewall/dbLayer"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/rs/cors"
)

type application struct {
	ctx      context.Context
	queries  *dbLayer.Queries
	port     string
	validate *validator.Validate
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Server is Up</h1>")
}

func main() {
	var app application
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=postgres dbname=redq")
	if err != nil {
		log.Fatalf("\n<FATAL>\t\t[posgresql connection failed]\n%s\n", err)
	}
	defer conn.Close(ctx)
	queries := dbLayer.New(conn)
	validate := validator.New(validator.WithRequiredStructEnabled())
	app.ctx = ctx
	app.queries = queries
	app.validate = validate
	app.port = "0.0.0.0:3333"
	log.Printf("<INFO>\t\t[server port on %s]\n\n", app.port)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", getRoot)
	mux.HandleFunc("POST /account/sign-up", app.postAccountSignUp)
	mux.HandleFunc("POST /account/sign-in", app.postAccountSignIn)
	mux.HandleFunc("GET /hardware/ussage", app.getHardwareUssage)
	mux.HandleFunc("GET /hardware/details", app.getHardwareDetails)
	mux.HandleFunc("POST /firewall/action", app.postFirewallAction)
	mux.HandleFunc("GET /firewall/current-total-bandwith", app.getTotalBandwidthUssage)

	handler := cors.Default().Handler(mux) //remove when in production
	if err := http.ListenAndServe(app.port, handler); err != nil {
		log.Fatalf("<FATAL>\t\t[unable to listen and serve]\n%s\n\n", err)
	}
}
