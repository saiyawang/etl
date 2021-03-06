package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/saiyawang/etl/util"
)

func renderHandler(w http.ResponseWriter, r *http.Request, templatepath string) {
	t, err := template.ParseFiles(templatepath)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	renderHandler(w, r, "template/index.html")
}

func viewWorldHandler(w http.ResponseWriter, r *http.Request) {

	result := getDeviceInContinent()

	fmt.Fprint(w, result)
}

func viewUSHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	devicetype := r.PostFormValue("devicetype")

	result := getDeviceInCountry(devicetype)

	fmt.Fprint(w, result)
}

func viewProvinceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	province := r.PostFormValue("province")
	devicetype := r.PostFormValue("devicetype")

	result := GetDeviceInProvince(strings.TrimPrefix(province, "US-"), devicetype)

	fmt.Fprint(w, result)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r.URL.Path[1:])

	http.ServeFile(w, r, r.URL.Path[1:])
}

func getDeviceInContinent() string {

	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.GetDeviceInContinent()
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func getDeviceInCountry(devicetype string) string {
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.GetDeviceInCountry(devicetype)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func GetDeviceInProvince(province string, devicetype string) string {

	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.GetDeviceInProvince(province, devicetype)
	if err != nil {
		log.Fatal(err)
	}

	//	log.Println(result)

	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func main() {
	http.HandleFunc("/template/", staticHandler)
	http.HandleFunc("/", viewHandler)

	http.HandleFunc("/world", viewWorldHandler)
	http.HandleFunc("/us", viewUSHandler)
	http.HandleFunc("/province", viewProvinceHandler)

	http.ListenAndServe(":137", nil)

}
