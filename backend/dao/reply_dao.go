package dao

import (
	"db/model"
	"log"
)

func GetReplies(ReplyToID string) ([]model.RepliesAndUserInfo, error) {

	rows, err := db.Query("SELECT r.id, r.reply_to_id, r.user_id, r.text, r.created_at, r.updated_at, u.name, u.email, u.icon FROM replies r JOIN users u ON r.user_id = u.id WHERE r.reply_to_id = ?", ReplyToID)

	replies := make([]model.RepliesAndUserInfo, 0)

	for rows.Next() {
		var r model.RepliesAndUserInfo
		if ServerErr := rows.Scan(&r.ID, &r.ReplyToID, &r.UserID, &r.Text, &r.CreatedAt, &r.UpdatedAt, &r.Name, &r.Email, &r.Icon); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		replies = append(replies, r)
	}

	return replies, err
}

func CreateReply(re model.Replies) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO replies(id, reply_to_id, user_id, text) values (?,?,?,?)", re.ID, re.ReplyToID, re.UserID, re.Text)
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

func EditReply(re model.Replies) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//UPDATEする
	_, err = tx.Exec("UPDATE replies SET text = ? WHERE id = ?", re.Text, re.ID)
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

func DeleteReply(re model.Replies) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//DELETEする
	_, err = tx.Exec("DELETE FROM replies WHERE id = ?", re.ID)
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
