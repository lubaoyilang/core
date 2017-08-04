package format

import "strconv"

// bool -> string
func Btos(param bool) string {
	var retStr string
	retStr = strconv.FormatBool(param)
	return retStr
}

// string -> bool
func Stob(param string) bool {
	var retBool bool
	var err error
	retBool, err = strconv.ParseBool(param)
	if err != nil {
		panic(err)
	}
	return retBool
}
