package dao

import (
	"db/model"
	"log"
)

func SearchRoomByID(roomID string) ([]model.Rooms, error) {

	rows, err := db.Query("SELECT id, workspace_id, name FROM rooms WHERE id = ?", roomID)

	rooms := make([]model.Rooms, 0)

	for rows.Next() {
		var u model.Rooms
		if ServerErr := rows.Scan(&u.ID, &u.WorkspaceID, &u.Name); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		rooms = append(rooms, u)
	}

	return rooms, err
}

func CreateRoom(room model.Rooms) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO rooms(id, workspace_id, name) values (?,?,?)", room.ID, room.WorkspaceID, room.Name)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		tx.Rollback()
		return err
	}

	//トランザクション終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return err
	}

	return nil
}
