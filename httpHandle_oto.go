package main

import (
	"net/http"
)

func HandleFuncAdd() {
	http.HandleFunc("/projeler.html", ProjelerHandler)
	http.HandleFunc("/proje.html", ProjeHandler)
	http.HandleFunc("/projekaydet", ProjeKaydetHandler)
	http.HandleFunc("/projesil", ProjeSilHandler)
	http.HandleFunc("/sinifler.html", SiniflerHandler)
	http.HandleFunc("/sinif.html", SinifHandler)
	http.HandleFunc("/sinifkaydet", SinifKaydetHandler)
	http.HandleFunc("/sinifsil", SinifSilHandler)
	http.HandleFunc("/alanler.html", AlanlerHandler)
	http.HandleFunc("/alan.html", AlanHandler)
	http.HandleFunc("/alankaydet", AlanKaydetHandler)
	http.HandleFunc("/alansil", AlanSilHandler)
	http.HandleFunc("/tabloekozellikler.html", TabloEkOzelliklerHandler)
	http.HandleFunc("/tabloekozellik.html", TabloEkOzellikHandler)
	http.HandleFunc("/tabloekozellikkaydet", TabloEkOzellikKaydetHandler)
	http.HandleFunc("/tabloekozelliksil", TabloEkOzellikSilHandler)
	http.HandleFunc("/anahtardegerler.html", AnahtarDegerlerHandler)
	http.HandleFunc("/anahtardeger.html", AnahtarDegerHandler)
	http.HandleFunc("/anahtardegerkaydet", AnahtarDegerKaydetHandler)
	http.HandleFunc("/anahtardegersil", AnahtarDegerSilHandler)
}

