// データベースへの接続

package models

import (
	"database/sql"
	"log"

	// コード内では使用しないので、_を最初につける(postgresqlの接続用のドライバー)
	_ "github.com/lib/pq"
)

// データベースの変数宣言
var Db *sql.DB

// エラーの変数宣言
var err error

func init() {
	// データベースに接続(データベースの作成はpostgresqlのコマンドからすでに行った)
	Db, err = sql.Open("postgres", "user=koheiterajima dbname=crud_database password=password sslmode=disable")
	// エラーハンドリング
	if err != nil {
		// Paniclnは致命的なエラーが発生した場合に使う
		log.Panicln(err)
	}
}
