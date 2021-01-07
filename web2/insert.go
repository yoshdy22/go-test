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

// データ格納用
type PATENT struct {
    fy int
    status string
    country string
    title string
    inventor string
    no string
    famiry_no string
    app_no string
    filed_date string
    pub_no string
    patent_no string
    patent_date string
}

func main() {
    // Open(driver,  sql 名(任意の名前))
    DbConnection, _ := sql.Open("sqlite3", dbPath)

    // Connection をクローズする。(defer で閉じるのが Golang の作法)
    defer DbConnection.Close()

    // データを挿入(? には、値が入る)
    cmd := "INSERT INTO patent (fy, status, country, patent_name, inventor, no) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := DbConnection.Exec(cmd, 00, "登録", "JP", "操作監視用表示装置", "桑谷　資一", "N990291")

    // エラーハンドリング
    if err != nil {
        // Fatalln は便利
        // エラーが発生した場合、以降の処理を実行しない
        log.Fatalln(err)
    }
}
