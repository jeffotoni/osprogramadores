/**
 *
 *
 *
 */

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

// Quem mais recebe e quem menos recebe na empresa e a média salarial da empresa
// Quem mais recebe e quem menos recebe em cada área e a média salarial em cada área
// A área com mais funcionários e a área com menos funcionários
// Das pessoas que têm o mesmo sobrenome, aquela que recebe mais (não inclua sobrenomes que apenas uma pessoa tem nos resultados)

func main() {

	C := getCorporacao()

	funcionarios := C.Funcionarios
	// VAreas := C.Areas

	smaior := 0
	smenor := 1000000
	ssoma := 0

	qfunc := len(funcionarios)

	Mafunc := make(map[int]string)
	Mefunc := make(map[int]string)

	var datJson map[string]interface{}

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
	// vetor de area media salario

	NomeSobreSal := make(map[string]string)

	t := 0

	// Quem mais recebe e quem menos recebe na empresa e a média salarial da empresa
	for _, Func := range funcionarios {

		if Func.Salario > smaior { // recebe mais

			smaior = Func.Salario

		} else if Func.Salario < smenor { //recebe menos

			smenor = Func.Salario
		}

		// media salarial
		ssoma += Func.Salario

		MapAreas[Func.Area] = append(MapAreas[Func.Area], `{"area":"`+Func.Area+`","nome":"`+Func.Nome+`","salario":"`+fmt.Sprintf("%d", Func.Salario)+`"}`)

		VNomeSob[Func.Nome] = Func.Sobrenome

		t++
	}

	// buscando os nomes de maior e menor salario
	// da corporacao
	for i, Func := range funcionarios {

		if Func.Salario == smaior {

			Mafunc[i] = Func.Nome + " " + Func.Sobrenome
		}

		if Func.Salario == smenor {

			Mefunc[i] = Func.Nome + " " + Func.Sobrenome
		}

		// buscando e montando vetor dos sobrenomes que repetem
		if ExistSobreNome(VNomeSob, Func.Nome, Func.Sobrenome) {

			NomeSobreSal[Func.Nome+" "+Func.Sobrenome] = fmt.Sprintf("%d", Func.Salario)
			//fmt.Println("Sobrenome Igual: ", Func.Nome, " ::: ", Func.Sobrenome)
		}
	}

	qtmp := 0
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
		stmp2 := 100000

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
	fmt.Println("funcionarios maior salario na area: ", MapAreaSalFuncMaior)
	fmt.Println("funcionarios menor salario na area: ", MapAreaSalFuncMenor)
	fmt.Println("area com maior quantidade funcionarios ", AreaQMa["maior"])
	fmt.Println("area com menor quantidade funcionarios ", AreaQMa["menor"])

	fmt.Println("Funcionarios com mesmo sobrenome: ", NomeSobreSal)
	fmt.Println("Funcionarios com mesmo sobrenome com maior salario: ", nomeRMS, " salario: ", slftmp)

	fmt.Println("")
	fmt.Println("#################################")

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
