package devices

import (
	"fmt"

	evdev "github.com/gvalkov/golang-evdev"
)

type MouseReport struct {
	Store []byte
}

func (obj *MouseReport) String() string {

	return fmt.Sprintf("==>m: %+v", obj.Store)

}

func NewMouseReport() *MouseReport {
	obj := &MouseReport{}
	obj.Store = []byte{0, 0, 0, 0, 0, 0}
	//byte[0] button ,uint16 x[1,2],uint16 y[3,4],byte[5] wheel, total 6
	return obj
}

func (obj *MouseReport) OnButton(scancode byte, value int32) []byte {
	obj.Store[1], obj.Store[2], obj.Store[3], obj.Store[4], obj.Store[5] = 0, 0, 0, 0, 0

	if value == 1 {
		obj.Store[0] |= scancode
	} else {
		obj.Store[0] &^= scancode
	}
	//if Debug {
	//logrus.Info(obj)
	//}
	return obj.Store
}
func (obj *MouseReport) OnMove(evcode uint16, value int32) []byte {

	obj.Store[1], obj.Store[2], obj.Store[3], obj.Store[4], obj.Store[5] = 0, 0, 0, 0, 0
	if evcode == evdev.REL_X {
		obj.Store[1] = byte(value)
		obj.Store[2] = byte(value >> 8)
	} else if evcode == evdev.REL_Y {
		obj.Store[3] = byte(value)
		obj.Store[4] = byte(value >> 8)
	}
	return obj.Store
}

func (obj *MouseReport) OnWheel(value int32) []byte {
	obj.Store[1], obj.Store[2], obj.Store[3], obj.Store[4], obj.Store[5] = 0, 0, 0, 0, 0
	obj.Store[5] = byte(value)
	//if Debug {
	//	logrus.Info(obj)
	//}
	return obj.Store

}
