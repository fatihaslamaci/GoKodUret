package main

import (

	"testing"
)

var sc = `
Kullanici           :           :kullanicilar		:
Id                  :int64		:id 				:INTEGER primary key autoincrement
Ad                  :string	 	:ad 				:VARCHAR(50)
KayitTarihi         :time.Time  :kayittarihi 		:DATE
HataliGirisSayisi   :int        :hataligirissayisi	:INTEGER
BlokeHesap			:bool       :blokehesap 		:BIT
`
var HedefKullanicistruct = `type Kullanici struct {
Id int64
Ad string
KayitTarihi time.Time
HataliGirisSayisi int
BlokeHesap bool

}`

var HedefKullanicitable = `CREATE TABLE IF NOT EXISTS kullanicilar(
id INTEGER primary key autoincrement
,ad VARCHAR(50)
,kayittarihi DATE
,hataligirissayisi INTEGER
,blokehesap BIT

);`

var HedefSelectIdTmp = `func KullaniciSelect(db *sql.DB, id int) Kullanici {
	item := Kullanici{}
	if id > 0 {
		row := db.QueryRow("Select id, ad, kayittarihi, hataligirissayisi, blokehesap from kullanicilar where id=?", id)
		err := row.Scan(&item.Id, &item.Ad, &item.KayitTarihi, &item.HataliGirisSayisi, &item.BlokeHesap)
		CheckErr(err)
	}
	return item
}`


func TestSelectTable(t *testing.T) {
	type args struct {
		value string
		tmpl  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Kullanici",
			args:args{ value:sc, tmpl: "./template/selectsql.tmpl" },
			want: HedefSelectIdTmp,
		},
		{
			name: "KullaniciCreate",
			args:args{ value:sc, tmpl: "./template/createtable.tmpl" },
			want: HedefKullanicitable,
		},
		{
			name: "KullaniciStruct",
			args:args{ value:sc, tmpl: "./template/struct.tmpl" },
			want: HedefKullanicistruct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TemplateExecute(tt.args.value, tt.args.tmpl); got != tt.want {
				t.Errorf("SelectTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
