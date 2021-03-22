package main

import (
	"context"
	"fmt"
	"github.com/Naist4869/awesomeProject1/ent"
	"github.com/Naist4869/awesomeProject1/ent/issuex1355"
	"github.com/Naist4869/awesomeProject1/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func main() {

	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	addr := os.Getenv("DATABASE_ADDR")
	dbName := os.Getenv("DATABASE_DBNAME")


	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		dbName,
		true,
		"Local")
	client, err := ent.Open("mysql", dsn)
	if err!=nil{
		panic(err)
	}
	ctx := context.Background()
	if err := client.Debug().Schema.Create(ctx, migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		panic(err)
	}
	_, err = client.IssueX1355.Create().Save(ctx)
	if err!=nil{
		panic(err)
	}
	var v []V
	err = client.IssueX1355.Query().Select(issuex1355.FieldInt32, issuex1355.FieldStr, issuex1355.FieldTime).Scan(ctx, &v)
	if err!=nil{
		// sql/scan: failed scanning rows: sql: Scan error on column index 0, name "int32": converting NULL to int32 is unsupported
		log.Fatal(err)
	}
	// want []V{
	//		Time:  time.Time{},
	//		Int32: 0,
	//		Str:   "abc",
	//	}  with no error

}


type XTime time.Time

func (x *XTime) Scan(value interface{}) error {
	if value == nil {
		return nil // switch there, value haven't type
	}
	return nil
}


type V struct {
// Time holds the value of the "time" field.
Time time.Time `json:"time,omitempty"`
// Int32 holds the value of the "int32" field.
Int32 int32 `json:"int32,omitempty"`
// Str holds the value of the "str" field.
Str string `json:"str,omitempty"`
}