package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"log"
	"fmt"
)

const TOKEN_AUTH  = "99524c616bc275d72b28c97f6c61b21669100621"

func GetRouter() *gin.Engine{

	s := gin.Default()

	s.LoadHTMLGlob("Pages/*")

	s.Static("/Static/", "./Static/")
	s.Static("/Public/", "./Public/")


	ctx := context.Background()

	s.POST("/api/upload/user/:name", func(c *gin.Context) {

		token := c.PostForm("token")

		if token != TOKEN_AUTH {
			c.String(http.StatusNotAcceptable, "No access")
		} else {
			name := c.Param("name")
			title := c.PostForm("title")
			path := c.PostForm("localPath")
			finalName := strings.ToUpper(string(name[0])) + name[1:]

			if err := UploadImage(ctx, path, finalName + "/" + title); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			} else {
				c.String(http.StatusOK, "Ok, uploaded")
			}

		}

	})

	s.POST("/api/download/user/:name", func(c *gin.Context) {

		token := c.PostForm("token")

		if token != TOKEN_AUTH {
			c.String(http.StatusNotAcceptable, "No access")

		} else {
			name := c.Param("name")
			title := c.PostForm("title")
			localPath := c.PostForm("localPath")
			finalName := strings.ToUpper(string(name[0])) + name[1:]

			if err := DownloadImage(ctx, finalName + "/" + title, localPath); err != nil {
				log.Fatalln(err)
				c.String(http.StatusInternalServerError, err.Error())
			} else {
				c.String(http.StatusOK, "Ok, saved in %v", localPath)
			}

		}

	})


	s.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		//Actualiza el contexto (El doctor que esta actualmente usando el servicio)
		//Desconosco el comportamiento que tendrá al ser concurrente y con múltiples conexiones
		user, err := GetUserById(id)
		if  err != nil {
			fmt.Println(err)
		}else{
			fmt.Println(user)
			GetCurrentContext().RefreshDefaultContext(user)

		}
		c.HTML(http.StatusOK, "dashboardTemplate.html", gin.H{
			"id" : id,
			"doctor" : user.Name + " " + user.LastName,
			"Patients" : user.Patients,
		})
	})

	s.GET("/user/:id/new_patient", func(c *gin.Context) {
		id := c.Param("id")

		user, err := GetUserById(id)

		if  err != nil {
			fmt.Println(err)
		}else{
			fmt.Println(user)
			GetCurrentContext().RefreshDefaultContext(user)

		}
		c.HTML(http.StatusOK, "createNewPatient.html", gin.H{
			"id" : id,
			"doctor" : user.Name + " " + user.LastName,
			"Patients" : user.Patients,
		})
	})

	s.POST("/user/:id/new_patient/submit", func(c *gin.Context) {
		id := c.Param("id")
		//dni := c.PostForm("dni")
		//Check if the patient already exist

		dni := c.PostForm("dni")
		nombre := c.PostForm("name")
		domicilio := c.PostForm("address")
		telefono := c.PostForm("phone")
		numero_de_HC := c.PostForm("numOfHC")
		sexo := c.PostForm("sex")

		newPatient := CreateNewPatient(dni, nombre, domicilio, telefono, numero_de_HC, sexo)
		currentUser, err := GetUserById(id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Id de usuario no encontrado")
		}

		currentUser.CreateANewPatient(newPatient)

		c.String(http.StatusOK, "ok")
	})

	s.POST("api/patients/remove/", func(c *gin.Context) {
		token := c.PostForm("token")
		userId := c.PostForm("userId")
		dni := c.PostForm("dni")

		if token != TOKEN_AUTH {
			c.String(http.StatusNotAcceptable, "No access")
		} else {
			user, err := GetUserById(userId)

			if err != nil {
				c.Error(err)
				return
			}

			err = user.RemovePatient(dni)

			if err != nil {
				c.Error(err)
				return
			}

		}
	})
	return s
}