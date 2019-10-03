package dao

import "fmt"

const (
	UserName = "sillyhat_word"
	Password = "sillyhat_word"
	Host     = "192.168.1.87"
	Port     = "3306"
	Schema   = "sillyhat_word"
)

func init() {
	err := Initial(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", UserName, Password, Host, Port, Schema), "")
	//err := Initial(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Asia/Singapore", UserName, Password, Host, Port, Schema))
	//err := Initial(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Asia%2FSingapore", UserName, Password, Host, Port, Schema))
	if err != nil {
		panic(err)
	}
}
