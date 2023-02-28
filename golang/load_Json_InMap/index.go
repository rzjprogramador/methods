package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Foo struct{
	Campo1 string
	Campo2 string
}

func main() {
	var fileJsonPath = "/home/rzj/..backend/golang/funcoesUteis/load_Json_InMap/arquivoJson.json"

	var fooVazio Foo
	fmt.Println(loadJson_InMap_v1(fileJsonPath, fooVazio))
}


// Esta funcao recebe um arquivoJson, e uma variavel de struct vazia
func loadJson_InMap_v1(f string, s interface{}) interface{} { 
	jsonFile, err := os.Open(f)
	if err != nil {
			fmt.Println("Ops! Nao foi encontrado o arquivo json.", err)
	} else {
		fmt.Println("Ok Arquivo Json Carregado com Sucesso!")
	}
	defer jsonFile.Close()

	myjson, _ := ioutil.ReadAll(jsonFile)
	
	if erro :=  json.Unmarshal([]byte(myjson), &s); erro != nil {
		log.Fatal(erro)
	}

	return s
	
}

	/*
	
	funcao UnMarshal() transforma o json em objStruct - recebe o objJsonalvo em slideDeByte, e a variavel EmMemoria que sera peenchida pelo struct ela tem que ser um structVazio
	por ultimo retorno o objetoStruct que passou pelo UnMarshal
	// obs: esta retornando o json em map

	// TODO: Tentar Retornar um struct a ser passado no Uso
	
	---
	anotacoes 
	// arrBytes, _ := json.Marshal(myjson) // devolveu bytes[]
	// objJSON := bytes.NewBuffer(arrBytes) // devolve o bytes[] em JSON

	// CRIANDO NOVO STRUCT E PREENCHENDO COM O JSON CAPTURADO DO ARQUIVO
	// var s Foo // objStruct vazio que sera preenchido
	*/