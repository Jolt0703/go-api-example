package handlers

import (
	"encoding/json"
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
	article1 := models.Article1
	jsonData, err := json.Marshal(article1)

	if err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
		return
	}

	log.Printf("Posting Article...\n")
	w.Write(jsonData)
}

// /article/list?page=[x]
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	page := 1

	if pageList, ok := queryMap["page"]; ok && len(pageList) > 0 {
		var err error
		page, err = strconv.Atoi(pageList[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	}

	article1, article2 := models.Article1, models.Article2
	articleList := []models.Article{article1, article2}
	jsonData, err := json.Marshal(articleList)

	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("Article List (page %d)\n", page)
	w.Write(jsonData)
}

// /article/[id]
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	article1 := models.Article1
	jsonData, err := json.Marshal(article1)

	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("Article No.%d\n", articleID)
	w.Write(jsonData)
}

// /article/nice
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
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
	comment1 := models.Comment1
	jsonData, err := json.Marshal(comment1)

	if err != nil {
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
	}

	log.Printf("Posting Comment...\n")
	w.Write(jsonData)
}
