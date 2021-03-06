package middlewares

import (
	"context"
	_entities "final-project-usamah/entities"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var ctxKey = &contextKey{"user"}

type contextKey struct {
	data string
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.Contains(header, "Bearer") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(header, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//validate jwt token
		token, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		userID := payload["id"].(float64)
		user := _entities.User{Id: int(userID)}

		ctx := context.WithValue(r.Context(), ctxKey, &user)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})

}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *_entities.User {
	raw, _ := ctx.Value(ctxKey).(*_entities.User)
	return raw
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
