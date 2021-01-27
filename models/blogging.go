package models

import (
	"time"

	"github.com/ramailh/mongodb-asset-management/builder/query"
	"github.com/renstrom/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
)

type MicroBlogging struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title,omitempty"`
	Text      string `json:"text" bson:"text,omitempty"`
	Author    string `json:"author" bson:"author,omitempty"`
	CreatedAt int64  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at,omitempty"`
}

func (mb MicroBlogging) GetFilterByID() bson.M {
	return query.NewMongoQueryBuilder().String("_id", mb.ID).Done()
}

func (mb MicroBlogging) GetUpdateDoc() bson.M {
	mb.UpdatedAt = time.Now().UnixNano()

	return bson.M{"$set": mb}
}

func (mb MicroBlogging) GetInsertDoc() interface{} {
	mb.CreatedAt = time.Now().UnixNano()
	mb.UpdatedAt = time.Now().UnixNano()
	mb.ID = shortuuid.New()

	return mb
}
