//go:build !appengine
// +build !appengine

package packages

import (
	"os"
	"reflect"

	"github.com/blazium-engine/anko/env"
)

func osNotAppEngine() {
	env.Packages["os"]["Getppid"] = reflect.ValueOf(os.Getppid)
}
