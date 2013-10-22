package hello

// #include "hello.h"
import "C"

func Random() int {
	return int(C.rnd())
}
