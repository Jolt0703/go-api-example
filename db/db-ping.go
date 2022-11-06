package main

import (
	"database/sql"
	"fmt"

	"github.com/Jolt0703/go-myapi/models"
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
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// クエリ結果を構造体に埋め込む
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		// Scanメソッドで取得レコードのデータをarticleの各フィールドに入れる
		err := rows.Scan(&article.ArticleID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

		// Nullかもしれいない値をsql.NullTime型の変数で検査
		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			// 読み出し結果を格納した変数articleをスライスに追加
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
