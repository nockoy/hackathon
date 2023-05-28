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

func SendMessage(m model.Messages) ([]byte, error) {

	m.ID = ulid.Make().String()

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
