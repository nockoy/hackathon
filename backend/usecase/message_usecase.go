package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"log"
)

type MessageID struct {
	ID string `json:"id"`
}

func GetMessages(RoomID string) ([]byte, error) {

	messages, err := dao.GetMessages(RoomID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(messages)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func SendMessage(m model.Messages) ([]byte, error) {

	m.MessageID = ulid.Make().String()

	err := dao.CreateMSG(m)
	if err != nil {
		return nil, err
	}

	//idを返す
	var messageID MessageID
	messageID.ID = m.MessageID

	bytes, err := json.Marshal(messageID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", m)

	return bytes, nil
}

func EditMessage(m model.Messages) ([]byte, error) {

	err := dao.EditMSG(m)
	if err != nil {
		return nil, err
	}

	//idを返す
	var messageID MessageID
	messageID.ID = m.MessageID

	bytes, err := json.Marshal(messageID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Edited: ", m)

	return bytes, nil
}

func DeleteMSG(m model.Messages) error {

	err := dao.DeleteMSG(m)
	if err != nil {
		return err
	}

	//Registerが成功したら知らせる
	fmt.Println("Message Deleted: ", m.MessageID)

	return nil
}
