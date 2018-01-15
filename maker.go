package main

import (
	"strings"
	"bytes"
	"os"
	"io/ioutil"
	"log"
	"os/exec"
	"text/template"
	"encoding/json"
)

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

func DataOku(files []os.FileInfo) []Sinif {
	dat, _ := ioutil.ReadFile("./kaynak/jsondata.json")
	var projeler  []Proje
	_ = json.Unmarshal(dat, &projeler)

	projeler[0].ProjeYolu="C:\\Users\\Fatih\\go\\src\\otoprj"

	b, _ := json.Marshal(projeler)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	ioutil.WriteFile("./kaynak/jsondata2.json",out.Bytes(),0644)

	return projeler[0].Siniflar
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

	dataArray := DataOku(files)

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
			HedefFile :=  strings.ToLower(data.SinifAdi) + "ler.html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/tablo.tmpl"), "tablo.tmpl")

			HedefFile =  strings.ToLower(data.SinifAdi) + ".html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/form.tmpl"), "form.tmpl")

			HedefFile =  strings.ToLower(data.SinifAdi) + "Field_oto.html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/formField.tmpl"), "formField.tmpl")

			HedefFile =  strings.ToLower(data.SinifAdi) + "lerField_oto.html"
			HedefeKaydet(data, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/tabloField.tmpl"), "tabloField.tmpl")
	}


	exec.Command("bash", "-c", "go fmt /home/fatih/gowork/src/otoprj/*.go").Run()

}

