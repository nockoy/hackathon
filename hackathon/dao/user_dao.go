package dao

import (
	"database/sql"
	"db/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Create(u model.User) error {
	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	//INSERTする
	_, err = tx.Exec("INSERT INTO user(id, name, age) values (?,?,?)", u.Id, u.Name, u.Age)
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

func SearchByName(name string) ([]model.User, error) {

	rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)

	users := make([]model.User, 0)

	for rows.Next() {
		var u model.User
		if ServerErr := rows.Scan(&u.Id, &u.Name, &u.Age); ServerErr != nil {
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
