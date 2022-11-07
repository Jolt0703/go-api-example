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
	PingMySQL(db)

	// クエリ実行(複数レコード)
	SelectArticles(db)

	// クエリ実行(単一レコード)
	SelectArticle(db, 3)

	// データの挿入
	InsertArticle(db)

	// クエリ実行(複数レコード)
	SelectArticles(db)
}

// Pingでの疎通確認
func PingMySQL(db *sql.DB) {
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected to DB")
	}
}

// Select文複数レコード(db.Query)
func SelectArticles(db *sql.DB) {
	const sqlStr = `select * from articles;`
	rows, err := db.Query(sqlStr)
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

// Select文単一レコード(db.QueryRow)
func SelectArticle(db *sql.DB, articleID int) {
	// クエリ実行
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
	err := row.Scan(&article.ArticleID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

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

//　Insert文(db.Exec)
func InsertArticle(db *sql.DB) {
	// 挿入する記事データ
	article := models.Article{
		Title:    "3rd Post",
		Contents: "This is the 3rd blog.",
		UserName: "Chomsky",
	}
	const sqlStr = `
		insert into articles(title, contents, username, nice, created_at)
		values (?, ?, ?, 0, now());
		`
	// データを挿入
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// result.LastInsertIdの実行結果から、記事ID が何番になったのかを調べる
	fmt.Println(result.LastInsertId())
	// result.RowsAffectedの実行結果から、クエリの影響範囲の広さを調べる
	fmt.Println(result.RowsAffected())
}
