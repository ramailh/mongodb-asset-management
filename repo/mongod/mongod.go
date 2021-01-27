package mongod

import (
	"context"
	"log"

	"github.com/ramailh/mongodb-asset-management/props"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	blogCollection = "micro-blogging-management"
)

type mongod struct {
	client *mongo.Client
}

func NewMongoClient() (*mongod, error) {
	opt := options.Client().ApplyURI("mongodb://" + props.MongoHost + "/?connect=direct")
	if props.MongoUsername != "" {
		opt.SetAuth(options.Credential{Username: props.MongoUsername, Password: props.MongoPassword})
	}

	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}

	if err = client.Connect(context.Background()); err != nil {
		return nil, err
	}

	return &mongod{client: client}, nil
}

func (mgo *mongod) Find(filter bson.M, opt ...*options.FindOptions) ([]map[string]interface{}, error) {
	defer mgo.client.Disconnect(context.Background())
	result, err := mgo.client.Database(props.MongoDB).Collection(blogCollection).Find(context.Background(), filter, opt...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var datas []map[string]interface{}
	if err = result.All(context.Background(), &datas); err != nil {
		log.Println(err)
		return nil, err
	}

	return datas, err
}

func (mgo *mongod) FindOne(filter bson.M) (map[string]interface{}, error) {
	defer mgo.client.Disconnect(context.Background())
	data := make(map[string]interface{})
	err := mgo.client.Database(props.MongoDB).Collection(blogCollection).FindOne(context.Background(), filter).Decode(&data)

	return data, err
}

func (mgo *mongod) Insert(doc interface{}) (interface{}, error) {
	defer mgo.client.Disconnect(context.Background())
	return mgo.client.Database(props.MongoDB).Collection(blogCollection).InsertOne(context.Background(), doc)
}

func (mgo *mongod) Update(updateDoc, filter bson.M) (interface{}, error) {
	defer mgo.client.Disconnect(context.Background())
	return mgo.client.Database(props.MongoDB).Collection(blogCollection).UpdateOne(context.Background(), filter, updateDoc)
}

func (mgo *mongod) Delete(filter bson.M) (interface{}, error) {
	defer mgo.client.Disconnect(context.Background())
	return mgo.client.Database(props.MongoDB).Collection(blogCollection).DeleteOne(context.Background(), filter)
}
