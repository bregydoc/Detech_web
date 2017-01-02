package main

import (

	_"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"github.com/gopherjs/gopherjs/js"
)

const TOKEN_AUTH  = "99524c616bc275d72b28c97f6c61b21669100621"

var jQuery = jquery.NewJQuery


func showLoader() {
	jQuery(".loader").Show()
}

func hideLoader() {
	jQuery(".loader").Hide()
}


func main() {

	materialize := js.Global.Get("Materialize")

	btnCreatePatient := jQuery("#btnSendNewPatient")
	btnClearForm := jQuery("#btnResetForm")

	jQuery().Ready(func() {
		println("Ready gopherJs")
		materialize.Call("updateTextFields")
		jQuery(".button-collapse").Call("sideNav")
		jQuery(".loader").Hide()
		jQuery(".sexSelector").Call("material_select")
	})

	btnClearForm.On(jquery.CLICK, func(e jquery.Event) {
		jQuery("#nombre_completo").SetVal("")
		jQuery("#dni").SetVal("")
		jQuery("#telefono").SetVal("")
		jQuery("#domicilio").SetVal("")
		jQuery("#numero_de_hc").SetVal("")
		materialize.Call("updateTextFields")

	})


	btnCreatePatient.On(jquery.CLICK, func(e jquery.Event) {
		//println("Submit")

		name := jQuery("#nombre_completo").Val()
		dni := jQuery("#dni").Val()
		phone := jQuery("#telefono").Val()
		address := jQuery("#domicilio").Val()
		hc := jQuery("#numero_de_hc").Val()
		sex := jQuery("input:radio[name ='sexo']:checked").Val()


		//Falta verificación de fidelidad de datos
		//... Espacio reservado ...
		showLoader()
		jquery.Post("new_patient/submit", map[string]string{
			"dni": dni,
			"name":name,
			"address":address,
			"phone":phone,
			"numOfHC": hc,
			"sex" : sex,

		}).Done(func(data jquery.Deferred) {
			jQuery("#dialogUserCreated").Call("modal", "open")
			materialize.Call("toast", "Usuario creado con exito!", 5000, "", func() {
				println("I'm a callback")
			})
			hideLoader()
		}).Fail(func() {
			println("Failed!")

		})


	})
}