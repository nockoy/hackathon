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

func EditIcon(u model.Users) ([]byte, error) {

	err := dao.UpdateIcon(u)
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
	fmt.Println("Icon Updated: ", u)

	return bytes, nil
}

func UpdateUserName(u model.Users) ([]byte, error) {

	err := dao.UpdateUserName(u)
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
	fmt.Println("Username Updated: ", u)

	return bytes, nil
}
