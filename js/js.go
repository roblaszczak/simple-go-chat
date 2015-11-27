package main

import "github.com/gopherjs/gopherjs/js"

// BindOnWindowResize binds provided function on browser window resize.
func BindOnWindowResize(doOnResize func()) {
	jQuery(js.Global.Get("window")).Resize(func() {
		doOnResize()
	})
}

// An AngularApp represents instance of angular application in js.
type AngularApp struct {
	*js.Object
}

// An AngularAppModules represents list of modules for Angular application.
type AngularAppModules []string

// CreateAngularApp creates new instance of AngularApp
func CreateAngularApp(name string, modules AngularAppModules) *AngularApp {
	app := js.Global.Get("angular").Call("module", name, modules)
	js.Global.Set("app", app)

	return &AngularApp{app}
}
