package main

import (
	"net/http"

	"github.com/taissiane/crudGo/crudGo/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
