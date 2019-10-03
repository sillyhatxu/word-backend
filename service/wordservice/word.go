package wordservice

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/sillyhatxu/word-backend/dao"
	"github.com/sillyhatxu/word-backend/dao/params"
	"github.com/sillyhatxu/word-backend/grpc/word"
)

func List(request *word.ListRequest) ([]*word.Word, int64, error) {
	p := params.VocabularyListParams{Offset: request.Offset, Limit: request.Limit}
	totalCount, err := dao.CountVocabulary(p)
	if err != nil {
		return nil, 0, err
	}
	if totalCount == 0 {
		return make([]*word.Word, 0), 0, nil
	}
	array, err := dao.FindVocabularyList(p)
	if err != nil {
		return nil, 0, err
	}
	if array == nil || len(array) == 0 {
		return make([]*word.Word, 0), 0, nil
	}
	var resultArray []*word.Word
	for _, w := range array {
		createdTime, err := ptypes.TimestampProto(w.CreatedTime)
		if err != nil {
			return nil, 0, err
		}
		lastModifiedTime, err := ptypes.TimestampProto(w.LastModifiedTime)
		if err != nil {
			return nil, 0, err
		}
		resultArray = append(resultArray, &word.Word{
			Id:               w.Id,
			Word:             w.Word,
			CreatedTime:      createdTime,
			LastModifiedTime: lastModifiedTime,
		})
	}
	return resultArray, totalCount, nil
}
