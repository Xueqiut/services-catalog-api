package test

import (
  "fmt"
  "log"
  "strings"

  "database/sql"
  _ "github.com/lib/pq"
)

type Test struct {
  db	*sql.DB
}

func NewTest (connStr string) (Test, error) {
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    log.Fatal(err)
    return Test{}, err
  }
  pingErr := db.Ping()
  if pingErr != nil {
    log.Fatal(pingErr)
    return Test{}, err
  }

  return Test{db}, nil
}

func (t Test) GetDb () *sql.DB {
  return t.db
}

func (t Test) ProvisionService(data [][]interface{}) error {
  sql := "INSERT INTO services (name, description, user_id) VALUES"
  if data == nil {
    log.Fatal("data cannot be empty")
  }

  for _, values := range data {
    sql = fmt.Sprintf("%s ('%v', '%v', %v),", sql, values[0], values[1], values[2])
  }
  
  sql = strings.TrimSuffix(sql, ",")
  t.db.QueryRow(sql)

  return nil
}

func (t Test) ProvisionVersion(data [][]interface{}) error {
  sql := "INSERT INTO versions (name, service_id, enabled) VALUES"
  if data == nil {
    log.Fatal("data cannot be empty")
  }

  for _, values := range data {
    sql = fmt.Sprintf("%s ('%v', %v, %v),", sql, values[0], values[1], values[2])
  }

  sql = strings.TrimSuffix(sql, ",")
  t.db.QueryRow(sql)

  return nil
}

func (t Test) Truncate() error {
  t.db.QueryRow("TRUNCATE services, versions;")
  return nil
}