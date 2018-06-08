package command

import (
	"github.com/urfave/cli"
	"github.com/gorilla/mux"
	"github.com/BeyondBankingDays/minions-api/api"
	"net/http"
	"log"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
)

var Server = cli.Command{
	Name:  "server",
	Usage: "Start the server",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "port",
			Value:  "8080",
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

		r.Handle("/v1/data", &api.DataHandler{Meta: meta}).Methods("POST", "OPTIONS")
		r.Handle("/v1/sources", &api.SourceListHandler{Meta: meta}).Methods("GET")
		r.Handle("/v1/sources/{id}", &api.SourceGetHandler{Meta: meta}).Methods("GET")
		r.Handle("/v1/sources", &api.SourcePostHandler{Meta: meta}).Methods("POST")

		r.Handle("/v1/challenges", &api.ChallengeListHandler{Meta: meta}).Methods("GET")
		r.Handle("/v1/challenges/{id}", &api.ChallengeGetHandler{Meta: meta}).Methods("GET")

		http.Handle("/", r)

		log.Println("Starting http server port " + c.String("port"))
		log.Fatal(http.ListenAndServe(":"+c.String("port"), nil))

		return nil
	},
}
