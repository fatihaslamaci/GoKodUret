package main

import "database/sql"


// bu iptal otomatik üretiliyor
func SinifSelectAllMasterId_OLD(db *sql.DB, Masterid int64) []Sinif {
	rows, err := db.Query("Select id, proje_id, sinif_adi, tablo_adi from siniflar where proje_id=?", Masterid)
	CheckErr(err)
	var result  []Sinif
	for rows.Next() {
		item :=Sinif{}
		err2 := rows.Scan(&item.Id, &item.ProjeId, &item.SinifAdi, &item.TabloAdi)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


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
		projeler := DataOku()

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