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



func WriteString(fhedef *os.File, s string) {
	_, err := fhedef.WriteString(s)
	check(err)
}

func HedefeKaydet(value []TmplData, hedefFile string, TemplateFile string) {
	fhedef, err := os.Create(hedefFile)
	check(err)
	defer fhedef.Close()
	s := TemplateExecuteArray(value, TemplateFile)
	WriteString(fhedef, s)
	fhedef.Sync()
}

var hedefklasor = "/home/fatih/gowork/src/otoprj"
func main() {

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
			HedefeKaydet(dataArray, (hedefklasor + "/" + HedefFile), ("./template/" + f.Name()))
		}
	}


	templatefiles, err = ioutil.ReadDir("./template/templates")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range templatefiles {
		HedefFile := strings.Replace(f.Name(), ".", "_", 1) + ".tmpl"
		HedefeKaydet(dataArray, (hedefklasor+"/templates/"+HedefFile), ("./template/templates/" +f.Name()))
	}




	exec.Command("bash", "-c", "go fmt /home/fatih/gowork/src/otoprj/*.go").Run()
}


