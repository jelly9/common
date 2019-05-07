package mysql

import (
	"fmt"
	"database/sql"

    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

type Options struct {
	DriverName	string
	Source 		string
	DefaultTable string
	ShowSQL 	bool
}

type MySQLDao struct{
	engine		*xorm.Engine
	defaultTb 	string
}

func NewMySQLDao (opt *Options) *MySQLDao {
	var driver string
	if opt.DriverName == "" {
		driver = "mysql"
	} else {
		driver = opt.DriverName
	}
	engine, err := xorm.NewEngine(driver, opt.Source)
	if err != nil {
		fmt.Println("error:", err)
		panic("MySQL init failed")
	}
	engine.ShowSQL(opt.ShowSQL)
	
	return &MySQLDao{
		engine:		engine,
		defaultTb:	opt.DefaultTable,
	}
}

// Where
func (dao *MySQLDao) Where(query interface{}, args ...interface{}) *xorm.Session {
	return dao.engine.Where(query, args...)
}

// Exec
func (dao *MySQLDao) Exec(sql string, args ...interface{}) (sql.Result, error) {
    params := make([]interface{}, len(args)+1)
    params[0] = sql
	copy(params[1:], args)
	return dao.engine.Exec(params...)
}

// Table
func (dao *MySQLDao) Table(tbname string) *xorm.Session {
    return dao.engine.Table(tbname)
}

// Cols
func (dao *MySQLDao) Cols(columns ...string) *xorm.Session {
    return dao.engine.Cols(columns...)
}

// Get
func (dao *MySQLDao) Get(bean interface{}) (bool, error) {
    return dao.engine.Get(bean)
}

// Find
func (dao *MySQLDao) Find(beans interface{}, condiBeans ...interface{}) error {
    return dao.engine.Find(beans, condiBeans...)
}

// Insert
func (dao *MySQLDao) Insert(beans ...interface{}) (int64, error) {
    return dao.engine.Insert(beans)
}

// Update
func (dao *MySQLDao) Update(beans ...interface{}) (int64, error) {
    return dao.engine.Update(beans...)
}

// NewSession
func (dao *MySQLDao) NewSession() *xorm.Session {
    return dao.engine.NewSession().Table(dao.defaultTb)
}

// ShowSQL
func (dao *MySQLDao) ShowSQL(b bool) {
    dao.engine.ShowSQL(b)
}