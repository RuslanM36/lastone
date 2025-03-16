package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
	"path/filepath"
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
    tmplPath := filepath.Join("index.html")
    tmpl := template.Must(template.ParseFiles(tmplPath))

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
    fs := http.FileServer(http.Dir("."))  
    http.Handle("/assets/style.css", fs)
    http.HandleFunc("/", handler)

    fmt.Println("Running demo app. Press Ctrl+C to exit...")
    log.Fatal(http.ListenAndServe(":8888", nil))
}
