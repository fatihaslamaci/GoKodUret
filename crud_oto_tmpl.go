package main

import (
	"database/sql"
)


//0. Proje  **************************************************************************************************

//{{ .Proje }}


func ProjeSelectAll(db *sql.DB) []Proje {
	rows, err := db.Query("Select id, proje_adi, proje_yolu from projeler")
	CheckErr(err)
	var result  []Proje
	for rows.Next() {
		item :=Proje{}
		err2 := rows.Scan(&item.Id, &item.ProjeAdi, &item.ProjeYolu)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


func ProjeSelect(db *sql.DB, id int) Proje {
	item := Proje{}
	if id > 0 {
		row := db.QueryRow("Select id, proje_adi, proje_yolu from projeler where id=?", id)
		err := row.Scan(&item.Id, &item.ProjeAdi, &item.ProjeYolu)
		CheckErr(err)
	}
	return item
}

func ProjeInsert(db *sql.DB, item Proje) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO projeler(proje_adi, proje_yolu) VALUES (?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.ProjeAdi, item.ProjeYolu)
	CheckErr(err)
	r,err = ret.LastInsertId()
	CheckErr(err)
	return r
}

func ProjeUpdate(db *sql.DB, item Proje) {
	stmt, err := db.Prepare("Update projeler set proje_adi=?, proje_yolu=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.ProjeAdi, item.ProjeYolu, item.Id)
	CheckErr(err2)
}


//1. Sinif  **************************************************************************************************

//{{ .Sinif }}


func SinifSelectAll(db *sql.DB) []Sinif {
	rows, err := db.Query("Select id, proje_id, sinif_adi, tablo_adi from siniflar")
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


func SinifSelect(db *sql.DB, id int) Sinif {
	item := Sinif{}
	if id > 0 {
		row := db.QueryRow("Select id, proje_id, sinif_adi, tablo_adi from siniflar where id=?", id)
		err := row.Scan(&item.Id, &item.ProjeId, &item.SinifAdi, &item.TabloAdi)
		CheckErr(err)
	}
	return item
}

func SinifInsert(db *sql.DB, item Sinif) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO siniflar(proje_id, sinif_adi, tablo_adi) VALUES (?,?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.ProjeId, item.SinifAdi, item.TabloAdi)
	CheckErr(err)
	r,err = ret.LastInsertId()
	CheckErr(err)
	return r
}

func SinifUpdate(db *sql.DB, item Sinif) {
	stmt, err := db.Prepare("Update siniflar set proje_id=?, sinif_adi=?, tablo_adi=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.ProjeId, item.SinifAdi, item.TabloAdi, item.Id)
	CheckErr(err2)
}


//2. Alan  **************************************************************************************************

//{{ .Alan }}


func AlanSelectAll(db *sql.DB) []Alan {
	rows, err := db.Query("Select id, is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu from alanlar")
	CheckErr(err)
	var result  []Alan
	for rows.Next() {
		item :=Alan{}
		err2 := rows.Scan(&item.Id, &item.IsId, &item.SinifId, &item.AlanAdi, &item.AlanVeriTuru, &item.DbAlanAdi, &item.DbAlanVeriTuru)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


func AlanSelect(db *sql.DB, id int) Alan {
	item := Alan{}
	if id > 0 {
		row := db.QueryRow("Select id, is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu from alanlar where id=?", id)
		err := row.Scan(&item.Id, &item.IsId, &item.SinifId, &item.AlanAdi, &item.AlanVeriTuru, &item.DbAlanAdi, &item.DbAlanVeriTuru)
		CheckErr(err)
	}
	return item
}

func AlanInsert(db *sql.DB, item Alan) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO alanlar(is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu) VALUES (?,?,?,?,?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.IsId, item.SinifId, item.AlanAdi, item.AlanVeriTuru, item.DbAlanAdi, item.DbAlanVeriTuru)
	CheckErr(err)
	r,err = ret.LastInsertId()
	CheckErr(err)
	return r
}

func AlanUpdate(db *sql.DB, item Alan) {
	stmt, err := db.Prepare("Update alanlar set is_id=?, sinif_id=?, alan_adi=?, alan_veri_turu=?, db_alan_adi=?, db_alan_veri_turu=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.IsId, item.SinifId, item.AlanAdi, item.AlanVeriTuru, item.DbAlanAdi, item.DbAlanVeriTuru, item.Id)
	CheckErr(err2)
}
