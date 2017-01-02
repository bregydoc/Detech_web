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


func removePatient(dni string) {
	materialize := js.Global.Get("Materialize")
	idOfUser := jQuery("#id_of_user").Attr("class")
	showLoader()
	jquery.Post("/api/patients/remove/", map[string]string{
		"dni": dni,
		"token" : TOKEN_AUTH,
		"userId":idOfUser,

	}).Done(func(data jquery.Deferred) {


		materialize.Call("toast", "Usuario eliminado", 5000, "", func() {
			//println("I'm a callback")
		})
		jQuery("#element"+dni).Remove()
		hideLoader()
	}).Fail(func() {
		materialize.Call("toast", "No se completo la operacion", 5000, "", func() {
			//println("I'm a callback")
		})

	})

}


func main() {
	jQuery().Ready(func() {

		hideLoader()

		jQuery(".button-collapse").Call("sideNav")



		jQuery(".deletePatient").Each(func (index int, btn interface{}) {
			//btn.Attr("id")
			btnJs := btn.(*js.Object)

			btnJs.Call("addEventListener", "click", func() {
				go func() {
					dni := btnJs.Get("id").String()[6:]
					removePatient(dni)

				}()
			})

		})

		jQuery(".detailsPatient").Each(func (index int, btn interface{}) {
			//btn.Attr("id")
			btnJs := btn.(*js.Object)

			btnJs.Call("addEventListener", "click", func() {
				go func() {
					dni := btnJs.Get("id").String()[6:]
					println(dni)
				}()
			})

		})
	})

}
