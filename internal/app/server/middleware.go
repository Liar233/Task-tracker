package server

import (
	"context"
	"io"
	"net/http"
)

func ProtocolMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		var body []byte

		body, err := io.ReadAll(request.Body)

		if err != nil {

			RenderResponse(writer, ErrorResponse)

			return
		}

		cmd, err := BuildRequestDto(body)

		if err != nil {

			RenderResponse(writer, ErrorResponse)

			return
		}

		ctx := context.WithValue(request.Context(), "cmd", *cmd)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
