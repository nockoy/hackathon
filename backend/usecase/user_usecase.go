package usecase

import (
	"db/dao"
	"db/model"
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid/v2"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func SearchUserByEmail(email string) ([]byte, error) {

	users, err := dao.SearchUserByEmail(email)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func SearchUserByUserID(userID string) ([]byte, error) {

	users, err := dao.SearchUserByUserID(userID)
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

	u.ID = ulid.Make().String()
	u.Icon = ""

	//日本の現在時刻を記録したいが日本の時刻にならなかった
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().In(jst)
	u.CreatedAt = nowJST
	u.UpdatedAt = nowJST

	err := dao.CreateUser(u)
	if err != nil {
		return nil, err
	}

	//id,name,iconを返す
	var user User
	user.ID = u.ID
	user.Name = u.Name
	user.Icon = u.Icon

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", u)

	return bytes, nil
}
