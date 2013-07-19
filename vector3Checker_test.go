package math

import (
	. "launchpad.net/gocheck"
)

var Vector3Check = &Vector3Checker{}

type Vector3Checker struct{}

func (checker *Vector3Checker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "Vector3Checker", Params: []string{"obtained", "expected"}}
}

func (checker *Vector3Checker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var v1, v2 Vector3
	var ok bool

	v1, ok = (params[0]).(Vector3)
	if ok == false {
		return false, "Param[0] not a Vector3 type"
	}
	v2, ok = (params[1]).(Vector3)
	if ok == false {
		return false, "Param[1] not a Vector3 type"
	}

	return NearlyEqualFloat32(v1.X, v2.X) && NearlyEqualFloat32(v1.Y, v2.Y) && NearlyEqualFloat32(v1.Z, v2.Z), ""
}
