package morm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModelI interface {
	Create(d interface{}) (*mongo.InsertOneResult, error)
	CreateMany(d interface{}, r []interface{}) (*mongo.InsertManyResult, error)
	FindOne(d interface{}) error
	FindOneBy(d, q interface{}) error
	FindManyBy(d, q interface{}, o *options.FindOptions) (*mongo.Cursor, error)
	CountBy(d, q interface{}) (int64, error)
	UpdateOne(d interface{}) (*mongo.UpdateResult, error)
	UpdateOneBy(d, q, s interface{}, o *options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateManyBy(d, q, s interface{}, o *options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(d interface{}) (*mongo.DeleteResult, error)
	DeleteBy(d, q interface{}) (*mongo.DeleteResult, error)
}

func FindOne(m ModelI) error {
	return m.FindOne(m)
}

func FindOneBy(m ModelI, q interface{}) error {
	return m.FindOneBy(m, q)
}

func FindManyBy(m ModelI, q interface{}, o *options.FindOptions) (*mongo.Cursor, error) {
	return m.FindManyBy(m, q, o)
}

func CountBy(m ModelI, q interface{}) (int64, error) {
	return m.CountBy(m, q)
}

func UpdateOne(m ModelI) (*mongo.UpdateResult, error) {
	return m.UpdateOne(m)
}

func UpdateOneBy(m ModelI, q, s interface{}, o *options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.UpdateOneBy(m, q, s, o)
}

func UpdateManyBy(m ModelI, q, s interface{}, o *options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.UpdateManyBy(m, q, s, o)
}

func Create(m ModelI) (*mongo.InsertOneResult, error) {
	return m.Create(m)
}

func CreateMany(m ModelI, d []interface{}) (*mongo.InsertManyResult, error) {
	return m.CreateMany(m, d)
}

func DeleteOne(m ModelI) (*mongo.DeleteResult, error) {
	return m.DeleteOne(m)
}

func DeleteBy(m ModelI, q interface{}) (*mongo.DeleteResult, error) {
	return m.DeleteBy(m, q)
}
