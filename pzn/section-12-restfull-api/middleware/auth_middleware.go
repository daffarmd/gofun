package middleware

import (
	"net/http"

	"github.com/daffarmd/gofun/helper"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "adalahpokoknya" == r.Header.Get("X-API-Key") {
		m.Handler.ServeHTTP(w, r)
	} else {

		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Error Unauthorized",
		}

		helper.WriteResponseBody(w, webResponse)

	}
}
