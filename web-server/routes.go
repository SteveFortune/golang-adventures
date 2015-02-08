
package main


type map[string] *RouteHandler RouteMap


/**
*	Our routing config... A big ol' dictionary of route keys paired
*	with handlers
*/
var RouteConfig = RouteMap{
	"/welcome": &WelcomeHandler{},
	"/about": &AboutHandler{},
	"/info": &InfoHandler{},
}


func (routeMap *RouteMap) RegisterRoutes(app *WebApplication) {}
