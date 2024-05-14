package database

import (
	"fmt"

	"github.com/BeepLoop/nearbyassist_seeder/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	conf *config.Config
	Conn *sqlx.DB
}

func NewMysqlDatabase(conf *config.Config) *Mysql {
	return &Mysql{
		conf: conf,
	}
}

func (m *Mysql) InitConnection() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.conf.DB_User, m.conf.DB_Pass, m.conf.DB_Host, m.conf.DB_Port, m.conf.DB_Name)

	if conn, err := sqlx.Connect("mysql", dsn); err != nil {
		return err
	} else {
		m.Conn = conn
	}

	return nil
}
