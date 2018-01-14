package main

import "net/http"

func HandleFuncAdd() {
	http.HandleFunc("/projeler.html", projelerHandler)
	http.HandleFunc("/proje.html", projeHandler)
	http.HandleFunc("/projekaydet", projeKaydetHandler)
}

func projelerHandler(response http.ResponseWriter, request *http.Request) {
	fData  := ProjeSelectAll(db)
	context := Context{Data: fData}
	render(response, request, "projeler", context)
}

func projeHandler(response http.ResponseWriter, request *http.Request) {
	context := Context{Data: ProjeSelect(db, getFormId(request))}
	render(response, request, "proje", context)
}


func projeKaydetHandler(response http.ResponseWriter, request *http.Request) {

	id := getFormId(request)
	item := ProjeSelect(db, id)
	item.ProjeAdi = request.FormValue("ProjeAdi")
	item.ProjeYolu = request.FormValue("ProjeYolu")
	context := Context{}

	if len(item.ProjeAdi) > 0 {
		if id > 0 {
			ProjeUpdate(db, item)
		} else {
			item.Id = ProjeInsert(db, item)
		}
		context.Message = "Kayıt yapıldı"
	} else {
		context.Message = "Lütfen Zorunlu alanları giriniz"
	}

	context.Data = item
	render(response, request, "proje", context)

}

