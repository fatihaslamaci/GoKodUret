package main

import (
	_"time"
)

type Proje struct {
  Id   int64
  ProjeAdi   string
  ProjeYolu   string
  Siniflar []Sinif
}

type Sinif struct {
  Id   int64
  ProjeId   int64
  SinifAdi   string
  TabloAdi string
  Alanlar []Alan
     
}

type Alan struct {
  Id   int64
  SinifId   int64
  AlanAdi   string
  AlanVeriTuru   string
  DbAlanAdi   string
  DbAlanVeriTuru   string
  IsForeignKey bool
  IsId   bool

}
