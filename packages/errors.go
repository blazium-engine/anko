package packages

import (
	"errors"
	"reflect"

	"github.com/blazium-engine/anko/env"
)

func init() {
	env.Packages["errors"] = map[string]reflect.Value{
		"New": reflect.ValueOf(errors.New),
	}
}
