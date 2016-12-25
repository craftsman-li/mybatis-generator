# mybatis-generator
mybatis xml生成脚本

## 简单、快速
window32位使用: bin/win86-mybatis-generator.exe
window64位使用: bin/win64-mybatis-generator.exe
linux使用: bin/linux-mybatis-generator
mac使用: bin/mac-mybatis-generator
### 支持的参数:
```json
  -db string
    	数据库名
  -f string
    	表名分隔符 (default "_")
  -h string
    	数据库地址(ip:port) (default "127.0.0.1:3306")
  -package string
    	包名(用于生成resultMap的type属性)
  -pass string
    	数据库密码
  -t string
    	数据表名
  -u string
    	数据库用户名 (default "root")
```