package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func atributeMake(anahtar string, deger string) string {
	return anahtar + `="` + deger + `" `
}

func MyAtribute(alan Alan) string {
	r := ""
	if (alan.AlanAdi != "") || (alan.DbAlanAdi != "") {

		if alan.HtmlInputType != "" {
			if alan.Requered {
				r += " required "
			}
			if alan.MinLength != "" {
				r += atributeMake("minlength", alan.MinLength)
			}
			if alan.MaxLength != "" {
				r += atributeMake("maxlength", alan.MaxLength)
			}
			if alan.MinValue != "" {
				r += atributeMake("min", alan.MinValue)
			}
			if alan.MaxValue != "" {
				r += atributeMake("max", alan.MaxValue)
			}

		}
	}
	return r
}

func TemplateExecuteArray(data interface{}, tmpl string, TemplateName string) string {
	funcMap := template.FuncMap{
		"ToLover":    strings.ToLower,
		"MyAtribute": MyAtribute,
	}

	t := template.Must(template.New(TemplateName).Funcs(funcMap).ParseFiles(tmpl))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, data)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}

func JsonDataOku() []Proje {
	dat, _ := ioutil.ReadFile("./kaynak/jsondata.json")
	var projeler []Proje
	_ = json.Unmarshal(dat, &projeler)

	projeler[0].ProjeYolu = "C:\\Users\\Fatih\\go\\src\\otoprj"

	b, _ := json.Marshal(projeler)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	ioutil.WriteFile("./kaynak/jsondata2.json", out.Bytes(), 0644)

	return projeler
}

func ProjelerJsonYedekKaydet(db *sql.DB) {

	projeler := ProjeSelectAll(db)
	for i, _ := range projeler {
		projeler[i] = DataOku2(db, projeler[i].Id)
	}

	b, _ := json.Marshal(projeler)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	ioutil.WriteFile("./kaynak/yedek2018_04_16.json", out.Bytes(), 0644)

}

func DataOku2(db *sql.DB, id int64) Proje {
	proje := ProjeSelect(db, id)
	proje.Siniflar = SinifSelectMasterId(db, id)
	for i, _ := range proje.Siniflar {
		proje.Siniflar[i].Alanlar = AlanSelectMasterId(db, "order by sira_no", proje.Siniflar[i].Id)
		proje.Siniflar[i].TabloEkOzellikler = TabloEkOzellikSelectMasterId(db, proje.Siniflar[i].Id)
		for j, _ := range proje.Siniflar[i].Alanlar {
			proje.Siniflar[i].Alanlar[j].AnahtarDegerler = AnahtarDegerSelectMasterId(db, proje.Siniflar[i].Alanlar[j].Id)
		}
	}

	return proje
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteString(fhedef *os.File, s string) {

	//Boş satırları silmek için regex
	regex, err := regexp.Compile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	check(err)
	s2 := regex.ReplaceAllString(s, "")
	_, err = fhedef.WriteString(s2)
	check(err)
}

func HedefeKaydet(data interface{}, hedefFile string, TemplateFile string, TemplateName string) {
	fhedef, err := os.Create(hedefFile)
	check(err)
	defer fhedef.Close()
	s := TemplateExecuteArray(data, TemplateFile, TemplateName)

	WriteString(fhedef, s)
	fhedef.Sync()
}

func DosyaKopyala(kaynak string, hedef string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(kaynak)
	check(err)
	// Write data to dst
	err = ioutil.WriteFile(hedef, data, 0644)
	check(err)
}

func HedefdeDosyaYokIseKlasorKopyala(kaynakKlasor string, hedefKlasor string, recursively bool) {

	os.MkdirAll(hedefKlasor, os.ModePerm)
	fmt.Println(kaynakKlasor)
	fmt.Println(hedefKlasor)

	files, err := ioutil.ReadDir(kaynakKlasor)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() == false {
			HedefFile := f.Name()
			if HedefdeDosyaYokIse((hedefKlasor + "/" + HedefFile)) {
				DosyaKopyala((kaynakKlasor + "/" + HedefFile), (hedefKlasor + "/" + HedefFile))
			}
		} else if recursively == true {
			HedefFile := f.Name()
			HedefdeDosyaYokIseKlasorKopyala((kaynakKlasor + "/" + HedefFile), (hedefKlasor + "/" + HedefFile), all)
		}
	}
}

func Makeproje(db *sql.DB, id int64) {

	proje := DataOku2(db, id)
	//proje :=JsonDataOku()[0]

	dataArray := proje.Siniflar

	hedefklasor := proje.ProjeYolu

	os.MkdirAll(hedefklasor, os.ModePerm)
	os.MkdirAll(hedefklasor+"/templates", os.ModePerm)

	HedefdeDosyaYokIseKlasorKopyala("./statics", (hedefklasor + "/statics"), true)

	//dataArray := DataOku()[0].Siniflar

	templatefiles, err := ioutil.ReadDir("./template")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range templatefiles {
		if f.IsDir() == false {
			HedefFile := strings.Replace(f.Name(), ".tmpl", ".go", 1)
			if HedefdeDosyaYokVeyaDosyaAdiOtoIse((hedefklasor + "/" + HedefFile)) {
				HedefeKaydet(dataArray, (hedefklasor + "/" + HedefFile), ("./template/" + f.Name()), f.Name())
			}
		}
	}

	HedefdeDosyaYokIseKlasorKopyala(("./template/direk_kopyalanacaklar"), (hedefklasor + "/templates"), false)

	for _, data := range dataArray {
		HedefFile := strings.ToLower(data.SinifAdi) + "ler.html"
		if HedefdeDosyaYokIse((hedefklasor + "/templates/" + HedefFile)) {
			HedefeKaydet(data, (hedefklasor + "/templates/" + HedefFile), ("./template/templates/tablo.tmpl"), "tablo.tmpl")
		}

		HedefFile = strings.ToLower(data.SinifAdi) + ".html"
		if HedefdeDosyaYokIse((hedefklasor + "/templates/" + HedefFile)) {
			HedefeKaydet(data, (hedefklasor + "/templates/" + HedefFile), ("./template/templates/form.tmpl"), "form.tmpl")
		}

		HedefFile = strings.ToLower(data.SinifAdi) + "Field_oto.html"
		HedefeKaydet(data, (hedefklasor + "/templates/" + HedefFile), ("./template/templates/formField.tmpl"), "formField.tmpl")

		HedefFile = strings.ToLower(data.SinifAdi) + "lerField_oto.html"
		HedefeKaydet(data, (hedefklasor + "/templates/" + HedefFile), ("./template/templates/tabloField.tmpl"), "tabloField.tmpl")
	}

	c := exec.Command("bash", "-c", "go fmt "+hedefklasor+"/*.go")
	//c :=exec.Command("cmd", "/C", "gofmt -w", hedefklasor)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

}
func HedefdeDosyaYokVeyaDosyaAdiOtoIse(HedefFile string) bool {
	return (Exists(HedefFile) == false) || (strings.Index(HedefFile, "_oto.") >= 0)
}

func HedefdeDosyaYokIse(HedefFile string) bool {
	return (Exists(HedefFile) == false)
}
