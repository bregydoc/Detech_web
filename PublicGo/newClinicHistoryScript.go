package main

import (
	"github.com/gopherjs/gopherjs/js"
	_ "github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"

	"strconv"
)

const TOKEN_AUTH = "99524c616bc275d72b28c97f6c61b21669100621"

var jQuery = jquery.NewJQuery
var materialize = js.Global.Get("Materialize")

func showToast(message string, callback func()) {
	materialize.Call("toast", message, 5000, "", callback)
}

func showLoader() {
	jQuery(".loader").Show()
}

func hideLoader() {
	jQuery(".loader").Hide()
}

func main() {
	btnCreate := jQuery("#btnCreateEvaluationFile")

	jQuery().Ready(func() {

		hideLoader()
		println("Ready Sir!")

		jQuery(".button-collapse").Call("sideNav")
		jQuery("select").Call("material_select")

		jQuery(".datepicker").Call("pickadate", map[string]interface{}{
			"selectMonths": true,
			"selectYears":  100,
		})

		btnCreate.On(jquery.CLICK, func(e jquery.Event) {
			peso := jQuery("#peso").Val()
			talla := jQuery("#talla").Val()
			hopDate := jQuery("#hopDate").Val()
			estadoBasal := jQuery("#estado_basal option:selected").Text()
			estadoNutricional := jQuery("#estado_nutricional option:selected").Text()
			frequeciaCambioPos := jQuery("#frecuencia_de_cambio_de_pos").Val()
			sondaVesical := jQuery("#sonda_vesical").Is(":checked")
			portadorSNG := jQuery("#portador_de_sng").Is(":checked")
			panal := jQuery("#panal").Is(":checked")
			albumina := jQuery("#albumina").Val()
			examenSangre := jQuery("#examenSangre").Val()
			creatinina := jQuery("#creatinina").Val()

			showLoader()
			jquery.Post("new-evaluation-file/submit", map[string]string{
				"peso":               peso,
				"talla":              talla,
				"hopDate":            hopDate,
				"estadoBasal":        estadoBasal,
				"estadoNutricional":  estadoNutricional,
				"frequeciaCambioPos": frequeciaCambioPos,
				"sondaVesical":       strconv.FormatBool(sondaVesical),
				"portadorSNG":        strconv.FormatBool(portadorSNG),
				"panal":              strconv.FormatBool(panal),
				"albumina":           albumina,
				"examenSangre":       examenSangre,
				"creatinina":         creatinina,
			}).Done(func(data jquery.Deferred) {
				jQuery("#dialogUserCreated").Call("modal", "open")
				showToast("Historia clínica creada", func() {
					println("Callback")
				})
				hideLoader()
			}).Fail(func() {
				println("Failed!")

			})

			//println(peso, talla, hopDate, estadoBasal, estadoNutricional, frequeciaCambioPos, sondaVesical,portadorSNG, panal,
			//	albumina, examenSangre, creatinina)
		})

	})

}
