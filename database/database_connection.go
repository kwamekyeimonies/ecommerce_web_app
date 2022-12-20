package database

import "go.mongodb.org/mongo-driver/mongo"

// "github.com/kwamekyeimonies/ecommerce_web_app/models"

func Database_Connection() *mongo.Client {
	client, err := mongo.NewClient()
}

func UserData(client *mongo.Client, collection string) *mongo.Collection {

}

func ProductsData(client *mongo.Client, collection string) *mongo.Collection {

}
