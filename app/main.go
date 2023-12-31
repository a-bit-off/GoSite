package main

import (
	"app/controller"
	"app/server"
	"github.com/julienschmidt/httprouter"

	"log"
	"net/http"
)

func main() {
	//инициализируем подключение к базе данных
	err := server.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	//создаем и запускаем в работу роутер для обслуживания запросов
	r := httprouter.New()
	routes(r)

	//прикрепляемся к хосту и свободному порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	err = http.ListenAndServe("localhost:4444", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
	r.POST("/user/add", controller.AddUser)
	r.DELETE("/user/delete/:userId", controller.DeleteUser)
	r.POST("/user/update/:userId", controller.UpdateUser)
}
