
package main


/**
*	Fires up our web application: instantiates our server and starts
*	its listening.
*/
func main() {

	server := NewWithRoutes("localhost:9999", RouteConfig)
	server.Start()

}
