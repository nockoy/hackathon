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

type Userid struct {
	Id string `json:"id"`
}

func RegisterUser(u model.User) ([]byte, error) {

	id := ulid.Make().String()
	u.Id = id
	err := dao.Create(u)
	if err != nil {
		return nil, err
	}

	//idを返す
	var u2 Userid
	u2.Id = id

	bytes, err := json.Marshal(u2)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", u)

	return bytes, nil
}
