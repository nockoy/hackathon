package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid/v2"
	"log"
	"time"
)

type ChannelID struct {
	ID string `json:"id"`
}

func GetJoinChannelsByUserID(userID string) ([]byte, error) {

	channels, err := dao.GetJoinChannelsByUserID(userID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(channels)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetNotJoinChannelsByUserID(userID string) ([]byte, error) {

	channels, err := dao.GetNotJoinChannelsByUserID(userID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(channels)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func RegisterChannel(channel model.Channels) ([]byte, error) {

	channel.ID = ulid.Make().String()

	//日本の現在時刻を記録したいが日本の時刻にならなかった
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().In(jst)

	channel.CreatedAt = nowJST
	channel.UpdatedAt = nowJST

	err := dao.CreateChannel(channel)
	if err != nil {
		return nil, err
	}

	//idを返す
	var channelID ChannelID
	channelID.ID = channel.ID

	bytes, err := json.Marshal(channelID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", channel)

	return bytes, nil
}
