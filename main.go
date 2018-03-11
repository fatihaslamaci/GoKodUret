package main

import (

	"database/sql"
	"log"
	"net/http"
	"text/template"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Gezgin struct{
	Link string
	Baslik string
}

type Context struct {
	Message string

	KayitId  string
	KayitId2 string

	MasterId int64

	Gezgin []Gezgin

	//AktifKayitId string
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

	if FileExists("./templates/" + tmpl + "Field_oto.html") {
		files = append(files, "./templates/"+tmpl+"Field_oto.html")
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


func indexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "index", Context{})
}

func main() {


	const dbpath = "./db/gomaker.sqlite"
	db = InitDB(dbpath)
	defer db.Close()

	CreateTable(db)
	AlterDb(db)

	ProjeDoldur(db)
	//ProjelerJsonYedekKaydet(db)

	Makeproje(db,1)

	http.HandleFunc("/", indexHandler)

	HandleFuncAdd()

	http.HandleFunc("/projecreate",ProjeCreateHandler)

	addStaticDirAll()
	http.ListenAndServe(":8000", nil)

}
func ProjeCreateHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := FormValueInt64(request,"id")
	item := ProjeSelect(db, id)
	Makeproje(db,id)

	context := Context{}
	if id > 0 {
	context.Message = item.ProjeYolu + " Proje Oluşturuldu"
	} else{
		context.Message = "Kayıt Bulunamadı"
	}

	context.Data = item
	render(response, request, "proje", context)

}



