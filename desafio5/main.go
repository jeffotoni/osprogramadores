/**
 *
 * @author      Jefferson Otoni Lima
 * @copyright   Copyright (c) 2017 - 2017
 * @since       Version 1.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
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

	t1 := time.Now()

	fmt.Println("")
	fmt.Println("########## inicio abrindo arquivo json #########")
	fmt.Println("time inicio: [" + fmt.Sprintf("%s", t1.Format("2006-01-02 15:04:05")) + "]")
	fmt.Println("")

	// colocando na memoria todo arquivo 1.1G
	// raw, err := ioutil.ReadFile("./files-json/funcionarios-14M.json")
	raw, err := ioutil.ReadFile("./funcionarios-2k.json")

	//
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	var C Corporacao
	json.Unmarshal(raw, &C)

	t2 := time.Now()

	timeSpent := t2.Sub(t1)

	fmt.Println("")
	fmt.Println("########## fim, json na memoria time fim #########")
	fmt.Println("time fim: [" + fmt.Sprintf("%s", timeSpent) + "]")
	fmt.Println("")

	return C
}

// Quem mais recebe e quem menos recebe na empresa e a média salarial da empresa
// Quem mais recebe e quem menos recebe em cada área e a média salarial em cada área
// A área com mais funcionários e a área com menos funcionários
// Das pessoas que têm o mesmo sobrenome, aquela que recebe mais (não inclua sobrenomes que apenas uma pessoa tem nos resultados)

func main() {

	var debugtime bool

	// definindo o ambiente
	// para executar somente
	// sem apresentar os resultados
	debugtime = true

	t1 := time.Now()

	if debugtime {

		fmt.Println("")
		fmt.Println("########## inicio da execucao code #########")
		fmt.Println("time inicio: [" + fmt.Sprintf("%s", t1.Format("2006-01-02 15:04:05")) + "]")
		fmt.Println("")
	}

	// Struct em Json
	C := getCorporacao()

	// Json Funcionarios
	funcionarios := C.Funcionarios

	// maior salario
	smaior := 0

	// menor salario
	smenor := 1000000

	//soma dos salarios
	ssoma := 0

	// quantidade de funcionarios
	// em toda corporacao
	qfunc := len(funcionarios)

	// mapa de funcionarios
	// maior salario
	Mafunc := make(map[int]string)

	// mapa de funcionarios
	// menor salario
	Mefunc := make(map[int]string)

	// Areas
	MapAreas := make(map[string][]string)

	// funcionarios que possui maior salario por area
	MapAreaSalFuncMaior := make(map[string][]string)

	// funcionarios que possui menor salario por area
	MapAreaSalFuncMenor := make(map[string][]string)

	// vetor de area salario maior
	AreaSMa := make(map[string]int)

	// vetor de area salario menor
	AreaSMe := make(map[string]int)

	// vetor de area salario medio
	AreaSSm := make(map[string]int)

	// vetor de area maior quantidade
	AreaQMa := make(map[string]string)

	// vetor de nomes, sobrenomes
	VNomeSob := make(map[string]string)

	// vetor de nome sobrenome e salrio
	// somente os que repetem
	NomeSobreSal := make(map[string]string)

	// dataJson responsavel por
	// Unmarshal em um string json
	var datJson map[string]interface{}

	// Quem mais recebe e quem menos
	// recebe na empresa e a média salarial da empresa
	for _, Func := range funcionarios {

		if Func.Salario > smaior { // recebe mais

			smaior = Func.Salario

		} else if Func.Salario < smenor { //recebe menos

			smenor = Func.Salario
		}

		// media salarial
		ssoma += Func.Salario

		// amarzenando area e seus funcionarios
		// agrupando area => funcionarios
		// [area] => [{funcionarios}]
		MapAreas[Func.Area] = append(MapAreas[Func.Area], `{"area":"`+Func.Area+`","nome":"`+Func.Nome+`","salario":"`+fmt.Sprintf("%d", Func.Salario)+`"}`)

		// Vetor de nome e sobrenome de
		// toda corporacao
		VNomeSob[Func.Nome] = Func.Sobrenome
	}

	// buscando os nomes de maior e menor salario
	// da corporacao
	for i, Func := range funcionarios {

		// montando uma lista de
		// nome que possuem o maior salario
		if Func.Salario == smaior {

			Mafunc[i] = Func.Nome + " " + Func.Sobrenome
		}

		// montando uma lista de
		// nome que possuem o menor salario
		if Func.Salario == smenor {

			Mefunc[i] = Func.Nome + " " + Func.Sobrenome
		}

		// buscando e montando vetor dos sobrenomes que repetem
		// os que repetem monta uma lista com nome sobrenome
		// e seu respectivo salario
		if ExistSobreNome(VNomeSob, Func.Nome, Func.Sobrenome) {

			NomeSobreSal[Func.Nome+" "+Func.Sobrenome] = fmt.Sprintf("%d", Func.Salario)
			//fmt.Println("Sobrenome Igual: ", Func.Nome, " ::: ", Func.Sobrenome)
		}
	}

	// quantidade
	// de colaboradores
	// nas areas
	qtmp := 0

	// quantidade
	// de colaboradores
	// nas areas
	qtmpe := 100000

	// Salarios por areas, medias
	for S, M := range MapAreas {

		quantFuncArea := len(M)

		if quantFuncArea > qtmp {

			qtmp = quantFuncArea
			AreaQMa["maior"] = `{"area":"` + S + `,"quantidade":"` + fmt.Sprintf("%d", qtmp) + `"}`
		}

		if quantFuncArea < qtmpe {

			qtmpe = quantFuncArea
			AreaQMa["menor"] = `{"area":"` + S + `,"quantidade":"` + fmt.Sprintf("%d", qtmpe) + `"}`
		}

		stmp := 0
		stmp2 := smaior + 100

		// menor e maior salario por area
		// media por area
		for _, V := range M {

			json.Unmarshal([]byte(V), &datJson)

			// assertions para converter
			s, _ := datJson["salario"].(string)

			// string to int
			st, _ := strconv.Atoi(s)

			if st > stmp {

				stmp = st
				AreaSMa[S] = st
			}

			if st < stmp2 {

				stmp2 = st
				AreaSMe[S] = st
			}

			// Media do salario
			AreaSSm[S] += st
		}

		AreaSSm[S] = AreaSSm[S] / len(M)
	}

	// varrendo e montando quem tem os menores salarios e os maiores por setor
	for S, M := range MapAreas {

		// menor e maior salario por area
		// media por area
		for _, V := range M {

			json.Unmarshal([]byte(V), &datJson)

			// assertions para converter
			s, _ := datJson["salario"].(string)
			nome, _ := datJson["nome"].(string)

			// string to int
			st, _ := strconv.Atoi(s)

			// buscar maior salario da regiao
			sr := AreaSMa[S]
			if st == sr {

				MapAreaSalFuncMaior[S] = append(MapAreaSalFuncMaior[S], `{"nome":"`+nome+`"}`)
			}

			// buscar maior salario da regiao
			sm := AreaSMe[S]
			if st == sm {

				MapAreaSalFuncMenor[S] = append(MapAreaSalFuncMenor[S], `{"nome":"`+nome+`"}`)
			}
		}
	}

	slftmp := 0
	nomeRMS := ""

	// saber que ganha mais dos funcionarios com mesmo sobrenome
	for NS, Sl := range NomeSobreSal {

		slf, _ := strconv.Atoi(Sl)

		if slf > slftmp {

			slftmp = slf
			nomeRMS = NS
		}
	}

	smedia := (ssoma / qfunc)

	if debugtime {

		t2 := time.Now()

		timeSpent := t2.Sub(t1)

		fmt.Println("")
		fmt.Println("########## fim, json na memoria time fim #########")
		fmt.Println("time fim: [" + fmt.Sprintf("%s", timeSpent) + "]")
		fmt.Println("")

	} else {

		fmt.Println("")
		fmt.Println("################ listando na tela o resultado #################")
		fmt.Println("")

		//// Quem mais recebe e quem menos recebe na empresa e a média salarial da empresa
		fmt.Println("Maior salario da Corporacao: ", smaior)
		fmt.Println("")
		fmt.Println("Menor salario da Corporacao: ", smenor)
		fmt.Println("")
		fmt.Println("Media de salario da Corporacao: ", smedia)
		fmt.Println("")
		fmt.Println("Funcionario(s) que mais recebe na Corporacao")

		for _, Nome := range Mafunc {

			fmt.Println("Nome: ", Nome)
		}

		fmt.Println("")
		fmt.Println("Funcionario(s) que menor recebe na Corporacao")

		for _, Nome := range Mefunc {

			fmt.Println("Nome: ", Nome)
		}

		// os que mais recebem por area e os que menos recebem

		fmt.Println("")
		fmt.Println("maior salario das areas: ")

		for A, S := range AreaSMa {

			fmt.Println("Area: ", A, " Salario: ", S)
		}

		fmt.Println("")
		fmt.Println("menor salario das areas: ", AreaSMe)

		for A, S := range AreaSMe {

			fmt.Println("Area: ", A, " Salario: ", S)
		}

		fmt.Println("")
		fmt.Println("media salario area: ", AreaSSm)

		for A, S := range AreaSSm {

			fmt.Println("Area: ", A, " Media Salario: ", S)
		}

		fmt.Println("")
		fmt.Println("Funcionario(s) com maior salario na Area: ")

		for A, N := range MapAreaSalFuncMaior {

			fmt.Println("Area: ", A, " Nome: ", N)
		}

		fmt.Println("")
		fmt.Println("Funcionarios com menor salario na Area: ")

		for A, N := range MapAreaSalFuncMenor {

			fmt.Println("Area: ", A, " Nome: ", N)
		}

		fmt.Println("")
		fmt.Println("area com maior quantidade funcionarios ", AreaQMa["maior"])

		fmt.Println("")
		fmt.Println("area com menor quantidade funcionarios ", AreaQMa["menor"])

		fmt.Println("")
		fmt.Println("Funcionario(s) com mesmo sobrenome: ")

		for A, S := range NomeSobreSal {

			fmt.Println("Nome: ", A, " Salario: ", S)
		}

		fmt.Println("")
		fmt.Println("Funcionario com mesmo sobrenome e com maior salario")
		fmt.Println("Nome: ", nomeRMS, " salario: ", slftmp)

		fmt.Println("")
		fmt.Println("#################################")
	}

}

//
//
//
func ExistSobreNome(V map[string]string, nome, sobrenome string) bool {

	for n, s := range V {

		if sobrenome == s && n != nome {

			return true
		}
	}
	return false
}
