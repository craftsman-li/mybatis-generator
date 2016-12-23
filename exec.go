package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mybatis-generator/constant"
	. "mybatis-generator/initialization"
	"mybatis-generator/model"
	"mybatis-generator/util"
	"mybatis-generator/write"
	"strings"
)

var Mapping = [...][4]string{
	{"VARCHAR", "String", "NO", "java.lang.String"},
	{"CHAR", "String", "NO", "java.lang.String"},
	{"BLOB", "byte[]", "NO", "java.lang.byte"},
	{"TEXT", "String", "NO", "java.lang.String"},
	{"INTEGER", "Long", "NO", "java.lang.Long"},
	{"TINYINT", "Integer", "NO", "java.lang.Integer"},
	{"SMALLINT", "Integer", "NO", "java.lang.Integer"},
	{"MEDIUMINT", "Integer", "NO", "java.lang.Integer"},
	{"BIT", "Boolean", "NO", "java.lang.Boolean"},
	// 替换成Long,标准为BigInteger
	{"BIGINT", "Long", "NO", "java.lang.Long"},
	// {"BIGINT", "BigInteger", "YES", "java.math.BigInteger"},
	{"FLOAT", "Float", "NO", "java.lang.Float"},
	{"DOUBLE", "Double", "NO", "java.lang.Double"},
	{"DECIMAL", "BigDecimal", "YES", "java.math.BigDecimal"},
	{"BOOLEAN", "Integer", "NO", "java.lang.Integer"},
	{"DATE", "Date", "YES", "java.util.Date"},
	{"TIME", "Time", "YES", "java.sql.Time"},
	{"DATETIME", "Date", "YES", "java.sql.Date"},
	// 	{"DATETIME", "Timestamp", "YES", "java.sql.Timestamp"},
	{"TIMESTAMP", "Timestamp", "YES", "java.sql.Timestamp"},
	{"YEAR", "Date", "YES", "java.sql.Date"},
}

func main() {
	Init()
	defer DB.Close()
	fields := getcolumnInfo(DB, Table)
	model := &model.Model{Fields: fields}
	model.Name = util.ToUpperWithSplitter(*Table, *Splitter, true)
	model.TableName = *Table
	model.Comment = getTableInfo(DB, Table)
	var mapperWrite write.Write = &write.MapperWrite{}
	var modelWrite write.Write = &write.ModelWrite{}
	var ch = make(chan string, 1)
	go mapperWrite.Write(model, ch)
	go modelWrite.Write(model, ch)

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}

// 获取表字段信息
func getcolumnInfo(db *sql.DB, table *string) *[]model.Field {
	result, err := db.Query(fmt.Sprintf("SHOW FULL columns FROM `%s`", *table))
	util.CheckError(err, constant.ShowFullColumnFailed)
	var fields []model.Field
	for result.Next() {
		var column model.Column
		result.Scan(&column.Name, &column.Type, &column.Collation, &column.Null, &column.Key, &column.Default, &column.Extra, &column.Privileges, &column.Comment)

		checkFiledAllowNull(column)
		var fieldMeta model.FieldMeta
		fieldMeta.FieldName = util.ToUpperWithSplitter(column.Name, *Splitter, false)
		fieldMeta.FieldType = mappingMYSQLToJava(column.Type)
		m := model.Field{Column: column, FieldMeta: fieldMeta}
		fields = append(fields, m)
	}
	return &fields
}

func getTableInfo(db *sql.DB, table *string) string {
	var tableName string
	var sql string
	err := db.QueryRow(fmt.Sprintf("show create table `%s`", *table)).Scan(&tableName, &sql)
	util.CheckError(err, constant.ShowFullColumnFailed)
	temp := strings.Split(sql, "=")
	length := len(temp)
	if length > 0 && strings.Contains(temp[length - 2], "COMMENT") {
		return strings.Replace(temp[len(temp) - 1], "'", "", -1)
	}
	return constant.EmptyString
}

// 最字段检查是否允许为空,如果允许为空,给出警告提示
// http://www.360doc.com/content/15/0208/22/1073512_447324234.shtml
// http://blog.csdn.net/z69183787/article/details/53318418
// 应该指定列为NOT NULL，除非你想存储NULL。在MySQL中，含有空值的列很难进行查询优化
// 因为它们使得索引、索引的统计信息以及比较运算更加复杂。你应该用0、一个特殊的值或者一个空串代替空值
func checkFiledAllowNull(column model.Column) {
	if model.YesOrNo.String(0) == column.Null {
		log.Printf("字段(%s)允许为空,如非必要,建议设置成NOT NULL.\n", column.Name)
	}
	// if !column.Default.Valid {
	// 	log.Printf("字段(%s)默认值为NULL,可调优.\n", column.Name)
	// }
}

// 将数据库字段转换为Java字段
func mappingMYSQLToJava(t string) [4]string {
	for _, v := range Mapping {
		if strings.Contains(strings.ToUpper(t), v[0]) {
			return v
		}
	}
	return [4]string{}
}
