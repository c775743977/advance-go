package main

import (
	"fmt"
)

type Option func(*DB)

type DB struct {
	Host string
	Port int
}

func WithHost(host string) Option {
	return func(db *DB) {
		db.Host = host
	}
}

func WithPort(port int) Option {
	return func(db *DB) {
		db.Port = port
	}
}

func NewDB(options ...Option) *DB {
	db := &DB{}
	for _, k := range options {
		k(db)
	}

	return db
}

func main() {
	db := NewDB(WithHost("localhost"), WithPort(3306))

	fmt.Println(db)
}