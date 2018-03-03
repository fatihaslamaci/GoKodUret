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
func FormValueBoolCheckbox(request *http.Request, s string) bool {
	r := false
	if len(request.Form[s]) > 0 {
		r = true
	}
	return r
}
func FormValueDate(request *http.Request, s string) time.Time {
	r, _ := time.Parse("2006-01-02", request.FormValue(s))
	return r
}
func ProjeFormValue(item *Proje, request *http.Request) {
	item.ProjeAdi = request.FormValue("projeadi")
	item.ProjeYolu = request.FormValue("projeyolu")
}
func SinifFormValue(item *Sinif, request *http.Request) {
	item.ProjeId = FormValueInt64(request, "projeid")
	item.SinifAdi = request.FormValue("sinifadi")
	item.TabloAdi = request.FormValue("tabloadi")
	item.DetailTablo = FormValueBool(request, "detailtablo")
}
func AlanFormValue(item *Alan, request *http.Request) {
	item.IsId = FormValueBoolCheckbox(request, "isid")
	item.SinifId = FormValueInt64(request, "sinifid")
	item.AlanAdi = request.FormValue("alanadi")
	item.AlanVeriTuru = request.FormValue("alanverituru")
	item.DbAlanAdi = request.FormValue("dbalanadi")
	item.DbAlanVeriTuru = request.FormValue("dbalanverituru")
	item.HtmlInputType = request.FormValue("htmlinputtype")
	item.IsForeignKey = FormValueBoolCheckbox(request, "isforeignkey")
}
func TabloEkOzellikFormValue(item *TabloEkOzellik, request *http.Request) {
	item.SinifId = FormValueInt64(request, "sinifid")
	item.Ozellik = request.FormValue("ozellik")
}
