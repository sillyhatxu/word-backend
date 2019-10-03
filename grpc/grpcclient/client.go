package grpcclient

import (
	"context"
	"github.com/sillyhatxu/word-backend/grpc/longman"
	"github.com/sillyhatxu/word-backend/grpc/word"
	"google.golang.org/grpc"
	"time"
)

const defaultTimeout = 10 * time.Second

type Client struct {
	address string
	config  *Config
}

func NewGRPCWordClient(address string, opts ...Option) *Client {
	//default
	config := &Config{
		timeout: defaultTimeout,
	}
	for _, opt := range opts {
		opt(config)
	}
	return &Client{
		address: address,
		config:  config,
	}
}

func (c Client) AddLongmanWord(req *longman.AddRequest) (*longman.Response, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := longman.NewLongmanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return client.Add(ctx, req)
}

func (c Client) UpdateLongmanWord(req *longman.UpdateRequest) (*longman.Response, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := longman.NewLongmanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return client.Update(ctx, req)
}

func (c Client) GetLongmanWordById(req *longman.GetByIdRequest) (*longman.WordResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := longman.NewLongmanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return client.GetById(ctx, req)
}

func (c Client) GetLongmanWordByWord(req *longman.GetByVocabularyIdRequest) (*longman.WordResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := longman.NewLongmanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return client.GetByVocabularyId(ctx, req)
}

func (c Client) LongmanWordList(req *longman.ListRequest) (*longman.WordListResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := longman.NewLongmanServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return client.List(ctx, req)
}

func (c Client) WordList(req *word.ListRequest) (*word.WordListResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := word.NewWordServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return client.List(ctx, req)
}
