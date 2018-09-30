package middleware

import (
	"context"
	"github.com/go-kit/kit/log"
	"icarus/blog"
	"icarus/blog/model"
	"time"
)

type Logger struct {
	Logger log.Logger
	Inner  blog.Service
}

func (mw Logger) GetByID(ctx context.Context, ID string) (res model.Entry, err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
			"ts", time.Now(),
			"method", "GetByID",
			"input", ID,
			"entry-title", res.Title,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())

	res, err = mw.Inner.GetByID(ctx, ID)

	return
}
func (mw Logger) GetByAuthor(ctx context.Context, author string) (res []model.Entry, err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
			"ts", time.Now(),
			"method", "GetByAuthor",
			"author", author,
			"count", len(res),
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())

	res, err = mw.Inner.GetByAuthor(ctx, author)

	return
}

func (mw Logger) Add(ctx context.Context, entry model.Entry) (id string, err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
			"ts", time.Now(),
			"method", "Add",
			"id", id,
			"success", err == nil,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())

	id, err = mw.Inner.Add(ctx, entry)

	return
}
func (mw Logger) UpdateByID(ctx context.Context, ID string, entry model.Entry) (res model.Entry, err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
			"ts", time.Now(),
			"method", "UpdateByID",
			"ID", entry.ID,
			"success", err == nil,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())

	res, err = mw.Inner.UpdateByID(ctx, ID, entry)

	return
}

func (mw Logger) DeleteByID(ctx context.Context, ID string) (err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
			"method", "DeleteByID",
			"ID", ID,
			"success", err == nil,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())

	err = mw.Inner.DeleteByID(ctx, ID)

	return
}

func (mw Logger) CleanUp() {
	defer func(start time.Time) {
		mw.Logger.Log(
			"method", "CleanUp",
			"took", time.Since(start),
		)
	}(time.Now())

	mw.Inner.CleanUp()

	return
}
