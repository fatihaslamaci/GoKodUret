package main

import (
	"database/sql"
)

func ProjeVarmi(db *sql.DB) bool {
	id := 0
	err := db.QueryRow("Select id from projeler where id>? Limit 1", 0).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		return false
	}
	return id > 0
}

func ProjeDoldur(db *sql.DB){

	if ProjeVarmi(db) ==false {
		projeler := JsonDataOku()

		for _, proje := range projeler {
			projeId:=ProjeInsert(db,proje)
			for _, sinif := range proje.Siniflar {
				sinif.ProjeId = projeId
				sinifId:=SinifInsert(db,sinif)
				for _, alan := range sinif.Alanlar {
					alan.SinifId = sinifId
					AlanInsert(db,alan)
				}
				for _, ekozellik := range sinif.TabloEkOzellikler {
					ekozellik.SinifId = sinifId
					TabloEkOzellikInsert(db,ekozellik)
				}
			}
		}
	}


}