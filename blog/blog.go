package blog

import (
	"context"
	"fmt"
	"icarus/blog/cfg"
	"icarus/blog/model"
	"icarus/blog/network"
	"time"
)

// Service interface that defines the methods the actual service will need to implement
type Service interface {
	GetByID(ctx context.Context, ID string) (*model.Entry, error)
	GetByAuthor(ctx context.Context, author string) ([]*model.Entry, error)
	Add(ctx context.Context, entry *model.Entry) (string, error)
	UpdateByID(ctx context.Context, ID string, entry *model.Entry) (*model.Entry, error)
	DeleteByID(ctx context.Context, ID string) error
	CleanUp()
}

// NewService returns a ready-to-use service object.
func NewService(opts cfg.ServiceOpts) (Service, error) {
	mongo, err := model.NewMongo(opts.Database)
	if err != nil {
		return nil, fmt.Errorf("service init failed: %v", err)
	}

	var db model.Database = mongo

	return svc{
		db: db,
	}, nil
}

type svc struct {
	db model.Database
}

func (s svc) GetByID(ctx context.Context, ID string) (*model.Entry, error) {
	return s.db.GetByID(ctx, ID)
}

func (s svc) GetByAuthor(ctx context.Context, author string) ([]*model.Entry, error) {
	return s.db.GetByAuthor(ctx, author)
}

func (s svc) Add(ctx context.Context, entry *model.Entry) (string, error) {
	return s.db.Add(ctx, entry)
}

func (s svc) UpdateByID(ctx context.Context, ID string, entry *model.Entry) (*model.Entry, error) {
	entry.UpdateDate = time.Now()

	return s.db.UpdateByID(ctx, ID, entry)
}

func (s svc) DeleteByID(ctx context.Context, ID string) error {
	return s.db.DeleteByID(ctx, ID)
}

func (s svc) CleanUp() {
	s.db.Close()
}

func ModelToNetwork(entry *model.Entry) *network.Entry {
	return &network.Entry{
		ID:          entry.ID,
		Title:       entry.Title,
		Subtitle:    entry.Subtitle,
		Content:     entry.Content,
		Author:      entry.Author,
		CreateDate:  entry.CreateDate,
		UpdateDate:  entry.UpdateDate,
		PublishDate: entry.PublishDate,
	}
}

func NetworkToModel(entry *network.Entry) *model.Entry {
	return &model.Entry{
		ID:          entry.ID,
		Title:       entry.Title,
		Subtitle:    entry.Subtitle,
		Content:     entry.Content,
		Author:      entry.Author,
		CreateDate:  entry.CreateDate,
		UpdateDate:  entry.UpdateDate,
		PublishDate: entry.PublishDate,
	}
}
