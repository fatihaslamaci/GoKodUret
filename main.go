package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type DbBilgi struct {
	StructName string
	TableName  string
	KaynakBilgi []KaynakBilgi
}

type KaynakBilgi struct {
	GoName string
	GoTip  string
	DbName string
	DbTip  string
}

var sc = `
Kullanici           :           :kullanicilar		:
Id                  :int64		:id 				:INTEGER primary key autoincrement
Ad                  :string	 	:ad 				:VARCHAR(50)
KayitTarihi         :time.Time  :kayittarihi 		:DATE
HataliGirisSayisi   :int        :hataligirissayisi	:INTEGER
BlokeHesap			:bool       :blokehesap 		:BIT
`

var HedefKullanicistruct = `
type Kullanici struct {
Id int64
Ad string
KayitTarihi time.Time
HataliGirisSayisi int
BlokeHesap bool

}`

var KullanicistructTmp = `
type {{.StructName}} struct {
{{ range $i, $e := .KaynakBilgi }}{{$e.GoName}} {{$e.GoTip}}
{{ end }}
}`



var TableTmp = `
CREATE TABLE IF NOT EXISTS {{.TableName}}(
{{ range $i, $e := .KaynakBilgi }}{{ if eq $i 0 }}{{$e.DbName}} {{$e.DbTip}}{{ else }},{{$e.DbName}} {{$e.DbTip}}{{ end }}
{{ end }}
);`



var HedefKullanicitable = `
CREATE TABLE IF NOT EXISTS kullanicilar(
id INTEGER primary key autoincrement
,ad VARCHAR(50)
,kayittarihi DATE
,hataligirissayisi INTEGER
,blokehesap BIT

);`



var SelectIdTmp = `
func {{.StructName}}Select(db *sql.DB, id int) {{.StructName}} {
	item := {{.StructName}}{}
	if id > 0 {
		row := db.QueryRow("Select {{ range $i, $e := .KaynakBilgi }}{{ if eq $i 0 }}{{$e.DbName}}{{ else }}, {{$e.DbName}}{{ end }}{{ end }} from {{.TableName}} where id=?", id)
		err := row.Scan({{ range $i, $e := .KaynakBilgi }}{{ if eq $i 0 }}&item.{{$e.GoName}}{{ else }}, &item.{{$e.GoName}}{{ end }}{{ end }})
		CheckErr(err)
	}
	return item
}`

var HedefSelectIdTmp = `
func KullaniciSelect(db *sql.DB, id int) Kullanici {
	item := Kullanici{}
	if id > 0 {
		row := db.QueryRow("Select id, ad, kayittarihi, hataligirissayisi, blokehesap from kullanicilar where id=?", id)
		err := row.Scan(&item.Id, &item.Ad, &item.KayitTarihi, &item.HataliGirisSayisi, &item.BlokeHesap)
		CheckErr(err)
	}
	return item
}`



func KaynakBilgiToDBBilgi (kaynakBilgi []KaynakBilgi) DbBilgi {
	r:=DbBilgi{}

	for index, element := range kaynakBilgi {
		if index==0{
			r.StructName=element.GoName
			r.TableName = element.DbName
		}else{
			r.KaynakBilgi=append(r.KaynakBilgi,element)
		}
	}
	return r
}



func ParsKaynakBilgi(value string) []KaynakBilgi {
	r := []KaynakBilgi{}
	line := strings.Split(strings.TrimSpace(value) ,"\n")
	for _, element := range line {
		if element!="" {
			if strings.Index(element,":")>0{
				kb:=KaynakBilgi{}
				linef:=strings.Split(element ,":")
				kb.GoName = strings.TrimSpace(linef[0])
				kb.GoTip = strings.TrimSpace(linef[1])
				kb.DbName = strings.TrimSpace(linef[2])
				kb.DbTip = strings.TrimSpace(linef[3])
				r = append(r, kb)
			}
		}
	}
	return r
}


func SelectTable(value string, tmpl string) string {

	dbb:=KaynakBilgiToDBBilgi(ParsKaynakBilgi(value));

	t := template.Must(template.New("selecttable").Parse(tmpl))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, dbb)
	if err != nil {
		panic(err)
	}
	return tpl.String()

}

func main() {

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
			args:args{ value:sc, tmpl: SelectIdTmp },
			want: HedefSelectIdTmp,
		},
		{
			name: "KullaniciCreate",
			args:args{ value:sc, tmpl: TableTmp },
			want: HedefKullanicitable,
		},
		{
			name: "KullaniciStruct",
			args:args{ value:sc, tmpl: KullanicistructTmp },
			want: HedefKullanicistruct,
		},
	}
	for _, tt := range tests {
			 fmt.Println(SelectTable(tt.args.value, tt.args.tmpl))
		     fmt.Println("-----")
	}
}
