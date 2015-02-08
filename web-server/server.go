
package main


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
	handlerMap map[string] *RouteHandler


	/**
	*	Starts the application running.
	*
	*	@see Start
	*	@private
	*/
	isRunning bool

}


/**
*	Adds a new route and handler that will be listened for on start.
*	Routes should be added before `Start` is called, otherwise an error
*	will be returned.
*
*	@param	route		A web route in the format `/web/route`
*	@param 	handler		Handler for the route
*/
func (app *WebApplication) AddRoute(route string, handler *RouteHandler) error {

}
