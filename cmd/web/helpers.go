package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ntiGideon/ecommerce-site/ent"
	"github.com/ntiGideon/ecommerce-site/ent/migrate"
	"net/http"
	"runtime/debug"

	_ "github.com/lib/pq"
)

func connectDb(connectionString string) (*ent.Client, error) {
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to the database")
	return client, nil
}

func OpenDriver(connectionString string) (*sql.DB, error) {
	drv, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return drv, nil
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1>An Error occurred at our side!</h1>"))
	app.logger.Error("An error occurred at our side!", "status", status, "trace", string(debug.Stack()))
	http.Error(w, http.StatusText(status), status)
}
