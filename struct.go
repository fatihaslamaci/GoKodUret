package main

import (
	//"time"
)

type Proje struct {
	Id   int64
	ProjeAdi   string
	ProjeYolu   string
}

type Sinif struct {
	Id   int
	ProjeId   int
	SinifAdi   string
	AlanAdi   string
	AlanVeriTuru   string
	DbAlanAdi   string
	DbAlanVeriTuru   string

}

type Alan struct {
	Id   int
	IsId   bool
	SinifId   int
	AlanAdi   string
	AlanVeriTuru   string
	DbAlanAdi   string
	AbAlanVeriTuru   string

}
