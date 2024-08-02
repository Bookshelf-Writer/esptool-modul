package main

import "os"

//###########################################################//

type MethodObj struct {
	log    *LogObj
	isTest bool
}

func (obj *MethodObj) End() {
	if obj.isTest {
		return
	}
	os.Exit(0)
}

func (obj *MethodObj) EndInvalid() {
	if obj.isTest {
		return
	}
	os.Exit(1)
}
