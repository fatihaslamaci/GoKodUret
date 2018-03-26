package main

import "net/http"

func IdAlaniEkle(item Sinif) {
	alan := Alan{}
	alan.SinifId = item.Id
	alan.AlanAdi = "Id"
	alan.AlanVeriTuru = "int64"
	alan.DbAlanAdi = "id"
	alan.DbAlanVeriTuru = "INTEGER primary key autoincrement"
	alan.HtmlInputType = "hidden"
	alan.IsForeignKey = false
	alan.IsId = true
	//alan.IsMasterId = false
	AlanInsert(db, alan)
}




func SinifKaydetHandler2(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	MasterId := int64(0)
	id := FormValueInt64(request, "id")
	item := SinifSelect(db, id)
	SinifFormValue(&item, request)
	MasterId = item.ProjeId
	context := Context{}

	if Message := SinifKaydetValidate(&item); len(Message) > 0 {
		if id > 0 {
			SinifUpdate(db, item)
		} else {
			item.Id = SinifInsert(db, item)
			IdAlaniEkle(item)
		}
		context.Message = "Kayıt yapıldı"
	}else{
		context.Message=Message[0]
	}



	context.Data = item
	context.Gezgin = GetGezgin(MasterId, "sinif")
	render(response, request, "sinif", context)
}
