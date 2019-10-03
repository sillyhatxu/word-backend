package dao

import (
	"github.com/sillyhatxu/word-backend/dao/params"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLongmanList(t *testing.T) {
	p := params.LongmanWordListParams{Offset: 0, Limit: 10}
	array, err := FindLongmanList(p)
	assert.Nil(t, err)
	logrus.Infof("%#v", array)
}
