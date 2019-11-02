package db

import (
  "conf"
  "fmt"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

type Mysql struct {
  Config conf.Config
  Db *sql.DB
}

func (mysql *Mysql) Connect() {
    connection, err := sql.Open("mysql", fmt.
      Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
        mysql.Config.DbUser,
        mysql.Config.DbPassword,
        mysql.Config.DbHost,
        mysql.Config.DbPort,
        mysql.Config.DbName))

  if err != nil {
    panic(err.Error())
  }

  mysql.Db = connection
  mysql.Db.Exec(fmt.
    Sprintf("set session time_zone = '%s'", mysql.Config.DbTimeZone))
}

func (mysql *Mysql) Select(sql string) *sql.Rows {
  if nil == (*mysql).Db {
    (*mysql).Connect()
  }

  rows, err := mysql.Db.Query(sql)

  if err != nil {
    panic(err.Error())
  }

  return rows
}

func (mysql *Mysql) Execute(query string, values ...interface {}) {
  if nil == (*mysql).Db {
    (*mysql).Connect()
  }

  prepare, err := (*mysql).Db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = prepare.Exec(values...)
  if err != nil {
    panic(err.Error())
  }
}

