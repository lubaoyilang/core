package format

import (
	"strconv"
)

// int -> string
func Itos(param int) string {
	var retStr string
	retStr = strconv.Itoa(param)
	return retStr
}

// int8 -> string
func I8tos(param int8) string {
	var retStr string
	retStr = strconv.Itoa(int(param))
	return retStr
}

// int16 -> string
func I16tos(param int16) string {
	var retStr string
	retStr = strconv.FormatInt(int64(param), 10)
	return retStr
}

// int32 -> string
func I32tos(param int32) string {
	var retStr string
	retStr = strconv.FormatInt(int64(param), 10)
	return retStr
}

// int64 -> string
func I64tos(param int64) string {
	var retStr string
	retStr = strconv.FormatInt(param, 10)
	return retStr
}

// uint -> string
func Uitos(param uint) string {
	var retStr string
	retStr = strconv.FormatUint(uint64(param), 10)
	return retStr
}

// uint8 -> string
func Ui8tos(param uint8) string {
	var retStr string
	retStr = strconv.FormatUint(uint64(param), 10)
	return retStr
}

// uint16 -> string
func Ui16tos(param uint16) string {
	var retStr string
	retStr = strconv.FormatUint(uint64(param), 10)
	return retStr
}

// uint32 -> string
func Ui32tos(param uint32) string {
	var retStr string
	retStr = strconv.FormatUint(uint64(param), 10)
	return retStr
}

// uint64 -> string
func Ui64tos(param uint64) string {
	var retStr string
	retStr = strconv.FormatUint(param, 10)
	return retStr
}

// string -> int
func Stoi(param string) int {
	var retInt int
	var err error
	retInt, err = strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	return retInt
}

// string -> int8
func Stoi8(param string) int8 {
	var retInt int8
	var retInt64 int64
	var err error
	retInt64, err = strconv.ParseInt(param, 10, 64)
	retInt = int8(retInt64)
	if err != nil {
		panic(err)
	}
	return retInt
}

// string -> int16
func Stoi16(param string) int16 {
	var retInt int16
	var retInt64 int64
	var err error
	retInt64, err = strconv.ParseInt(param, 10, 64)
	retInt = int16(retInt64)
	if err != nil {
		panic(err)
	}
	return retInt
}

// string -> int32
func Stoi32(param string) int32 {
	var retInt int32
	var retInt64 int64
	var err error
	retInt64, err = strconv.ParseInt(param, 10, 64)
	retInt = int32(retInt64)
	if err != nil {
		panic(err)
	}
	return retInt
}

// string -> int64
func Stoi64(param string) int64 {
	var retInt int64
	var err error
	retInt, err = strconv.ParseInt(param, 10, 64)
	if err != nil {
		panic(err)
	}
	return retInt
}

// string -> uint64
func Stoui64(param string) uint64 {
	var retUint uint64
	var err error
	retUint, err = strconv.ParseUint(param, 10, 64)
	if err != nil {
		panic(err)
	}
	return retUint
}
