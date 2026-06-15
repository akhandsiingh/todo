package sqlc

import "context"

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (int64, error) {
    res, err := q.db.ExecContext(ctx, `INSERT INTO tasks (user_id, group_id, title, description, priority, due_at) VALUES (?, ?, ?, ?, ?, ?)`, arg.UserID, arg.GroupID, arg.Title, arg.Description, arg.Priority, arg.DueAt)
    if err != nil { return 0, err }
    return res.LastInsertId()
}

func (q *Queries) ListTasksByUser(ctx context.Context, arg ListTasksByUserParams) ([]Task, error) {
    rows, err := q.db.QueryContext(ctx, `SELECT id, user_id, group_id, title, description, status, priority, due_at, completed_at, created_at, updated_at FROM tasks WHERE user_id = ? AND (? = '' OR status = ?) AND (? = 0 OR group_id = ?) ORDER BY COALESCE(due_at, created_at) ASC, created_at DESC LIMIT ? OFFSET ?`, arg.UserID, arg.Status, arg.Status, arg.GroupID, arg.GroupID, arg.Limit, arg.Offset)
    if err != nil { return nil, err }
    defer rows.Close()
    var items []Task
    for rows.Next() { var t Task; if err := rows.Scan(&t.ID,&t.UserID,&t.GroupID,&t.Title,&t.Description,&t.Status,&t.Priority,&t.DueAt,&t.CompletedAt,&t.CreatedAt,&t.UpdatedAt); err != nil { return nil, err }; items = append(items, t) }
    return items, rows.Err()
}

func (q *Queries) GetTaskByID(ctx context.Context, arg GetTaskByIDParams) (Task, error) {
    row := q.db.QueryRowContext(ctx, `SELECT id, user_id, group_id, title, description, status, priority, due_at, completed_at, created_at, updated_at FROM tasks WHERE id = ? AND user_id = ? LIMIT 1`, arg.ID, arg.UserID)
    var t Task
    err := row.Scan(&t.ID,&t.UserID,&t.GroupID,&t.Title,&t.Description,&t.Status,&t.Priority,&t.DueAt,&t.CompletedAt,&t.CreatedAt,&t.UpdatedAt)
    return t, err
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
    _, err := q.db.ExecContext(ctx, `UPDATE tasks SET group_id = ?, title = ?, description = ?, status = ?, priority = ?, due_at = ?, completed_at = ? WHERE id = ? AND user_id = ?`, arg.GroupID, arg.Title, arg.Description, arg.Status, arg.Priority, arg.DueAt, arg.CompletedAt, arg.ID, arg.UserID)
    return err
}

func (q *Queries) DeleteTask(ctx context.Context, arg DeleteTaskParams) error {
    _, err := q.db.ExecContext(ctx, `DELETE FROM tasks WHERE id = ? AND user_id = ?`, arg.ID, arg.UserID)
    return err
}
