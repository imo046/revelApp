package controllers

import "github.com/revel/revel"

/*
Implementation of the Revel controller, should have Controller in the first slot
*/
type App struct {
	*revel.Controller
}

/*
Any method which is exported and returns revel.Result could serve as an Action.
In the method below we call Render() which renders a template as a response with 200 OK
*/
func (a App) Home() revel.Result {
	message := "Test"
	return a.Render(message)
}

func (a App) Hello(inputInfo string) revel.Result {
	return a.Render(inputInfo)
}
