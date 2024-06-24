package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func randFloats(min, max float64) float64 {
	res := min + rand.Float64()*(max-min)
	return res
}

// later implement caching of all the templates
func getRoot(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/root.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func getUpdateGraph(w http.ResponseWriter, r *http.Request) {
	colSize := randFloats(0, 1)
	ts, err := template.ParseFiles("./ui/graph.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, colSize)
	if err != nil {
		log.Print(err)
	}
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=postgres dbname=test3")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	var name string
	var price int64
	err = conn.QueryRow(context.Background(), "select name, price from widgets where id=$1", 2).Scan(&name, &price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, price)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", getRoot)
	mux.HandleFunc("GET /update-graph", getUpdateGraph)

	if http.ListenAndServe(":3333", mux) != nil {
		log.Println(err)
	}
}
