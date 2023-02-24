package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type Modelo struct {
	Campo1   string `json:"campo1"`
	Campo2   string `json:"campo2"`
}

var arquivoCopyPath = "/home/rzj/..git/methods/golang/ler_arquivo_json/data_json/data.golang.json"

func UseLoadJsonV1(f string) (*bytes.Buffer){
    jsonFile, err := os.Open(f)
    if err != nil {
        fmt.Println("Ops! Nao foi encontrado o arquivo json.", err)
    } else {
			fmt.Println("Ok Arquivo Json Carregado com Sucesso!")
		}
    defer jsonFile.Close()

    jsonForSliceBytes, _ := ioutil.ReadAll(jsonFile)

		objectModel := bytes.NewBuffer(jsonForSliceBytes)
		return objectModel
}

func main() {
    fmt.Println(UseLoadJsonV1(arquivoCopyPath))
}
