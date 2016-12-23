package write

import (
	"fmt"
	"gobackup/mapper/util"
	"math/rand"
	. "mybatis-generator/constant"
	"mybatis-generator/initialization"
	"mybatis-generator/model"
	"strings"
	"time"
)

// ModelWrite 接口实现
type ModelWrite struct{}

var modelTabNum int

func (m *ModelWrite) Write(model *model.Model, ch chan string) {
	ch <- fmt.Sprintf("开始写入model- %s", time.Now().String())
	file := util.CreateFile(fmt.Sprintf("%s.java", model.Name))
	defer file.Close()
	var content string
	m.writeDoc(&content, model.Comment, GeneratorInfo+time.Now().Format(TimeFormat)+DocEnd)
	m.writeInfo(&content, DataAnnotation)
	m.writeInfo(&content, ClassDefine+model.Name+ImplementSerializable)
	m.writeSerialVersionUID(&content)
	for _, v := range *model.Fields {
		m.writeField(&content, v)
	}
	m.writeInfo(&content, BraceEnd)
	m.writePackage(&content, *initialization.Packages)

	file.WriteString(content)
	ch <- fmt.Sprintf("写入model完成- %s", time.Now().String())
}

func (m *ModelWrite) writeField(content *string, field model.Field) {
	if field.FieldType[2] == Yes && !strings.Contains(*content, field.FieldType[3]) {
		pushContent := Import + Blank + field.FieldType[3] + LineEnd
		util.PrependLine(content, pushContent)
	}
	m.writeDoc(content, field.Comment.String)
	m.writeInfo(content, Private+field.FieldType[1]+Blank+field.FieldName+LineEnd)
}

func (m *ModelWrite) writePackage(f *string, packageInfo string) {
	if EmptyString != packageInfo {
		util.PrependLine(f, Package+packageInfo+LineEnd+NewLine)
	}
}

func (m *ModelWrite) writeDoc(f *string, docs ...string) {
	tab := strings.Join(make([]string, modelTabNum), Blank)
	first := NewLine + tab + BackSlant + Star + Star
	util.AppendLine(f, first)
	for _, doc := range docs {
		writeDoc := tab + Blank + Star + Blank + doc
		util.AppendLine(f, writeDoc)
	}
	end := tab + Blank + Star + BackSlant
	util.AppendLine(f, end)
}

// 写入serialVersionUID
func (m *ModelWrite) writeSerialVersionUID(content *string) {
	uid := time.Now().UnixNano()

	flag := rand.New(rand.NewSource(uid)).Int63n(100)
	if flag > 10 {
		uid *= (flag / 10 * uid)
	}
	if flag > 50 {
		uid = 0 - uid
	}

	m.writeInfo(content, fmt.Sprintf(SerialVersionUID, uid))
}

// 写入具体信息,所有写入均调用此方法
func (m *ModelWrite) writeInfo(f *string, info string) {
	if strings.Contains(info, BraceEnd) {
		modelTabNum -= TabNum
	}

	info = strings.Join(make([]string, modelTabNum), Blank) + info
	util.AppendLine(f, info)
	if strings.Contains(info, BraceStart) {
		modelTabNum += TabNum
	}
}
