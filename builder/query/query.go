package query

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type query struct {
	queries []bson.M
}

func NewMongoQueryBuilder() *query {
	return &query{}
}

func (qb *query) String(key, value string) *query {
	if key != "" && value != "" {
		qb.queries = append(qb.queries, bson.M{key: value})
	}

	return qb
}

func (qb *query) Int(key string, value int) *query {
	if key != "" && value != 0 {
		qb.queries = append(qb.queries, bson.M{key: value})
	}

	return qb
}

func (qb *query) Search(key, value string) *query {
	if key != "" && value != "" {
		qb.queries = append(qb.queries, bson.M{key: primitive.Regex{Pattern: fmt.Sprintf(".*%s.*", value), Options: "i"}})
	}

	return qb
}

func (qb *query) Range(key string, from, to int) *query {
	if from != 0 && to != 0 {
		qb.queries = append(qb.queries, bson.M{key: bson.M{"$gte": from, "$lte": to}})
	}

	return qb
}

func (qb *query) Done() bson.M {
	if len(qb.queries) > 0 {
		return bson.M{"$and": qb.queries}
	}

	return bson.M{}
}

func (qb *query) DoneD() bson.D {
	if len(qb.queries) > 0 {
		return bson.D{{"$match", bson.D{{"$and", qb.queries}}}}
	}

	return bson.D{{}}
}
