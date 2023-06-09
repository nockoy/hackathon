package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid/v2"
	"log"
)

type ReplyID struct {
	ID string `json:"id"`
}

func SendReply(re model.Replies) ([]byte, error) {

	re.ID = ulid.Make().String()

	err := dao.CreateReply(re)
	if err != nil {
		return nil, err
	}

	//idを返す
	var replyID ReplyID
	replyID.ID = re.ID

	bytes, err := json.Marshal(replyID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", re)

	return bytes, nil
}
