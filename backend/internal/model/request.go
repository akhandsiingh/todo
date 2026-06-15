package model

type RegisterRequest struct { Name string `json:"name"`; Email string `json:"email"`; Password string `json:"password"` }
type LoginRequest struct { Email string `json:"email"`; Password string `json:"password"` }
type GroupRequest struct { Name string `json:"name"`; Color string `json:"color"` }
type TaskRequest struct { GroupID *int64 `json:"group_id"`; Title string `json:"title"`; Description string `json:"description"`; Status string `json:"status"`; Priority string `json:"priority"`; DueAt string `json:"due_at"` }
type ReminderRequest struct { TaskID int64 `json:"task_id"`; RemindAt string `json:"remind_at"`; Message string `json:"message"` }
