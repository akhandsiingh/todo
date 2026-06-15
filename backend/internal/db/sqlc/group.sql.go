package sqlc

import "context"

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (int64, error) {
    res, err := q.db.ExecContext(ctx, `INSERT INTO task_groups (user_id, name, color) VALUES (?, ?, ?)`, arg.UserID, arg.Name, arg.Color)
    if err != nil { return 0, err }
    return res.LastInsertId()
}

func (q *Queries) ListGroupsByUser(ctx context.Context, userID int64) ([]TaskGroup, error) {
    rows, err := q.db.QueryContext(ctx, `SELECT id, user_id, name, color, created_at, updated_at FROM task_groups WHERE user_id = ? ORDER BY name ASC`, userID)
    if err != nil { return nil, err }
    defer rows.Close()
    var items []TaskGroup
    for rows.Next() { var g TaskGroup; if err := rows.Scan(&g.ID,&g.UserID,&g.Name,&g.Color,&g.CreatedAt,&g.UpdatedAt); err != nil { return nil, err }; items = append(items, g) }
    return items, rows.Err()
}

func (q *Queries) GetGroupByID(ctx context.Context, arg GetGroupByIDParams) (TaskGroup, error) {
    row := q.db.QueryRowContext(ctx, `SELECT id, user_id, name, color, created_at, updated_at FROM task_groups WHERE id = ? AND user_id = ? LIMIT 1`, arg.ID, arg.UserID)
    var g TaskGroup
    err := row.Scan(&g.ID,&g.UserID,&g.Name,&g.Color,&g.CreatedAt,&g.UpdatedAt)
    return g, err
}

func (q *Queries) UpdateGroup(ctx context.Context, arg UpdateGroupParams) error {
    _, err := q.db.ExecContext(ctx, `UPDATE task_groups SET name = ?, color = ? WHERE id = ? AND user_id = ?`, arg.Name, arg.Color, arg.ID, arg.UserID)
    return err
}

func (q *Queries) DeleteGroup(ctx context.Context, arg DeleteGroupParams) error {
    _, err := q.db.ExecContext(ctx, `DELETE FROM task_groups WHERE id = ? AND user_id = ?`, arg.ID, arg.UserID)
    return err
}
