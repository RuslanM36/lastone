package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"html/template"
)


var flooringOptions = map[string]float64{
	"Ламинат":    500,
	"Паркет":     1000,
	"Ковролин":   700,
	"Плитка":     1200,
	"Линолеум":   400,
}

type CalculationResult struct {
	FlooringType string
	Area         float64
	TotalPrice   float64
}

var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Калькулятор покрытия</title>
	<link rel="stylesheet" type="text/css" href="style.css">
</head>
<body>
	<h1>Калькулятор покрытия</h1>
	<form method="POST" action="/">
		<label for="flooring">Выберите тип покрытия:</label>
		<select name="flooring" id="flooring">
			<option value="Ламинат">Ламинат</option>
			<option value="Паркет">Паркет</option>
			<option value="Ковролин">Ковролин</option>
			<option value="Плитка">Плитка</option>
			<option value="Линолеум">Линолеум</option>
		</select>
		<br>
		<label for="area">Введите площадь (в м²):</label>
		<input type="text" name="area" id="area" required>
		<br>
		<input type="submit" value="Рассчитать">
	</form>
	{{if .FlooringType}}
	<h2>Результат:</h2>
	<p>Тип покрытия: {{.FlooringType}}</p>
	<p>Площадь: {{.Area}} м²</p>
	<p>Общая стоимость: {{.TotalPrice}} рублей</p>
	{{end}}
</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}