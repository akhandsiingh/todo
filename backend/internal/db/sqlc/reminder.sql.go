package sqlc

import "context"

func (q *Queries) CreateReminder(ctx context.Context, arg CreateReminderParams) (int64, error) {
    res, err := q.db.ExecContext(ctx, `INSERT INTO reminders (user_id, task_id, remind_at, message) VALUES (?, ?, ?, ?)`, arg.UserID, arg.TaskID, arg.RemindAt, arg.Message)
    if err != nil { return 0, err }
    return res.LastInsertId()
}

func (q *Queries) ListRemindersByUser(ctx context.Context, userID int64) ([]Reminder, error) {
    rows, err := q.db.QueryContext(ctx, `SELECT id, user_id, task_id, remind_at, message, sent, created_at, updated_at FROM reminders WHERE user_id = ? ORDER BY remind_at ASC`, userID)
    if err != nil { return nil, err }
    defer rows.Close()
    var items []Reminder
    for rows.Next() { var r Reminder; if err := rows.Scan(&r.ID,&r.UserID,&r.TaskID,&r.RemindAt,&r.Message,&r.Sent,&r.CreatedAt,&r.UpdatedAt); err != nil { return nil, err }; items = append(items, r) }
    return items, rows.Err()
}

func (q *Queries) ListDueReminders(ctx context.Context, arg ListDueRemindersParams) ([]Reminder, error) {
    rows, err := q.db.QueryContext(ctx, `SELECT id, user_id, task_id, remind_at, message, sent, created_at, updated_at FROM reminders WHERE sent = FALSE AND remind_at <= ? ORDER BY remind_at ASC LIMIT ?`, arg.RemindAt, arg.Limit)
    if err != nil { return nil, err }
    defer rows.Close()
    var items []Reminder
    for rows.Next() { var r Reminder; if err := rows.Scan(&r.ID,&r.UserID,&r.TaskID,&r.RemindAt,&r.Message,&r.Sent,&r.CreatedAt,&r.UpdatedAt); err != nil { return nil, err }; items = append(items, r) }
    return items, rows.Err()
}

func (q *Queries) MarkReminderSent(ctx context.Context, id int64) error {
    _, err := q.db.ExecContext(ctx, `UPDATE reminders SET sent = TRUE WHERE id = ?`, id)
    return err
}

func (q *Queries) DeleteReminder(ctx context.Context, arg DeleteReminderParams) error {
    _, err := q.db.ExecContext(ctx, `DELETE FROM reminders WHERE id = ? AND user_id = ?`, arg.ID, arg.UserID)
    return err
}
