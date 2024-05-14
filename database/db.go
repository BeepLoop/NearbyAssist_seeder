package database

type Database interface {
	InitConnection() error
}
