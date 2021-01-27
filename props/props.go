package props

import "os"

var (
	MongoHost, MongoUsername, MongoPassword, MongoDB string

	Port string
)

func Setup() {
	MongoHost = os.Getenv("MONGO_HOST")
	MongoUsername = os.Getenv("MONGO_USERNAME")
	MongoPassword = os.Getenv("MONGO_PASSWORD")
	MongoDB = os.Getenv("MONGO_DB")

	Port = os.Getenv("PORT")
}
