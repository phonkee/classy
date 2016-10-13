package classy

import (
	"strings"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

/*
New instantiates new classy view, that is able to introspect view for http methods.
*/
func New(view Viewer) Classy {
	result := &classy{
		chain: alice.New(),
		view:  view,
	}
	return result
}

/*
Implementation of classy
*/
type classy struct {
	chain alice.Chain
	view  Viewer
}

/*
Register routes to router
*/
func (c *classy) Register(router *mux.Router, name ...string) (result []*mux.Route) {
	result = []*mux.Route{}

	// check if name was provided
	forcename := ""
	if len(name) > 0 {
		forcename = strings.TrimSpace(name[0])
	}


	for _, route := range c.view.GetRoutes() {
		bound, _ := route.GetMethodMap().Bind(c.view)
		for httpmethod, bounditem := range bound {
			routename := route.GetFullName(c.view)

			// set name that was given as optional argument
			if forcename != "" {
				routename = forcename
			}

			registeredRoute := router.Handle(
				route.GetPath(),
				c.chain.ThenFunc(bounditem.handlerfunc),
			).Methods(string(httpmethod)).Name(routename)

			//logger.Info("classy, method: %+v, routename: %+v, path: %+v, func: %+v", httpmethod, routename, route.GetPath(),
			//	bounditem.method,
			//)

			result = append(result, registeredRoute)
		}
	}
	return
}

/*
Add Middleware
*/
func (c *classy) Use(m ...alice.Constructor) Classy {
	c.chain = c.chain.Append(m...)
	return c
}
