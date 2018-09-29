package blog

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetByIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		result, err := svc.GetByID(ctx, req.ID)
		if err != nil {
			return GetByIDResponse{Entry: EntryTransport{}, Error: err.Error()}, nil
		}
		return GetByIDResponse{Entry: result.ToTransport()}, nil
	}
}

func MakeGetByAuthorEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByAuthorRequest)
		result, err := svc.GetByAuthor(ctx, req.Author)
		if err != nil {
			return GetByAuthorResponse{Entries: nil, Error: err.Error()}, nil
		}

		entries := make([]EntryTransport, len(result))
		for index, entry := range result {
			entries[index] = entry.ToTransport()
		}

		return GetByAuthorResponse{Entries: entries}, nil
	}
}

func MakeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		err := svc.Add(ctx, req.Entry.ToModel())
		if err != nil {
			return AddResponse{Error: err.Error()}, nil
		}
		return AddResponse{}, nil
	}
}

func MakeUpdateByIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateByIDRequest)
		result, err := svc.UpdateByID(ctx, req.ID, req.Entry.ToModel())
		if err != nil {
			return UpdateByIDResponse{Error: err.Error()}, nil
		}
		return UpdateByIDResponse{Entry: result.ToTransport()}, nil
	}
}

func MakeDeleteByIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteByIDRequest)
		err := svc.DeleteByID(ctx, req.ID)
		if err != nil {
			return DeleteByIDResponse{Error: err.Error()}, nil
		}
		return DeleteByIDResponse{}, nil
	}
}
