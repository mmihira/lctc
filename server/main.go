package main

import (
	_ "database/sql"
	"flag"
	"fmt"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	models "lctc/models"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connectDb() *sqlx.DB {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
      "host=%s port=%s user=postgres dbname=lctc password=exampl sslmode=disable",
      viper.Get("PG_HOST_URL"),
      viper.Get("PG_HOST_PORT")),
	)
	if err != nil {
		log.Print(err)
	}

	return db
}

func main() {
	log.AddHook(filename.NewHook())
	log.Print("Starting server")

	var config_filename = flag.String("configfile", "", "Path to config")
	flag.Parse()

	if len(*config_filename) == 0 {
		log.Print("Loading config from ./config.yml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	} else {
		log.Print("Loading config from ", *config_filename)
		viper.SetConfigFile(*config_filename)
	}

	viper.SetDefault("PG_HOST_URL", "127.0.0.1")
	viper.SetDefault("PG_HOST_PORT", "5432")
	viper.BindEnv("PG_HOST_URL", "PG_HOST_URL")
	viper.BindEnv("PG_HOST_PORT", "PG_HOST_PORT")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Fatal error config file")
		panic("Could not read config file")
	}

	// Create database
	db := connectDb()
	defer db.Close()

	// Initialise models
	models.InitTweet(db)

	// Start tweet stream processing
	tweet_stream := TweetStream{
		keyword:             "btc",
		consumer_key:        viper.Get("CONSUMER_KEY").(string),
		consumer_secret:     viper.Get("CONSUMER_SECRET").(string),
		access_token:        viper.Get("ACCESS_TOKEN").(string),
		access_token_secret: viper.Get("ACCESS_TOKEN_SECRET").(string),
	}

	// Start the tweet stream co-routine
	go tweet_stream.crawl(db)

	// Start server
	r := mux.NewRouter()
	r.HandleFunc("/ws", WsHandler{db}.index)
	http.Handle("/", r)
	log.Print("Running server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
