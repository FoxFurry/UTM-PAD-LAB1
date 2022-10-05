package store

import (
	"context"
	"time"
)

type Listing struct {
	ID           uint   `db:"id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	ThumbnailURL string `db:"thumbnail_url"`
	AuthorID     uint   `db:"author_id"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (c *catalogue) GetAllListings(ctx context.Context) ([]Listing, error) {
	var buffer []Listing

	if err := c.db.SelectContext(ctx, &buffer, `SELECT * from listings`); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (c *catalogue) AddListing(ctx context.Context, l *Listing) error {
	_, err := c.db.ExecContext(ctx, `INSERT INTO listings (title, description, thumbnail_url, author_id) VALUES (?, ?, ?, ?)`,
		l.Title,
		l.Description,
		l.ThumbnailURL,
		l.AuthorID)
	return err
}

func (c *catalogue) GetListingByTitle(ctx context.Context, title string) (*Listing, error) {
	var buffer Listing
	if err := c.db.GetContext(ctx, &buffer, `SELECT * from listings WHERE title = ?`, title); err != nil {
		return nil, err
	}
	return &buffer, nil
}
