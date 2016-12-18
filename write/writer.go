package write

import "mybatis-generator/model"

// Write 写接口
type Write interface {
	Write(m *model.Model, ch chan string)
}
