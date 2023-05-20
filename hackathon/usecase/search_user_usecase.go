package usecase

import (
	"db/dao"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func SearchUser(name string) ([]byte, error) {

	users, err := dao.SearchByName(name)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
