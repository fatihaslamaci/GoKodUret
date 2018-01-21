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

    http.HandleFunc("/tabloekozellikler.html",TabloEkOzelliklerHandler)
	http.HandleFunc("/tabloekozellik.html",TabloEkOzellikHandler)
	http.HandleFunc("/tabloekozellikkaydet",TabloEkOzellikKaydetHandler)

}
//-----------------------------------------------------------------------




func ProjelerHandler(response http.ResponseWriter, request *http.Request) {
	fData  := ProjeSelectAll(db)
	context := Context{Data: fData}
	render(response, request, "projeler", context)
}



func ProjeHandler(response http.ResponseWriter, request *http.Request) {
    request.ParseForm()
	id := FormValueInt64(request,"Id")
    item := ProjeSelect(db, id)
    
    context := Context{Data: item}
	render(response, request, "proje", context)
}

func ProjeKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := FormValueInt64(request,"Id")
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
	request.ParseForm()
	MasterId:=FormValueInt64(request,"id")
	fData  := SinifSelectMasterId(db,MasterId)
    context := Context{Data: fData, MasterId:MasterId}

	render(response, request, "sinifler", context)
}



func SinifHandler(response http.ResponseWriter, request *http.Request) {
    request.ParseForm()
	id := FormValueInt64(request,"Id")
    item := SinifSelect(db, id)
    
    item.ProjeId=FormValueInt64(request,"MasterId")
    
    context := Context{Data: item}
	render(response, request, "sinif", context)
}

func SinifKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := FormValueInt64(request,"Id")
	item := SinifSelect(db, id)

	
	
	item.ProjeId =  FormValueInt64(request,"projeid")
	
	
	item.SinifAdi =  request.FormValue("sinifadi")
	
	
	item.TabloAdi =  request.FormValue("tabloadi")
	
	
	//item.DetailTablo =  request.FormValue("detailtablo")
	

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
	request.ParseForm()
	MasterId:=FormValueInt64(request,"id")
	fData  := AlanSelectMasterId(db,MasterId)
    context := Context{Data: fData, MasterId:MasterId}

	render(response, request, "alanler", context)
}



func AlanHandler(response http.ResponseWriter, request *http.Request) {
    request.ParseForm()
	id := FormValueInt64(request,"Id")
    item := AlanSelect(db, id)
    
    item.SinifId=FormValueInt64(request,"MasterId")
    
    context := Context{Data: item}
	render(response, request, "alan", context)
}

func AlanKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := FormValueInt64(request,"Id")
	item := AlanSelect(db, id)

	
	
	//item.IsId =  request.FormValue("isid")
	
	
	item.SinifId =  FormValueInt64(request,"sinifid")
	
	
	item.AlanAdi =  request.FormValue("alanadi")
	
	
	item.AlanVeriTuru =  request.FormValue("alanverituru")
	
	
	item.DbAlanAdi =  request.FormValue("dbalanadi")
	
	
	item.DbAlanVeriTuru =  request.FormValue("dbalanverituru")
	
	
	item.HtmlInputType =  request.FormValue("htmlinputtype")
	

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




func TabloEkOzelliklerHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId:=FormValueInt64(request,"id")
	fData  := TabloEkOzellikSelectMasterId(db,MasterId)
    context := Context{Data: fData, MasterId:MasterId}

	render(response, request, "tabloekozellikler", context)
}



func TabloEkOzellikHandler(response http.ResponseWriter, request *http.Request) {
    request.ParseForm()
	id := FormValueInt64(request,"Id")
    item := TabloEkOzellikSelect(db, id)
    
    item.SinifId=FormValueInt64(request,"MasterId")
    
    context := Context{Data: item}
	render(response, request, "tabloekozellik", context)
}

func TabloEkOzellikKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := FormValueInt64(request,"Id")
	item := TabloEkOzellikSelect(db, id)

	
	
	item.SinifId =  FormValueInt64(request,"sinifid")
	
	
	item.Ozellik =  request.FormValue("ozellik")
	

	context := Context{}

	//if len(item.ProjeAdi) > 0 {
		if id > 0 {
			TabloEkOzellikUpdate(db, item)
		} else {
			item.Id = TabloEkOzellikInsert(db, item)
		}
		context.Message = "Kayıt yapıldı"
	//} else {
	//	context.Message = "Lütfen Zorunlu alanları giriniz"
	//}

	context.Data = item
	render(response, request, "tabloekozellik", context)

}

