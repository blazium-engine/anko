//go:build go1.10
// +build go1.10

package packages

import (
	"reflect"
	"strings"

	"github.com/blazium-engine/anko/env"
)

func stringsGo110() {
	env.PackageTypes["strings"] = map[string]reflect.Type{
		"Builder": reflect.TypeOf(strings.Builder{}),
	}
}
