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
	fmt.Println("Register: channel[", m.ChannelID, "], user[", m.UserID, "]")

	return bytes, nil
}
