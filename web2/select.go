package main

import (
    // ビルド時のみ使用する
    "database/sql"
    "log"
    "fmt"

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

    // シングルセレクト
    cmd := "SELECT * FROM patent where no = ?"
    row := DbConnection.QueryRow(cmd, "N990291")
    var pat PATENT
    err := row.Scan(&pat.fy, &pat.status, &pat.country, &pat.title, &pat.inventor, &pat.no, &pat.famiry_no, &pat.app_no, &pat.filed_date, &pat.pub_no, &pat.patent_no, &pat.patent_date)
    if err != nil {
                // シングルセレクトの場合は、エラーハンドリングが異なる
        if err == sql.ErrNoRows {
            log.Println("There is no row!!!")
        } else {
            log.Println(err)
        }
    }
    fmt.Println(pat)
}
