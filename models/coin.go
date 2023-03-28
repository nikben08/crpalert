package models

type Coin struct {
	Id     string `bson:"_id,omitempty"`
	Symbol string `bson:"symbol,omitempty;default:"`
}
