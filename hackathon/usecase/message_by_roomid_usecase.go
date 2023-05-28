package usecase

import (
	"db/dao"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func GetMessage(RoomID string) ([]byte, error) {

	users, err := dao.GetMessage(RoomID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
