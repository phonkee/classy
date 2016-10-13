package classy

import (
	"fmt"
	"strings"
)

/*
Routes is list of routes
*/
type Routes []ViewRoute

func NewRoutes(routes ...ViewRoute) (result Routes) {
	result = make(Routes, 0, len(routes))
	result = append(result, routes...)
	return
}

func (r Routes) Get(name string) (result ViewRoute, err error) {
	for _, route := range r {
		if route.GetName() == name {
			result = route
			break
		}
	}
	return
}

/*
Route interface
API for routes
*/
type ViewRoute interface {

	// Setters for properties
	// Sets name of route, such as "detail", "list"
	Name(string) ViewRoute

	// Set path for route, can contain backreferences to variables (gorilla mux is used)
	Path(string) ViewRoute

	// GetPath is getter method to return path of ViewRoute
	GetPath() string

	// Set Suffix to be appended after route name.
	Suffix(string) ViewRoute

	// Getters of properties
	// Getter for name
	GetName() string

	// return methodmap by given http method
	GetMethodMap() MethodMap

	// return methodmap by given http method
	GetMethodSet(httpMethod string, create ...bool) (MethodSet, bool)

	// return full name for route name
	GetFullName(view Viewer) string
}

/*
NewRoute returns new route
*/
func NewViewRoute(path string, methods ...MethodMap) ViewRoute {
	result := &viewroute{
		methodmap: NewMethodMap(),
	}
	result.Path(path)
	for _, method := range methods {
		result.methodmap.Update(method)
	}
	return result
}

/*
Route Implementation
*/
type viewroute struct {
	path      string
	methodmap MethodMap
	name      string
	suffix    string
}

func (r *viewroute) GetName() string {
	return r.name
}

func (r *viewroute) GetMethodMap() MethodMap {
	return r.methodmap
}

func (r *viewroute) GetMethodSet(method string, create ...bool) (result MethodSet, ok bool) {
	var shouldcreate bool
	if len(create) > 1 {
		shouldcreate = create[0]
	}

	//_, _ = r.methodmap[method]
	_ = shouldcreate

	return
}

// Name sets name of route
func (r *viewroute) Name(name string) ViewRoute {
	r.name = strings.TrimSpace(name)
	return r
}
func (r *viewroute) Path(path string) ViewRoute {
	r.path = strings.TrimSpace(path)
	return r
}

func (r *viewroute) Suffix(suffix string) ViewRoute {
	r.suffix = strings.TrimSpace(suffix)
	return r
}

func (r *viewroute) GetPath() string {
	return r.path
}

func (r *viewroute) GetFullName(view Viewer) (result string) {
	result = GetViewName(view, true)
	if r.name != "" {
		result = fmt.Sprintf("%v:%v", result, r.name)
	}
	return
}
