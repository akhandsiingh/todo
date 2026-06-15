package middleware

import (
    "log"
    "net/http"
    "todo-app/backend/internal/util"
)

func Recovery(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() { if rec := recover(); rec != nil { log.Printf("panic: %v", rec); util.Error(w, http.StatusInternalServerError, "internal server error") } }()
        next.ServeHTTP(w, r)
    })
}
