
package main

import (
	"net/http"
	"error"
)


type map[string] *RouteHandler RouteMap


/**
*	Our web application. There is
*/
type WebApplication struct {


	/**
	*	A map of web routes to handler for those routes. To register a new
	*	handler for a route, use the `AddRoute` method
	*
	*	@see AddRoute
	*	@private
	*/
	routeMap RouteMap


	/**
	*	Starts the application running.
	*
	*	@see Start
	*	@private
	*/
	isRunning bool


	/**
	*	The application's base path, e.g. `localhost:8080`
	*/
	basePath string

}


/**
*	Has a handler already been registered against the given route?
*
*	@param 	route		The given route to query the underlying route
*						map with
*	@return	bool		Whether a handler exists for the route
*/
func (app *WebApplication) RouteExists(route string) bool {
	_, isOk = app.routeMap[route]
	return isOk
}


/**
*	Adds a new route and handler that will be listened for on start.
*	Routes should be added before `Start` is called, otherwise an error
*	will be returned.
*
*	@see 	AddRoutes
*	@param	route		A web route in the format `/web/route`
*	@param 	handler		Handler for the route
*	@return	error		Not nil if the application has started already or
*						a handler exists for the route already
*/
func (app *WebApplication) AddRoute(route string, handler *RouteHandler) error {

	var err error

	if app.isRunning {
		err = error.New("Web application is already running")
	} else if app.routeExists(route) {
		err = error.New("Route already exists, please remove it first")
	} else {
		app.routeMap[route] = handler
	}

	return err
}


/**
*	Merges the given route map into the existing map.
*
*	@see 	AddRoute
*	@param 	routeMap	The route map to merge into the underlying map
*	@return	error		Not nil if the application has started already or
*						a handler exists for any one of the routes already
*/
func (app *WebApplication) AddRoutes(routeMap RouteMap) error {

	if app.isRunning {
		return error.New("Web application is already running")
	}

	for route, handler := range routeMap {
		if app.routeExists(route) {
			return error.New("Route in map already exists.")
		}
	}

	for route, handler := range routeMap {
		app.routeMap[route] = handler
	}

	return nil

}


func (app *WebApplication) Start() error {

	for route, handler := range app.routeMap {
		err = http.ListenAndServer(app.bastPath)
	}

	return http.ListenAndServer(app.bastPath)
}


func (app *WebApplication) Stop() error {

	///... hmmm how do I do this... do I even want to do this?
}


func (app *WebApplication) Restart() error {

}


func NewWithRoutes(basePath string, routeMap *RouteMap) *WebApplication {
	return &WebApplication{
		routeMap: routeMap,
		basePath: basePath,
	}
}


func New(basePath string) *WebApplication {
	return NewWithRoutes(basePath, RouteMap{})
}
