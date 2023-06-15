package controller

import (
	"app/model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//указываем пути к файлам с шаблонами
	main := filepath.Join("public", "html", "usersDynamicPage.html")
	common := filepath.Join("public", "html", "common.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// получааем данные из строки запроса
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	// проверям на валидность
	if name == "" || surname == "" {
		http.Error(rw, "Имя и фамилия не могут быть пустыми", 400)
		return
	}

	// создаем новый user
	user := model.NewUser(name, surname)

	// добавляем user в bd
	err := user.Add()
	if err != nil {
		http.Error(rw, fmt.Sprintf("Add err: %s", err.Error()), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// получаем id из строки запроса
	userId, err := strconv.Atoi(p.ByName("userId"))
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	// получаем user по id
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Пользователь не был найден: %s", err.Error()), 400)
		return
	}

	// удаляем user из bd
	err = user.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь успешно удален!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//	получаем id из строки запроса
	userId, err := strconv.Atoi(p.ByName("userId"))
	if err != nil {
		http.Error(rw, fmt.Sprintf("httprouter error: %s", err.Error()), 400)
		return
	}
	userName := r.FormValue("name")
	userSurname := r.FormValue("surname")

	// получаем user по id
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Пользователь не был найден: %s", err.Error()), 400)
		return
	}

	// обновляем данные user
	user.Name = userName
	user.Surname = userSurname
	err = user.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь успешно обновлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
