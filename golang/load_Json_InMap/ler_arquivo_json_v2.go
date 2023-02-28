package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// #TODO Fazer esta funcao voltar um arquivo de um determinado Modelo struct

func UseLoadJsonV2(f string) interface{}{
    jsonFile, err := os.Open(f)
    if err != nil {
        fmt.Println("Ops! Nao foi encontrado o arquivo json.", err)
    } else {
			fmt.Println("Ok Arquivo Json Carregado com Sucesso!")
		}
    defer jsonFile.Close()

		myjson, _ := ioutil.ReadAll(jsonFile)
		obj, _ := json.Marshal(myjson)
		res := obj
		return res
}
