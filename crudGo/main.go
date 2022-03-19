package main

import (
	"net/http"

	"github.com/taissiane/crudGo/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
