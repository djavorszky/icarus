package blog

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type EntryTransport struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	CreateDate  time.Time `json:"create-date"`
	UpdateDate  time.Time `json:"update-date"`
	PublishDate time.Time `json:"publish-date"`
}

func (et EntryTransport) ToModel() EntryModel {
	return EntryModel{
		ID:          et.ID,
		Title:       et.Title,
		Subtitle:    et.Subtitle,
		Content:     et.Content,
		Author:      et.Author,
		CreateDate:  et.CreateDate,
		UpdateDate:  et.UpdateDate,
		PublishDate: et.PublishDate,
	}
}

type GetByIDRequest struct {
	ID string `json:"id"`
}

type GetByIDResponse struct {
	Entry EntryTransport `json:"entry"`
	Error string         `json:"err"`
}

func DecodeGetByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type GetByAuthorRequest struct {
	Author string `json:"author"`
}

type GetByAuthorResponse struct {
	Entries []EntryTransport `json:"entries"`
	Error   string           `json:"err"`
}

func DecodeGetByAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetByAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type AddRequest struct {
	Entry EntryTransport `json:"entry"`
}

type AddResponse struct {
	Error string `json:"err"`
}

func DecodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type UpdateByIDRequest struct {
	ID    string         `json:"id"`
	Entry EntryTransport `json:"entry"`
}

type UpdateByIDResponse struct {
	Entry EntryTransport `json:"entry"`
	Error string         `json:"err"`
}

func DecodeUpdateByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type DeleteByIDRequest struct {
	ID string `json:"ID"`
}

type DeleteByIDResponse struct {
	Error string `json:"err"`
}

func DecodeDeleteByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
