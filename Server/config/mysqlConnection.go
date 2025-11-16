package config

// import (
// 	"database/sql"
// 	"fmt"
// )

// //type DB *sql.DB

// func InitDB() (*sql.DB, error) {
// 	// MySQL configuration
// 	//dbConfig := "root:root123@tcp(142.171.228.250:3306)/k8s?charset=utf8mb4&parseTime=True&loc=Local"
// 	dbConfig := "root:86868686mM@tcp(192.168.1.213:3306)/k8s?charset=utf8mb4&parseTime=True&loc=Local"

// 	// open database connection
// 	db, err := sql.Open("mysql", dbConfig)
// 	if err != nil {
// 		return nil, fmt.Errorf("Database connection error: %v", err)
// 	}
// 	defer db.Close()
// 	// test the database connection
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("Error is : ", err)
// 		return nil, fmt.Errorf("Failure to connect to database: %v", err)
// 	}

// 	return db, nil
// }
