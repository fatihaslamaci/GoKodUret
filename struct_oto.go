package main

import (
	_"time"
)

type Proje struct {
	Id        int64
	ProjeAdi  string
	ProjeYolu string
	Siniflar  []Sinif
}
type Sinif struct {
	Id                int64
	ProjeId           int64
	SinifAdi          string
	TabloAdi          string
	DetailTablo       bool
	Alanlar           []Alan
	TabloEkOzellikler []TabloEkOzellik
}
type Alan struct {
	Id              int64
	IsId            bool
	SinifId         int64
	AlanAdi         string
	AlanVeriTuru    string
	DbAlanAdi       string
	DbAlanVeriTuru  string
	HtmlInputType   string
	IsForeignKey    bool
	AnahtarDegerler []AnahtarDeger
	Requered        bool
	MinLength       string
	MaxLength       string
	MinValue        string
	MaxValue        string
	RegexPatern     string
	MasterTableName string
}
type TabloEkOzellik struct {
	Id      int64
	SinifId int64
	Ozellik string
}
type AnahtarDeger struct {
	Id      int64
	AlanId  int64
	Anahtar string
	Deger   string
}
