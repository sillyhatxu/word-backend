package dao

import (
	"bytes"
	"database/sql"
	"github.com/sillyhatxu/word-backend/dao/params"
	"github.com/sillyhatxu/word-backend/model"
)

const insertVocabularySQL = `
insert into vocabulary (word) VALUES 
`

func BatchAddVocabularyRule(inputArray []*model.Vocabulary) error {
	return client.Transaction(func(tx *sql.Tx) error {
		var args []interface{}
		var sqlBuffer bytes.Buffer
		sqlBuffer.WriteString(insertVocabularySQL)
		if inputArray != nil && len(inputArray) > 0 {
			for i, input := range inputArray {
				if i == 0 {
					sqlBuffer.WriteString("(?)")
				} else {
					sqlBuffer.WriteString(", (?)")
				}
				args = append(args, input.Word)
			}
		}
		_, err := tx.Exec(sqlBuffer.String(), args...)
		if err != nil {
			return err
		}
		return nil
	})
}

const findVocabularyListSQL = `
select * from vocabulary limit ?,?
`

func FindVocabularyList(params params.VocabularyListParams) ([]*model.Vocabulary, error) {
	var res []*model.Vocabulary
	err := client.FindList(findVocabularyListSQL, &res, params.Offset, params.Limit)
	return res, err
}

const countVocabularySQL = `
select count(1) from vocabulary
`

func CountVocabulary(params params.VocabularyListParams) (int64, error) {
	return client.Count(countVocabularySQL)
}
