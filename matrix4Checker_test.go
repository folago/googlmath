package math

import (
	. "launchpad.net/gocheck"
)

var Matrix4Check = &Matrix4Checker{}

type Matrix4Checker struct{}

func (checker *Matrix4Checker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "Matrix4Checker", Params: []string{"obtained", "expected"}}
}

func (checker *Matrix4Checker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var m1, m2 *Matrix4
	var ok bool

	m1, ok = (params[0]).(*Matrix4)
	if ok == false {
		return false, "Param[0] not a *Matrix4 type"
	}
	m2, ok = (params[1]).(*Matrix4)
	if ok == false {
		return false, "Param[1] not a *Matrix4 type"
	}

	return Float32Equals(m1.M11, m2.M11) && Float32Equals(m1.M12, m2.M12) && Float32Equals(m1.M13, m2.M13) && Float32Equals(m1.M14, m2.M14) &&
		Float32Equals(m1.M21, m2.M21) && Float32Equals(m1.M22, m2.M22) && Float32Equals(m1.M23, m2.M23) && Float32Equals(m1.M24, m2.M24) &&
		Float32Equals(m1.M31, m2.M31) && Float32Equals(m1.M32, m2.M32) && Float32Equals(m1.M33, m2.M33) && Float32Equals(m1.M34, m2.M34) &&
		Float32Equals(m1.M41, m2.M41) && Float32Equals(m1.M42, m2.M42) && Float32Equals(m1.M43, m2.M43) && Float32Equals(m1.M44, m2.M44), ""
}

func Float32Equals(a, b float32) bool {
	if a == b {
		return true
	}
	return Floor(a+0.5) == Floor(b+0.5)
}
