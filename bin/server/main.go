package main

import (
	"net/http"

	gsi "github.com/SpinaNico/go-service-invoice"
)

func main() {
	http.ListenAndServe(":8000", gsi.Handler())
}
