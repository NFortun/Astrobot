package db

import (
	"github.com/NFortun/Astrobot/config"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

var conn *pgx.Conn

func Connect() {

	var err error
	cfg := pgx.ConnConfig{
		Host:     config.Data.Addr,
		Port:     uint16(config.Data.Port),
		Database: config.Data.DBName,
		User:     config.Data.User,
		Password: config.Data.Passwd,
	}

	conn, err = pgx.Connect(cfg)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %v\n", err.Error())
	}

}
