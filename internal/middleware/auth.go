package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/cafaray/kvs-service/pkg/claim"
	"github.com/cafaray/kvs-service/pkg/response"
)

type key string

const (
	UserIDKey key = "id"
)

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("Authorization token is required")
	}
	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("Invalid authorization format, please use bearer tokens")
	}
	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("Invalid authorization format")
	}
	return l[1], nil
}

func Authorizator(next http.Handler) http.Handler {
	signingString := os.Getenv("SIGNING_STRING")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authorization)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}
		c, err := claim.GetFromToken(tokenString, signingString)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}
		ctx := r.Context()
		// Here, we are creating a new context to use into the request pointer, and we are including
		// the userId to be used in next handler. This is very useful when we need to transport
		// some values from middleware to handlers
		ctx = context.WithValue(ctx, UserIDKey, c.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
