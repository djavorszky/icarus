package main

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"icarus/blog"
	"icarus/blog/cfg"
	"icarus/blog/middleware"
	"icarus/blog/network"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc, err := blog.NewService(cfg.ServiceOpts{
		Database: cfg.DatabaseOpts{
			SkipAuth:     true,
			DatabaseName: "blog",
			Address:      "127.0.0.1:27017",
			User:         "root",
			Pass:         "root",
		},
	})
	if err != nil {
		logger.Log(
			"level", "fatal",
			"error", err.Error(),
			"exit-code", "1",
		)
		os.Exit(1)
	}
	defer svc.CleanUp()

	svc = middleware.Logger{Logger: logger, Inner: svc}

	getByIDHandler := httptransport.NewServer(
		blog.MakeGetByIDEndpoint(svc),
		network.DecodeGetByIDRequest,
		encodeResponse,
	)

	getByAuthorHandler := httptransport.NewServer(
		blog.MakeGetByAuthorEndpoint(svc),
		network.DecodeGetByAuthorRequest,
		encodeResponse,
	)

	addHandler := httptransport.NewServer(
		blog.MakeAddEndpoint(svc),
		network.DecodeAddRequest,
		encodeResponse,
	)

	updateByIDHandler := httptransport.NewServer(
		blog.MakeUpdateByIDEndpoint(svc),
		network.DecodeUpdateByIDRequest,
		encodeResponse,
	)

	deleteByIDHandler := httptransport.NewServer(
		blog.MakeDeleteByIDEndpoint(svc),
		network.DecodeDeleteByIDRequest,
		encodeResponse,
	)

	http.Handle("/api/v1/get-by-id", getByIDHandler)
	http.Handle("/api/v1/get-by-author", getByAuthorHandler)
	http.Handle("/api/v1/add", addHandler)
	http.Handle("/api/v1/update-by-id", updateByIDHandler)
	http.Handle("/api/v1/delete-by-id", deleteByIDHandler)

	logger.Log(
		"level", "info",
		"event", "startup",
		"listen", ":8080",
	)

	err = http.ListenAndServe(":8080", nil)
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
