package db

import (
	"fmt"
	"go-api-report2/config"
	"go-api-report2/ent"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username, password, host, db string
	port                         int
)

func CreateConnection() (*ent.Client, error) {

	dbConfig := config.RootConfig.Database

	username := dbConfig.Username
	password := dbConfig.Password
	host := dbConfig.Endpoint
	db := dbConfig.Name
	port := dbConfig.Port

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, db, "parseTime=true")

	client, err := ent.Open("mysql", connectString)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// func CreateConnection2() (*ent.Client, error) {
// 	client, err := ent.Open("mysql", "root:some-secret@tcp(localhost:3306)/reportserver?parseTime=true")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }
