package service

import (
	"context"
	"net/http"
)

func NewProtocolMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		var body []byte

		n, err := request.Body.Read(body)

		if err != nil || n == 0 {

			RenderResponse(writer, ERROR_RESPONSE)

			return
		}

		cmd, err := BuildCommand(body)

		if err != nil {

			RenderResponse(writer, ERROR_RESPONSE)

			return
		}

		ctx := context.WithValue(request.Context(), "cmd", cmd)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
