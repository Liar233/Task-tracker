package server

import (
	"fmt"
	"net/http"

	"github.com/braintree/manners"
)

type HttpServerAdapterInterface interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
	SetHandler(handler http.Handler)
	Close()
}

type HttpServerAdapter struct {
	server *manners.GracefulServer
}

func (hsa *HttpServerAdapter) ListenAndServe() error {

	return hsa.server.ListenAndServe()
}

func (hsa *HttpServerAdapter) ListenAndServeTLS(certFile, keyFile string) error {

	return hsa.server.ListenAndServeTLS(certFile, keyFile)
}

func (hsa *HttpServerAdapter) Close() {

	hsa.server.BlockingClose()
}

func (hsa *HttpServerAdapter) SetHandler(handler http.Handler) {

	hsa.server.Handler = handler
}

func NewHttpServerAdapter(host string, port uint64) *HttpServerAdapter {

	httpServer := http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}

	server := manners.NewWithServer(&httpServer)

	return &HttpServerAdapter{
		server: server,
	}
}
