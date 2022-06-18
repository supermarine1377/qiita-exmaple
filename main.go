package main

import (
	"fmt"
	"net/http"
	"qiita/config"
)

func main() {
	var file, err = config.GetConfigFile()
	if err != nil {
		panic(err)
	}
	conf, err := config.GetConfig(file)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/qiita", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, world!")
	})
	http.ListenAndServe(":"+conf.Port, nil)
}
