package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func RegisterMember(m model.Members) ([]byte, error) {

	err := dao.CreateMember(m)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: channel_id[", m.ChannelID, "], user_id[", m.UserID, "]")

	return bytes, nil
}

func DeleteMember(m model.Members) ([]byte, error) {

	err := dao.DeleteMember(m)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Delete: channel_id[", m.ChannelID, "], user_id[", m.UserID, "]")

	return bytes, nil
}
