package controllers

import (
	"appointy-task/models"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session *mgo.Session
}

type PostController struct{
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController{
	return &UserController{session: session}
}

func NewPostController(session *mgo.Session) *PostController{
	return &PostController{session: session}
}

func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	objectID := bson.ObjectIdHex(id)
	user := models.Users{}

	if error := userController.session.DB("instagram").C("users").FindId(objectID).One(&user); error != nil{
		w.WriteHeader(404)
		return
	}

	userJSON, error := json.Marshal(user)
	if error != nil{
		fmt.Println(error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJSON)
}

func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	user := models.Users{}

	json.NewDecoder(r.Body).Decode(&user)

	user.Id = bson.NewObjectId()

	userController.session.DB("instagram").C("users").Insert(user)

	userJSON, error := json.Marshal(user)

	if error != nil{
		panic(error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userJSON)
}

func (postController PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	objectID := bson.ObjectIdHex(id)
	posts := models.Posts{}

	if error := postController.session.DB("instagram").C("posts").FindId(objectID).One(&posts); error != nil{
		w.WriteHeader(404)
		return
	}

	userJSON, error := json.Marshal(posts)
	if error != nil{
		fmt.Println(error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJSON)
}

func (postController PostController) CreatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	post := models.Posts{}

	json.NewDecoder(r.Body).Decode(&post)

	post.Id = bson.NewObjectId()

	postController.session.DB("instagram").C("posts").Insert(post)

	userJSON, error := json.Marshal(post)

	if error != nil{
		panic(error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userJSON)
}