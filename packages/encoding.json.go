package packages

import (
	"encoding/json"
	"reflect"

	"github.com/blazium-engine/anko/env"
)

func init() {
	env.Packages["encoding/json"] = map[string]reflect.Value{
		"Marshal":   reflect.ValueOf(json.Marshal),
		"Unmarshal": reflect.ValueOf(json.Unmarshal),
	}
}
