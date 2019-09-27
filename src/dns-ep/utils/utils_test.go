package utils

import "testing"

func TestStrConvFloat(t *testing.T) {

	x, y, z, vel, err := StrConvFloat("hi", "10", "20", "30.111")
	if err == nil {
		t.Errorf("Test Failed for invalid input. x %+v, y %+v, z %+v, vel %+v, err %+v", x, y, z, vel, err)
	}
	x, y, z, vel, err = StrConvFloat("30.433", "yTest", "20", "30.111")
	if err == nil {
		t.Errorf("Test Failed for invalid input. x %+v, y %+v, z %+v, vel %+v, err %+v", x, y, z, vel, err)
	}
	x, y, z, vel, err = StrConvFloat("30.433", "1111.444", "zTest", "30.111")
	if err == nil {
		t.Errorf("Test Failed for invalid input. x %+v, y %+v, z %+v, vel %+v, err %+v", x, y, z, vel, err)
	}
	x, y, z, vel, err = StrConvFloat("30.433", "344.3223", "131331", "velTest")
	if err == nil {
		t.Errorf("Test Failed for invalid input. x %+v, y %+v, z %+v, vel %+v, err %+v", x, y, z, vel, err)
	}

	x, y, z, vel, err = StrConvFloat("30.433", "344.3223", "131331.22999999999999999", "20.22")
	if err != nil {
		t.Errorf("Test Failed for invalid input. x %+v, y %+v, z %+v, vel %+v, err %+v", x, y, z, vel, err)
	}
	if x != 30.433 && y != 344.3223 && z != 131331.22999999999999999 && vel != 20.22 {
		t.Errorf("Test Failed. Can't convert Str to float64 x %+v, y %+v, z %+v, vel %+v, err %+v", x, y, z, vel, err)
	}
}
