package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticationMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, path := range common.NOT_REQUIRED_AUTH_CHECK_URLS {
			if path == r.URL.Path {
				h.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenHeaderList := strings.Split(tokenHeader, " ")
		if len(tokenHeaderList) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
		}

		tokenStr := tokenHeaderList[1]
		claims := &common.Claim{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return []byte(common.JWT_KEY), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "CLAIMS", claims)

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
