package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rokuosan/image-server/internal/handler"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Configure(e *echo.Echo) {
	handlers := handler.Handlers()
	for _, h := range handlers {
		e.GET(h.Route(), h.GET)
		e.POST(h.Route(), h.POST)
		e.PUT(h.Route(), h.PUT)
		e.DELETE(h.Route(), h.DELETE)
	}
}
