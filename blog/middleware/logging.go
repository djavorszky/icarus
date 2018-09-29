package middleware

import (
	"context"
	"github.com/go-kit/kit/log"
	"icarus/blog"
	"time"
)

type LoggingMW struct {
	Logger log.Logger
	Inner  blog.Service
}

func (mw LoggingMW) GetByID(ctx context.Context, ID string) (res blog.EntryModel, err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
			"method", "GetByID",
			"input", ID,
			"entry-title", res.Title,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())

	time.Sleep(300 * time.Millisecond)

	res, err = mw.Inner.GetByID(ctx, ID)

	return
}
func (mw LoggingMW) GetByAuthor(ctx context.Context, author string) (res []blog.EntryModel, err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
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
func (mw LoggingMW) Add(ctx context.Context, entry blog.EntryModel) (err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
			"method", "Add",
			"inputTitle", entry.Title,
			"success", err == nil,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())

	err = mw.Inner.Add(ctx, entry)

	return
}
func (mw LoggingMW) UpdateByID(ctx context.Context, ID string, entry blog.EntryModel) (res blog.EntryModel, err error) {
	defer func(start time.Time) {
		mw.Logger.Log(
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

func (mw LoggingMW) DeleteByID(ctx context.Context, ID string) (err error) {
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
