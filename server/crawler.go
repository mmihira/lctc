package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	models "lctc/models"
)

type TweetStream struct {
	keyword             string
	consumer_key        string
	consumer_secret     string
	access_token        string
	access_token_secret string
}

func handleTweet(db *sqlx.DB) func(t *twitter.Tweet) {
	return func(t *twitter.Tweet) {
		m := models.NewTweet(
			t.CreatedAt,
			t.ID,
			t.IDStr,
			t.Text,
		)
		m.Create(db)
		log.Print("--", t.Text)
	}
}

/*
 Start a stream and handle the messages
*/
func (this *TweetStream) crawl(db *sqlx.DB) {
	log.Print("Starting tweet stream for ", this.keyword)
	config := oauth1.NewConfig(this.consumer_key, this.consumer_secret)
	token := oauth1.NewToken(this.access_token, this.access_token_secret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	params := &twitter.StreamFilterParams{
		Track:         []string{this.keyword},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)
  if err != nil {
    log.Error(err);
    panic("Error creating stream")
  }
	demux := twitter.NewSwitchDemux()
	demux.Tweet = handleTweet(db)
	demux.HandleChan(stream.Messages)
  log.Print("Twitter stream ending")
}
