package main

import (
	"database/sql"
	"fmt"
	"strings"
)

const sql_table = `
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
        ,detail_tablo bit
        ,FOREIGN KEY(proje_id) REFERENCES projeler(id)
);
CREATE TABLE IF NOT EXISTS alanlar(
        id INTEGER primary key autoincrement
        ,is_id BIT
        ,sinif_id INTEGER
        ,alan_adi VARCHAR(50)
        ,alan_veri_turu VARCHAR(50)
        ,db_alan_adi varchar(50)
        ,db_alan_veri_turu varchar(100)
        ,html_input_type varchar(50)
        ,is_foreign_key BIT
        ,requered BIT
        ,minlength VARCHAR(50)
        ,maxlength VARCHAR(50)
        ,minvalue VARCHAR(50)
        ,maxvalue VARCHAR(50)
        ,regexpatern VARCHAR(100)
        ,master_table_name VARCHAR(50)
        ,like_field BIT
        ,sira_no INTEGER
        ,FOREIGN KEY(sinif_id) REFERENCES siniflar(id)
);
CREATE TABLE IF NOT EXISTS tablo_ek_ozellikler(
        id INTEGER primary key autoincrement
        ,sinif_id INTEGER
        ,ozellik VARCHAR(150)
        ,FOREIGN KEY(sinif_id) REFERENCES siniflar(id)
);
CREATE TABLE IF NOT EXISTS anahtardegerler(
        id INTEGER primary key autoincrement
        ,alan_id INTEGER
        ,anahtar VARCHAR(50)
        ,deger VARCHAR(50)
        ,FOREIGN KEY(alan_id) REFERENCES alanlar(id)
);
`

type tableProp struct {
	Cid        int64
	Name       string
	Type       string
	Notnull    bool
	Dflt_value sql.NullString
	Pk         bool
}

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
func getTableProp(db *sql.DB, tableName string) []tableProp {
	query := "PRAGMA TABLE_INFO(" + tableName + ")"
	rows, err := db.Query(query, nil)
	CheckErr(err)
	var result []tableProp
	for rows.Next() {
		item := tableProp{}
		err2 := rows.Scan(&item.Cid, &item.Name, &item.Type, &item.Notnull, &item.Dflt_value, &item.Pk)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}
func indexArtiLen(s, substr string) int {
	r := -1
	i := strings.Index(s, substr)
	if i >= 0 {
		r = i + len(substr)
	}
	return r
}
func AlterDb(db *sql.DB) {
	var s = strings.Split(sql_table, "\n")
	var tableName = ""
	tableProps := []tableProp{}
	for _, satir := range s {
		satir = strings.Trim(satir, " ")
		satir = strings.Trim(satir, "\t")
		if (satir == "") || (strings.Index(satir, "FOREIGN KEY(") >= 0) || (satir == ");") {
			continue
		}
		satir = strings.Trim(satir, ",")
		indexof := indexArtiLen(satir, "CREATE TABLE IF NOT EXISTS")
		if indexof >= 0 {
			tableName = strings.Trim(((strings.Trim(satir, "("))[indexof+1:]), "")
			tableProps = getTableProp(db, tableName)
		} else {
			if tableName != "" {
				i := strings.Index(satir, " ")
				field := satir[0:i]
				bulundu := false
				for _, tableProp := range tableProps {
					if tableProp.Name == field {
						bulundu = true
						break
					}
				}
				if bulundu == false {
					cmd := "ALTER TABLE " + tableName + " ADD COLUMN " + satir
					fmt.Println(cmd)
					_, err := db.Exec(cmd)
					CheckErr(err)
				}
			}
		}
	}
}
func CreateTable(db *sql.DB) {
	// create table if not exists
	_, err := db.Exec(sql_table)
	CheckErr(err)
}
