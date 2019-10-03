package longmanservice

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/sillyhatxu/word-backend/dao"
	"github.com/sillyhatxu/word-backend/dao/params"
	"github.com/sillyhatxu/word-backend/grpc/longman"
)

func Add(request *longman.AddRequest) error {

	return nil
}

func Update(request *longman.UpdateRequest) error {
	return nil
}

func GetById(request *longman.GetByIdRequest) (*longman.Word, error) {
	return nil, nil
}

func GetByWord(request *longman.GetByWordRequest) (*longman.Word, error) {
	return nil, nil
}

func List(request *longman.ListRequest) ([]*longman.Word, int64, error) {
	p := params.LongmanWordListParams{Offset: request.Offset, Limit: request.Limit}
	totalCount, err := dao.CountLongman(p)
	if err != nil {
		return nil, 0, err
	}
	if totalCount == 0 {
		return make([]*longman.Word, 0), 0, nil
	}
	array, err := dao.FindLongmanList(p)
	if err != nil {
		return nil, 0, err
	}
	if array == nil || len(array) == 0 {
		return make([]*longman.Word, 0), 0, nil
	}
	var resultArray []*longman.Word
	for _, w := range array {
		createdTime, err := ptypes.TimestampProto(w.CreatedTime)
		if err != nil {
			return nil, 0, err
		}
		lastModifiedTime, err := ptypes.TimestampProto(w.LastModifiedTime)
		if err != nil {
			return nil, 0, err
		}
		resultArray = append(resultArray, &longman.Word{
			Id:               w.Id,
			VocabularyId:     w.VocabularyId,
			Description:      w.Description,
			CreatedTime:      createdTime,
			LastModifiedTime: lastModifiedTime,
		})
	}
	return resultArray, totalCount, nil
}
