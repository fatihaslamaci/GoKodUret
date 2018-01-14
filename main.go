package main

import (

	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
	"strconv"

	"os"
)

type Context struct {
	Message string

	KayitId  string
	KayitId2 string

	AktifKayitId string
	Ara          string

	ValueList []interface{}
	Data      interface{}
}


func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}


func render(w http.ResponseWriter, r *http.Request, tmpl string, context Context) {
	files := []string{
		"./templates/base.html", "./templates/" + tmpl + ".html",
	}

	if FileExists("./templates/" + tmpl + "Filed.html") {
		files = append(files, "./templates/"+tmpl+"Field.oto.html")
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", context)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func addStaticDir(s string) {
	http.Handle("/"+s+"/", http.StripPrefix("/"+s, http.FileServer(http.Dir("./statics/"+s))))
}

func addStaticDirAll() {
	addStaticDir("css")
}

var db *sql.DB


func getFormId(request *http.Request) int {
	request.ParseForm()
	id := request.FormValue("id")
	var i int
	i, _ = strconv.Atoi(id)
	return i
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "index", Context{})
}

func main() {

	Makeproje()


	const dbpath = "./db/gomaker.sqlite"
	db = InitDB(dbpath)
	defer db.Close()
	CreateTable(db)

	http.HandleFunc("/", indexHandler)
	HandleFuncAdd()

	addStaticDirAll()
	http.ListenAndServe(":8000", nil)

}
