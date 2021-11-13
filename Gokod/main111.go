package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//пишем чтобы делать запросы из браузера
		w.Header().Set("Access-Control-Allow-Origin", "*")

		//считываю данные из файла как срез байт
		dataFromFile, _ := ioutil.ReadFile("data.json")

		//перводим срез байт в строку и выводим
		fmt.Fprintf(w, string(dataFromFile))
		return
	})
	http.ListenAndServe(":8080", nil)

}
