package util

import (
	"log"
	"mybatis-generator/constant"
	"os"
)

// CheckError 检查错误信息并打印出来
func CheckError(err error, info string) {
	if nil != err {
		log.Fatalln(info, err)
	}
}

// CreateFile 创建文件
func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	CheckError(err, constant.CreateFileFail+fileName)
	return file
}

// WriteLine 写入一行文件
func WriteLine(f *os.File, line string) {
	_, err := f.WriteString(line + constant.NewLine)
	CheckError(err, constant.WriteFileFail)
}

// AppendLine 对字符串追加一行
func AppendLine(f *string, line string) {
	*f = *f + line + constant.NewLine
}

// PrependLine 在开头追加
func PrependLine(f *string, line string) {
	*f = line + constant.NewLine + *f
}

// ToUpperWithSplitter 根据分隔符将分隔符后第一个字符转换为大写字符
// @param value 要分割的字符串
// @param splitter 分隔符
// @param first 第一个字母是否转换为大写字母
func ToUpperWithSplitter(value string, splitter string, first bool) string {
	bytes := []rune(value)
	flag := first
	var result string
	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		if splitter == string(b) {
			flag = true
			continue
		}
		if flag && b > 96 && b < 123 {
			b -= 32
			flag = false
		}
		result += string(b)
	}
	return result
}
