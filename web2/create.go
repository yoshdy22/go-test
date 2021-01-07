package main

import (
    // ビルド時のみ使用する
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

// DB Path(相対パスでも大丈夫かと思うが、筆者の場合、絶対パスでないと実行できなかった)
const dbPath = "db.sql"

// コネクションプールを作成
var DbConnection *sql.DB


func main() {
    // Open(driver,  sql 名(任意の名前))
    DbConnection, _ := sql.Open("sqlite3", dbPath)

    // Connection をクローズする。(defer で閉じるのが Golang の作法)
    defer DbConnection.Close()

    // blog テーブルの作成
    cmd := `CREATE TABLE IF NOT EXISTS patent(
        fy INT,
	status STRING,
	country STRING,
	title STRING,
	inventor STRING,
	no STRING,
	famiry_no STRING,
	app_no STRING,
	filed_date STRING,
	pub_no STRING,
	patent_no STRING,
	patent_date STRING)`

    // cmd を実行
    // _ -> 受け取った結果に対して何もしないので、_ にする
    _, err := DbConnection.Exec(cmd)

    // エラーハンドリング(Go だと大体このパターン)
    if err != nil {
        // Fatalln は便利
        // エラーが発生した場合、以降の処理を実行しない
        log.Fatalln(err)
    }
}
