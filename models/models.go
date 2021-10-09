package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Users struct{
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Posts struct{
	Id bson.ObjectId `json:"id" bson:"_id"`
	Caption string `json:"caption" bson:"caption"`
	ImgURL string `json:"img" bson:"img"`
	TimestampPost string `json:"time" bson:"time"`
}