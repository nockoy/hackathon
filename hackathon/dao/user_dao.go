package dao

import (
	"database/sql"
	"db/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func Create(u model.Users) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO users(id, name, created_at, updated_at) values (?,?,?,?)", u.ID, u.Name, u.CreatedAt, u.UpdatedAt)
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

func SearchUserByName(name string) ([]model.Users, error) {

	rows, err := db.Query("SELECT id, name FROM users WHERE name = ?", name)

	users := make([]model.Users, 0)

	for rows.Next() {
		var u model.Users
		if ServerErr := rows.Scan(&u.ID, &u.Name); ServerErr != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if ServerErr := rows.Close(); ServerErr != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, ServerErr
		}
		users = append(users, u)
	}

	return users, err
}

/*
func Delete(w http.ResponseWriter) {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return
	}

	//DELETEする
	_, err1 := tx.Exec("DELETE FROM user WHERE name = ?", "")
	if err1 != nil {
		log.Printf("fail: tx.Exec, %v\n", err1)
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	//トランザクション終了
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		return
	}
}
*/
