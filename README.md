# osprogramadores

Grupo os Programdores, desafios em programação, é um site destinado a disseminar o conhecimento de programação e assuntos ligados aos computadores em geral.

Existe alguns desafios interessantes em programção em:
- https://osprogramadores.com/desafios


# link para acessar o desafio 5 no site

 - https://osprogramadores.com/desafios/d05/

O desafio 5 é para desenvolver um programa que leia um arquivo no formato json e consiga abstrair as seguintes respostas:

 - Quem mais recebe e quem menos recebe na empresa e a média salarial da empresa

 - Quem mais recebe e quem menos recebe em cada área e a média salarial em cada área
 
 - A área com mais funcionários e a área com menos funcionários

 - Das pessoas que têm o mesmo sobrenome, aquela que recebe mais (não inclua sobrenomes que apenas uma pessoa tem nos resultados)


# Codigo

```go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//
//
//
type Corporacao struct {
	Funcionarios []struct {
		ID        int    `json:"id"`
		Nome      string `json:"nome"`
		Sobrenome string `json:"sobrenome"`
		Salario   int    `json:"salario"`
		Area      string `json:"area"`
	} `json:"funcionarios"`

	Areas []struct {
		Codigo string `json:"codigo"`
		Nome   string `json:"nome"`
	} `json:"areas"`
}

//
//
//
func getCorporacao() Corporacao {

	raw, err := ioutil.ReadFile("./funcionarios.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	var C Corporacao
	json.Unmarshal(raw, &C)
	return C
}

```
