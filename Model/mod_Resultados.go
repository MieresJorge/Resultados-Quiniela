package mod_resultados

import (
	resultados "new.com/events/Scraping"
	//"fmt"
)

type User struct {
	Posicion  string            `json:"Quiniela"`
	Resultado map[string]string `json:"Cabeza"`
}

//Arrays con resultados de las quinielas nacionales
var Ciudad []string = resultados.Resultados_quiniela_Ciuad()
var Provincia []string = resultados.Resultados_quiniela_Provincia()
var Cordoba []string = resultados.Resultados_quiniela_Cordoba()
var EntreRios []string = resultados.Resultados_quiniela_EntreRios()
var Mendoza []string = resultados.Resultados_quiniela_Mendoza()
var SantaFe []string = resultados.Resultados_quiniela_SantaFe()

func generateResultMap(numbers []string, labels []string) map[string]string {
	resultMap := make(map[string]string)
	for i, number := range numbers {
		if i < len(labels) {
			resultMap[labels[i]] = number
		}
	}
	return resultMap
}

var Resultados_quiniela_Ciuad_1 = User{
		Posicion:  "Ciudad",
		Resultado: generateResultMap(Ciudad, []string{"Previa", "Primera","Matutina","Vespertina", "Nocturna"}), 
	}

var Resultados_quiniela_Provincia_1 = User{
	Posicion:  "Provincia",
	Resultado: generateResultMap(Provincia, []string{"Previa", "Primera","Matutina","Vespertina", "Nocturna"}),
	}

var Resultados_quiniela_EntreRios_1 = User{
	Posicion:  "Entre Rios",
	Resultado: generateResultMap(EntreRios, []string{"Previa", "Primera","Matutina","Vespertina", "Nocturna"}),
	}

var Resultados_quiniela_Cordoba_1 = User{
	Posicion:  "Cordoba",
	Resultado: generateResultMap(Cordoba, []string{"Previa", "Primera","Matutina","Vespertina", "Nocturna"}),
	}

var Resultados_quiniela_SantaFe_1 = User{
	Posicion:  "SantaFe",
	Resultado: generateResultMap(SantaFe, []string{"Previa", "Primera","Matutina","Vespertina", "Nocturna"}),
	}

var Resultados_quiniela_Mendoza_1 = User{
	Posicion:  "Mendoza",
	Resultado: generateResultMap(Mendoza, []string{"Previa", "Primera","Matutina","Vespertina", "Nocturna"}),
	}