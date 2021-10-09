package main

import (
	"appointy-task/controllers"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main(){
	req := httprouter.New()
	userController := controllers.NewUserController(getSession())
	postController := controllers.NewPostController(getSession())

	req.GET("/users/:id", userController.GetUser)
	req.POST("/users", userController.CreateUser)
	req.GET("/posts/:id", postController.GetPost)
	req.POST("/posts", postController.CreatePost)
	http.ListenAndServe("localhost:8081", req)
}

func getSession() *mgo.Session{
	session, error := mgo.Dial("mongodb://localhost:27017")
	if error != nil{
		panic(error)
	}
	return session
}