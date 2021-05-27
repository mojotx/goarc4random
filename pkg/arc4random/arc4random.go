package arc4random

/*
#include <stdlib.h>
*/
import "C"

type Arc4Random struct{}

func (c *Arc4Random) Arc4Random() uint32 {
	return uint32(C.arc4random())
}

// ub = upper bound
func (c *Arc4Random) Uniform(ub uint32) uint32 {
	return uint32(C.arc4random_uniform(C.uint(ub)))
}
