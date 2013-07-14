package math

import (
	. "launchpad.net/gocheck"
)

var Matrix3Check = &Matrix4Checker{}

type Matrix3Checker struct{}

func (checker *Matrix3Checker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "Matrix3Checker", Params: []string{"obtained", "expected"}}
}

func (checker *Matrix3Checker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var m1, m2 *Matrix3
	var ok bool

	m1, ok = (params[0]).(*Matrix3)
	if ok == false {
		return false, "Param[0] not a *Matrix3 type"
	}
	m2, ok = (params[1]).(*Matrix3)
	if ok == false {
		return false, "Param[1] not a *Matrix3 type"
	}

	return Float32Equals(m1.M11, m2.M11) && Float32Equals(m1.M12, m2.M12) && Float32Equals(m1.M13, m2.M13) &&
		Float32Equals(m1.M21, m2.M21) && Float32Equals(m1.M22, m2.M22) && Float32Equals(m1.M23, m2.M23) &&
		Float32Equals(m1.M31, m2.M31) && Float32Equals(m1.M32, m2.M32) && Float32Equals(m1.M33, m2.M33), ""
}
