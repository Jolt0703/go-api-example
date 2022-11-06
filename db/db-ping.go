package main

import (
	"database/sql"
	"fmt"

	"github.com/Jolt0703/go-api-example/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := ""
	dbPassword := ""
	dbDatabase := ""
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// MySQLへ接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Pingでの疎通確認
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected to DB")
	}

	// クエリ実行
	articleID := 1
	const sqlStr = `select * from articles where article_id = ?;`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// クエリ結果を構造体に埋め込む
	var article models.Article
	var createdTime sql.NullTime
	// Scanメソッドで取得レコードのデータをarticleの各フィールドに入れる
	err = row.Scan(&article.ArticleID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Nullかもしれいない値をsql.NullTime型の変数で検査
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)
}
