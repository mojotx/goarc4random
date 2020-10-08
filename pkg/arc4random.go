package arc4random

/*
#include <stdlib.h>
*/
import "C"

func Arc4Random() uint32 {
	return uint32(C.arc4random())
}

// ub = upper bound
func Arc4RandomUniform(ub uint32) uint32 {
	return uint32(C.arc4random_uniform(C.uint(ub)))
}
