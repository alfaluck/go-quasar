package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/alfaluck/go-quasar/api"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

const clientBuildPath = `./client/dist/spa-mat`
const jwtSecret = `myVerySecret`

func main() {
	rpc.Register(&api.Auth{})
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	rpcHandler := jwtMiddleware.Handler(jsonrpc2.HTTPHandler(nil))
	http.Handle(`/rpc`, rpcHandler)

	http.Handle(`/css/`, http.FileServer(http.Dir(clientBuildPath+`/css`)))
	http.Handle(`/fonts/`, http.FileServer(http.Dir(clientBuildPath+`/fonts`)))
	http.Handle(`/img/`, http.FileServer(http.Dir(clientBuildPath+`/img`)))
	http.Handle(`/js/`, http.FileServer(http.Dir(clientBuildPath+`/js`)))
	http.Handle(`/statics/`, http.FileServer(http.Dir(clientBuildPath+`/statics`)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, clientBuildPath+`/index.html`)
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}