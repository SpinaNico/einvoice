package invoice

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler get http handler of invoices service.
func Handler() http.Handler {

	router := mux.NewRouter()

	return router
}
