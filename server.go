package main

import (
	"log"
	"net/http"
	"text/template"
)

// URL  string
// Body string

// Error         error
// Output        string
// IsValidPython bool

var IndexTemplate = `<html>
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
<title>IsItPython v0.1</title>

<div class="container">
<h1>IsItPython! v0.1</h1>
{{ with .Snippets }}
	{{ range . }}
		<div class="card" style="width: 100%;">
			<div class="card-body">
				<h5 class="card-title">IsItPython?
				{{if .IsValidPython}}
					<img style="height:25px" src="http://icons.iconarchive.com/icons/bokehlicia/captiva/256/checkbox-icon.png"
				{{ else }} 
					<img style="height:25px" src="http://icons.iconarchive.com/icons/ampeross/qetto-2/256/no-icon.png"	 
				{{ end }}
				</h5>
				<h6 class="card-subtitle mb-2 text-muted">{{.URL}}</h6>
				<p class="card-text">
					<pre><code>{{.Body}}</code></pre>

					<strong>Output</strong>
					<pre<code>{{.Output}}{{.Error}}</code></pre>
				</p>
			</div>
	</div>
	{{ end }} 
{{ end }}




</div>
</html>`

func Handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("index").Parse(IndexTemplate)

	data := struct {
		Snippets []*Snippet
	}{
		LoadSnippets(),
	}

	t.Execute(w, data)
}

func RunWebServer() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
