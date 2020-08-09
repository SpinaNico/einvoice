package einvoice

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type sdiService struct {
	UnimplementedSDIServer
}

// func (s *sdiService) Validate(ctx contex.Context, in *Document)(*Response, error){
// 	return nil,nil
// }

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
		var e FatturaElettronica
		err = xml.Unmarshal(data, &e)
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
