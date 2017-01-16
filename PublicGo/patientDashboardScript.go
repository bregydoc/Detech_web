package main

import (
	"github.com/gopherjs/gopherjs/js"
	_ "github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

const TOKEN_AUTH = "99524c616bc275d72b28c97f6c61b21669100621"

var jQuery = jquery.NewJQuery

func showLoader() {
	jQuery(".loader").Show()
}

func hideLoader() {
	jQuery(".loader").Hide()
}

func removeEvaluationFile(id string) {
	materialize := js.Global.Get("Materialize")
	idOfUser := jQuery("#id_of_user").Attr("class")
	dniOfPatient := jQuery("#dni_of_patient").Attr("class")
	showLoader()
	jquery.Post("/api/evaluation-file/remove/", map[string]string{
		"dni":          dniOfPatient,
		"token":        TOKEN_AUTH,
		"userId":       idOfUser,
		"idOfEvalFile": id,
	}).Done(func(data jquery.Deferred) {

		materialize.Call("toast", "Ficha de evaluacion eliminada", 5000, "", func() {
			//println("I'm a callback")
		})
		jQuery("#element" + id).Remove()
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
		println("Ready Sir!")

		jQuery(".button-collapse").Call("sideNav")

		jQuery(".deleteEvaluationFile").Each(func(index int, btn interface{}) {
			//btn.Attr("id")
			btnJs := btn.(*js.Object)

			btnJs.Call("addEventListener", "click", func() {
				go func() {
					id := btnJs.Get("id").String()[6:]
					removeEvaluationFile(id)

				}()
			})

		})
	})

}
