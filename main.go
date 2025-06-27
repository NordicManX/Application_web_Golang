package main

import (
	"net/http"

	"gihub.com/nordicmanx/app_web_API/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
