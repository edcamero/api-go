package controllers

import (
	//"encoding/json"
	"context"
	"fmt"
	"log"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/edcamero/api-go/db"
	"github.com/edcamero/api-go/environment"
	"github.com/edcamero/api-go/models"
	"github.com/edcamero/api-go/util"
)

type Token struct {
	TokenUser string `json:"tokenUser"`
}

func Login(ctx iris.Context) {

	user := models.User{}
	//err := json.NewDecoder(ctx.Request().Body).Decode(&filter)
	username := ctx.FormValue("username")
	password := util.Encrypt([]byte(ctx.FormValue("password")))
	filter := bson.M{"username": username, "password": password}
	fmt.Println(username)
	fmt.Println(password)
	//fmt.Println(ctx.FormValue("password"))

	conexion := db.GetConnection()
	collection := conexion.Database(environment.DATABASE).Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println(user)
		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"rol":      user.Rol.Nombre,
			"username": user.UserName,
		})
		tokenString, _ := token.SignedString([]byte(environment.SECRETKEY))
		user.Password = ""
		ctx.JSON(iris.Map{"user": user, "tokenUser": tokenString})

	}

}

func AuthenticatedAdoptante(ctx iris.Context) {

	user := ctx.Values().Get("jwt").(*jwt.Token)

	foobar := user.Claims.(jwt.MapClaims)
	resp := false
	for key, value := range foobar {

		ctx.Writef("%s = %s", key, value)
		if key == "rol" && value == "adoptante" {
			resp = true
		}
	}

	if !resp {

		user = nil
	}

}

func AuthenticatedAdmin(ctx iris.Context) {
	//ctx.StopWithStatus(iris.StatusUnauthorized)

	user := ctx.Values().Get("jwt").(*jwt.Token)
	item := user.Claims.(jwt.MapClaims)
	if item["rol"] == "admin" {
		ctx.Next()
	} else {
		ctx.StopWithStatus(iris.StatusUnauthorized)
	}

	return
}
