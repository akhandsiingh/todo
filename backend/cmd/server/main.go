package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "todo-app/backend/internal/controller"
    "todo-app/backend/internal/db"
    "todo-app/backend/internal/db/sqlc"
    "todo-app/backend/internal/middleware"
    "todo-app/backend/internal/repository"
    "todo-app/backend/internal/routes"
    "todo-app/backend/internal/scheduler"
    "todo-app/backend/internal/service"
)

func main() {
    db.LoadEnv(".env")
    database, err := db.Connect(db.FromEnv())
    if err != nil { log.Fatalf("database connection failed: %v", err) }
    defer database.Close()
    if err := db.RunMigrations(database, filepath.Join("internal", "db", "migrations")); err != nil { log.Fatalf("migrations failed: %v", err) }

    q := sqlc.New(database)
    userRepo := repository.NewUserRepository(q)
    groupRepo := repository.NewGroupRepository(q)
    taskRepo := repository.NewTaskRepository(q)
    reminderRepo := repository.NewReminderRepository(q)

    secret := env("JWT_SECRET", "change-this-secret")
    authSvc := service.NewAuthService(userRepo, secret)
    groupSvc := service.NewGroupService(groupRepo)
    taskSvc := service.NewTaskService(taskRepo)
    reminderSvc := service.NewReminderService(reminderRepo, taskRepo)

    scheduler.NewReminderScheduler(reminderSvc).Start(context.Background())

    handler := routes.New(routes.Controllers{
        Auth: controller.NewAuthController(authSvc),
        Tasks: controller.NewTaskController(taskSvc),
        Groups: controller.NewGroupController(groupSvc),
        Reminders: controller.NewReminderController(reminderSvc),
    }, secret)
    handler = middleware.Recovery(middleware.Logging(middleware.CORS(env("CORS_ALLOWED_ORIGIN", "http://localhost:3000"))(handler)))

    addr := ":" + env("APP_PORT", "8080")
    log.Printf("backend listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, handler))
}

func env(key, fallback string) string { if v := os.Getenv(key); v != "" { return v }; return fallback }
