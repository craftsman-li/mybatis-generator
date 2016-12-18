package write

import (
	"fmt"
	"mybatis-generator/model"
	"time"
)

type MapperWrite struct{}

func (w MapperWrite) Write(model *model.Model, ch chan string) {
	ch <- fmt.Sprintf("开始写入Mapper- %s\n", time.Now().String())

	fmt.Println("ddd")
	ch <- fmt.Sprintf("写入Mapper完成- %s\n", time.Now().String())
}
