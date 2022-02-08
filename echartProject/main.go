package main

import (
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	nameItems := []string{"衣", "食", "住"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "my示例图"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", []int{10, 20, 30}).
		AddYAxis("商家B", []int{15, 25, 35})
	f, err := os.Open("mybar.html")
	if err != nil {
		f, _ = os.Create("mybar.html")
	}
	bar.Render(w, f)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
