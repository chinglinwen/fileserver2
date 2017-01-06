package main

import (
	"html/template"
	"net/http"
	"strings"
)

func uploadPageHandler(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<html>
<title>Go upload</title>
<body>
<form action="{{.}}/uploadapi" method="post" enctype="multipart/form-data">
<label for="file">Files:</label>
<input type="file" name="file" id="file" multiple> <br>
Optional Filename:
<input type="text" name="file" >
<input type="submit" name="submit" value="Submit">
</form>
</body>
</html>
`
	t, err := template.New("page").Parse(tpl)
	checkError(err)
	t.Execute(w, strings.TrimSuffix((r.RequestURI), "/upload"))
}
