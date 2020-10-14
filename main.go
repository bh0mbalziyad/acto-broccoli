package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// i, _ := strconv.Atoi(v)
	start, _ := strconv.Atoi(params["start"])
	count, _ := strconv.Atoi(params["count"])
	var paginatedPosts = append(posts[start:(start + count)])
	json.NewEncoder(w).Encode(paginatedPosts)
}
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(&item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func searchPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	var query = params["q"]
	for _, item := range posts {
		if strings.Contains(item.Title, query) {
			json.NewEncoder(w).Encode(&item)
			return
		}
		if strings.Contains(item.Body, query) {
			json.NewEncoder(w).Encode(&item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})

}

func main() {
	router := mux.NewRouter()
	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "2", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "3", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "4", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "5", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "6", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "7", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "8", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "9", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "10", Title: "My first post", Body: "This is the content of my first post"})
	router.HandleFunc("/posts", getPosts).Methods("GET").Queries("start", "{start:[0-9]+}", "count", "{count:[0-9]+}")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/search", searchPost).Methods("GET").Queries("q", "{q}")
	router.HandleFunc("/posts/{id:[0-9]+}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id:[0-9]+}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
