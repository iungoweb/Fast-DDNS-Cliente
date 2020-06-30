package main

import (
	"encoding/json"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Response struct {
	Status bool
	Msg string
}

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	trataErro(err, "Erro ao recuperar local de execução")

	println(dir + "/config.ini")

	cfg, err := ini.Load(dir + "/config.ini")
	trataErro(err, "Erro ao ler arquivo de configuração")

	var url string
	url  = "http://ddns.lo/v1/"
	url += cfg.Section("usuario").Key("dominio").String()
	url += "/"
	url += cfg.Section("usuario").Key("key").String()

	//println(url)

	for {

		agora := time.Now()

		println("[" + agora.Format("02/01/2006 15:04:05") + "] Fazendo requisição")

		resp, err := http.Get(url)
		trataErro(err, "Erro ao fazer requisição")

		println("[" + agora.Format("02/01/2006 15:04:05") + "] Requisição concluída")

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var response Response
		_ = json.Unmarshal([]byte(body), &response)

		println("[" + agora.Format("02/01/2006 15:04:05") + "] Retorno: " + response.Msg)

		println("[" + agora.Format("02/01/2006 15:04:05") + "] Aguardando para fazer a próxima requisição...")
		println("")
		println("")

		time.Sleep(1 * time.Minute)
	}
}

func trataErro(err error, msg string) {

	if err != nil {
		log.Fatal(msg)
	}
}