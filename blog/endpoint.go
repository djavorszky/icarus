package blog

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"icarus/blog/network"
)

func MakeGetByIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(network.GetByIDRequest)
		result, err := svc.GetByID(ctx, req.ID)
		if err != nil {
			return network.GetByIDResponse{Success: false, Error: err.Error()}, nil
		}
		return network.GetByIDResponse{Success: true, Entry: ModelToNetwork(result)}, nil
	}
}

func MakeGetByAuthorEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(network.GetByAuthorRequest)
		result, err := svc.GetByAuthor(ctx, req.Author)
		if err != nil {
			return network.GetByAuthorResponse{Success: false, Error: err.Error()}, nil
		}

		entries := make([]*network.Entry, len(result))
		for index, entry := range result {
			entries[index] = ModelToNetwork(entry)
		}

		return network.GetByAuthorResponse{Success: true, Entries: entries}, nil
	}
}

func MakeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(network.AddRequest)
		id, err := svc.Add(ctx, NetworkToModel(req.Entry))
		if err != nil {
			return network.AddResponse{Success: false, Error: err.Error()}, nil
		}
		return network.AddResponse{Success: true, ID: id}, nil
	}
}

func MakeUpdateByIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(network.UpdateByIDRequest)
		result, err := svc.UpdateByID(ctx, req.ID, NetworkToModel(req.Entry))
		if err != nil {
			return network.UpdateByIDResponse{Success: false, Error: err.Error()}, nil
		}
		return network.UpdateByIDResponse{Success: true, Entry: ModelToNetwork(result)}, nil
	}
}

func MakeDeleteByIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(network.DeleteByIDRequest)
		err := svc.DeleteByID(ctx, req.ID)
		if err != nil {
			return network.DeleteByIDResponse{Success: false, Error: err.Error()}, nil
		}
		return network.DeleteByIDResponse{Success: true}, nil
	}
}
