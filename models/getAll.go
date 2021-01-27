package models

import (
	"time"

	"github.com/ramailh/mongodb-asset-management/builder/query"
	"go.mongodb.org/mongo-driver/bson"
)

type GetAll struct {
	From     string `json:"from" form:"from" example:""`
	To       string `json:"to" form:"to" example:""`
	SortType string `json:"sort_type" form:"sort_type" example:""`
	SortBy   string `json:"sort_by" form:"sort_by" example:""`
	Search   string `json:"search" form:"search" example:""`
}

func (ga GetAll) GetFilter() bson.M {
	from, _ := time.Parse("2006010215", ga.From)
	to, _ := time.Parse("2006010215", ga.To)

	return query.NewMongoQueryBuilder().Search("title", ga.Search).Range("created_at", int(from.UnixNano()), int(to.UnixNano())).Done()
}