//-----------------------------------------------------------------------
func ProjelerHandler(response http.ResponseWriter, request *http.Request) {
	fData := ProjeSelectAll(db)
	context := Context{Data: fData}
	context.Gezgin = GetGezgin(0, "projeler")
	render(response, request, "projeler", context)
}
func ProjeHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := ProjeSelect(db, id)
	context := Context{Data: item}
	context.Gezgin = GetGezgin(MasterId, "proje")
	render(response, request, "proje", context)
}
func ProjeKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := ProjeSelect(db, id)
	ProjeFormValue(&item, request)
	context := Context{}
	if Message := ProjeKaydetValidate(&item); len(Message) == 0 {
		if id > 0 {
			ProjeUpdate(db, item)
			context.Message = "Kayıt güncellendi"
		} else {
			item.Id = ProjeInsert(db, item)
			context.Message = "Yeni kayıt yapıldı"
		}
	} else {
		context.Message = Message[0]
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "proje")
	render(response, request, "proje", context)
}
func ProjeSilHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := ProjeSelect(db, id)
	context := Context{}
	if id > 0 {
		ProjeDelete(db, id)
		context.Message = "Kayıt Silindi"
	} else {
		context.Message = "Kayıt Bulunamadı"
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "proje")
	render(response, request, "proje", context)
}
func SiniflerHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := FormValueInt64(request, "id")
	fData := SinifSelectMasterId(db, MasterId)
	context := Context{Data: fData, MasterId: MasterId}
	context.Gezgin = GetGezgin(MasterId, "sinif")
	render(response, request, "sinifler", context)
}
func SinifHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := SinifSelect(db, id)
	MasterId = item.ProjeId
	if item.ProjeId == 0 {
		item.ProjeId = FormValueInt64(request, "MasterId")
		MasterId = item.ProjeId
	}
	context := Context{Data: item}
	context.Gezgin = GetGezgin(MasterId, "sinif")
	render(response, request, "sinif", context)
}
func SinifKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := SinifSelect(db, id)
	SinifFormValue(&item, request)
	MasterId = item.ProjeId
	context := Context{}
	if Message := SinifKaydetValidate(&item); len(Message) == 0 {
		if id > 0 {
			SinifUpdate(db, item)
			context.Message = "Kayıt güncellendi"
		} else {
			item.Id = SinifInsert(db, item)
			context.Message = "Yeni kayıt yapıldı"
		}
	} else {
		context.Message = Message[0]
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "sinif")
	render(response, request, "sinif", context)
}
func SinifSilHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := SinifSelect(db, id)
	MasterId = item.ProjeId
	context := Context{}
	if id > 0 {
		SinifDelete(db, id)
		context.Message = "Kayıt Silindi"
	} else {
		context.Message = "Kayıt Bulunamadı"
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "sinif")
	render(response, request, "sinif", context)
}
func AlanlerHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := FormValueInt64(request, "id")
	fData := AlanSelectMasterId(db, MasterId)
	context := Context{Data: fData, MasterId: MasterId}
	context.Gezgin = GetGezgin(MasterId, "alan")
	render(response, request, "alanler", context)
}
func AlanHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := AlanSelect(db, id)
	MasterId = item.SinifId
	if item.SinifId == 0 {
		item.SinifId = FormValueInt64(request, "MasterId")
		MasterId = item.SinifId
	}
	context := Context{Data: item}
	context.Gezgin = GetGezgin(MasterId, "alan")
	render(response, request, "alan", context)
}
func AlanKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := AlanSelect(db, id)
	AlanFormValue(&item, request)
	MasterId = item.SinifId
	context := Context{}
	if Message := AlanKaydetValidate(&item); len(Message) == 0 {
		if id > 0 {
			AlanUpdate(db, item)
			context.Message = "Kayıt güncellendi"
		} else {
			item.Id = AlanInsert(db, item)
			context.Message = "Yeni kayıt yapıldı"
		}
	} else {
		context.Message = Message[0]
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "alan")
	render(response, request, "alan", context)
}
func AlanSilHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := AlanSelect(db, id)
	MasterId = item.SinifId
	context := Context{}
	if id > 0 {
		AlanDelete(db, id)
		context.Message = "Kayıt Silindi"
	} else {
		context.Message = "Kayıt Bulunamadı"
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "alan")
	render(response, request, "alan", context)
}
func TabloEkOzelliklerHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := FormValueInt64(request, "id")
	fData := TabloEkOzellikSelectMasterId(db, MasterId)
	context := Context{Data: fData, MasterId: MasterId}
	context.Gezgin = GetGezgin(MasterId, "tabloekozellik")
	render(response, request, "tabloekozellikler", context)
}
func TabloEkOzellikHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := TabloEkOzellikSelect(db, id)
	MasterId = item.SinifId
	if item.SinifId == 0 {
		item.SinifId = FormValueInt64(request, "MasterId")
		MasterId = item.SinifId
	}
	context := Context{Data: item}
	context.Gezgin = GetGezgin(MasterId, "tabloekozellik")
	render(response, request, "tabloekozellik", context)
}
func TabloEkOzellikKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := TabloEkOzellikSelect(db, id)
	TabloEkOzellikFormValue(&item, request)
	MasterId = item.SinifId
	context := Context{}
	if Message := TabloEkOzellikKaydetValidate(&item); len(Message) == 0 {
		if id > 0 {
			TabloEkOzellikUpdate(db, item)
			context.Message = "Kayıt güncellendi"
		} else {
			item.Id = TabloEkOzellikInsert(db, item)
			context.Message = "Yeni kayıt yapıldı"
		}
	} else {
		context.Message = Message[0]
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "tabloekozellik")
	render(response, request, "tabloekozellik", context)
}
func TabloEkOzellikSilHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := TabloEkOzellikSelect(db, id)
	MasterId = item.SinifId
	context := Context{}
	if id > 0 {
		TabloEkOzellikDelete(db, id)
		context.Message = "Kayıt Silindi"
	} else {
		context.Message = "Kayıt Bulunamadı"
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "tabloekozellik")
	render(response, request, "tabloekozellik", context)
}
func AnahtarDegerlerHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := FormValueInt64(request, "id")
	fData := AnahtarDegerSelectMasterId(db, MasterId)
	context := Context{Data: fData, MasterId: MasterId}
	context.Gezgin = GetGezgin(MasterId, "anahtardeger")
	render(response, request, "anahtardegerler", context)
}
func AnahtarDegerHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := AnahtarDegerSelect(db, id)
	MasterId = item.AlanId
	if item.AlanId == 0 {
		item.AlanId = FormValueInt64(request, "MasterId")
		MasterId = item.AlanId
	}
	context := Context{Data: item}
	context.Gezgin = GetGezgin(MasterId, "anahtardeger")
	render(response, request, "anahtardeger", context)
}
func AnahtarDegerKaydetHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := AnahtarDegerSelect(db, id)
	AnahtarDegerFormValue(&item, request)
	MasterId = item.AlanId
	context := Context{}
	if Message := AnahtarDegerKaydetValidate(&item); len(Message) == 0 {
		if id > 0 {
			AnahtarDegerUpdate(db, item)
			context.Message = "Kayıt güncellendi"
		} else {
			item.Id = AnahtarDegerInsert(db, item)
			context.Message = "Yeni kayıt yapıldı"
		}
	} else {
		context.Message = Message[0]
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "anahtardeger")
	render(response, request, "anahtardeger", context)
}
func AnahtarDegerSilHandler(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := AnahtarDegerSelect(db, id)
	MasterId = item.AlanId
	context := Context{}
	if id > 0 {
		AnahtarDegerDelete(db, id)
		context.Message = "Kayıt Silindi"
	} else {
		context.Message = "Kayıt Bulunamadı"
	}
	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "anahtardeger")
	render(response, request, "anahtardeger", context)
}
