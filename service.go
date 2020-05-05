package invoice

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Handler get http handler of invoices service.
func Handler() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/{type:(?:json|xml)}/validate", validateHTTP).Methods("POST")
	return router
}

func validateHTTP(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	file, _, err := r.FormFile("file")
	if err != nil {
		// error open files
		print(err, "\n")
	}
	defer file.Close()
	f, err := os.OpenFile("ff.xml", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		// error in open file directory store
		print(err, "\n")
	}
	fmt.Fprintf(w, `{"type":"%s"}`, vars["type"])
	io.Copy(f, file)

}
