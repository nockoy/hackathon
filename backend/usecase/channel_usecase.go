package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid/v2"
	"log"
)

type ChannelID struct {
	ID string `json:"id"`
}

func GetChannelByChannelID(channelID string) ([]byte, error) {

	channels, err := dao.GetChannelByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(channels)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetUserChannelsByUserID(userID string) ([]byte, error) {

	channels, err := dao.GetUserChannelsByUserID(userID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(channels)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetOtherChannelsByUserID(userID string) ([]byte, error) {

	channels, err := dao.GetOtherChannelsByUserID(userID)
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
