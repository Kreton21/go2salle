package main

import (
	// "go2salle/worker"

	"fmt"
	"go2salle/bolted"
	"go2salle/webapp"
	// "github.com/gorilla/mux"
)

func main() {
	webapp.Main()
	// bolted.Main()
	a := []byte{1, 2, 3}
	bolted.Wdb([]byte("bucket"), []byte("test"), a)
	b := bolted.Rdb("bucket", "test")
	fmt.Println(b)
}
