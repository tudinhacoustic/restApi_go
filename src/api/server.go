package api

import (
	"api/auto"
	"api/config"
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}
