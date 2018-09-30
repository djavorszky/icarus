package model

import (
	"context"
	"time"
)

type Database interface {
	GetByID(ctx context.Context, ID string) (*Entry, error)
	GetByAuthor(ctx context.Context, author string) ([]*Entry, error)
	Add(ctx context.Context, entry *Entry) (string, error)
	UpdateByID(ctx context.Context, ID string, entry *Entry) (*Entry, error)
	DeleteByID(ctx context.Context, ID string) error
	Exists(ctx context.Context, ID string) (bool, error)
	Close()
}

// Entry represents a blog entry in the database
type Entry struct {
	ID          string
	Title       string
	Subtitle    string
	Content     string
	Author      string
	CreateDate  time.Time
	UpdateDate  time.Time
	PublishDate time.Time
}
