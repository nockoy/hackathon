package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"log"
	"time"
)

type MessageID struct {
	ID string `json:"id"`
}

func GetMessage(RoomID string) ([]byte, error) {

	users, err := dao.GetMessage(RoomID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func SendMessage(m model.Messages) ([]byte, error) {

	m.ID = ulid.Make().String()

	//日本の現在時刻を記録したいが日本の時刻にならなかった
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().In(jst)
	m.CreatedAt = nowJST
	m.UpdatedAt = nowJST

	err := dao.CreateMSG(m)
	if err != nil {
		return nil, err
	}

	//idを返す
	var messageID MessageID
	messageID.ID = m.ID

	bytes, err := json.Marshal(messageID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", m)

	return bytes, nil
}
