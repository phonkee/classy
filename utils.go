package classy

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/serenize/snaker"
)

var (
	// regexp pattern to remove "view" occurences in word with trailing/leading underscores
	REMOVE_VIEW_PATTERN, _ = regexp.Compile("(?i)([_]*view[_]*)")
)

/*
IsEnabledOption returns whether wildcard option is enabled
*/
func IsEnabledOption(option []bool) bool {
	return len(option) > 0 && option[0]
}

/*
GetStructName returns struct name in snake case
*/
func GetStructName(target interface{}, snake ...bool) (result string) {
	fullname := fmt.Sprintf("%T", target)
	splitted := strings.Split(fullname, ".")
	sl := len(splitted)

	result = splitted[sl-1]

	if IsEnabledOption(snake) {
		result = snaker.CamelToSnake(result)
	}

	return
}

/*
GetViewName returns structname with stripped "view"
*/
func GetViewName(target interface{}, snake ...bool) (result string) {
	structname := GetStructName(target, snake...)
	result = REMOVE_VIEW_PATTERN.ReplaceAllString(structname, "")

	if len(result) == 0 {
		result = structname
	}
	return
}
