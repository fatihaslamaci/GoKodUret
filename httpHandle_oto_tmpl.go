package main

import (
	"net/http"
	"strconv"
	"time"
)

func FormValueInt64(request *http.Request, s string) int64 {
	r, _ := strconv.ParseInt(request.FormValue(s), 10, 64)
	return r
}

func FormValueInt(request *http.Request, s string) int {
	r, _ := strconv.Atoi(request.FormValue(s))
	return r
}

func FormValueDouble(request *http.Request, s string) float64 {
	r, _ := strconv.ParseFloat(request.FormValue(s), 32)
	return r
}

func FormValueBool(request *http.Request, s string) bool {
	r, _ := strconv.ParseBool(request.FormValue(s))
	return r
}

func FormValueDate(request *http.Request, s string) time.Time {
	r, _ := time.Parse("2006-01-02", request.FormValue(s))
	return r
}

func HandleFuncAdd() {

    http.HandleFunc("/projeler.html",ProjelerHandler)
	http.HandleFunc("/proje.html",ProjeHandler)
	http.HandleFunc("/projekaydet",ProjeKaydetHandler)

    http.HandleFunc("/sinifler.html",SiniflerHandler)
	http.HandleFunc("/sinif.html",SinifHandler)
	http.HandleFunc("/sinifkaydet",SinifKaydetHandler)

    http.HandleFunc("/alanler.html",AlanlerHandler)
	http.HandleFunc("/alan.html",AlanHandler)
	http.HandleFunc("/alankaydet",AlanKaydetHandler)

}
//-----------------------------------------------------------------------



func ProjelerHandler(response http.ResponseWriter, request *http.Request) {
	fData  := ProjeSelectAll(db)
	context := Context{Data: fData}
	render(response, request, "projeler", context)
}

func ProjeHandler(response http.ResponseWriter, request *http.Request) {
	context := Context{Data: ProjeSelect(db, getFormId(request))}
	render(response, request, "proje", context)
}

func ProjeKaydetHandler(response http.ResponseWriter, request *http.Request) {
	id := getFormId(request)
	item := ProjeSelect(db, id)

	
	
	item.ProjeAdi =  request.FormValue("projeadi")
	
	
	item.ProjeYolu =  request.FormValue("projeyolu")
	

	context := Context{}

	//if len(item.ProjeAdi) > 0 {
		if id > 0 {
			ProjeUpdate(db, item)
		} else {
			item.Id = ProjeInsert(db, item)
		}
		context.Message = "Kayıt yapıldı"
	//} else {
	//	context.Message = "Lütfen Zorunlu alanları giriniz"
	//}

	context.Data = item
	render(response, request, "proje", context)

}



func SiniflerHandler(response http.ResponseWriter, request *http.Request) {
	fData  := SinifSelectAll(db)
	context := Context{Data: fData}
	render(response, request, "sinifler", context)
}

func SinifHandler(response http.ResponseWriter, request *http.Request) {
	context := Context{Data: SinifSelect(db, getFormId(request))}
	render(response, request, "sinif", context)
}

func SinifKaydetHandler(response http.ResponseWriter, request *http.Request) {
	id := getFormId(request)
	item := SinifSelect(db, id)

	
	
	item.ProjeId =  FormValueInt64(request,"projeid")
	
	
	item.SinifAdi =  request.FormValue("sinifadi")
	

	context := Context{}

	//if len(item.ProjeAdi) > 0 {
		if id > 0 {
			SinifUpdate(db, item)
		} else {
			item.Id = SinifInsert(db, item)
		}
		context.Message = "Kayıt yapıldı"
	//} else {
	//	context.Message = "Lütfen Zorunlu alanları giriniz"
	//}

	context.Data = item
	render(response, request, "sinif", context)

}



func AlanlerHandler(response http.ResponseWriter, request *http.Request) {
	fData  := AlanSelectAll(db)
	context := Context{Data: fData}
	render(response, request, "alanler", context)
}

func AlanHandler(response http.ResponseWriter, request *http.Request) {
	context := Context{Data: AlanSelect(db, getFormId(request))}
	render(response, request, "alan", context)
}

func AlanKaydetHandler(response http.ResponseWriter, request *http.Request) {
	id := getFormId(request)
	item := AlanSelect(db, id)

	
	
	//item.IsId =  request.FormValue("isid")
	
	
	item.SinifId =  FormValueInt64(request,"sinifid")
	
	
	item.AlanAdi =  request.FormValue("alanadi")
	
	
	item.AlanVeriTuru =  request.FormValue("alanverituru")
	
	
	item.DbAlanAdi =  request.FormValue("dbalanadi")
	
	
	item.DbAlanVeriTuru =  request.FormValue("dbalanverituru")
	

	context := Context{}

	//if len(item.ProjeAdi) > 0 {
		if id > 0 {
			AlanUpdate(db, item)
		} else {
			item.Id = AlanInsert(db, item)
		}
		context.Message = "Kayıt yapıldı"
	//} else {
	//	context.Message = "Lütfen Zorunlu alanları giriniz"
	//}

	context.Data = item
	render(response, request, "alan", context)

}

