package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

var flooringOptions = map[string]float64{
    "Ламинат":  500,
    "Паркет":   1000,
    "Ковролин": 700,
    "Плитка":   1200,
    "Линолеум": 400,
}

type CalculationResult struct {
    FlooringType string
    Area         float64
    TotalPrice   float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
    if r.Method == http.MethodPost {
        flooring := r.FormValue("flooring")
        areaStr := r.FormValue("area")
        area, err := strconv.ParseFloat(areaStr, 64)
        if err != nil {
            http.Error(w, "Неверный ввод площади", http.StatusBadRequest)
            return
        }
        pricePerSqM := flooringOptions[flooring]
        totalPrice := area * pricePerSqM
        result := CalculationResult{
            FlooringType: flooring,
            Area:         area,
            TotalPrice:   totalPrice,
        }
        tmpl.Execute(w, result)
    } else {
        tmpl.Execute(w, nil)
    }
}

func main() {
    http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    fmt.Println("Running demo app. Press Ctrl+C to exit...")
	og.Println("Server started at http://localhost:8888")
    log.Fatal(http.ListenAndServe(":8888", nil))
}
