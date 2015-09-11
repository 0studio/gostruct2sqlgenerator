package example

// go build main.go -o gostruct2sqlgenerator
// ./gostruct2sqlgenerator example/example_1.go

import (
	"time"
)

type Hello struct {
	Id   int    `json:"id" mysql:"pk,defalut=1,type=bigint"` // id
	Name string `mysql:"pk,defalut='hello',name=helloName,type=varchar(10)"`
	Age  int
	T    time.Time `mysql:"type=int"`
	T2   time.Time
}
type World struct {
	Id int
}
