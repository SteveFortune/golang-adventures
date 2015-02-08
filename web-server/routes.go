
package main


/**
*	Our routing config... A big ol' dictionary of route keys paired
*	with handlers
*/
var RouteConfig = RouteMap{
	"/welcome": WelcomeHandler{
		title: "Welcome! To my first ever GoLang web server",
		greeting: "I hope you enjoy your stay.",
	},
	"/about": AboutHandler("About us.."),
	"/info": InfoHandler("Info.."),
}
