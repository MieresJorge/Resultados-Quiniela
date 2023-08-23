package resultados

import (
	//"fmt"
	"log"
	"net/http"
	//"strings"

	"github.com/PuerkitoBio/goquery"
)

func Resultados_quiniela_Ciuad() []string{
	url := "https://www.jugandoonline.com.ar/rHome.aspx"
	//url := "https://www.jugandoonline.com.ar/SorteosQuiniela3.aspx"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Error: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resultado_Previa []string
	doc.Find( "#MainContent_CabezasHoyCiudad .enlaces-numeros"/*"#MainContent_CabezasHoyCiudadDiv .container.Ciudad .no-enlaces-numeros"*/).Each(func(index int, element *goquery.Selection) {
		text := element.Text()
		//fmt.Println(text)
		resultado_Previa = append(resultado_Previa, text)
	})
	return resultado_Previa
}

func Resultados_quiniela_Provincia() []string{
	url := "https://www.jugandoonline.com.ar/rHome.aspx"
	//url := "https://www.jugandoonline.com.ar/SorteosQuiniela3.aspx"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Error: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resultado_Provincia []string
	doc.Find( "#MainContent_CabezasHoyProvBsAsDiv .enlaces-numeros"/*"#MainContent_CabezasHoyCiudadDiv .container.Ciudad .no-enlaces-numeros"*/).Each(func(index int, element *goquery.Selection) {
		text := element.Text()
		//fmt.Println(text)
		resultado_Provincia = append(resultado_Provincia, text)
	})
	return resultado_Provincia
}


func Resultados_quiniela_SantaFe() []string{
	url := "https://www.jugandoonline.com.ar/rHome.aspx"
	//url := "https://www.jugandoonline.com.ar/SorteosQuiniela3.aspx"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Error: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resultado_SantaFe []string
	doc.Find( "#MainContent_CabezasHoySantaFeDiv .enlaces-numeros"/*"#MainContent_CabezasHoyCiudadDiv .container.Ciudad .no-enlaces-numeros"*/).Each(func(index int, element *goquery.Selection) {
		text := element.Text()
		//fmt.Println(text)
		resultado_SantaFe = append(resultado_SantaFe, text)
	})
	return resultado_SantaFe
}


func Resultados_quiniela_Cordoba() []string{
	url := "https://www.jugandoonline.com.ar/rHome.aspx"
	//url := "https://www.jugandoonline.com.ar/SorteosQuiniela3.aspx"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Error: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resultado_Cordoba []string
	doc.Find( "#MainContent_CabezasHoyCordoba .enlaces-numeros"/*"#MainContent_CabezasHoyCiudadDiv .container.Ciudad .no-enlaces-numeros"*/).Each(func(index int, element *goquery.Selection) {
		text := element.Text()
		//fmt.Println(text)
		resultado_Cordoba = append(resultado_Cordoba, text)
	})
	return resultado_Cordoba
}

func Resultados_quiniela_EntreRios() []string{
	url := "https://www.jugandoonline.com.ar/rHome.aspx"
	//url := "https://www.jugandoonline.com.ar/SorteosQuiniela3.aspx"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Error: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resultado_EntreRios []string
	doc.Find( "#MainContent_CabezasHoyEntreRios .enlaces-numeros"/*"#MainContent_CabezasHoyCiudadDiv .container.Ciudad .no-enlaces-numeros"*/).Each(func(index int, element *goquery.Selection) {
		text := element.Text()
		//fmt.Println(text)
		resultado_EntreRios = append(resultado_EntreRios, text)
	})
	return resultado_EntreRios
}

func Resultados_quiniela_Mendoza() []string{
	url := "https://www.jugandoonline.com.ar/rHome.aspx"
	//url := "https://www.jugandoonline.com.ar/SorteosQuiniela3.aspx"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Error: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resultado_Mendoza []string
	doc.Find( "#MainContent_CabezasHoyMendoza .enlaces-numeros"/*"#MainContent_CabezasHoyCiudadDiv .container.Ciudad .no-enlaces-numeros"*/).Each(func(index int, element *goquery.Selection) {
		text := element.Text()
		//fmt.Println(text)
		resultado_Mendoza = append(resultado_Mendoza, text)
	})
	return resultado_Mendoza
}