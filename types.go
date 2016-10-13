/*
All interfaces used in classy module
*/
package classy

import (
	"net/http"
	"reflect"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

var (
	// List of available (supported) http methods. You can extend this with new methods
	AVAILABLE_METHODS = []Method{"GET", "POST", "PUT", "PATCH", "OPTIONS", "TRACE", "HEAD"}
)

/*
View

Interface for ClassyView, structs will be so-called class based view.
*/
type Viewer interface {
	Before(w http.ResponseWriter, r *http.Request) error

	// Return used route map.
	GetRoutes() Routes
}

/*
Classy is struct that wraps classy view and can register views to gorilla mux
*/
type Classy interface {
	Register(*mux.Router, ...string) []*mux.Route

	Use(middlewares ...alice.Constructor) Classy
}

/*
GetFuncName returns function name (primarily for logging reasons)
*/
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
