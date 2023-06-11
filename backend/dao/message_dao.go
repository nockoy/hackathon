package dao

import (
	"db/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetMessages(ChannelID string) ([]model.MessagesAndUserInfo, error) {

	rows, err := db.Query("SELECT m.id, m.channel_id, m.user_id, m.text, m.created_at, m.updated_at, u.name, u.email, u.icon FROM messages m JOIN users u ON m.user_id = u.id WHERE m.channel_id = ?", ChannelID)

	messages := make([]model.MessagesAndUserInfo, 0)

	for rows.Next() {
		var m model.MessagesAndUserInfo
		if ServerErr := rows.Scan(&m.ID, &m.ChannelID, &m.UserID, &m.Text, &m.CreatedAt, &m.UpdatedAt, &m.Name, &m.Email, &m.Icon); ServerErr != nil {
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

func GetMSGByMSGID(MessageID string) ([]model.MessagesAndUserInfo, error) {

	rows, err := db.Query("SELECT m.id, m.channel_id, m.user_id, m.text, m.created_at, m.updated_at, u.name, u.email, u.icon FROM messages m JOIN users u ON m.user_id = u.id WHERE m.id = ?", MessageID)

	messages := make([]model.MessagesAndUserInfo, 0)

	for rows.Next() {
		var m model.MessagesAndUserInfo
		if ServerErr := rows.Scan(&m.ID, &m.ChannelID, &m.UserID, &m.Text, &m.CreatedAt, &m.UpdatedAt, &m.Name, &m.Email, &m.Icon); ServerErr != nil {
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

func CreateMSG(m model.Messages) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO messages(id, channel_id, user_id, text) values (?,?,?,?)", m.MessageID, m.ChannelID, m.UserID, m.Text)
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

func EditMSG(m model.Messages) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//UPDATEする
	_, err = tx.Exec("UPDATE messages SET text = ? WHERE id = ?", m.Text, m.MessageID)
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

func DeleteMSG(m model.Messages) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//DELETEする
	_, err = tx.Exec("DELETE FROM messages WHERE id = ?", m.MessageID)
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
