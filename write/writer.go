package write

import "mybatis-generator/model"

// Write 写接口
type Write interface {
	Write(m *model.Model, ch chan string)

	// 在写入文件的过程中,需要根据表示符进行缩进,自行实现
	writeInfo(f *string, info string)
}
