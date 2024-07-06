package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type firewallArg struct {
	Type   string `validate:"oneof=dns macaddr"`
	Action string `validate:"oneof=block unblock"`
	Value  string `validate:"required,min=4,max=50"`
}

func (app *application) postFirewallAction(w http.ResponseWriter, r *http.Request) {
	var arg firewallArg
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		log.Printf("<ERROR>\t\t[(Firewall-Action)json decode failed]\n%s\n\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := app.validate.Struct(arg); err != nil {
		log.Printf("<WARNING>\t\t[(Firewall-Action)json fields failed to match struct requirements]\n%s\n\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print(arg)
}
