package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"todo-app/config"
)

// HTTPレスポンスを介して、HTMLを生成(テンプレートファイルを組み合わせてHTMLを作成)
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// ファイルを格納するスライスの作成
	var files []string
	// for文にてテンプレートファイルを格納していく
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	// ここはよくわからない
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// var validPath = regexp.MustCompile("^/delete/([0-9]+)$")
var validPath = regexp.MustCompile(`^/(delete|edit|update)/([0-9a-f\-]+)$`)

func parseURL(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}

		// q[2]の部分が([0-9a-f\-]+)にあてはまる
		fn(w, r, q[2])
	}
}

// サーバー立ち上げ
func StartMainServer() error {
	// URLの登録
	http.HandleFunc("/", index)
	http.HandleFunc("/todonew", todoNew)
	http.HandleFunc("/save", todoSave)
	http.HandleFunc("/edit/", parseURL(todoEdit))
	http.HandleFunc("/delete/", parseURL(todoDelete))
	http.HandleFunc("/update/", parseURL(todoUpdate))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
