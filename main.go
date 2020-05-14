package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go func() {
			err := http.ListenAndServe(":8077", nil)
			if err != nil {
				panic(err)
		}
	}()
}

