package main

import (
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//пишем чтобы делать запросы из браузера
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//получаем данные с фронта
		data := r.URL.Query().Get("data")

		//кароче не записывается get параметр потому что как то оно прогоняется два раза
		//первый раз он видит get параметр и создает файл
		//и второй раз get парметр - ПУСТОЙ и файл создается с пустотой (если отправлять руками из браузера)
		//поэтому просто добавим проверку парметра на пустоту
		if data != "" {

			//открываем файл для работы с ним
			file, _ := os.Create("data.json")

			defer file.Close()

			//записываем в файл данные
			file.WriteString(data)

		}

	})
	http.ListenAndServe(":8081", nil)

}
