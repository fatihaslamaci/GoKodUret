package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type TmplData struct {
	StructName  string
	TableName   string
	KaynakBilgi []KaynakBilgi
}

type KaynakBilgi struct {
	GoName string
	GoTip  string
	DbName string
	DbTip  string
}

func KaynakBilgiToTmplData(kaynakBilgi []KaynakBilgi) TmplData {
	r := TmplData{}

	for index, element := range kaynakBilgi {
		if index == 0 {
			r.StructName = element.GoName
			r.TableName = element.DbName
		} else {
			r.KaynakBilgi = append(r.KaynakBilgi, element)
		}
	}
	return r
}

func ParsKaynakBilgi(value string) []KaynakBilgi {
	r := []KaynakBilgi{}
	line := strings.Split(strings.TrimSpace(value), "\n")
	for _, element := range line {
		if element != "" {
			if strings.Index(element, ":") > 0 {
				kb := KaynakBilgi{}
				linef := strings.Split(element, ":")
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

	dbb := KaynakBilgiToTmplData(ParsKaynakBilgi(value))

	t := template.Must(template.ParseFiles(tmpl))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, dbb)
	if err != nil {
		panic(err)
	}
	return tpl.String()

}

func TemplateExecuteArray(value []TmplData, tmpl string) string {
	t := template.Must(template.ParseFiles(tmpl))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, value)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}

func FileToDataArray(files []os.FileInfo) []TmplData {
	r := []TmplData{}
	for _, f := range files {
		dat, err := ioutil.ReadFile(("./kaynak/" + f.Name()))
		check(err)
		dbb := KaynakBilgiToTmplData(ParsKaynakBilgi(string(dat)))
		r = append(r, dbb)
	}
	return r
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var hedefklasor = "/home/fatih/gowork/src/otoprj"

func main() {

	os.MkdirAll(hedefklasor, os.ModePerm)

	files, err := ioutil.ReadDir("./kaynak")
	if err != nil {
		log.Fatal(err)
	}

	dataArray := FileToDataArray(files)

	MainOlustur(dataArray)
	InitDBOlustur(dataArray)
	StructOlustur(dataArray)
	CrudOlustur(dataArray)

	exec.Command("bash", "-c", "go fmt /home/fatih/gowork/src/otoprj/*.go").Run()

}
func WriteString(fhedef *os.File, s string) {
	_, err := fhedef.WriteString(s)
	check(err)
}

func MainOlustur(value []TmplData) {
	fhedef, err := os.Create((hedefklasor + "/" + "main.go"))
	check(err)
	defer fhedef.Close()
	s := TemplateExecuteArray(value, "./template/main.tmpl")
	WriteString(fhedef, s)
	fhedef.Sync()
}

func InitDBOlustur(value []TmplData) {
	fhedef, err := os.Create((hedefklasor + "/" + "InitDB.go"))
	check(err)
	defer fhedef.Close()
	s := TemplateExecuteArray(value, "./template/InitDB.tmpl")
	WriteString(fhedef, s)
	fhedef.Sync()
}

func StructOlustur(value []TmplData) {
	fhedef, err := os.Create(hedefklasor + "/entity.go")
	check(err)
	defer fhedef.Close()
	s := TemplateExecuteArray(value, "./template/struct.tmpl")
	WriteString(fhedef, s+"\n\n")
	fhedef.Sync()
}

func CrudOlustur(value []TmplData) {
	fhedef, err := os.Create(hedefklasor + "/crud.go")
	check(err)
	defer fhedef.Close()
	s := TemplateExecuteArray(value, "./template/crud.tmpl")
	WriteString(fhedef, s+"\n\n")
	fhedef.Sync()
}
