package server

import (
	"context"
	"net/http"
)

func NewProtocolMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		var body []byte

		n, err := request.Body.Read(body)

		if err != nil || n == 0 {

			RenderResponse(writer, ErrorResponse)

			return
		}

		cmd, err := BuildCommand(body)

		if err != nil {

			RenderResponse(writer, ErrorResponse)

			return
		}

		ctx := context.WithValue(request.Context(), "Cmd", cmd)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
