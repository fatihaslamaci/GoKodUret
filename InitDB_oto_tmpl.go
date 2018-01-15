package main

import (
	"database/sql"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	CheckErr(err)

	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `

CREATE TABLE IF NOT EXISTS projeler(
   id INTEGER primary key autoincrement  
  ,proje_adi VARCHAR(50)   
  ,proje_yolu VARCHAR(500)  
);

CREATE TABLE IF NOT EXISTS siniflar(
   id INTEGER primary key autoincrement  
  ,proje_id INTEGER   
  ,sinif_adi VARCHAR(50)   
  ,tablo_adi VARCHAR(50)   
  , FOREIGN KEY(proje_id) REFERENCES projeler(id)  
);

CREATE TABLE IF NOT EXISTS alanlar(
   id INTEGER primary key autoincrement  
  ,is_id BIT   
  ,sinif_id INTEGER   
  ,alan_adi VARCHAR(50)   
  ,alan_veri_turu VARCHAR(50)   
  ,db_alan_adi varchar(50)   
  ,db_alan_veri_turu varchar(100)   
  , FOREIGN KEY(sinif_id) REFERENCES siniflar(id)  
);

`
	_, err := db.Exec(sql_table)

	CheckErr(err)

}

