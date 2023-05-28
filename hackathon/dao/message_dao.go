package dao

import (
	"db/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetMessage(RoomID string) ([]model.Messages, error) {

	rows, err := db.Query("SELECT * FROM messages WHERE room_id = ?", RoomID)

	messages := make([]model.Messages, 0)

	for rows.Next() {
		var m model.Messages
		if ServerErr := rows.Scan(&m.ID, &m.ReplyToID, &m.RoomID, &m.From, &m.Text, &m.CreatedAt, &m.UpdatedAt); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		messages = append(messages, m)
	}

	return messages, err
}
