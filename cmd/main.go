package main

import "url-shortener/internal/transport"

func main() {
	router := transport.InitRouter()
	err := router.Run(":5361")
	if err != nil {
		panic(err)
	}
}
