package generator

import (
	"errors"
	"strings"
)

var DefaultMysqlTypeMap map[string]string = map[string]string{
	"int":       "int",
	"int8":      "tinyint",
	"int16":     "smallint",
	"int32":     "int",
	"int64":     "bigint",
	"uint8":     "tinyint unsigned",
	"uint16":    "smallint unsigned",
	"uint32":    "int unsigned",
	"uint64":    "bigint unsigned",
	"float32":   "float",
	"float64":   "double",
	"string":    "varchar(255)",
	"time.Time": "timestamp",
}
var DefaultMysqlDefaultValueMap map[string]string = map[string]string{
	"int":       "0",
	"int8":      "0",
	"int16":     "0",
	"int32":     "0",
	"int64":     "0",
	"uint8":     "0",
	"uint16":    "0",
	"uint32":    "0",
	"uint64":    "0",
	"float32":   "0",
	"float64":   "0",
	"string":    "''",
	"time.Time": "0",
}

type FieldDescriptoin struct {
	FieldName         string
	FieldGoType       string
	TagString         string
	MysqlTagFieldList TagFieldList
}

func (fd FieldDescriptoin) IsPK() bool {
	return fd.MysqlTagFieldList.Contains("pk")
}
func (fd FieldDescriptoin) GetMysqlType() string {
	mysqlType := fd.MysqlTagFieldList.GetValue("type")
	if mysqlType != "" {
		return mysqlType
	}
	return DefaultMysqlTypeMap[fd.FieldGoType]
}
func (fd FieldDescriptoin) GetMysqlDefalutValue() string {
	mysqlDefault := fd.MysqlTagFieldList.GetValue("default")
	if mysqlDefault != "" {
		return mysqlDefault
	}
	return DefaultMysqlDefaultValueMap[fd.FieldGoType]

}
func (fd FieldDescriptoin) GetMysqlFieldName() string {
	mysqlFieldName := fd.MysqlTagFieldList.GetValue("name")
	if mysqlFieldName != "" {
		return mysqlFieldName
	}
	return fd.FieldName
}

type StructDescription struct {
	StructName string
	Fields     []FieldDescriptoin
}

func (sd StructDescription) GetPK() (pkList []string) {
	for _, field := range sd.Fields {
		if field.IsPK() {
			pkList = append(pkList, field.GetMysqlFieldName())
		}
	}
	return

}

func (sd StructDescription) GenerateCreateTableSql() (sql string, err error) {
	if len(sd.Fields) == 0 {
		return "", errors.New("no filed found ,generate create table sql error")
	}
	sql += "create table if not exists `" + sd.StructName + "`(\n"
	for idx, fieldD := range sd.Fields {
		sql += "`" + fieldD.GetMysqlFieldName() + "` " + fieldD.GetMysqlType() + " NOT NULL DEFAULT " + fieldD.GetMysqlDefalutValue()
		if idx != len(sd.Fields)-1 {
			sql += ",\n"
		} else {
			sql += "\n"
		}
	}
	pkList := sd.GetPK()
	if len(pkList) != 0 {
		sql += ",primary key (" + strings.Join(pkList, ",") + ")\n"
	}

	sql += ");"
	return
}
