package math

import (
	. "launchpad.net/gocheck"
)

var EqualsFloat32 = &Float32EqualChecker{}

type Float32EqualChecker struct{}

func (checker *Float32EqualChecker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "Float32EqualChecker", Params: []string{"obtained", "expected"}}
}

func (checker *Float32EqualChecker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var f1, f2 float32
	var ok bool

	f1, ok = (params[0]).(float32)
	if ok == false {
		return false, "Param[0] not a float32 type"
	}
	f2, ok = (params[1]).(float32)
	if ok == false {
		return false, "Param[1] not a float32 type"
	}

	return NearlyEqualFloat32(f1, f2), ""
}

func NearlyEqualFloat32(a, b float32) bool {
	if a == b {
		return true
	}
	diff := Abs(a - b)
	return diff/(Abs(a)+Abs(b)) < 0.0001
}
