package blog

import (
	"context"
	"time"
)

// Service interface that defines the methods the actual service will need to implement
type Service interface {
	GetByID(ctx context.Context, ID string) (EntryModel, error)
	GetByAuthor(ctx context.Context, author string) ([]EntryModel, error)
	Add(ctx context.Context, entry EntryModel) error
	UpdateByID(ctx context.Context, ID string, entry EntryModel) (EntryModel, error)
	DeleteByID(ctx context.Context, ID string) error
}

// EntryModel represents a blog entry
type EntryModel struct {
	ID          string
	Title       string
	Subtitle    string
	Content     string
	Author      string
	CreateDate  time.Time
	UpdateDate  time.Time
	PublishDate time.Time
}

func (em EntryModel) ToTransport() EntryTransport {
	return EntryTransport{
		ID:          em.ID,
		Title:       em.Title,
		Subtitle:    em.Subtitle,
		Content:     em.Content,
		Author:      em.Author,
		CreateDate:  em.CreateDate,
		UpdateDate:  em.UpdateDate,
		PublishDate: em.PublishDate,
	}
}

type SVC struct {
}

func (SVC) GetByID(ctx context.Context, ID string) (EntryModel, error) {
	return EntryModel{
		ID:          "some-id",
		Title:       "some-title",
		Subtitle:    "some-subtitle",
		Content:     "this is a content, yay",
		Author:      "javdaniel",
		CreateDate:  time.Now(),
		UpdateDate:  time.Now().AddDate(0, 1, 1),
		PublishDate: time.Now().AddDate(0, 0, 1),
	}, nil
}
func (SVC) GetByAuthor(ctx context.Context, author string) ([]EntryModel, error) {
	return nil, nil
}
func (SVC) Add(ctx context.Context, entry EntryModel) error {
	return nil
}
func (SVC) UpdateByID(ctx context.Context, ID string, entry EntryModel) (EntryModel, error) {
	return EntryModel{}, nil
}
func (SVC) DeleteByID(ctx context.Context, ID string) error {
	return nil
}
