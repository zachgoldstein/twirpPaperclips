package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zachgoldstein/twirpPaperclips/paperclipserver"
	"github.com/zachgoldstein/twirpPaperclips/rpc"
)

func main() {
	fmt.Printf("Starting Universal Paperclip Service on :6666")

	server := paperclipserver.NewServer()
	twirpHandler := rpc.NewUniversalPaperclipsServer(server, nil)

	log.Fatal(http.ListenAndServe(":6666", twirpHandler))
}
