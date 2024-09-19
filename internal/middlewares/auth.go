package middlewares

import (
    "net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Дополнительная логика проверки токена

        next.ServeHTTP(w, r)
    })
}