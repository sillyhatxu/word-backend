package dao

import (
	"github.com/sillyhatxu/word-backend/dao/params"
	"github.com/sillyhatxu/word-backend/model"
)

const insertLongmanSQL = `
INSERT INTO longman_word (vocabulary_id, description, status) VALUES(? , ? , ?)
`

func InsertLongman(input *model.Longman) (err error) {
	_, err = client.Insert(insertLongmanSQL, input.VocabularyId, input.Description, input.Status)
	return err
}

const updateLongmanSQL = `
update longman_word set description = ? where id = ?
`

func UpdateLongman(input *model.Longman) error {
	_, err := client.Update(updateLongmanSQL, input.Description, input.Id)
	return err
}

const findLongmanListSQL = `
select * from longman_word limit ?,?
`

func FindLongmanList(p params.LongmanWordListParams) ([]*model.Longman, error) {
	var res []*model.Longman
	err := client.FindFirst(findLongmanListSQL, &res, p.Offset, p.Limit)
	return res, err
}

const countLongmanSQL = `
select count(1) from longman_word
`

func CountLongman(p params.LongmanWordListParams) (int64, error) {
	return client.Count(countLongmanSQL)
}
