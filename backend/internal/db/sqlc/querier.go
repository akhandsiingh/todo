package sqlc

import (
    "context"
    "database/sql"
    "time"
)

type Querier interface {
    CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
    GetUserByEmail(ctx context.Context, email string) (User, error)
    GetUserByID(ctx context.Context, id int64) (User, error)
    CreateGroup(ctx context.Context, arg CreateGroupParams) (int64, error)
    ListGroupsByUser(ctx context.Context, userID int64) ([]TaskGroup, error)
    GetGroupByID(ctx context.Context, arg GetGroupByIDParams) (TaskGroup, error)
    UpdateGroup(ctx context.Context, arg UpdateGroupParams) error
    DeleteGroup(ctx context.Context, arg DeleteGroupParams) error
    CreateTask(ctx context.Context, arg CreateTaskParams) (int64, error)
    ListTasksByUser(ctx context.Context, arg ListTasksByUserParams) ([]Task, error)
    GetTaskByID(ctx context.Context, arg GetTaskByIDParams) (Task, error)
    UpdateTask(ctx context.Context, arg UpdateTaskParams) error
    DeleteTask(ctx context.Context, arg DeleteTaskParams) error
    CreateReminder(ctx context.Context, arg CreateReminderParams) (int64, error)
    ListRemindersByUser(ctx context.Context, userID int64) ([]Reminder, error)
    ListDueReminders(ctx context.Context, arg ListDueRemindersParams) ([]Reminder, error)
    MarkReminderSent(ctx context.Context, id int64) error
    DeleteReminder(ctx context.Context, arg DeleteReminderParams) error
}

type CreateUserParams struct { Name, Email, PasswordHash string }
type CreateGroupParams struct { UserID int64; Name, Color string }
type GetGroupByIDParams struct { ID, UserID int64 }
type UpdateGroupParams struct { Name, Color string; ID, UserID int64 }
type DeleteGroupParams struct { ID, UserID int64 }
type CreateTaskParams struct { UserID int64; GroupID sql.NullInt64; Title string; Description sql.NullString; Priority string; DueAt sql.NullTime }
type ListTasksByUserParams struct { UserID int64; Status string; GroupID int64; Limit int32; Offset int32 }
type GetTaskByIDParams struct { ID, UserID int64 }
type UpdateTaskParams struct { GroupID sql.NullInt64; Title string; Description sql.NullString; Status string; Priority string; DueAt sql.NullTime; CompletedAt sql.NullTime; ID, UserID int64 }
type DeleteTaskParams struct { ID, UserID int64 }
type CreateReminderParams struct { UserID, TaskID int64; RemindAt time.Time; Message sql.NullString }
type ListDueRemindersParams struct { RemindAt time.Time; Limit int32 }
type DeleteReminderParams struct { ID, UserID int64 }
