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

type Engine struct{
	engine		*xorm.Engine
	defaultTb 	string
}

func NewEngine (opt *Options) *Engine {
	driver := opt.DriverName
	if driver == "" {
		driver = "mysql"
	}
	engine, err := xorm.NewEngine(driver, opt.Source)
	if err != nil {
		fmt.Println("error:", err)
		panic("MySQL init failed")
	}
	engine.ShowSQL(opt.ShowSQL)
	return &Engine{
		engine:		engine,
		defaultTb:	opt.DefaultTable,
	}
}

// Where
func (e *Engine) Where(query interface{}, args ...interface{}) *xorm.Session {
	return e.engine.Where(query, args...)
}

// Exec
func (e *Engine) Exec(sql string, args ...interface{}) (sql.Result, error) {
    params := make([]interface{}, len(args)+1)
    params[0] = sql
	copy(params[1:], args)
	return e.engine.Exec(params...)
}

// Table
func (e *Engine) Table(tbname string) *xorm.Session {
    return e.engine.Table(tbname)
}

// Cols
func (e *Engine) Cols(columns ...string) *xorm.Session {
    return e.engine.Cols(columns...)
}

// Get
func (e *Engine) Get(bean interface{}) (bool, error) {
    return e.engine.Get(bean)
}

// Find
func (e *Engine) Find(beans interface{}, condiBeans ...interface{}) error {
    return e.engine.Find(beans, condiBeans...)
}

// Insert
func (e *Engine) Insert(beans ...interface{}) (int64, error) {
    return e.engine.Insert(beans)
}

// Update
func (e *Engine) Update(beans ...interface{}) (int64, error) {
    return e.engine.Update(beans...)
}

// NewSession
func (e *Engine) NewSession(tbname ...string) *xorm.Session {
	tb := e.defaultTb
	if len(tbname) != 0 {
		tb = tbname[0]
	}
	return e.engine.NewSession().Table(tb)
}

// ShowSQL
func (e *Engine) ShowSQL(b bool) {
    e.engine.ShowSQL(b)
}