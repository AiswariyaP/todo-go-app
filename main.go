package main

import (
	"GOLANG/todo/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	middleware.InitMiddleWare(r)

	//run the server at port number 3039
	s := &http.Server{
		Addr:    ":3039",
		Handler: r,
	}
	e := s.ListenAndServe()
	if e != nil {
		panic(e)
	}

}
