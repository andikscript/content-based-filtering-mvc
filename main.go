package main

import (
	"contentbasedfilteringmvc/cmd/app"
	"fmt"
	"net/http"
)

func main() {
	address := fmt.Sprintf(":%d", 8787)
	server := &http.Server{
		Addr: address,
	}

	app.RouterHandler()

	fmt.Println("start server port", address)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error when start server")
	}
}
