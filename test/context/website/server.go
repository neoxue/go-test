package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/neoxue/go-test/test/context/libs"
	"time"
	"golang.org/x/blog/content/context/google"
	"golang.org/x/blog/content/context/userip"
	"html/template"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
	var ctx 	libs.Mycontext
	var cancel 	libs.CancelFunc

	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		ctx, cancel = libs.WithTimeout(libs.Background(), 1)
	} else {
		ctx, cancel = libs.WithCancel(libs.Background())
	}

	time.Sleep(100 * time.Second)
	defer cancel()

	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
	}

	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)


	start := time.Now()
	results, err := google.Search(ctx, query)
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := resultsTemplate.Execute(w, struct {
		Results google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		logrus.Print(err)
		return
	}
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>`))