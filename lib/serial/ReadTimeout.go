package serial

import "time"

//###########################################################//

type ReadTimeoutObj struct {
	ss *SerialObj
}

func (obj *ReadTimeoutObj) Set(duration time.Duration) error {
	obj.ss.readTimeout = duration
	return obj.ss.serial.SetReadTimeout(obj.ss.readTimeout)
}

func (obj *ReadTimeoutObj) Get() time.Duration {
	return obj.ss.readTimeout
}
