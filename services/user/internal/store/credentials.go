package store

import (
	"context"
)

type Credentials struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

func (u *user) CreateCredentials(ctx context.Context, c *Credentials) error {
	_, err := u.db.ExecContext(ctx, `INSERT INTO credentials (name, password) VALUES (?, ?)`,
		c.Name,
		c.Password)

	return err
}

func (u *user) GetCredentials(ctx context.Context, name string) (*Credentials, error) {
	var buffer Credentials

	if err := u.db.GetContext(ctx, &buffer, `SELECT * FROM credentials WHERE name = ?`, name); err != nil {
		return nil, err
	}

	return &buffer, nil
}
