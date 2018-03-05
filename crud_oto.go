package main

import (
	"database/sql"
)

//0. Proje  **************************************************************************************************
//{{ .Proje }}
func ProjeSelectAll(db *sql.DB) []Proje {
	rows, err := db.Query(`SELECT
	            id
	            ,proje_adi
	            ,proje_yolu
	 FROM projeler`)
	CheckErr(err)
	var result []Proje
	for rows.Next() {
		item := Proje{}
		err2 := rows.Scan(
			&item.Id,
			&item.ProjeAdi,
			&item.ProjeYolu,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}
func ProjeSelect(db *sql.DB, id int64) Proje {
	item := Proje{}
	if id > 0 {
		row := db.QueryRow(`SELECT
		            id
		            ,proje_adi
		            ,proje_yolu
		FROM projeler WHERE id=?`, id)
		err := row.Scan(
			&item.Id,
			&item.ProjeAdi,
			&item.ProjeYolu,
		)
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
	r, err = ret.LastInsertId()
	CheckErr(err)
	return r
}
func ProjeUpdate(db *sql.DB, item Proje) {
	stmt, err := db.Prepare("UPDATE projeler SET proje_adi=?, proje_yolu=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.ProjeAdi, item.ProjeYolu, item.Id)
	CheckErr(err2)
}
func ProjeDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("DELETE FROM projeler WHERE id=?")
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
	rows, err := db.Query(`SELECT
	            id
	            ,proje_id
	            ,sinif_adi
	            ,tablo_adi
	            ,detail_tablo
	 FROM siniflar`)
	CheckErr(err)
	var result []Sinif
	for rows.Next() {
		item := Sinif{}
		err2 := rows.Scan(
			&item.Id,
			&item.ProjeId,
			&item.SinifAdi,
			&item.TabloAdi,
			&item.DetailTablo,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}
func SinifSelect(db *sql.DB, id int64) Sinif {
	item := Sinif{}
	if id > 0 {
		row := db.QueryRow(`SELECT
		            id
		            ,proje_id
		            ,sinif_adi
		            ,tablo_adi
		            ,detail_tablo
		FROM siniflar WHERE id=?`, id)
		err := row.Scan(
			&item.Id,
			&item.ProjeId,
			&item.SinifAdi,
			&item.TabloAdi,
			&item.DetailTablo,
		)
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
	r, err = ret.LastInsertId()
	CheckErr(err)
	return r
}
func SinifUpdate(db *sql.DB, item Sinif) {
	stmt, err := db.Prepare("UPDATE siniflar SET proje_id=?, sinif_adi=?, tablo_adi=?, detail_tablo=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.ProjeId, item.SinifAdi, item.TabloAdi, item.DetailTablo, item.Id)
	CheckErr(err2)
}
func SinifDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("DELETE FROM siniflar WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect
}

//----deneme
func SinifSelectMasterId(db *sql.DB, Masterid int64) []Sinif {
	rows, err := db.Query(`SELECT
	            id
	            ,proje_id
	            ,sinif_adi
	            ,tablo_adi
	            ,detail_tablo
	FROM siniflar WHERE proje_id=?`, Masterid)
	CheckErr(err)
	var result []Sinif
	for rows.Next() {
		item := Sinif{}
		err2 := rows.Scan(
			&item.Id,
			&item.ProjeId,
			&item.SinifAdi,
			&item.TabloAdi,
			&item.DetailTablo,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}

//2. Alan  **************************************************************************************************
//{{ .Alan }}
func AlanSelectAll(db *sql.DB) []Alan {
	rows, err := db.Query(`SELECT
	            id
	            ,is_id
	            ,sinif_id
	            ,alan_adi
	            ,alan_veri_turu
	            ,db_alan_adi
	            ,db_alan_veri_turu
	            ,html_input_type
	            ,is_foreign_key
	 FROM alanlar`)
	CheckErr(err)
	var result []Alan
	for rows.Next() {
		item := Alan{}
		err2 := rows.Scan(
			&item.Id,
			&item.IsId,
			&item.SinifId,
			&item.AlanAdi,
			&item.AlanVeriTuru,
			&item.DbAlanAdi,
			&item.DbAlanVeriTuru,
			&item.HtmlInputType,
			&item.IsForeignKey,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}
func AlanSelect(db *sql.DB, id int64) Alan {
	item := Alan{}
	if id > 0 {
		row := db.QueryRow(`SELECT
		            id
		            ,is_id
		            ,sinif_id
		            ,alan_adi
		            ,alan_veri_turu
		            ,db_alan_adi
		            ,db_alan_veri_turu
		            ,html_input_type
		            ,is_foreign_key
		FROM alanlar WHERE id=?`, id)
		err := row.Scan(
			&item.Id,
			&item.IsId,
			&item.SinifId,
			&item.AlanAdi,
			&item.AlanVeriTuru,
			&item.DbAlanAdi,
			&item.DbAlanVeriTuru,
			&item.HtmlInputType,
			&item.IsForeignKey,
		)
		CheckErr(err)
	}
	return item
}
func AlanInsert(db *sql.DB, item Alan) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO alanlar(is_id, sinif_id, alan_adi, alan_veri_turu, db_alan_adi, db_alan_veri_turu, html_input_type, is_foreign_key) VALUES (?,?,?,?,?,?,?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.IsId, item.SinifId, item.AlanAdi, item.AlanVeriTuru, item.DbAlanAdi, item.DbAlanVeriTuru, item.HtmlInputType, item.IsForeignKey)
	CheckErr(err)
	r, err = ret.LastInsertId()
	CheckErr(err)
	return r
}
func AlanUpdate(db *sql.DB, item Alan) {
	stmt, err := db.Prepare("UPDATE alanlar SET is_id=?, sinif_id=?, alan_adi=?, alan_veri_turu=?, db_alan_adi=?, db_alan_veri_turu=?, html_input_type=?, is_foreign_key=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.IsId, item.SinifId, item.AlanAdi, item.AlanVeriTuru, item.DbAlanAdi, item.DbAlanVeriTuru, item.HtmlInputType, item.IsForeignKey, item.Id)
	CheckErr(err2)
}
func AlanDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("DELETE FROM alanlar WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect
}

//----deneme
func AlanSelectMasterId(db *sql.DB, Masterid int64) []Alan {
	rows, err := db.Query(`SELECT
	            id
	            ,is_id
	            ,sinif_id
	            ,alan_adi
	            ,alan_veri_turu
	            ,db_alan_adi
	            ,db_alan_veri_turu
	            ,html_input_type
	            ,is_foreign_key
	FROM alanlar WHERE sinif_id=?`, Masterid)
	CheckErr(err)
	var result []Alan
	for rows.Next() {
		item := Alan{}
		err2 := rows.Scan(
			&item.Id,
			&item.IsId,
			&item.SinifId,
			&item.AlanAdi,
			&item.AlanVeriTuru,
			&item.DbAlanAdi,
			&item.DbAlanVeriTuru,
			&item.HtmlInputType,
			&item.IsForeignKey,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}

//3. TabloEkOzellik  **************************************************************************************************
//{{ .TabloEkOzellik }}
func TabloEkOzellikSelectAll(db *sql.DB) []TabloEkOzellik {
	rows, err := db.Query(`SELECT
	            id
	            ,sinif_id
	            ,ozellik
	 FROM tablo_ek_ozellikler`)
	CheckErr(err)
	var result []TabloEkOzellik
	for rows.Next() {
		item := TabloEkOzellik{}
		err2 := rows.Scan(
			&item.Id,
			&item.SinifId,
			&item.Ozellik,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}
func TabloEkOzellikSelect(db *sql.DB, id int64) TabloEkOzellik {
	item := TabloEkOzellik{}
	if id > 0 {
		row := db.QueryRow(`SELECT
		            id
		            ,sinif_id
		            ,ozellik
		FROM tablo_ek_ozellikler WHERE id=?`, id)
		err := row.Scan(
			&item.Id,
			&item.SinifId,
			&item.Ozellik,
		)
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
	r, err = ret.LastInsertId()
	CheckErr(err)
	return r
}
func TabloEkOzellikUpdate(db *sql.DB, item TabloEkOzellik) {
	stmt, err := db.Prepare("UPDATE tablo_ek_ozellikler SET sinif_id=?, ozellik=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.SinifId, item.Ozellik, item.Id)
	CheckErr(err2)
}
func TabloEkOzellikDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("DELETE FROM tablo_ek_ozellikler WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect
}

//----deneme
func TabloEkOzellikSelectMasterId(db *sql.DB, Masterid int64) []TabloEkOzellik {
	rows, err := db.Query(`SELECT
	            id
	            ,sinif_id
	            ,ozellik
	FROM tablo_ek_ozellikler WHERE sinif_id=?`, Masterid)
	CheckErr(err)
	var result []TabloEkOzellik
	for rows.Next() {
		item := TabloEkOzellik{}
		err2 := rows.Scan(
			&item.Id,
			&item.SinifId,
			&item.Ozellik,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}

//4. AnahtarDeger  **************************************************************************************************
//{{ .AnahtarDeger }}
func AnahtarDegerSelectAll(db *sql.DB) []AnahtarDeger {
	rows, err := db.Query(`SELECT
	            id
	            ,alan_id
	            ,anahtar
	            ,deger
	 FROM anahtardegerler`)
	CheckErr(err)
	var result []AnahtarDeger
	for rows.Next() {
		item := AnahtarDeger{}
		err2 := rows.Scan(
			&item.Id,
			&item.AlanId,
			&item.Anahtar,
			&item.Deger,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}
func AnahtarDegerSelect(db *sql.DB, id int64) AnahtarDeger {
	item := AnahtarDeger{}
	if id > 0 {
		row := db.QueryRow(`SELECT
		            id
		            ,alan_id
		            ,anahtar
		            ,deger
		FROM anahtardegerler WHERE id=?`, id)
		err := row.Scan(
			&item.Id,
			&item.AlanId,
			&item.Anahtar,
			&item.Deger,
		)
		CheckErr(err)
	}
	return item
}
func AnahtarDegerInsert(db *sql.DB, item AnahtarDeger) int64 {
	var r int64
	stmt, err := db.Prepare("INSERT INTO anahtardegerler(alan_id, anahtar, deger) VALUES (?,?,?)")
	CheckErr(err)
	defer stmt.Close()
	ret, err := stmt.Exec(item.AlanId, item.Anahtar, item.Deger)
	CheckErr(err)
	r, err = ret.LastInsertId()
	CheckErr(err)
	return r
}
func AnahtarDegerUpdate(db *sql.DB, item AnahtarDeger) {
	stmt, err := db.Prepare("UPDATE anahtardegerler SET alan_id=?, anahtar=?, deger=? WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	_, err2 := stmt.Exec(item.AlanId, item.Anahtar, item.Deger, item.Id)
	CheckErr(err2)
}
func AnahtarDegerDelete(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("DELETE FROM anahtardegerler WHERE id=?")
	CheckErr(err)
	defer stmt.Close()
	res, err := stmt.Exec(id)
	CheckErr(err)
	affect, err := res.RowsAffected()
	CheckErr(err)
	return affect
}

//----deneme
func AnahtarDegerSelectMasterId(db *sql.DB, Masterid int64) []AnahtarDeger {
	rows, err := db.Query(`SELECT
	            id
	            ,alan_id
	            ,anahtar
	            ,deger
	FROM anahtardegerler WHERE alan_id=?`, Masterid)
	CheckErr(err)
	var result []AnahtarDeger
	for rows.Next() {
		item := AnahtarDeger{}
		err2 := rows.Scan(
			&item.Id,
			&item.AlanId,
			&item.Anahtar,
			&item.Deger,
		)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}
