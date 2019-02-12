package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var tweetSchema = `
tweets (
    created_db int,
    created_at text,
    id bigint PRIMARY KEY,
    id_str text,
    text text
);`

type Tweet struct {
	CreatedDb int32  `db:"created_db"`
	CreatedAt string `db:"created_at"`
	ID        int64  `db:"id"`
	IDStr     string `db:"id_str"`
	Text      string `db:"text"`
}

/*
Insert a new tweet into the database
*/
func (t *Tweet) Create(db *sqlx.DB) {
	if _, err := db.NamedExec(
		`INSERT INTO tweets VALUES (:created_db, :created_at, :id, :id_str, :text) `,
		t,
	); err != nil {
		log.Error(err)
	}
}

func NewTweet(created_at string, id int64, iDStr string, text string) Tweet {
	return Tweet{
		CreatedDb: int32(time.Now().Unix()),
		CreatedAt: created_at,
		ID:        id,
		IDStr:     iDStr,
		Text:      text,
	}
}

func InitTweet(db *sqlx.DB) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("Recovered in f", r)
			os.Exit(3)
		}
	}()
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s", tweetSchema)
	result := db.MustExec(statement)
	if no, _ := result.RowsAffected(); no > 0 {
		log.Print("Created tweet db")
	}
}

func TweetsSinceTimestamp(created_at_db int32, db *sqlx.DB) []string {
	statement := fmt.Sprintf("SELECT * FROM public.tweets WHERE tweets.created_db > %d", created_at_db)
	var ret []string

	rows, err := db.Queryx(statement)
	if err != nil {
		log.Error(err)
	}

	for rows.Next() {
		var t Tweet
		rows.StructScan(&t)
		ret = append(ret, t.IDStr)
	}
	return ret
}
