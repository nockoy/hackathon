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

type RoomID struct {
	ID string `json:"id"`
}

func SearchRoom(roomID string) ([]byte, error) {

	users, err := dao.SearchRoomByID(roomID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func RegisterRoom(room model.Rooms) ([]byte, error) {

	room.ID = ulid.Make().String()

	//日本の現在時刻を記録したいが日本の時刻にならなかった
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().In(jst)

	room.CreatedAt = nowJST
	room.UpdatedAt = nowJST

	err := dao.CreateRoom(room)
	if err != nil {
		return nil, err
	}

	//idを返す
	var roomID RoomID
	roomID.ID = room.ID

	bytes, err := json.Marshal(roomID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", room)

	return bytes, nil
}
