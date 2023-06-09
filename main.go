package main

import (
	"fmt"
	"log"
	"net/http"

	Handlers "github.com/spicypumpkin666/eth-go/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
    // create client instance to connect to provider
    client, err := ethclient.Dial("http://localhost:7545")
    if err != nil {
        fmt.Println(err)
    }

    // creat mux router
    r := mux.NewRouter()

    // define endpoint
    r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})
    log.Fatal(http.ListenAndServe(":8080", r))
}