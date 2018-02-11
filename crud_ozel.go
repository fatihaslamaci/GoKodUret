package main

import (
	"database/sql"
	"strings"
	"fmt"
)

type TableProp struct {
	Cid   int64
	Name   string
	Type   string
	Notnull bool
	Dflt_value sql.NullString
	Pk bool
}


func GetTableProp(db *sql.DB, tableName string) []TableProp {

	query:="PRAGMA TABLE_INFO("+tableName+")"

	fmt.Println(query)

	rows, err := db.Query(query,nil)
	CheckErr(err)
	var result  []TableProp
	for rows.Next() {
		item :=TableProp{}
		err2 := rows.Scan(&item.Cid, &item.Name, &item.Type, &item.Notnull, &item.Dflt_value,&item.Pk)
		CheckErr(err2)
		result = append(result, item)
	}
	return result
}


func AlterDb(db *sql.DB){
	var s = strings.Split(sql_table,"\n")
	var tableName=""
	tableProps:=[]TableProp{}

	for _,satir := range s{

		satir = strings.Trim(satir," ")
		satir = strings.Trim(satir,"\t")


		if (satir==""){
			continue
		}
		if (strings.Index(satir,"FOREIGN KEY(")>=0){
			continue
		}
		if (satir==");"){
			continue
		}

		satir = strings.Trim(satir,",")


		slen:=len("CREATE TABLE IF NOT EXISTS")

		indexof:=strings.LastIndex(satir,"CREATE TABLE IF NOT EXISTS")
		if (indexof>=0){
			tableName = strings.Trim(((strings.Trim(satir,"("))[indexof+slen+1:]),"")
			tableProps = GetTableProp(db,tableName)
		}else{
			if (tableName !=""){

				i:=strings.Index(satir," ")

				field:=satir[0:i]

				bulundu :=false

				//fmt.Println(field)

				//fmt.Println(tableProps)

				for _,tableProp:= range(tableProps){


					if (tableProp.Name==field){
						fmt.Print("'")
						fmt.Print(tableProp.Name)
						fmt.Print("'")
						fmt.Print("=")
						fmt.Print("'")
						fmt.Print(field)
						fmt.Println("'")
						bulundu=true;
						break;
					}
				}
				if (bulundu==false){

					cmd:="ALTER TABLE "+tableName+" ADD COLUMN "+satir
					fmt.Println(cmd)
					_, err := db.Exec(cmd)
					CheckErr(err)
				}
			}
		}
	}
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