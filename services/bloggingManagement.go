package services

import (
	"encoding/json"
	"fmt"

	"github.com/ramailh/mongodb-asset-management/repo/mongod"
)

func Find(param find) (interface{}, error) {
	client, err := mongod.NewMongoClient()
	if err != nil {
		return nil, err
	}

	filter := param.GetFilter()
	filterJson, _ := json.MarshalIndent(filter, "", "  ")
	fmt.Println(string(filterJson))
	return client.Find(filter)
}

func FindByID(param findByID) (interface{}, error) {
	client, err := mongod.NewMongoClient()
	if err != nil {
		return nil, err
	}

	filter := param.GetFilterByID()
	return client.FindOne(filter)
}

func Insert(param insert) (interface{}, error) {
	client, err := mongod.NewMongoClient()
	if err != nil {
		return nil, err
	}

	doc := param.GetInsertDoc()
	return client.Insert(doc)
}

func Update(param update) (interface{}, error) {
	client, err := mongod.NewMongoClient()
	if err != nil {
		return nil, err
	}

	updateDoc := param.GetUpdateDoc()
	filter := param.GetFilterByID()

	return client.Update(updateDoc, filter)
}

func Delete(param findByID) (interface{}, error) {
	client, err := mongod.NewMongoClient()
	if err != nil {
		return nil, err
	}

	filter := param.GetFilterByID()
	return client.Delete(filter)
}
