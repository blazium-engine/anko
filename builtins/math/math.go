// Package math implements math interface for anko script.
package math

import (
	"github.com/mattn/anko/vm"
	t "math"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("math")
	m.Define("Abs", reflect.ValueOf(t.Abs))
	m.Define("Acos", reflect.ValueOf(t.Acos))
	m.Define("Acosh", reflect.ValueOf(t.Acosh))
	m.Define("Asin", reflect.ValueOf(t.Asin))
	m.Define("Asinh", reflect.ValueOf(t.Asinh))
	m.Define("Atan", reflect.ValueOf(t.Atan))
	m.Define("Atan2", reflect.ValueOf(t.Atan2))
	m.Define("Atanh", reflect.ValueOf(t.Atanh))
	m.Define("Cbrt", reflect.ValueOf(t.Cbrt))
	m.Define("Ceil", reflect.ValueOf(t.Ceil))
	m.Define("Copysign", reflect.ValueOf(t.Copysign))
	m.Define("Cos", reflect.ValueOf(t.Cos))
	m.Define("Cosh", reflect.ValueOf(t.Cosh))
	m.Define("Dim", reflect.ValueOf(t.Dim))
	m.Define("Erf", reflect.ValueOf(t.Erf))
	m.Define("Erfc", reflect.ValueOf(t.Erfc))
	m.Define("Exp", reflect.ValueOf(t.Exp))
	m.Define("Exp2", reflect.ValueOf(t.Exp2))
	m.Define("Expm1", reflect.ValueOf(t.Expm1))
	m.Define("Float32bits", reflect.ValueOf(t.Float32bits))
	m.Define("Float32frombits", reflect.ValueOf(t.Float32frombits))
	m.Define("Float64bits", reflect.ValueOf(t.Float64bits))
	m.Define("Float64frombits", reflect.ValueOf(t.Float64frombits))
	m.Define("Floor", reflect.ValueOf(t.Floor))
	m.Define("Frexp", reflect.ValueOf(t.Frexp))
	m.Define("Gamma", reflect.ValueOf(t.Gamma))
	m.Define("Hypot", reflect.ValueOf(t.Hypot))
	m.Define("Ilogb", reflect.ValueOf(t.Ilogb))
	m.Define("Inf", reflect.ValueOf(t.Inf))
	m.Define("IsInf", reflect.ValueOf(t.IsInf))
	m.Define("IsNaN", reflect.ValueOf(t.IsNaN))
	m.Define("J0", reflect.ValueOf(t.J0))
	m.Define("J1", reflect.ValueOf(t.J1))
	m.Define("Jn", reflect.ValueOf(t.Jn))
	m.Define("Ldexp", reflect.ValueOf(t.Ldexp))
	m.Define("Lgamma", reflect.ValueOf(t.Lgamma))
	m.Define("Log", reflect.ValueOf(t.Log))
	m.Define("Log10", reflect.ValueOf(t.Log10))
	m.Define("Log1p", reflect.ValueOf(t.Log1p))
	m.Define("Log2", reflect.ValueOf(t.Log2))
	m.Define("Logb", reflect.ValueOf(t.Logb))
	m.Define("Max", reflect.ValueOf(t.Max))
	m.Define("Min", reflect.ValueOf(t.Min))
	m.Define("Mod", reflect.ValueOf(t.Mod))
	m.Define("Modf", reflect.ValueOf(t.Modf))
	m.Define("NaN", reflect.ValueOf(t.NaN))
	m.Define("Nextafter", reflect.ValueOf(t.Nextafter))
	m.Define("Pow", reflect.ValueOf(t.Pow))
	m.Define("Pow10", reflect.ValueOf(t.Pow10))
	m.Define("Remainder", reflect.ValueOf(t.Remainder))
	m.Define("Signbit", reflect.ValueOf(t.Signbit))
	m.Define("Sin", reflect.ValueOf(t.Sin))
	m.Define("Sincos", reflect.ValueOf(t.Sincos))
	m.Define("Sinh", reflect.ValueOf(t.Sinh))
	m.Define("Sqrt", reflect.ValueOf(t.Sqrt))
	m.Define("Tan", reflect.ValueOf(t.Tan))
	m.Define("Tanh", reflect.ValueOf(t.Tanh))
	m.Define("Trunc", reflect.ValueOf(t.Trunc))
	m.Define("Y0", reflect.ValueOf(t.Y0))
	m.Define("Y1", reflect.ValueOf(t.Y1))
	m.Define("Yn", reflect.ValueOf(t.Yn))
	return m
}
