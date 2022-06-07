package main

import (
	// "go2salle/worker"

	"go2salle/bolted"
	"go2salle/webapp"
	// "github.com/gorilla/mux"
)

func main() {
	webapp.Main()
	// bolted.Main()
	bolted.InitDay(5)
}
