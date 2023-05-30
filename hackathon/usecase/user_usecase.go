package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid/v2"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserID struct {
	ID string `json:"id"`
}

func SearchUser(name string) ([]byte, error) {

	users, err := dao.SearchUserByName(name)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func RegisterUser(u model.Users) ([]byte, error) {

	id := ulid.Make().String()
	u.ID = id
	err := dao.CreateUser(u)
	if err != nil {
		return nil, err
	}

	//idを返す
	var userID UserID
	userID.ID = id

	bytes, err := json.Marshal(userID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", u)

	return bytes, nil
}
