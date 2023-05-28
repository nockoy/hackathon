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

type UserID struct {
	ID string `json:"id"`
}

func RegisterUser(u model.Users) ([]byte, error) {

	u.ID = ulid.Make().String() //作成日時は自動で入る
	err := dao.Create(u)
	if err != nil {
		return nil, err
	}

	//idを返す
	var userID UserID
	userID.ID = u.ID

	bytes, err := json.Marshal(userID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", u)

	return bytes, nil
}
