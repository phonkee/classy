package classy

import (
	"fmt"
	"net/http"
	"reflect"
	"sync"
)

/*
Method
alias to string (for future methods possibility)
*/
type Method string

/*
MethodMap map methods to aliases (custom method names)
*/
type MethodMap interface {

	// Add method aliases (view method names)
	Add(method Method, to ...Method) MethodMap

	// Bind returns bound methods
	Bind(Viewer) (map[Method]BindItem, error)

	// returns whether MethodMap contains http method (key of map) or alias
	Contains(method Method) bool

	// Delete given aliases (view methods) for http method. If no aliases given, flush is performed.
	Delete(method Method, to ...Method) MethodMap

	// Return data
	GetData() map[Method][]Method

	// Return methodset for given http method
	GetMethodSet(method Method) (MethodSet, bool)

	// Update udpates data from other method map
	Update(MethodMap) MethodMap
}

/*
MethodMap revamped
*/
func NewMethodMap() MethodMap {
	return &methodmap{
		storage: map[Method]MethodSet{},
		mutex:   &sync.RWMutex{},
	}
}

/*
MethodMap implementation
*/
type methodmap struct {
	storage map[Method]MethodSet
	mutex   *sync.RWMutex
}

/*
Add aliases to method
*/
func (m *methodmap) Add(method Method, aliases ...Method) MethodMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.storage[method]; !ok {
		m.storage[method] = NewMethodSet()
	}

	for _, alias := range aliases {
		m.storage[method].Add(alias)
	}

	return m
}

/*
BindItem is extracted handler from class based view along with name of method (method)
 */
type BindItem struct {
	method      Method
	handlerfunc http.HandlerFunc
}

/*
Bind methods.
*/
func (m *methodmap) Bind(view Viewer) (result map[Method]BindItem, err error) {

	result = map[Method]BindItem{}

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, method := range AVAILABLE_METHODS {

		var (
			ms MethodSet
			ok bool
		)

		if ms, ok = m.storage[method]; !ok {
			continue
		}

		for _, vm := range ms.ToSlice() {

			var (
				viewmethodvalue reflect.Value
				viewmethod      http.HandlerFunc
				okvm            bool
			)

			viewmethodvalue = reflect.ValueOf(view).MethodByName(string(vm))
			if !viewmethodvalue.IsValid() {
				continue
			}

			if viewmethod, okvm = viewmethodvalue.Interface().(func(w http.ResponseWriter, r *http.Request)); okvm {
				if _, isinresult := result[method]; isinresult {
					panic(fmt.Sprintf("http method %+v is already registered to %+v", method, ms))
				}

				// prepare handlerfunc that first calls Before method and then calls view method.
				final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// call before method that is required
					view.Before(w, r)

					// call actual viewmethod
					viewmethod(w, r)
				})

				// assign bind item
				result[method] = BindItem{method: vm, handlerfunc: final}
			}
		}
	}

	return
}

func (m *methodmap) Contains(method Method) (result bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, ok := m.storage[method]; ok {
		result = true
	}

	return
}

/*
Delete aliases from method
*/
func (m *methodmap) Delete(method Method, aliases ...Method) MethodMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.storage[method]; !ok {
		return m
	}

	if len(aliases) > 0 {
		for _, alias := range aliases {
			m.storage[method].Remove(alias)
		}
	} else {
		delete(m.storage, method)
	}

	return m
}

/*
GetData returns mapping data
*/
func (m *methodmap) GetData() (result map[Method][]Method) {
	result = map[Method][]Method{}

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for method, aliasesSet := range m.storage {
		result[method] = make([]Method, 0, aliasesSet.Cardinality())
		for _, alias := range aliasesSet.ToSlice() {
			result[method] = append(result[method], alias)
		}
	}

	return
}

/*
Update updates methodmap with other
@TODO: Check once more
*/
func (m *methodmap) Update(other MethodMap) MethodMap {
	for method, aliases := range other.GetData() {
		for _, alias := range aliases {
			m.Add(method, Method(alias))
		}
	}
	return m
}

func (m *methodmap) GetMethodSet(method Method) (result MethodSet, ok bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	result, ok = m.storage[method]
	return
}
