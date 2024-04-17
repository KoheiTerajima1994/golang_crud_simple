package models

import (
	"log"

	"github.com/google/uuid"
)

// struct
type Todo struct {
	Content string
	Done    bool
	ID      string
}

// CreateTodo
func CreateTodo(content string) (err error) {
	// UUID生成
	id := uuid.New()

	// データベースに書き込む処理
	cmd := "INSERT INTO todo (content, done, id) VALUES ($1, $2, $3)"
	// Execにてコマンド実行
	_, err = Db.Exec(cmd, content, false, id)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ReadTodo
func ReadTodo() (todos []Todo, err error) {
	cmd := "SELECT * FROM todo"
	// Queryにより、全て取得
	rows, _ := Db.Query(cmd)
	// structを作成
	var tt []Todo
	// 取得したデータをループでスライスに追加
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.Content, &t.Done, &t.ID)
		if err != nil {
			log.Println(err)
		}
		tt = append(tt, t)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	rows.Close()
	return tt, err
}

// GetTodo(todoを削除、編集するために個別にIDを取得する必要がある)
func GetTodo(id string) (todo Todo, err error) {
	cmd := "SELECT * FROM todo WHERE id = $1"
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.Content,
		&todo.Done,
		&todo.ID)

	return todo, err
}

// UpdateTodo
func (t *Todo) UpdateTodo() error {
	cmd := "UPDATE todo SET content = $1 WHERE id = $2"
	_, err := Db.Exec(cmd, t.Content, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// DeleteTodo
func (t *Todo) DeleteTodo() error {
	cmd := "DELETE FROM todo WHERE id = $1"
	_, err := Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
