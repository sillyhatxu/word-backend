package dao

import (
	"bufio"
	"fmt"
	"github.com/sillyhatxu/word-backend/dao/params"
	"github.com/sillyhatxu/word-backend/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestPage(t *testing.T) {
	array := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	rows := len(array)
	pageSize := 10
	pageSum := (rows + pageSize - 1) / pageSize
	for page := 0; page < pageSum; page++ {
		start := page * pageSize
		end := (page + 1) * pageSize
		if end > rows {
			end = rows
		}
		fmt.Println(array[start:end])
		//err := BatchAddVocabularyRule()
		//if err != nil {
		//	panic(err)
		//}
	}
}
func TestBatchAddVocabularyRule(t *testing.T) {
	//filePath := "/Users/shikuanxu/go/src/github.com/sillyhatxu/word-backend/txt/words_alpha.txt"
	filePath := "/Users/cookie/go/gopath/src/github.com/sillyhatxu/word-backend/txt/words_alpha.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var inputArray []*model.Vocabulary
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArray = append(inputArray, &model.Vocabulary{Word: scanner.Text()})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rows := len(inputArray)
	pageSize := 5000
	pageSum := (rows + pageSize - 1) / pageSize
	logrus.Infof("rows : %d", rows)
	logrus.Infof("pageSize : %d", pageSize)
	logrus.Infof("pageSum : %d", pageSum)
	for page := 0; page < pageSum; page++ {
		logrus.Infof("page : %d", page)
		start := page * pageSize
		end := (page + 1) * pageSize
		if end > rows {
			end = rows
		}
		err := BatchAddVocabularyRule(inputArray[start:end])
		if err != nil {
			panic(err)
		}
	}
}

func TestFindVocabularyList(t *testing.T) {
	p := params.VocabularyListParams{Offset: 0, Limit: 10}
	array, err := FindVocabularyList(p)
	assert.Nil(t, err)
	for _, v := range array {
		logrus.Infof("created time : %v;%#v", v.CreatedTime.Format("2006-01-02T15:04:05"), v)
	}
}
