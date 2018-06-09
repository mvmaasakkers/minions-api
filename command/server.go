package command

import (
	"github.com/urfave/cli"
	"github.com/gorilla/mux"
	"github.com/BeyondBankingDays/minions-api/api"
	"net/http"
	"log"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
	"github.com/rs/cors"
)

var Server = cli.Command{
	Name:  "server",
	Usage: "Start the server",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "port",
			Value:  "8081",
			Usage:  "Port number",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   "mongodb",
			Value:  "mongodb://localhost:27017",
			Usage:  "MongoDB connection string",
			EnvVar: "MONGODB",
		},
		cli.StringFlag{
			Name:   "database",
			Value:  "hackathon_api",
			Usage:  "MongoDB database string",
			EnvVar: "DATABASE",
		},
	},
	Action: func(c *cli.Context) error {

		db := mongodb.DB{}
		db.Settings.Host = c.String("mongodb")
		db.Settings.Database = c.String("database")

		log.Println("Opening mongodb connection to", db.Settings.Host)
		log.Println("Using mongodb database", db.Settings.Database)

		db.Open()
		defer db.Close()

		r := mux.NewRouter()

		meta := api.Meta{DB: db}

		r.HandleFunc("/v1/data", meta.DataHandler).Methods("POST", "OPTIONS")
		r.Handle("/v1/sources", api.Auth(meta.SourceListHandler)).Methods("GET")
		r.Handle("/v1/sources/{id}", api.Auth(meta.SourceGetHandler)).Methods("GET")
		r.Handle("/v1/sources", api.Auth(meta.SourcePostHandler)).Methods("POST")

		r.Handle("/v1/challenges", api.Auth(meta.ChallengeListHandler)).Methods("GET")
		r.Handle("/v1/challenges/{id}", api.Auth(meta.ChallengeGetHandler)).Methods("GET")

		r.Handle("/v1/bank", api.Auth(meta.BankGetData)).Methods("GET")

		r.Handle("/v1/user/bankuser", api.Auth(meta.AddBankUser)).Methods("POST")

		r.Handle("/v1/user", api.Auth(meta.GetUserHandler)).Methods("GET")
		r.Handle("/v1/user", &api.CreateUserHandler{Meta: meta}).Methods("POST")
		r.Handle("/v1/token", &api.LoginHandler{Meta: meta}).Methods("POST")


		handler := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowCredentials: true,
			AllowedHeaders: []string{
				"authorization",
				"content-type",
			},
			AllowedMethods: []string{
				"GET",
				"POST",
				"PATCH",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS",
			},
		}).Handler(r)

		http.Handle("/", handler)

		log.Println("Starting http server port " + c.String("port"))
		log.Fatal(http.ListenAndServe(":"+c.String("port"), nil))

		return nil
	},
}
