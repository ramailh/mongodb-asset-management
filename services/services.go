package services

import "go.mongodb.org/mongo-driver/bson"

type (
	find interface {
		GetFilter() bson.M
	}

	findByID interface {
		GetFilterByID() bson.M
	}

	insert interface {
		GetInsertDoc() interface{}
	}

	update interface {
		findByID
		GetUpdateDoc() bson.M
	}
)
