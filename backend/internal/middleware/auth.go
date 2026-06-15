package middleware

import (
    "context"
    "net/http"
    "strings"
    "todo-app/backend/internal/util"
)

type contextKey string
const UserIDKey contextKey = "userID"

func Auth(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            header := r.Header.Get("Authorization")
            if !strings.HasPrefix(header, "Bearer ") { util.Error(w, http.StatusUnauthorized, "missing bearer token"); return }
            claims, err := util.VerifyToken(strings.TrimPrefix(header, "Bearer "), secret)
            if err != nil { util.Error(w, http.StatusUnauthorized, "invalid token"); return }
            next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserIDKey, claims.UserID)))
        })
    }
}

func UserID(r *http.Request) int64 { id, _ := r.Context().Value(UserIDKey).(int64); return id }
