package routes

import (
    "net/http"
    "todo-app/backend/internal/controller"
    "todo-app/backend/internal/middleware"
)

type Controllers struct { Auth *controller.AuthController; Tasks *controller.TaskController; Groups *controller.GroupController; Reminders *controller.ReminderController }

func New(c Controllers, secret string) http.Handler {
    mux := http.NewServeMux()
    auth := middleware.Auth(secret)
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); _, _ = w.Write([]byte("ok")) })
    mux.HandleFunc("/api/auth/register", method(http.MethodPost, c.Auth.Register))
    mux.HandleFunc("/api/auth/login", method(http.MethodPost, c.Auth.Login))
    mux.Handle("/api/auth/me", auth(http.HandlerFunc(method(http.MethodGet, c.Auth.Me))))
    mux.Handle("/api/tasks", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { if r.Method==http.MethodGet { c.Tasks.List(w,r); return }; if r.Method==http.MethodPost { c.Tasks.Create(w,r); return }; http.NotFound(w,r) })))
    mux.Handle("/api/tasks/", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { if r.Method==http.MethodPut || r.Method==http.MethodPatch { c.Tasks.Update(w,r); return }; if r.Method==http.MethodDelete { c.Tasks.Delete(w,r); return }; http.NotFound(w,r) })))
    mux.Handle("/api/groups", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { if r.Method==http.MethodGet { c.Groups.List(w,r); return }; if r.Method==http.MethodPost { c.Groups.Create(w,r); return }; http.NotFound(w,r) })))
    mux.Handle("/api/groups/", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { if r.Method==http.MethodPut || r.Method==http.MethodPatch { c.Groups.Update(w,r); return }; if r.Method==http.MethodDelete { c.Groups.Delete(w,r); return }; http.NotFound(w,r) })))
    mux.Handle("/api/reminders", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { if r.Method==http.MethodGet { c.Reminders.List(w,r); return }; if r.Method==http.MethodPost { c.Reminders.Create(w,r); return }; http.NotFound(w,r) })))
    mux.Handle("/api/reminders/", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { if r.Method==http.MethodDelete { c.Reminders.Delete(w,r); return }; http.NotFound(w,r) })))
    return mux
}

func method(expected string, h http.HandlerFunc) http.HandlerFunc { return func(w http.ResponseWriter, r *http.Request) { if r.Method != expected { w.WriteHeader(http.StatusMethodNotAllowed); return }; h(w,r) } }
