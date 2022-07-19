package datasqlite

//DataSqlite 接口
type DataSqlite interface {
	CreateTable(tableName string) error
	DeleleTable(tableName string) error
	InsertData(tableName string, destDataKeys []string, destDataValues []interface{}) error
	ReadData(tableName string, selectKeys []string, selectValues []interface{}, destDataKeys []string) ([][]string, error)
	UpData(tableName string, selectKeys []string, selectValues []interface{}, destDataKeys []string, destDataValues []interface{}) error
	DeleteData(tableName string, selectKeys []string, selectValues []interface{}) error
}
