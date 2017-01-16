package main

import (
	"github.com/gopherjs/gopherjs/js"
	_ "github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
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
	jQuery().Ready(func() {
		hideLoader()
		println("Ready Sir!")

		jQuery(".button-collapse").Call("sideNav")
		jQuery("select").Call("material_select")
	})
}
