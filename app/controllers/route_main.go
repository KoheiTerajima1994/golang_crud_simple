package controllers

import (
	"log"
	"net/http"
	"todo-app/app/models"
)

// ハンドラの作成
func index(w http.ResponseWriter, r *http.Request) {
	// modelsパッケージからReadTodoを呼び出し、返り値であるtodosとerrを変数としている
	todos, err := models.ReadTodo()
	if err != nil {
		log.Fatalln(err)
	}
	// dataにて、テンプレートファイルにデータベースのtodoを渡す
	generateHTML(w, todos, "layout", "index")
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "todo_new")
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	// textareaに入力した内容を持ってきて、変数にする
	content := r.PostFormValue("content")
	if err := models.CreateTodo(content); err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 302)
}

func todoEdit(w http.ResponseWriter, r *http.Request, id string) {
	t, err := models.GetTodo(id)
	if err != nil {
		log.Println(err)
	}
	generateHTML(w, t, "layout", "todo_edit")
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id string) {
	updateContent := r.PostFormValue("content")
	t := &models.Todo{ID: id, Content: updateContent}
	if err := t.UpdateTodo(); err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 302)
}

func todoDelete(w http.ResponseWriter, r *http.Request, id string) {
	t, err := models.GetTodo(id)
	if err != nil {
		log.Println(err)
	}
	if err := t.DeleteTodo(); err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 302)
}
