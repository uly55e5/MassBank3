package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func InitMongoDB(config DBConfig, files map[string]string) (MB3Database, error) {
	db, err := NewMongoDB(config)
	if err != nil {
		return nil, err
	}
	if err = db.Connect(); err != nil {
		return nil, err
	}
	if err = db.database.Drop(context.Background()); err != nil {
		return nil, err
	}
	for col, file := range files {

		buf, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		jsonStr := string(buf)
		var m []interface{}
		if err := bson.UnmarshalExtJSON([]byte(jsonStr), false, &m); err != nil {
			return nil, err
		}
		if _, err = db.database.Collection(col).InsertMany(context.Background(), m); err != nil {
			return nil, err
		}
	}
	return db, nil
}
