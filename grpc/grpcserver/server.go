package grpcserver

import (
	"context"
	"github.com/sillyhatxu/consul-client"
	"github.com/sillyhatxu/word-backend/grpc/longman"
	"github.com/sillyhatxu/word-backend/grpc/word"
	"github.com/sillyhatxu/word-backend/service/longmanservice"
	"github.com/sillyhatxu/word-backend/service/wordservice"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("initial grpc server")
	server := grpc.NewServer()

	healthServer := health.NewServer()
	healthServer.SetServingStatus(consul.DefaultHealthCheckGRPCServerName, hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(server, healthServer)

	longman.RegisterLongmanServiceServer(server, &Longman{})
	word.RegisterWordServiceServer(server, &Word{})
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type Longman struct{}

func (l Longman) Add(ctx context.Context, request *longman.AddRequest) (*longman.Response, error) {
	err := longmanservice.Add(request)
	if err != nil {
		logrus.Errorf("longmanservice.Add(%#v) response error. %#v", request, err)
		return nil, err
	}
	return &longman.Response{
		Code:    uint32(codes.OK),
		Message: "Success",
	}, nil
}

func (l Longman) Update(ctx context.Context, request *longman.UpdateRequest) (*longman.Response, error) {
	err := longmanservice.Update(request)
	if err != nil {
		logrus.Errorf("longmanservice.Update(%#v) response error. %#v", request, err)
		return nil, err
	}
	return &longman.Response{
		Code:    uint32(codes.OK),
		Message: "Success",
	}, nil
}

func (l Longman) GetById(ctx context.Context, request *longman.GetByIdRequest) (*longman.WordResponse, error) {
	res, err := longmanservice.GetById(request)
	if err != nil {
		logrus.Errorf("longmanservice.GetById(%#v) response error. %#v", request, err)
		return nil, err
	}
	return &longman.WordResponse{
		Code:    uint32(codes.OK),
		Data:    res,
		Message: "Success",
	}, nil
}

func (l Longman) GetByVocabularyId(ctx context.Context, request *longman.GetByVocabularyIdRequest) (*longman.WordResponse, error) {
	res, err := longmanservice.GetByVocabularyId(request)
	if err != nil {
		logrus.Errorf("longmanservice.GetByWord(%#v) response error. %#v", request, err)
		return nil, err
	}
	return &longman.WordResponse{
		Code:    uint32(codes.OK),
		Data:    res,
		Message: "Success",
	}, nil
}

func (l Longman) List(ctx context.Context, request *longman.ListRequest) (*longman.WordListResponse, error) {
	res, totalCount, err := longmanservice.List(request)
	if err != nil {
		logrus.Errorf("longmanservice.List(%#v) response error. %#v", request, err)
		return nil, err
	}
	return &longman.WordListResponse{
		Code:    uint32(codes.OK),
		Data:    res,
		Message: "Success",
		Extra:   &longman.Extra{TotalCount: totalCount},
	}, nil
}

type Word struct{}

func (Word) List(ctx context.Context, request *word.ListRequest) (*word.WordListResponse, error) {
	res, totalCount, err := wordservice.List(request)
	if err != nil {
		logrus.Errorf("wordservice.List(%#v) response error. %#v", request, err)
		return nil, err
	}
	return &word.WordListResponse{
		Code:    uint32(codes.OK),
		Data:    res,
		Message: "Success",
		Extra:   &word.Extra{TotalCount: totalCount},
	}, nil
}
