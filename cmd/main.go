package main

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"icarus/blog"
	"icarus/blog/middleware"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc blog.Service
	svc = blog.SVC{}
	svc = middleware.LoggingMW{logger, svc}

	getByIDHandler := httptransport.NewServer(
		blog.MakeGetByIDEndpoint(svc),
		blog.DecodeGetByIDRequest,
		encodeResponse,
	)

	getByAuthorHandler := httptransport.NewServer(
		blog.MakeGetByAuthorEndpoint(svc),
		blog.DecodeGetByAuthorRequest,
		encodeResponse,
	)

	addHandler := httptransport.NewServer(
		blog.MakeAddEndpoint(svc),
		blog.DecodeAddRequest,
		encodeResponse,
	)

	updateByIDHandler := httptransport.NewServer(
		blog.MakeUpdateByIDEndpoint(svc),
		blog.DecodeUpdateByIDRequest,
		encodeResponse,
	)

	deleteByIDHandler := httptransport.NewServer(
		blog.MakeDeleteByIDEndpoint(svc),
		blog.DecodeDeleteByIDRequest,
		encodeResponse,
	)

	http.Handle("/get-by-id", getByIDHandler)
	http.Handle("/get-by-author", getByAuthorHandler)
	http.Handle("/add", addHandler)
	http.Handle("/update-by-id", updateByIDHandler)
	http.Handle("/delete-by-id", deleteByIDHandler)

	logger.Log(
		"level", "info",
		"event", "startup",
		"listen", ":8080",
	)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log(
			"level", "fatal",
			"error", err.Error(),
			"exit-code", "1",
		)
		os.Exit(1)
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
