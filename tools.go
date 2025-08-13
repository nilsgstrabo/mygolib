package mygolib

import "strings"

func FooV1(arg string) string {
	return arg
}

func Bar(arg string) string {
	return strings.ToUpper(arg)
}
