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

func GetReplies(replyToID string) ([]byte, error) {

	replies, err := dao.GetReplies(replyToID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(replies)
	if err != nil {
		return nil, err
	}

	return bytes, nil
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

func EditReply(re model.Replies) ([]byte, error) {

	err := dao.EditReply(re)
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
	fmt.Println("Edited: ", re)

	return bytes, nil
}

func DeleteReply(re model.Replies) error {

	err := dao.DeleteReply(re)
	if err != nil {
		return err
	}

	//Registerが成功したら知らせる
	fmt.Println("Reply Deleted: ", re.ID)

	return nil
}
