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

type WorkspaceID struct {
	ID string `json:"id"`
}

func RegisterWorkspace(workspace model.Workspaces) ([]byte, error) {

	workspace.ID = ulid.Make().String()

	//日本の現在時刻を記録したいが日本の時刻にならなかった
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().In(jst)

	workspace.CreatedAt = nowJST
	workspace.UpdatedAt = nowJST

	err := dao.CreateWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	//idを返す
	var workspaceID WorkspaceID
	workspaceID.ID = workspace.ID

	bytes, err := json.Marshal(workspaceID)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		return nil, err
	}

	//Registerが成功したら知らせる
	fmt.Println("Register: ", workspace)

	return bytes, nil
}
