package httpgateway

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

const (
	HTTP_SERVER_ADDRESS string = "localhost:9001"
)

func NewMuxServer() *runtime.ServeMux {
	return runtime.NewServeMux()
}

func StartHttpGateway(mux *runtime.ServeMux) {
	log.Println("HTTP PROXY Server started on ", HTTP_SERVER_ADDRESS)

	if err := http.ListenAndServe(HTTP_SERVER_ADDRESS, mux); err != nil {
		panic(err)
	}
}
