package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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
	// リクエストヘッダのContent-Lengthからbodyの長さを取得
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "couldn't get content length\n", http.StatusBadRequest)
		return
	}

	// リクエストボディのjsonデータを[]byteスライスに格納
	reqBodybuffer := make([]byte, length)
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "failed to get request body\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// jsonデータをGoの構造体に変換
	var reqArticle models.Article
	if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	// レスポンスのためにGoの構造体をjsonデータに再度変換
	article := reqArticle
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
		return
	}

	log.Printf("Posting Article...\n")
	w.Write(jsonData)
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

	// レスポンスのためにGoの構造体をjsonデータに変換
	article1, article2 := models.Article1, models.Article2
	articleList := []models.Article{article1, article2}
	jsonData, err := json.Marshal(articleList)

	if err != nil {
		errMsg := fmt.Sprintf("failed to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("Article List (page %d)\n", page)
	w.Write(jsonData)
}

// /article/[id]
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストに含まれるパスパラメターのマップを取得しidを数値に変換
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	// レスポンスのためにGoの構造体をjsonデータに変換
	article1 := models.Article1
	jsonData, err := json.Marshal(article1)

	if err != nil {
		errMsg := fmt.Sprintf("failed to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("Article No.%d\n", articleID)
	w.Write(jsonData)
}

// /article/nice
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// レスポンスのためにGoの構造体をjsonデータに変換
	article1 := models.Article1
	jsonData, err := json.Marshal(article1)

	if err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
	}

	log.Printf("Posting Nice...\n")
	w.Write(jsonData)
}

// /comment
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// レスポンスのためにGoの構造体をjsonデータに変換
	comment1 := models.Comment1
	jsonData, err := json.Marshal(comment1)

	if err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
	}

	log.Printf("Posting Comment...\n")
	w.Write(jsonData)
}
