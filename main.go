// packageとはどのパッケージに属するか？実行可能なプログラムはpackage mainで始まるもの
package main

import (
	"log"
	"todo-app/app/controllers"
)

// struct
type Todo struct {
	Content string
	Done    bool
}

func main() {
	log.Println("Starting the main server...")
	controllers.StartMainServer()
}
