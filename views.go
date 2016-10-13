package classy

import "net/http"

/*
BaseView is blank implementation of basic view
*/
type BaseView struct{}

/*
Before blank implementation
*/
func (v BaseView) Before(w http.ResponseWriter, r *http.Request) error {
	return nil
}

/*
GetRoutes blank implementation
*/
func (v *BaseView) GetRoutes() Routes {
	return NewRoutes()
}

/*
ListView is predefined struct for list views
*/
type View struct {
	BaseView
}

/*
Routes returns list of routes with predefined method maps
*/
func (l View) GetRoutes() Routes {
	return NewRoutes(
		NewViewRoute("/",
			NewMethodMap().
				Add("GET", "Get").
				Add("POST", "Post").
				Add("DELETE", "Delete").
				Add("PUT", "Put").
				Add("HEAD", "Head").
				Add("TRACE", "Trace").
				Add("OPTIONS", "Options"),
		),
	)
}



/*
ListView is predefined struct for list views
*/
type ListView struct {
	BaseView
}

/*
Routes returns list of routes with predefined method maps
*/
func (l ListView) GetRoutes() Routes {
	return NewRoutes(
		NewViewRoute("/",
			NewMethodMap().
				Add("GET", "List").
				Add("OPTIONS", "Metadata").
				Add("POST", "Create"),
		),
	)
}

/*
DetailView is predefined struct for detail
*/
type DetailView struct {
	BaseView
}

/*
Routes returns list of routes with predefined method maps
*/
func (d DetailView) GetRoutes() Routes {
	return NewRoutes(
		NewViewRoute("/{pk:[0-9]+}/",
			NewMethodMap().
				Add("DELETE", "Delete").
				Add("GET", "Retrieve").
				Add("OPTIONS", "Metadata").
				Add("POST", "Update"),
		),
	)
}

/*
SlugDetailView is predefined struct for detail that handles id as slug
*/
type SlugDetailView struct {
	BaseView
}

/*
Routes returns list of routes with predefined method maps
*/
func (d SlugDetailView) GetRoutes() Routes {
	return NewRoutes(
		NewViewRoute("/{slug}/",
			NewMethodMap().
				Add("DELETE", "Delete").
				Add("GET", "Retrieve").
				Add("OPTIONS", "Metadata").
				Add("POST", "Update"),
		),
	)
}

/*
ViewSet is combination of list and detail views.
*/

type ViewSet struct {
	ListView
	DetailView
}

/*
Before is blank implementation for ViewSet
*/
func (v ViewSet) Before(w http.ResponseWriter, r *http.Request) error {
	return nil
}

/*
GetRoutes returns combination of list and detail routes
*/
func (v ViewSet) GetRoutes() Routes {

	result := Routes{}

	detailroutes := v.DetailView.GetRoutes()
	for _, route := range detailroutes {
		route.Name("detail")
		result = append(result, route)
	}

	listroutes := v.ListView.GetRoutes()
	for _, route := range listroutes {
		route.Name("list")
		result = append(result, route)
	}

	return result
}

type SlugViewSet struct {
	ListView
	SlugDetailView
}

/*
Before is blank implementation for ViewSet
*/
func (v SlugViewSet) Before(w http.ResponseWriter, r *http.Request) error {
	return nil
}

/*
GetRoutes returns combination of list and detail routes
*/
func (v SlugViewSet) GetRoutes() Routes {

	result := Routes{}

	detailroutes := v.SlugDetailView.GetRoutes()
	for _, route := range detailroutes {
		route.Name("detail")
		result = append(result, route)
	}

	listroutes := v.ListView.GetRoutes()
	for _, route := range listroutes {
		route.Name("list")
		result = append(result, route)
	}

	return result
}
