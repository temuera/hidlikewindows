package devices

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type KeyboardReport struct {
	buff []byte
	KEYS []byte
	//ReportCH chan []byte
}

func (obj *KeyboardReport) String() string {

	return fmt.Sprintf("==>k: %d,0,%d|%d|%d|%d|%d|%d", obj.buff[0], obj.KEYS[0], obj.KEYS[1], obj.KEYS[2], obj.KEYS[3], obj.KEYS[4], obj.KEYS[5])

}

func NewKeyboardReport() *KeyboardReport {
	obj := &KeyboardReport{}
	obj.KEYS = []byte{0, 0, 0, 0, 0, 0}
	obj.buff = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	return obj
}

func (obj *KeyboardReport) BuildBuff() {
	obj.buff[2] = byte(obj.KEYS[0])
	obj.buff[3] = byte(obj.KEYS[1])
	obj.buff[4] = byte(obj.KEYS[2])
	obj.buff[5] = byte(obj.KEYS[3])
	obj.buff[6] = byte(obj.KEYS[4])
	obj.buff[7] = byte(obj.KEYS[5])
}

func (obj *KeyboardReport) OnKey(scancode byte, value int32) []byte {
	if value == 1 {
		free_slot := -1
		for i, k := range obj.KEYS {
			if k == scancode {
				return nil
			} else if k == 0 {
				if free_slot == -1 {
					free_slot = i
				}
			}

		}
		if free_slot == -1 {
			return nil
		}
		obj.KEYS[free_slot] = scancode

		obj.BuildBuff()
	} else if value == 0 {
		for i, k := range obj.KEYS {
			if k == scancode {
				obj.KEYS[i] = 0
			}
		}
		obj.BuildBuff()
	}

	//if Debug {
	//	logrus.Info(obj)
	//}

	return obj.buff

}
func (obj *KeyboardReport) OnMod(scancode byte, value int32) []byte {
	logrus.Info("press mod key", scancode, "\t", value)
	if value == 1 {
		obj.buff[0] |= byte(scancode)
	} else if value == 0 {
		obj.buff[0] &^= byte(scancode)
	}

	//if Debug {
	//	logrus.Info(obj)
	//}
	obj.BuildBuff()
	return obj.buff

}
func (obj *KeyboardReport) ModStatus(scancode byte) bool {
	return obj.buff[0]&byte(scancode) == scancode
}
