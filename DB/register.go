package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := "root:123@tcp(mysql:3306)/russian_D"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	log.Println("Подключение к базе данных установлено")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var name string
		query := "SELECT name FROM users WHERE email = ? AND pass = ?"
		err := db.QueryRow(query, email, password).Scan(&name)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Redirect(w, r, "http://static-site:8080/bad_pass", http.StatusSeeOther)
			}
			return
		}

		http.Redirect(w, r, "http://static-site:8080/lkp", http.StatusSeeOther)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен запрос: %s %s, от клиента: %s", r.Method, r.URL, r.RemoteAddr)

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		log.Printf("Поля формы: name=%s, email=%s, password=%s", name, email, password)

		query := "INSERT INTO users (name, email, pass) VALUES (?, ?, ?)"
		_, err := db.Exec(query, name, email, password)
		if err != nil {
			log.Printf("Ошибка добавления пользователя: %v", err)
			http.Error(w, "Ошибка регистрации пользователя", http.StatusInternalServerError)
			return
		}

		log.Println("Регистрация успешна")
		http.Redirect(w, r, "http://static-site:8080/lkp", http.StatusSeeOther)
	}
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	log.Println("Сервер регистрации запущен на порту 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
