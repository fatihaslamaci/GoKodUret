package main

import (
	"strings"
	"bytes"
	"os"
	"io/ioutil"
	"log"
	"os/exec"
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


func TemplateExecuteArray(data interface{}, tmpl string,TemplateName string) string {
	funcMap := template.FuncMap{
		"ToLover": strings.ToLower,
	}

	t := template.Must(template.New(TemplateName).Funcs(funcMap).ParseFiles(tmpl))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, data)
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



func WriteString(fhedef *os.File, s string) {
	_, err := fhedef.WriteString(s)
	check(err)
}

func HedefeKaydet(data interface{}, hedefFile string, TemplateFile string,TemplateName string) {
	fhedef, err := os.Create(hedefFile)
	check(err)
	defer fhedef.Close()
	s := TemplateExecuteArray(data, TemplateFile,TemplateName)
	WriteString(fhedef, s)
	fhedef.Sync()
}




var hedefklasor = "C:\\Users\\Fatih\\go\\src\\otoprj"


func Makeproje(){
	os.MkdirAll(hedefklasor, os.ModePerm)
	os.MkdirAll(hedefklasor+"/templates", os.ModePerm)


	files, err := ioutil.ReadDir("./kaynak")
	if err != nil {
		log.Fatal(err)
	}

	dataArray := FileToDataArray(files)

	templatefiles, err := ioutil.ReadDir("./template")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range templatefiles {
		if f.IsDir()==false {

			HedefFile := strings.Replace(f.Name(), ".", "_", 1) + ".go"
			HedefeKaydet(dataArray, (hedefklasor + "/" + HedefFile), ("./template/" + f.Name()),f.Name())
		}
	}

	for _, data := range dataArray{
			HedefFile :=  strings.ToLower(data.StructName) + "ler.html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/tablo.tmpl"), "tablo.tmpl")

			HedefFile =  strings.ToLower(data.StructName) + ".html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/form.tmpl"), "form.tmpl")

			HedefFile =  strings.ToLower(data.StructName) + "Field_oto.html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/formField.tmpl"), "formField.tmpl")

			HedefFile =  strings.ToLower(data.StructName) + "lerField_oto.html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/tabloField.tmpl"), "tabloField.tmpl")
	}


	exec.Command("bash", "-c", "go fmt /home/fatih/gowork/src/otoprj/*.go").Run()

}

