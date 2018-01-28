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


func ProjeSelect(db *sql.DB, id int64) Proje {
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

func ProjeDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("Delete projeler WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect

}

//1. Sinif  **************************************************************************************************

//{{ .Sinif }}


func SinifSelectAll(db *sql.DB) []Sinif {
	rows, err := db.Query("Select id, proje_id, sinif_adi, tablo_adi, detail_tablo from siniflar")
	CheckErr(err)
	var result  []Sinif
	for rows.Next() {
		item :=Sinif{}
		err2 := rows.Scan(&item.Id, &item.ProjeId, &item.SinifAdi, &item.TabloAdi, &item.DetailTablo)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


func SinifSelect(db *sql.DB, id int64) Sinif {
	item := Sinif{}

	if id > 0 {
		row := db.QueryRow("Select id, proje_id, sinif_adi, tablo_adi, detail_tablo from siniflar where id=?", id)
		err := row.Scan(&item.Id, &item.ProjeId, &item.SinifAdi, &item.TabloAdi, &item.DetailTablo)
		CheckErr(err)
	}
	return item
}

func SinifInsert(db *sql.DB, item Sinif) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO siniflar(proje_id, sinif_adi, tablo_adi, detail_tablo) VALUES (?,?,?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.ProjeId, item.SinifAdi, item.TabloAdi, item.DetailTablo)
	CheckErr(err)
	r,err = ret.LastInsertId()
	CheckErr(err)
	return r
}

func SinifUpdate(db *sql.DB, item Sinif) {
	stmt, err := db.Prepare("Update siniflar set proje_id=?, sinif_adi=?, tablo_adi=?, detail_tablo=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.ProjeId, item.SinifAdi, item.TabloAdi, item.DetailTablo, item.Id)
	CheckErr(err2)
}

func SinifDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("Delete siniflar WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect

}


func SinifSelectMasterId(db *sql.DB, Masterid int64) []Sinif {
	rows, err := db.Query("Select id, proje_id, sinif_adi, tablo_adi, detail_tablo from siniflar where proje_id=?", Masterid)
	CheckErr(err)
	var result  []Sinif
	for rows.Next() {
		item :=Sinif{}
		err2 := rows.Scan(&item.Id, &item.ProjeId, &item.SinifAdi, &item.TabloAdi, &item.DetailTablo)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


//2. Alan  **************************************************************************************************

//{{ .Alan }}


func AlanSelectAll(db *sql.DB) []Alan {
	rows, err := db.Query("Select id, is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu, html_input_type from alanlar")
	CheckErr(err)
	var result  []Alan
	for rows.Next() {
		item :=Alan{}
		err2 := rows.Scan(&item.Id, &item.IsId, &item.SinifId, &item.AlanAdi, &item.AlanVeriTuru, &item.DbAlanAdi, &item.DbAlanVeriTuru, &item.HtmlInputType)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


func AlanSelect(db *sql.DB, id int64) Alan {
	item := Alan{}

	if id > 0 {
		row := db.QueryRow("Select id, is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu, html_input_type from alanlar where id=?", id)
		err := row.Scan(&item.Id, &item.IsId, &item.SinifId, &item.AlanAdi, &item.AlanVeriTuru, &item.DbAlanAdi, &item.DbAlanVeriTuru, &item.HtmlInputType)
		CheckErr(err)
	}
	return item
}

func AlanInsert(db *sql.DB, item Alan) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO alanlar(is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu, html_input_type) VALUES (?,?,?,?,?,?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.IsId, item.SinifId, item.AlanAdi, item.AlanVeriTuru, item.DbAlanAdi, item.DbAlanVeriTuru, item.HtmlInputType)
	CheckErr(err)
	r,err = ret.LastInsertId()
	CheckErr(err)
	return r
}

func AlanUpdate(db *sql.DB, item Alan) {
	stmt, err := db.Prepare("Update alanlar set is_id=?, sinif_id=?, alan_adi=?, alan_veri_turu=?, db_alan_adi=?, db_alan_veri_turu=?, html_input_type=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.IsId, item.SinifId, item.AlanAdi, item.AlanVeriTuru, item.DbAlanAdi, item.DbAlanVeriTuru, item.HtmlInputType, item.Id)
	CheckErr(err2)
}

func AlanDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("Delete FROM alanlar WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect

}


func AlanSelectMasterId(db *sql.DB, Masterid int64) []Alan {
	rows, err := db.Query("Select id, is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu, html_input_type from alanlar where sinif_id=?", Masterid)
	CheckErr(err)
	var result  []Alan
	for rows.Next() {
		item :=Alan{}
		err2 := rows.Scan(&item.Id, &item.IsId, &item.SinifId, &item.AlanAdi, &item.AlanVeriTuru, &item.DbAlanAdi, &item.DbAlanVeriTuru, &item.HtmlInputType)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


//3. TabloEkOzellik  **************************************************************************************************

//{{ .TabloEkOzellik }}


func TabloEkOzellikSelectAll(db *sql.DB) []TabloEkOzellik {
	rows, err := db.Query("Select id, sinif_id, ozellik from tablo_ek_ozellikler")
	CheckErr(err)
	var result  []TabloEkOzellik
	for rows.Next() {
		item :=TabloEkOzellik{}
		err2 := rows.Scan(&item.Id, &item.SinifId, &item.Ozellik)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


func TabloEkOzellikSelect(db *sql.DB, id int64) TabloEkOzellik {
	item := TabloEkOzellik{}

	if id > 0 {
		row := db.QueryRow("Select id, sinif_id, ozellik from tablo_ek_ozellikler where id=?", id)
		err := row.Scan(&item.Id, &item.SinifId, &item.Ozellik)
		CheckErr(err)
	}
	return item
}

func TabloEkOzellikInsert(db *sql.DB, item TabloEkOzellik) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO tablo_ek_ozellikler(sinif_id, ozellik) VALUES (?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.SinifId, item.Ozellik)
	CheckErr(err)
	r,err = ret.LastInsertId()
	CheckErr(err)
	return r
}

func TabloEkOzellikUpdate(db *sql.DB, item TabloEkOzellik) {
	stmt, err := db.Prepare("Update tablo_ek_ozellikler set sinif_id=?, ozellik=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.SinifId, item.Ozellik, item.Id)
	CheckErr(err2)
}

func TabloEkOzellikDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("Delete tablo_ek_ozellikler WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect

}


func TabloEkOzellikSelectMasterId(db *sql.DB, Masterid int64) []TabloEkOzellik {
	rows, err := db.Query("Select id, sinif_id, ozellik from tablo_ek_ozellikler where sinif_id=?", Masterid)
	CheckErr(err)
	var result  []TabloEkOzellik
	for rows.Next() {
		item :=TabloEkOzellik{}
		err2 := rows.Scan(&item.Id, &item.SinifId, &item.Ozellik)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}

