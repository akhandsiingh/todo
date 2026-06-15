package sqlc

import "context"

const createUser = `INSERT INTO users (name, email, password_hash) VALUES (?, ?, ?)`
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
    res, err := q.db.ExecContext(ctx, createUser, arg.Name, arg.Email, arg.PasswordHash)
    if err != nil { return 0, err }
    return res.LastInsertId()
}

const getUserByEmail = `SELECT id, name, email, password_hash, created_at, updated_at FROM users WHERE email = ? LIMIT 1`
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
    row := q.db.QueryRowContext(ctx, getUserByEmail, email)
    var u User
    err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
    return u, err
}

const getUserByID = `SELECT id, name, email, password_hash, created_at, updated_at FROM users WHERE id = ? LIMIT 1`
func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
    row := q.db.QueryRowContext(ctx, getUserByID, id)
    var u User
    err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
    return u, err
}
