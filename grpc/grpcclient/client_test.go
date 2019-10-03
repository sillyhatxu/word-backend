package grpcclient

import (
	"github.com/sillyhatxu/word-backend/grpc/word"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	address = "localhost:8080"
	timeout = 30 * time.Second
)

func TestClient_WordList(t *testing.T) {
	client := NewGRPCWordClient(address, Timeout(timeout))
	res, err := client.WordList(&word.ListRequest{Offset: 0, Limit: 10})
	assert.Nil(t, err)
	logrus.Infof("%d", len(res.Data))
	logrus.Infof("%#v", res)
}
