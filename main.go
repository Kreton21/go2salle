package main

import (
	// "go2salle/worker"

	"fmt"
	"go2salle/bolted"
	"go2salle/webapp"
	"time"
	// "github.com/gorilla/mux"
)

func main() {
	webapp.Main()
	day := time.Now()
	date := day.Format("2006-01-02")
	bolted.InitDay(5)
	bolted.Reserv(date, "22-09-2005", "12-21", "2", "Kreton")

	fmt.Println("reserved")
	bolted.Test(date, "2")
	// bolted.Main()
	bolted.InitDay(5)
}
