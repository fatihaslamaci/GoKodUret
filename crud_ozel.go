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
