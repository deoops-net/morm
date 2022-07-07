package morm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
	"time"
)

var db *mongo.Database

type Model struct {
	ID        string `json:"id" bson:"_id,omitempty" url:"id,omitempty"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt,omitempty" url:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt,omitempty" url:"updatedAt,omitempty"`
}

func Init(d *mongo.Database) {
	db = d
}

func (m *Model) FindOneBy(d interface{}, query interface{}) error {
	tableName := getTableName(d)

	return db.Collection(tableName).FindOne(context.Background(), query).Decode(d)
}

func (m *Model) FindManyBy(d interface{}, q interface{}, o *options.FindOptions) (*mongo.Cursor, error) {
	tableName := getTableName(d)
	return db.Collection(tableName).Find(context.Background(), q, o)
}

func (m *Model) CountBy(d interface{}, q interface{}) (int64, error) {
	tableName := getTableName(d)
	return db.Collection(tableName).CountDocuments(context.Background(), q)
}

func (m *Model) Create(d interface{}) (*mongo.InsertOneResult, error) {
	tableName := getTableName(d)
	m.SetCreateFields()

	return db.Collection(tableName).InsertOne(context.Background(), d)
}

func (m *Model) CreateMany(d interface{}, r []interface{}) (*mongo.InsertManyResult, error) {
	tableName := getTableName(d)
	// TODO add createdAt
	return db.Collection(tableName).InsertMany(context.Background(), r)
}

func (m *Model) FindOne(d interface{}) error {
	tableName := getTableName(d)
	id, err := primitive.ObjectIDFromHex(m.ID)
	if err != nil {
		return err
	}

	return db.Collection(tableName).FindOne(context.Background(), map[string]interface{}{"_id": id}).Decode(d)
}

func (m *Model) DeleteOne(d interface{}) (*mongo.DeleteResult, error) {
	tableName := getTableName(d)
	id, err := primitive.ObjectIDFromHex(m.ID)
	if err != nil {
		return nil, err
	}

	return db.Collection(tableName).DeleteOne(context.Background(), map[string]interface{}{"_id": id})
}

func (m *Model) DeleteBy(d, q interface{}) (*mongo.DeleteResult, error) {
	tableName := getTableName(d)

	return db.Collection(tableName).DeleteMany(context.Background(), q)
}

func (m *Model) UpdateOne(d interface{}) (*mongo.UpdateResult, error) {
	tableName := getTableName(d)
	id, err := primitive.ObjectIDFromHex(m.ID)
	if err != nil {
		return nil, err
	}

	query := bson.M{"_id": id}
	set := bson.M{"$set": d}
	m.SetUpdateFields()

	return db.Collection(tableName).UpdateOne(context.Background(), query, set)
}

func (m *Model) UpdateOneBy(d, q, s interface{}, o *options.UpdateOptions) (*mongo.UpdateResult, error) {
	tableName := getTableName(d)
	m.SetUpdateFields()

	// TODO if s is not bson.M
	s.(bson.M)["$set"].(bson.M)["updatedAt"] = time.Now().UnixMilli()
	return db.Collection(tableName).UpdateOne(context.Background(), q, s, o)
}

func (m *Model) UpdateManyBy(d, q, s interface{}, o *options.UpdateOptions) (*mongo.UpdateResult, error) {
	tableName := getTableName(d)
	return db.Collection(tableName).UpdateMany(context.Background(), q, s, o)
}

func (m *Model) SetID(id string) {
	m.ID = id
}

func (m *Model) GetID() string {
	return m.ID
}

func (m *Model) GetObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(m.ID)
}

func (m *Model) UnsetID() {
	m.ID = ""
}

func (m *Model) SetUpdateFields() {
	m.UnsetID()
	m.UpdatedAt = time.Now().UnixMilli()
}

func (m *Model) SetCreateFields() {
	m.CreatedAt = time.Now().UnixMilli()
}

func getTableName(d interface{}) string {
	var tableName string
	// data := reflect.Indirect(reflect.ValueOf(d))
	// t := reflect.TypeOf(data)
	source := reflect.ValueOf(d).Elem()
	sourceType := source.Type()
	// default
	tableName = sourceType.Name()
	tableName = strings.ToLower(string(tableName[0])) + tableName[1:] + "s"
	// check custom table name
	for i := 0; i < source.NumField(); i++ {
		tag := sourceType.Field(i).Tag.Get("morm")
		kvs := strings.Split(tag, "&")
		for _, v := range kvs {
			param := strings.Split(v, "=")
			if len(param) == 2 && param[0] == "colName" {
				tableName = param[1]
			}
		}
	}

	return tableName
}
