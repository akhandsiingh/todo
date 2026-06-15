package sqlc

import (
    "database/sql"
    "time"
)

type User struct {
    ID           int64     `json:"id"`
    Name         string    `json:"name"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type TaskGroup struct {
    ID        int64     `json:"id"`
    UserID    int64     `json:"user_id"`
    Name      string    `json:"name"`
    Color     string    `json:"color"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Task struct {
    ID          int64          `json:"id"`
    UserID      int64          `json:"user_id"`
    GroupID     sql.NullInt64  `json:"group_id"`
    Title       string         `json:"title"`
    Description sql.NullString `json:"description"`
    Status      string         `json:"status"`
    Priority    string         `json:"priority"`
    DueAt       sql.NullTime   `json:"due_at"`
    CompletedAt sql.NullTime   `json:"completed_at"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
}

type Reminder struct {
    ID        int64          `json:"id"`
    UserID    int64          `json:"user_id"`
    TaskID    int64          `json:"task_id"`
    RemindAt  time.Time      `json:"remind_at"`
    Message   sql.NullString `json:"message"`
    Sent      bool           `json:"sent"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
}
