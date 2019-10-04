package enums

import (
	"fmt"
	"github.com/sillyhatxu/word-backend/grpc/longman"
)

type LongmanStatus string

const (
	LongmanStatusError   LongmanStatus = "ERROR"
	LongmanStatusSuccess LongmanStatus = "SUCCESS"
)

func SwitchLongmanStatus(status longman.Status) (LongmanStatus, error) {
	if status == longman.Status_Error {
		return LongmanStatusError, nil
	} else if status == longman.Status_Error {
		return LongmanStatusSuccess, nil
	} else {
		return "", fmt.Errorf("unsupported status")
	}
}
