package seeds

import (
	"context"
	"crpalert/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetSeeds(db *mongo.Database) {
	coll := db.Collection("symbols")
	var coin models.Coin
	filter := bson.M{"symbol": bson.M{"$eq": "BTC"}}
	coll.FindOne(context.TODO(), filter).Decode(&coin)
	if len(coin.Symbol) == 0 {
		symbols := GetAllSymbols()
		for i := 0; i < len(symbols); i++ {
			var symbol = models.Coin{Symbol: symbols[i]}
			coll.InsertOne(context.TODO(), symbol)
		}
	}

}
