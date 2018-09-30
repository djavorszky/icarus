package network

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Entry struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	CreateDate  time.Time `json:"create-date"`
	UpdateDate  time.Time `json:"update-date"`
	PublishDate time.Time `json:"publish-date"`
}

type GetByIDRequest struct {
	ID string `json:"id"`
}

type GetByIDResponse struct {
	Entry   Entry  `json:"entry,omitempty"`
	Success bool   `json:"success"`
	Error   string `json:"err,omitempty"`
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
	Entries []Entry `json:"entries,omitempty"`
	Success bool    `json:"success"`
	Error   string  `json:"err,omitempty"`
}

func DecodeGetByAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetByAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type AddRequest struct {
	Entry Entry `json:"entry,omitempty"`
}

type AddResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"id,omitempty"`
	Error   string `json:"err,omitempty"`
}

func DecodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type UpdateByIDRequest struct {
	ID    string `json:"id"`
	Entry Entry  `json:"entry,omitempty"`
}

type UpdateByIDResponse struct {
	Entry   Entry  `json:"entry,omitempty"`
	Success bool   `json:"success"`
	Error   string `json:"err,omitempty"`
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
	Success bool   `json:"success"`
	Error   string `json:"err,omitempty"`
}

func DecodeDeleteByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
