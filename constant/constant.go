package constant

// 常量文件
const (
	NewLine               = "\n"
	CreateFileFail        = "创建文件失败"
	LinkDBFail            = "连接数据库失败"
	WriteFileFail         = "写入文件失败"
	ShowFullColumnFailed  = "获取表字段失败"
	EmptyString           = ""
	Tab                   = "    "
	BraceStart            = "{"
	BraceEnd              = "}"
	Private               = "private "
	Public                = "public "
	SerialVersionUID      = Private + "static final long serialVersionUID = %dL;"
	TimeFormat            = "2006/01/02 15:04:05"
	GeneratorInfo         = "Created by generator on "
	DocEnd                = "."
	Package               = "package "
	LineEnd               = ";"
	DataAnnotation        = "@Data"
	Blank                 = " "
	ClassDefine           = Public + "class "
	ImplementSerializable = " implements Serializable " + BraceStart
	Star                  = "*"
	BackSlant             = "/"
	Yes                   = "YES"
	Import                = "import"
	TabNum                = 5
	Lt                    = "<"
	Gt                    = ">"
	DoubleQuotes          = "\""
	CreateTimeDesc        = "创建时间"
	UpdateTimeDesc        = "更新时间"
)
