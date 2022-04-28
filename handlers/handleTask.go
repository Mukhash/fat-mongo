package handlers

import (
	"go-mongo/db"
	"go-mongo/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, obj interface{}) {
	templates, err := template.ParseFiles("index.html", "task.html")
	if err != nil {
		log.Fatal(err)
	}
	err = templates.ExecuteTemplate(w, tmpl+".html", obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tasks, err := db.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "index", tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	task := models.Task{Title: title, Body: body}

	err := db.InsertTask(&task)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Path[9:])
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+err.Error(), http.StatusInternalServerError)
		return
	}
	if task == nil {
		task = &models.Task{}
	}

	renderTemplate(w, "task", *task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Path[8:])
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}

	action := r.PostFormValue("button")
	switch action {
	case "delete":
		err = db.DeleteTask(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+err.Error(), http.StatusInternalServerError)
			return
		}
	case "update":
		body := r.PostFormValue("body")
		title := r.PostFormValue("title")

		task := models.Task{ID: int64(id), Title: title, Body: body}
		err = db.UpdateTask(&task)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
