package dao

import (
	"db/model"
	"log"
)

func GetJoinChannelsByUserID(userID string) ([]model.Channels, error) {

	rows, err := db.Query("SELECT c.id, c.name, c.description, c.created_at, c.updated_at FROM members m JOIN channels c ON m.channel_id = c.id WHERE m.user_id = ?", userID)

	channels := make([]model.Channels, 0)

	for rows.Next() {
		var c model.Channels
		if ServerErr := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		channels = append(channels, c)
	}

	return channels, err
}

func GetNotJoinChannelsByUserID(userID string) ([]model.Channels, error) {

	rows, err := db.Query("SELECT c.id, c.name, c.description, c.created_at, c.updated_at FROM channels c WHERE NOT EXISTS(SELECT 1 FROM members m WHERE m.channel_id = c.id AND m.user_id = ?)", userID)

	channels := make([]model.Channels, 0)

	for rows.Next() {
		var c model.Channels
		if ServerErr := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		channels = append(channels, c)
	}

	return channels, err
}

func CreateChannel(channel model.Channels) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO channels(id, name, description) values (?,?,?)", channel.ID, channel.Name, channel.Description)
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
