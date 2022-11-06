package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Jolt0703/go-myapi/models"
	"github.com/gorilla/mux"
)

// /hello
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /article
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストボディから記事の内容を取得
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	// Goの構造体をjsonデータに再度変換しレスポンスストリームに書き込み
	article := reqArticle
	log.Printf("Posting Article...\n")
	json.NewEncoder(w).Encode(article)
}

// /article/list?page=[x]
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// クエリパラメータのマップValuesを取得(type Values map[string][]string)
	queryMap := req.URL.Query()
	page := 1

	// クエリパラメータから"page"キーの値を取得
	if pageList, ok := queryMap["page"]; ok && len(pageList) > 0 {
		var err error
		page, err = strconv.Atoi(pageList[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	}

	// Goの構造体をjsonデータに変換しレスポンスストリームに書き込み
	articleList := []models.Article{models.Article1, models.Article2}
	log.Printf("Article List (page %d)\n", page)
	json.NewEncoder(w).Encode(articleList)
}

// /article/[id]
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストに含まれるパスパラメターのマップを取得しidを数値に変換
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	// Goの構造体をjsonデータに変換しレスポンスストリームに書き込み
	article := models.Article1
	log.Printf("Article No.%d\n", articleID)
	json.NewEncoder(w).Encode(article)
}

// /article/nice
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストボディからいいねする記事の内容を取得
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	// Goの構造体をjsonデータに再変換しレスポンスストリームに書き込み
	article := reqArticle
	log.Printf("Posting Nice...\n")
	json.NewEncoder(w).Encode(article)
}

// /comment
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストボディからコメントの内容を取得
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	// Goの構造体をjsonデータに再変換しレスポンスストリームに書き込み
	comment := reqComment
	log.Printf("Posting Comment...\n")
	json.NewEncoder(w).Encode(comment)
}
