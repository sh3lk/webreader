package datastore

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"webreader/config"
	"webreader/ent"
)

// New returns data source name
func New() string {
	mc := mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
			"charset":   config.C.Database.Params.Charset,
			"loc":       config.C.Database.Params.Loc,
		},
	}

	return mc.FormatDSN()
}

// NewClient returns an orm client
func NewClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	//dsn := New()

	//return ent.Open(dialect.MySQL, dsn, entOptions...)
	return ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
}
