package repositories

import (
	"context"
	"crpalert/db"
	"crpalert/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database = db.Init()

type Indicator models.BidAlert

func CreateEmptyIndicator(chatId int64) {
	coll := DB.Collection("bid_alerts")
	indicator := models.BidAlert{ChatId: chatId}
	_, err := coll.InsertOne(context.TODO(), indicator)
	if err != nil {
		panic(err)
	}
}

func ValidateSymbol(symbol string) bool {
	coll := DB.Collection("symbols")
	var coin models.Coin
	filter := bson.M{"symbol": bson.M{"$eq": symbol}}
	coll.FindOne(context.TODO(), filter).Decode(&coin)
	if len(coin.Symbol) != 0 {
		return true
	}
	return false
}

func SpecifyMsgType(chatId int64, msg string) string {
	coll := DB.Collection("bid_alerts")

	var indicator models.BidAlert
	filter := bson.M{"chat_id": bson.M{"$eq": chatId}, "$and": []bson.M{bson.M{"coin": bson.M{"$eq": ""}}}}
	coll.FindOne(context.TODO(), filter).Decode(&indicator)
	if indicator.ChatId != 0 {
		return "symbol"
	}

	filter = bson.M{"chat_id": bson.M{"$eq": chatId}, "$and": []bson.M{bson.M{"indicator": bson.M{"$eq": ""}}}}
	coll.FindOne(context.TODO(), filter).Decode(&indicator)
	if len(indicator.Indicator) != 0 {
		return "indicator"
	}

	filter = bson.M{"chat_id": bson.M{"$eq": chatId}, "$and": []bson.M{bson.M{"alert_value": bson.M{"$eq": ""}}}}
	coll.FindOne(context.TODO(), filter).Decode(&indicator)
	if len(indicator.Indicator) != 0 {
		return "indicator"
	}

	return "unknown"
}

func SetCoin(chatId int64, coin string) {
	coll := DB.Collection("bid_alerts")
	var oldIndicator Indicator
	filter := bson.M{"chat_id": bson.M{"$eq": chatId}, "$and": []bson.M{bson.M{"coin": bson.M{"$eq": ""}}}}
	coll.FindOne(context.TODO(), filter).Decode(&oldIndicator)
	update := bson.D{{"$set", bson.D{{"coin", coin}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func SetIndicator(chatId int64, indicator string) {
	coll := DB.Collection("bid_alerts")
	var oldIndicator Indicator
	filter := bson.M{"chat_id": bson.M{"$eq": chatId}, "$and": []bson.M{bson.M{"indicator": bson.M{"$eq": ""}}}}
	coll.FindOne(context.TODO(), filter).Decode(&oldIndicator)
	update := bson.D{{"$set", bson.D{{"indicator", indicator}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func SetFrame(chatId int64, frame string) {
	coll := DB.Collection("bid_alerts")
	var oldIndicator Indicator
	filter := bson.M{"chat_id": bson.M{"$eq": chatId}, "$and": []bson.M{bson.M{"frame": bson.M{"$eq": ""}}}}
	coll.FindOne(context.TODO(), filter).Decode(&oldIndicator)
	update := bson.D{{"$set", bson.D{{"frame", frame}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func SetCondition(chatId int64, condition string) {
	coll := DB.Collection("bid_alerts")
	var oldIndicator Indicator
	filter := bson.M{"chat_id": bson.M{"$eq": chatId}, "$and": []bson.M{bson.M{"condition": bson.M{"$eq": ""}}}}
	coll.FindOne(context.TODO(), filter).Decode(&oldIndicator)
	update := bson.D{{"$set", bson.D{{"condition", condition}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func GetAllIndicators(chatId int64) []Indicator {
	coll := DB.Collection("bid_alerts")

	cursor, err := coll.Find(context.TODO(), bson.D{{"chat_id", chatId}})
	if err != nil {
		panic(err)
	}

	var indicators []Indicator
	if err := cursor.All(context.TODO(), &indicators); err != nil {
		panic(err)
	}

	return indicators
}

func DeleteAllIndicators(chatId int64) []Indicator {
	coll := DB.Collection("bid_alerts")

	filter := bson.D{{"chat_id", chatId}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var indicators []Indicator
	if err := cursor.All(context.TODO(), &indicators); err != nil {
		panic(err)
	}

	_, err = coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return indicators
}

func DeleteIndicator(IndicatorId string) {
	coll := DB.Collection("bid_alerts")
	idPrimitive, err := primitive.ObjectIDFromHex(IndicatorId)
	filter := bson.M{"_id": idPrimitive}
	_, err = coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}
