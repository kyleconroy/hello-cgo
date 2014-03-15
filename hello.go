package main

/*
#include <stdlib.h>
#include "hello.h"
*/
import "C"

func main() {
	C.hello()
}
