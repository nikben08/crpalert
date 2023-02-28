package repositories

import (
	"crpalert/db"
	"crpalert/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var DB *redis.Client = db.Init()

func GetAllIndicators(chatId int64) []map[string]string {
	var keys []string
	var err error
	var cursor uint64
	keys, cursor, err = DB.Scan(cursor, "chatId="+strconv.FormatInt(int64(chatId), 10)+" indicatorId=*", 0).Result()
	if err != nil {
		panic(err)
	}

	response := []map[string]string{}

	for i := 0; i < len(keys); i++ {
		value, err := DB.Get(keys[i]).Result()
		if err != nil {
			fmt.Println(err)
		}
		params := strings.Split(value, " ")
		_, indicatorId := utils.GetKeyValue(strings.Split(keys[i], " ")[1])
		_, coin := utils.GetKeyValue(params[0])
		_, indicator := utils.GetKeyValue(params[1])
		_, interval := utils.GetKeyValue(params[2])
		row := map[string]string{"indicatorId": indicatorId, "coin": coin, "indicator": indicator, "interval": interval}
		response = append(response, row)
	}
	return response
}

func DeleteAllIndicators(chatId int64) {
	var rows []string
	var err error
	var cursor uint64
	rows, cursor, err = DB.Scan(cursor, "chatId="+strconv.FormatInt(int64(chatId), 10)+" indicatorId=*", 0).Result()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(rows); i++ {
		DB.Del(rows[i])
	}
}

func DeleteIndicator(chatId int64, indicatorId string) {
	err := DB.Del("chatId=" + strconv.FormatInt(int64(chatId), 10) + " indicatorId=" + indicatorId).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("chatId=" + strconv.FormatInt(int64(chatId), 10) + " indicatorId=" + indicatorId)
}

func GetSetIndicatorCmd(chatId int64) string {
	var keys []string
	var err error
	var cursor uint64
	keys, cursor, err = DB.Scan(cursor, "cmd=setIndicator chatId="+strconv.FormatInt(int64(chatId), 10), 0).Result()
	if err != nil {
		fmt.Print(err)
	}
	if len(keys) == 0 {
		return "not found"
	}

	value, err := DB.Get(keys[0]).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func CreateSetIndicatorCmd(chatId int64, coin string) {
	err := DB.Set("cmd=setIndicator chatId="+strconv.FormatInt(int64(chatId), 10), "coin="+coin, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateSetIndicatorCmd(chatId int64, value string) {
	cmd := GetSetIndicatorCmd(chatId)
	DB.Del(cmd)
	DB.Set("cmd=setIndicator chatId="+strconv.FormatInt(int64(chatId), 10), value, 0)
}

func AddIndicatorForSetIndicatorCmd(chatId int64, indicator string) {
	cmd := GetSetIndicatorCmd(chatId)
	cmd += " indicator=" + indicator
	UpdateSetIndicatorCmd(chatId, cmd)
}

func AddFrameForSetIndicatorCmd(chatId int64, frame string) {
	value := GetSetIndicatorCmd(chatId)
	value += " frame=" + frame
	SetIndicator(chatId, value)
}

func DeleteSetIndicatorCmd(chatId int64) {
	var keys []string
	var err error
	var cursor uint64
	keys, cursor, err = DB.Scan(cursor, "cmd=setIndicator chatId="+strconv.FormatInt(int64(chatId), 10)+"*", 0).Result()
	if err != nil {
		fmt.Print(err)
	}
	if len(keys) == 0 {
		fmt.Println("not found")
	}
	DB.Del(keys[0])
}

func SetIndicator(chatId int64, value string) {
	err := DB.Set("chatId="+strconv.FormatInt(int64(chatId), 10)+" indicatorId="+(uuid.New()).String(), value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	DeleteSetIndicatorCmd(chatId)
}

func CreateChat(chatId string) {
	err := DB.Set("chatId="+chatId, "30.01.2023", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func WhetherChatExist(chatId string) bool {
	val, err := DB.Get("chatId=" + chatId).Result()
	if err != nil {
		fmt.Println(err)
	}
	if len(val) == 0 {
		return false
	}
	return true
}
