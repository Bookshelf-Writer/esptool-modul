package main

import (
	"os"
)

//###########################################################//

type MethodObj struct {
	log *LogObj
}

func (obj *MethodObj) End() {
	os.Exit(0)
}

func (obj *MethodObj) EndInvalid() {
	os.Exit(1)
}
