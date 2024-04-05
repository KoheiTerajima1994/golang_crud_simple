// packageとはどのパッケージに属するか？実行可能なプログラムはpackage mainで始まるもの
package main

import (
	"database/sql"
	"log"
	"time"

	// コード内では使用しないので、_を最初につける
	_ "github.com/lib/pq"
)

// struct
type Todo struct {
	Content string
	Done    bool
	Until   time.Time
}

// データベースの変数宣言
var Db *sql.DB

// エラーの変数宣言
var err error

func main() {
	// データベースに接続
	Db, err = sql.Open("postgres", "user=koheiterajima dbname=test_crud password=password sslmode=disable")
	// エラーハンドリング
	if err != nil {
		// Paniclnは致命的なエラーが発生した場合に使う
		log.Panicln(err)
	}
	// データベース接続を閉じる命令、遅延実行処理
	defer Db.Close()

	/*
		// Create
		// データベースに入力
		cmd := "INSERT INTO todo (content, done, until) VALUES ($1, $2, $3)"
		untilTime := time.Date(2024, 4, 5, 10, 0, 0, 0, time.UTC)
		// Execにて、コマンドやプログラム実行を行う
		_, err := Db.Exec(cmd, "宿題", false, untilTime)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		// Read
		// content列が"宿題"と一致する一番最初のデータの取得
		cmd := "SELECT * FROM todo where content = $1"
		// QueryRow 1レコード取得
		row := Db.QueryRow(cmd, "宿題")
		// todoの変数宣言
		var t Todo
		// row.Scanにて、取得した行の各列の値をTodo構造体の対応するフィールドにスキャンする
		err = row.Scan(&t.Content, &t.Done, &t.Until)
		if err != nil {
			// データがなかったら
			if err == sql.ErrNoRows {
				log.Println("No Row")
				// それ以外のエラー
			} else {
				log.Println(err)
			}
		}
		// 取得したデータを出力
		fmt.Println(t.Content, t.Done, t.Until)

		// todoテーブルの全てを取得(ここで変数の更新をしている)
		cmd = "SELECT * FROM todo"
		// Queryは条件に合うものを全て取得
		rows, _ := Db.Query(cmd)
		defer rows.Close()
		// structを作成(空の構造体のスライスを宣言)
		var tt []Todo
		// 取得したデータをループでスライスに追加
		// rows.Next()はデータベースからの結果セット内で次の行が存在するかどうかを判定するメソッド
		for rows.Next() {
			var t Todo
			// scanデータ追加
			err := rows.Scan(&t.Content, &t.Done, &t.Until)
			// 一つずつエラーハンドリングver
			if err != nil {
				log.Println(err)
			}
			// 空のスライスにforループを使って、要素を追加していく
			tt = append(tt, t)
		}
		// まとめてエラーハンドリングver
		err = rows.Err()
		if err != nil {
			log.Fatalln(err)
		}
		// 表示
		for _, t := range tt {
			fmt.Println(t.Content, t.Done, t.Until)
		}
	*/

	/*
		// Update
		// どの内容をUPDATEさせるか、それは何の条件のものか
		cmd := "UPDATE todo SET content = $1, done = $2, until = $3 WHERE content = $4"
		untilTime := time.Date(2024, 4, 10, 10, 0, 0, 0, time.UTC)
		_, err := Db.Exec(cmd, "試験", true, untilTime, "宿題")
		// エラーハンドリング
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// Delete
	cmd := "DELETE FROM todo WHERE content = $1"
	_, err := Db.Exec(cmd, "試験")
	if err != nil {
		log.Fatalln(err)
	}

}
