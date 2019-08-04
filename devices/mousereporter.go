package devices

import (
	"fmt"

	evdev "github.com/gvalkov/golang-evdev"
)

type MouseReport struct {
	buff [6]byte
}

func (obj *MouseReport) String() string {

	return fmt.Sprintf("==>m: %+v", obj.buff)

}

func NewMouseReport() *MouseReport {
	obj := &MouseReport{}
	//obj.buff = [6]byte{0, 0, 0, 0, 0, 0}
	//byte[0] button ,uint16 x[1,2],uint16 y[3,4],byte[5] wheel, total 6
	return obj
}

func (obj *MouseReport) OnButton(scancode byte, value int32) []byte {
	obj.buff[1], obj.buff[2], obj.buff[3], obj.buff[4], obj.buff[5] = 0, 0, 0, 0, 0

	if value == 1 {
		obj.buff[0] |= scancode
	} else {
		obj.buff[0] &^= scancode
	}
	//if Debug {
	//logrus.Info(obj)
	//}
	return obj.buff[:]
}
func (obj *MouseReport) OnMove(evcode uint16, value int32) []byte {

	obj.buff[1], obj.buff[2], obj.buff[3], obj.buff[4], obj.buff[5] = 0, 0, 0, 0, 0
	if evcode == evdev.REL_X {
		obj.buff[1] = byte(value)
		obj.buff[2] = byte(value >> 8)
	} else if evcode == evdev.REL_Y {
		obj.buff[3] = byte(value)
		obj.buff[4] = byte(value >> 8)
	}
	return obj.buff[:]
}

func (obj *MouseReport) OnWheel(value int32) []byte {
	obj.buff[1], obj.buff[2], obj.buff[3], obj.buff[4], obj.buff[5] = 0, 0, 0, 0, 0
	obj.buff[5] = byte(value)
	//if Debug {
	//	logrus.Info(obj)
	//}
	return obj.buff[:]

}

func (obj *MouseReport) CheckButtonStatus(scancode byte) bool {
	return obj.buff[0]&byte(scancode) == scancode
}
func (obj *MouseReport) GetButtonStatus() byte {
	return obj.buff[0]
}
