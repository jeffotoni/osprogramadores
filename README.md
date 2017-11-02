# osprogramadores
Grupo os Programdores, desafios em programação


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