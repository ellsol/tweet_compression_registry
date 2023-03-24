package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type App struct {
	Controller Controller

	DB *sql.DB
}

func NewApp() (*App, error) {
	c, err := NewConfig()
	if err != nil {
		return nil, err
	}

	db, err := OpenSqlDB(c)
	if err != nil {
		return nil, err
	}
	service := NewChecksumService(db)

	return &App{
		Controller: NewController(&service),
		DB:         db,
	}, nil
}

func OpenSqlDB(c *Config) (*sql.DB, error) {
	fmt.Println(c.Info.Info())
	db, err := sql.Open("postgres", c.Info.Info())

	if err != nil {
		return nil, err
	}

	//_, err = db.Query("select * from auththree.user")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	db.SetMaxIdleConns(20)
	return db, nil
}
