package routes

import (
	"github.com/fuckthinkpad/internal/controllers"
	"github.com/fuckthinkpad/internal/middlewares"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	var r *mux.Router = mux.NewRouter()
	r.Use(middlewares.LoggerMiddleware)
	r.HandleFunc("/channel-create", controllers.CreateChannel)
	r.HandleFunc("/channel-join", controllers.GetChannels)

	return r
}
