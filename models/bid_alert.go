package models

type BidAlert struct {
	Id         string `bson:"_id,omitempty"`
	ChatId     int64  `bson:"chat_id,omitempty"`
	Coin       string `bson:"coin,omitempty;default:"`
	Indicator  string `bson:"indicator,omitempty;default:"`
	AlertValue int64  `bson:"alert_value,omitempty;default:"`
	Frame      string `bson:"frame,omitempty;default:"`
	Condition  string `bson:"condition,omitempty;default:"`
}
