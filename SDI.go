package einvoice

import (
	context "context"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
)

type sdiService struct {
	UnimplementedSDIServer
}

func (s *sdiService) Validate(ctx context.Context, in *Document) (*Response, error) {

	var e FatturaElettronica
	var resp Response
	err := xml.Unmarshal(in.Xml, &e)
	if err != nil {
		return &Response{Ok: false}, status.Error(codes.InvalidArgument, err.Error())
	}
	err = e.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "isntSDIPec":
				resp.Errors = append(resp.Errors, &Error{
					Code:        "00330",
					Description: ErrorsMap["00330"],
				})
				break

			}
		}
	}
	return &resp, status.Error(codes.OK, "ok")
}

func NewSDIService() SDIServer {
	return &sdiService{}
}

func HandlerHTTP(service SDIServer) http.Handler {
	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		res, err := service.Validate(r.Context(), &Document{
			Xml: data,
		})
		if err != nil {
			w.WriteHeader(500)
			return
		}
		data, _ = json.Marshal(res)
		w.Write(data)
	})
	return server
}
