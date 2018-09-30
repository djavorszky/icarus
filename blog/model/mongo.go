package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"icarus/blog/cfg"
)

func NewMongo(opts cfg.DatabaseOpts) (*mongo, error) {
	if err := validateOpts(opts); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	session, err := mgo.Dial(opts.Address)
	if err != nil {
		return nil, fmt.Errorf("dial failed: %v", err)
	}

	m := &mongo{session: session}

	if opts.SkipAuth {
		return m, nil
	}

	err = m.session.Login(&mgo.Credential{
		Username: opts.User,
		Password: opts.Pass,
	})
	if err != nil {
		return nil, fmt.Errorf("auth failed: %v", err)
	}

	return m, nil
}

func validateOpts(opts cfg.DatabaseOpts) error {
	if opts.DatabaseName == "" {
		return fmt.Errorf("empty database name")
	}

	if !opts.SkipAuth && opts.User == "" {
		return fmt.Errorf("empy username")
	}

	return nil
}

const (
	collection = "blog-entries"
)

var (
	ErrNotExists = errors.New("entry does not exist")
)

// mongo is a type that implements model.Database
type mongo struct {
	session *mgo.Session
	name    string
}

// Close closes the database connection
func (m *mongo) Close() {
	m.session.Close()
}

// GetByID returns the entry with a given ID. If no entry with the given ID exists, returns ErrNotExists
func (m *mongo) GetByID(ctx context.Context, ID string) (*Entry, error) {
	if ID == "" {
		return nil, fmt.Errorf("id is empty")
	}

	var entry Entry
	err := m.session.DB(m.name).C(collection).Find(by(ID)).One(&entry)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotExists
		}

		return nil, fmt.Errorf("failed query by ID: %v", err)
	}

	return &entry, nil
}

// GetByAuthor returns the entries with the given Author. If no entry with the given author exists, returns ErrNotExists
func (m *mongo) GetByAuthor(ctx context.Context, author string) ([]*Entry, error) {

	return nil, nil
}

// Add adds a new entry. It will generate a unique ID for the entry, even if specified.
func (m *mongo) Add(ctx context.Context, entry *Entry) (string, error) {
	if err := validateEntry(entry); err != nil {
		return "", fmt.Errorf("invalid entry: %v", err)
	}

	entry.ID = uuid.New().String()

	err := m.session.DB(m.name).C(collection).Insert(&entry)
	if err != nil {
		return "", fmt.Errorf("failed inserting entry: %v", err)
	}

	return entry.ID, nil
}

// UpdateByID updates the Entry specified by the ID with the provided Entry.
func (m *mongo) UpdateByID(ctx context.Context, ID string, entry *Entry) (*Entry, error) {
	if ID == "" {
		return nil, fmt.Errorf("id is empty")
	}

	if err := validateEntry(entry); err != nil {
		return nil, fmt.Errorf("invalid entry: %v", err)
	}

	if exists, _ := m.Exists(ctx, ID); !exists {
		return nil, ErrNotExists
	}

	entry.ID = ID

	err := m.session.DB(m.name).C(collection).Update(by(ID), entry)
	if err != nil {
		return nil, fmt.Errorf("update: %v", err)
	}

	return entry, nil
}

// Exists checks whether an entry with the given ID exists
func (m *mongo) Exists(ctx context.Context, ID string) (bool, error) {
	count, err := m.session.DB(m.name).C(collection).Find(by(ID)).Count()
	if err != nil {
		return false, fmt.Errorf("exists: %v", err)
	}

	return count != 0, nil
}

// DeleteByID deletes the entry specified by the provided ID.
func (m *mongo) DeleteByID(ctx context.Context, ID string) error {

	return nil
}

func validateEntry(entry *Entry) error {
	if entry.Title == "" {
		return fmt.Errorf("empty title")
	}

	if entry.Author == "" {
		return fmt.Errorf("empty author")
	}

	if entry.Content == "" {
		return fmt.Errorf("empty content")
	}

	if entry.CreateDate.IsZero() {
		return fmt.Errorf("create date is zero")
	}

	if entry.UpdateDate.IsZero() {
		return fmt.Errorf("update date is zero")
	}

	if entry.PublishDate.IsZero() {
		return fmt.Errorf("publish date is zero")
	}

	return nil
}

func by(id string) bson.M {
	return bson.M{"id": id}
}
