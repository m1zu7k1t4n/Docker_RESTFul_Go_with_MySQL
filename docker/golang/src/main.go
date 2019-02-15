package main

import (
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open("mysql","user:password@protocol(host:port)/dbname")
  db, err := sql.Open("mysql", "test:test@tcp(mysql_host:3306)/test")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close() // 関数がリターンする直前に呼び出される

  // ***************************
  // create table (DDL)
  // ***************************

  _, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS user(
      id INT PRIMARY KEY AUTO_INCREMENT,
      name VARCHAR(100),
      address VARCHAR(1000),
      created_at TIMESTAMP NOT NULL default current_timestamp,
      updated_at TIMESTAMP NOT NULL default current_timestamp on update current_timestamp
    );
  `)
  if err != nil {
    panic(err.Error())
  }

  // ***************************
  // test database insert
  // ***************************
  _, err = db.Exec(`
    INSERT INTO user(name, address) VALUES('takashi', 'nihon-1-1-1')
  `)
  if err != nil {
    panic(err.Error())
  }


  // ***************************
  // test database select
  // ***************************
  rows, err := db.Query("SELECT * FROM tasks;") //
  if err != nil {
    panic(err.Error())
  }

  columns, err := rows.Columns() // カラム名を取得
  if err != nil {
    panic(err.Error())
  }

  values := make([]sql.RawBytes, len(columns))

  //  rows.Scan は引数に `[]interface{}`が必要.

  scanArgs := make([]interface{}, len(values))
  for i := range values {
    scanArgs[i] = &values[i]
  }

  for rows.Next() {
    err = rows.Scan(scanArgs...)
    if err != nil {
      panic(err.Error())
    }

    var value string
    for i, col := range values {
      // Here we can check if the value is nil (NULL value)
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      fmt.Println(columns[i], ": ", value)
    }
    fmt.Println("-----------------------------------")
  }

  rows, err = db.Query("SELECT * FROM user;") //
  if err != nil {
    panic(err.Error())
  }

  columns, err = rows.Columns() // カラム名を取得
  if err != nil {
    panic(err.Error())
  }

  values = make([]sql.RawBytes, len(columns))

  //  rows.Scan は引数に `[]interface{}`が必要.

  scanArgs = make([]interface{}, len(values))
  for i := range values {
    scanArgs[i] = &values[i]
  }

  for rows.Next() {
    err = rows.Scan(scanArgs...)
    if err != nil {
      panic(err.Error())
    }

    var value string
    for i, col := range values {
      // Here we can check if the value is nil (NULL value)
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      fmt.Println(columns[i], ": ", value)
    }
    fmt.Println("-----------------------------------")
  }

}
