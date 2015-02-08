
package main

import (
	"net/http"
	"errors"
)


type RouteMap map[string] http.Handler


type WebServer struct {


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
	*
	*	@private
	*/
	basePath string

}


/**
*	Has a handler already been registered against the given route?
*
*	@public
*	@param 	route		The given route to query the underlying route
*						map with
*	@return	bool		Whether a handler exists for the route
*/
func (app *WebServer) RouteExists(route string) bool {
	_, isOk := app.routeMap[route]
	return isOk
}


/**
*	Adds a new route and handler that will be listened for on start.
*	Routes should be added before `Start` is called, otherwise an error
*	will be returned.
*
*	@public
*	@see 	AddRoutes
*	@param	route		A web route in the format `/web/route`
*	@param 	handler		Handler for the route
*	@return	error		Not nil if the application has started already or
*						a handler exists for the route already
*/
func (app *WebServer) AddRoute(route string, handler http.Handler) error {

	var err error

	if app.isRunning {
		err = errors.New("Web application is already running")
	} else if app.RouteExists(route) {
		err = errors.New("Route already exists, please remove it first")
	} else {
		app.routeMap[route] = handler
	}

	return err
}


/**
*	Merges the given route map into the existing map.
*
*	@public
*	@see 	AddRoute
*	@param 	routeMap	The route map to merge into the underlying map
*	@return	error		Not nil if the application has started already or
*						a handler exists for any one of the routes already
*/
func (app *WebServer) AddRoutes(routeMap RouteMap) error {

	if app.isRunning {
		return errors.New("Web application is already running")
	}

	for route, _ := range routeMap {
		if app.RouteExists(route) {
			return errors.New("Route in map already exists.")
		}
	}

	for route, handler := range routeMap {
		app.routeMap[route] = handler
	}

	return nil

}


/**
*	Starts the listening of the web server
*
*	@public
*	@return	error		Not nil if an error occurs on listen
*/
func (app *WebServer) Start() error {

	for route, handler := range app.routeMap {
		http.Handle(route, handler)
	}

	return http.ListenAndServe(app.basePath, nil)
}


/**
*	Is the web server current running ?
*
*	@return	bool
*/
func (app *WebServer) IsRunning() bool {
	return app.isRunning
}


/**
*	Inits the a web application with a route map.
*
*	@param	baseBath
*	@param 	routeMap
*	@return *WebServer		The new web application
*/
func NewWithRoutes(basePath string, routeMap RouteMap) *WebServer {
	return &WebServer{
		routeMap: routeMap,
		basePath: basePath,
	}
}


/**
*	Inits a web application with an empty route map
*
*	@param	baseBath
*	@return *WebServer		The new web application
*/
func New(basePath string) *WebServer {
	return NewWithRoutes(basePath, RouteMap{})
}
