package invoice

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvoiceValidate(t *testing.T) {

	pr, pw := io.Pipe()
	//This writers is going to transform
	//what we pass to it to multipart form data
	//and write it to our io.Pipe
	writer := multipart.NewWriter(pw)

	defer writer.Close()
	go func() {

		part, err := writer.CreateFormFile("fileupload", "examples/fake_invoice/IT01234567890_FPA01.xml")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(part, "ciao mondo")
	}()

	req, err := http.NewRequest("POST", "/xml/validate", pr)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		panic(err)
	}

	rec := httptest.NewRecorder()

	Handler().ServeHTTP(rec, req)
}
