package math

import (
	. "launchpad.net/gocheck"
)

var BoundingBoxCheck = &BoundingBoxChecker{}

type BoundingBoxChecker struct{}

func (checker *BoundingBoxChecker) Info() *CheckerInfo {
	return &CheckerInfo{Name: "BoundingBoxChecker", Params: []string{"obtained", "expected"}}
}

func (checker *BoundingBoxChecker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Param length not 2"
	}
	var v1, v2 *BoundingBox
	var ok bool

	v1, ok = (params[0]).(*BoundingBox)
	if ok == false {
		return false, "Param[0] not a *BoundingBoxtype"
	}
	v2, ok = (params[1]).(*BoundingBox)
	if ok == false {
		return false, "Param[1] not a *BoundingBox type"
	}

	return Vector3NearlyEqual(v1.Min, v2.Min) && Vector3NearlyEqual(v1.Max, v2.Max), ""
}
