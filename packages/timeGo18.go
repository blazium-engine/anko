//go:build go1.8
// +build go1.8

package packages

import (
	"reflect"
	"time"

	"github.com/blazium-engine/anko/env"
)

func timeGo18() {
	env.Packages["time"]["Until"] = reflect.ValueOf(time.Until)
}
