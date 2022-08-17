package morm

import "go.mongodb.org/mongo-driver/mongo"

type Conn struct {
	DB *mongo.Database
}

func NewConn(db *mongo.Database) Conn {
	return Conn{DB: db}
}
