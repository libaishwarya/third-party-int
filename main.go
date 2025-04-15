package main

import (
	"fmt"
	"third-party-int/routes"
	realtime "third-party-int/service/real-time"
	"third-party-int/store/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := mysql.NewDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	jokeStore := mysql.NewJokeStore(db)
	jokehandler := &realtime.RealTimeJoke{}
	route := routes.SetupRouter(jokeStore, jokehandler)

	route.Run(":3000")

}
