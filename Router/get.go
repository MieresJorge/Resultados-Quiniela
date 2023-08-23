package get_router

import ( 
	"encoding/json"
    "net/http"
	resultados "new.com/events/Scraping"
    "github.com/gorilla/mux"
)


//resultados de quiniela 
var mostrar []string=resultados.Resultados_quiniela_Ciuad()

// Estructura de ejemplo para los datos
type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

/*var items = []mostrar{
    {ID: "1", Name: "Item 1"},
    {ID: "2", Name: "Item 2"},
}*/

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	datos := map[string]interface{}{ 
        "Titulo" : "resultado1", 
        "Resultado" : mostrar[1],
        /*        "Previa": mostrar[0],
        "Primera": mostrar[1],
        "Matutina": mostrar[2],
        "Vespertina": mostrar[3],
        "Nocturna": mostrar[4],*/
    }
    json.NewEncoder(w).Encode(datos)
}

func Get_router() {
    r := mux.NewRouter()

    // Define una ruta para la API GET
    r.HandleFunc("/api/Resultados", GetItemsHandler).Methods("GET")

    http.Handle("/", r)

    // Inicia el servidor en el puerto 8080
    http.ListenAndServe(":8080", nil)
}