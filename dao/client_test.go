package dao

import (
	"fmt"
	"testing"
)

//const (
//	UserName = "sillyhat_word"
//	Password = "sillyhat_word"
//	Host     = "192.168.1.87"
//	Port     = "3306"
//	Schema   = "sillyhat_word"
//)
//const (
//	UserName = "sillyhat_word"
//	Password = "sillyhat_word"
//	Host     = "161.117.82.136"
//	Port     = "3306"
//	Schema   = "sillyhat_word"
//)
const (
	UserName = "sillyhat_word"
	Password = "sillyhat_word"
	Host     = "127.0.0.1"
	Port     = "3308"
	Schema   = "sillyhat_word"
)

func init() {
	err := Initial(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", UserName, Password, Host, Port, Schema)+"?loc=Asia%2FSingapore&parseTime=true", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/word-backend/db/migration")
	//err := Initial(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Asia/Singapore", UserName, Password, Host, Port, Schema))
	//err := Initial(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Asia%2FSingapore", UserName, Password, Host, Port, Schema))
	if err != nil {
		panic(err)
	}
}

func TestInitial(t *testing.T) {

}
