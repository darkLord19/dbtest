package main

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func main() {
  connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    panic(err)
  }
  query := `
    SELECT
      table_name
    FROM
      information_schema.tables
    WHERE
      table_schema='public'
    AND
      table_type='BASE TABLE';
  `
  rows, err := db.Query(query)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var table_name string
    if err := rows.Scan(&table_name); err != nil {
      panic(err)
    }
    fmt.Println(table_name)
  }
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS test (col1 varchar(10), col2 int, col3 bigint)")
  if err != nil {
    panic(err)
  }
  _, err = db.Exec(`INSERT INTO test(col1, col2, col3) VALUES ($1, $2, $3)`,"hello", 1, 111)
  if err != nil {
    panic(err)
  }
  rows, err = db.Query("SELECT * FROM test")
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    type test struct {
      col1 string
      col2 int
      col3 int64
    }
    var testdata test
    if err := rows.Scan(&testdata.col1, &testdata.col2, &testdata.col3); err != nil {
        panic(err)
    }
    fmt.Println(testdata)
  }
  _, err = db.Exec("DELETE FROM test WHERE col1=$1", "hello")
  if err != nil {
    panic(err)
  }
}
