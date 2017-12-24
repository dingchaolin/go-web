package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func main(){
	http.HandleFunc( "/", Hey )

	http.ListenAndServe(":8080", nil)

}




func Hey( w http.ResponseWriter, r *http.Request ){

	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<title>Hey</title>
	</head>
	<body>
		<form method="post" action="/">
			Username: <input type="text" name="uname">
			Password: <input type="password" name="pwd">
			<button type="submit">Submit</button>
		</form>
	</body>
</html>`

	if r.Method == "GET" {
		fmt.Println( "GET" )
		t, err := template.New("hey").Parse(tpl)
		fmt.Println( err )
		t.Execute(w, nil)
	}else{
		fmt.Println( "POST" )
		r.ParseForm()
		fmt.Println( r.FormValue("uname"))
	}

}
