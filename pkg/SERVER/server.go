package server

import "net/http"

type Server struct {
}

// creating a constructing a new constructor calling the main.go
func New() *Server {
	return &Server{}
}

var indexPage = `
<!DOCTYPE html>
<html>
    <body>
	    <h1> My first paragraph </h1>
		<p> My FIRST paragraph</p>
	</body>
</html>
`
var userInfo = `{
    "name": "testuser"
	"age": 21
}`

// handling the index path "/"
func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexPage))
}

// handling the user path "/user"
func (s *Server) HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userInfo))
}
