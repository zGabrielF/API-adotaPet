package main

import (
	"apiAdotaPet/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("rodando a API")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8080", r))
}
