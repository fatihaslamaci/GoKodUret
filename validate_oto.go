package main

import ()

func ProjeKaydetValidate(item *Proje) []string {
	var hata []string
	if item.ProjeAdi == "" {
		hata = append(hata, "ProjeAdi alanı boş geçilemez.")
	}
	if len(item.ProjeAdi) < 1 {
		hata = append(hata, "ProjeAdi alanı 1 karakterden az olamaz.")
	}
	if len(item.ProjeAdi) > 50 {
		hata = append(hata, "ProjeAdi alanı 50 karakterden çok olamaz.")
	}
	if item.ProjeYolu == "" {
		hata = append(hata, "ProjeYolu alanı boş geçilemez.")
	}
	if len(item.ProjeYolu) < 1 {
		hata = append(hata, "ProjeYolu alanı 1 karakterden az olamaz.")
	}
	if len(item.ProjeYolu) > 500 {
		hata = append(hata, "ProjeYolu alanı 500 karakterden çok olamaz.")
	}
	return hata
}
func SinifKaydetValidate(item *Sinif) []string {
	var hata []string
	return hata
}
func AlanKaydetValidate(item *Alan) []string {
	var hata []string
	return hata
}
func TabloEkOzellikKaydetValidate(item *TabloEkOzellik) []string {
	var hata []string
	return hata
}
func AnahtarDegerKaydetValidate(item *AnahtarDeger) []string {
	var hata []string
	return hata
}
