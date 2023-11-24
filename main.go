package main

import (
	"net/http"
	"xmas-list/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
