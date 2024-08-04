package main

import (
	"github.com/Bookshelf-Writer/esptool-modul/lib/output"
	"os"
)

//###########################################################//

type MethodObj struct {
	log    *output.LogObj
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

////

const (
	MethodRequiredParameterMissing = "Required parameter missing"
	MethodDevNotAvailable          = "Device not available"
	MethodFileNotFound             = "File not found"
)
