package main

import (
	"log"
	"net/http"

	"github.com/BeepNBoop/pokepaste"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", pokepaste.Handler))
}
