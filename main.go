package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"io/ioutil"
	"log"
	"os"
)

type TmplData struct {
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




func KaynakBilgiToTmplData(kaynakBilgi []KaynakBilgi) TmplData {
	r:=TmplData{}

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


func TemplateExecute(value string, tmpl string) string {

	dbb:= KaynakBilgiToTmplData(ParsKaynakBilgi(value));

	t := template.Must(template.ParseFiles(tmpl))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, dbb)
	if err != nil {
		panic(err)
	}
	return tpl.String()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	files, err := ioutil.ReadDir("./kaynak")
	if err != nil {
		log.Fatal(err)
	}

	fhedef, err := os.Create("./hedef/entity.txt")
	check(err)
	defer fhedef.Close()

	_, err = fhedef.WriteString(`package entity`+"\n\n")
	check(err)
	_, err = fhedef.WriteString(`import "time"`+"\n\n")
	check(err)


	for _, f := range files {
		dat, err := ioutil.ReadFile("./kaynak/"+f.Name())
		check(err)
		fmt.Print()
		s:=TemplateExecute(string(dat),"./template/struct.tmpl")
		fmt.Println(s)
		_, err = fhedef.WriteString(s+"\n\n")
		check(err)
	}

	fhedef.Sync()

	//--------

	for _, f := range files {
		dat, err := ioutil.ReadFile("./kaynak/"+f.Name())
		check(err)

		fhedef2, err := os.Create("./hedef/"+f.Name())
		check(err)
		defer fhedef2.Close()

		s:=TemplateExecute(string(dat),"./template/selectsql.tmpl")
		fmt.Println(s)
		_, err = fhedef2.WriteString(s+"\n\n")
		check(err)
		fhedef2.Sync()


	}




}
