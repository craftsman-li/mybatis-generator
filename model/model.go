package model

import "database/sql"

// Column 数据表
type Column struct {
	Name string // 字段名
	Type string // 字段类型
	Null string // 是否允许为null
	nullField
	Privileges string // 授权 select,insert,update,references
}

// NullField 对数据表描述中可为null的字段进行处理
// golang 不允许字段出现null,故碰到null字段需要单独处理
type nullField struct {
	Collation sql.NullString // 定序名称
	Key       sql.NullString // PRI   primary key  表示主键,唯一/uni   UNIQUE  表示唯一/mul  添加了索引/显示顺序  PRI>UNI>MUL
	Default   sql.NullString // 默认值
	Extra     sql.NullString //
	Comment   sql.NullString // 描述
}

// FieldMeta 字段实体类
type FieldMeta struct {
	FieldName string    // 属性名
	FieldType [4]string // 属性类型
}

// Field 根据数据表列名生成的实体类
type Field struct {
	Column
	FieldMeta
}

// Model 表对象
type Model struct {
	Fields    *[]Field // 字段列表
	Name      string   // 实体类名
	TableName string   // 表名
	Comment   string   // Comment 信息
}

// YesOrNo 是否允许
type YesOrNo int

const (
	// YES 允许为空
	YES = iota
	// NO 不允许为空
	NO

	//
	VARCHAR = iota
	CHAR
	BLOB
	TEXT
	INTEGER
	TINYINT
	SMALLINT
	MEDIUMINT
	BIT
	BIGINT
	FLOAT
	DOUBLE
	DECIMAL
	BOOLEAN
	DATE
	TIME
	DATETIME
	TIMESTAMP
	YEAR
)

var yesOrNo = [...]string{
	"YES",
	"NO",
}

func (y YesOrNo) String() string {
	return yesOrNo[y]
}
