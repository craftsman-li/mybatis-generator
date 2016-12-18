package write

import (
	"fmt"
	"mybatis-generator/model"
	"time"
)

type ModelWrite struct{}

func (m ModelWrite) Write(model *model.Model, ch chan string) {
	ch <- fmt.Sprintf("开始写入model- %s\n", time.Now().String())
	fmt.Println("fff")
	ch <- fmt.Sprintf("写入model完成- %s\n", time.Now().String())
}
