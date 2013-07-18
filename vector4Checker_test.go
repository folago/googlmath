package math

import (
	. "launchpad.net/gocheck"
)

var Vector4Check = &Vector4Checker{}

type Vector4Checker struct{}

func (checker *Vector4Checker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "Vector4Checker", Params: []string{"obtained", "expected"}}
}

func (checker *Vector4Checker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var v1, v2 Vector4
	var ok bool

	v1, ok = (params[0]).(Vector4)
	if ok == false {
		return false, "Param[0] not a Vector4 type"
	}
	v2, ok = (params[1]).(Vector4)
	if ok == false {
		return false, "Param[1] not a Vector4 type"
	}

	return NearlyEqualFloat32(v1.X, v2.X) && NearlyEqualFloat32(v1.Y, v2.Y) && NearlyEqualFloat32(v1.Z, v2.Z) && NearlyEqualFloat32(v1.W, v2.W), ""
}
