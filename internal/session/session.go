package session

import (
	"encoding/gob"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/antihax/goesi"
	"golang.org/x/oauth2"
)

var Scs *scs.SessionManager

func init() {
	gob.Register(goesi.VerifyResponse{})
	gob.Register(oauth2.Token{})
}

func Setup() {
	Scs = scs.New()
	// sessionManager.Store = pgxstore.New(db.Queries)

	Scs.Lifetime = 24 * time.Hour

	// var err error
	// DB, err = sql.Open("sqlite3", "./test.db")
	// if err != nil {
	// 	panic("Could not connect to db")
	// }
	// err = DB.Ping()
	// if err != nil {
	// 	fmt.Println("Could not ping db", err)
	// }
}
