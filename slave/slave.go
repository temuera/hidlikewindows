package main

import (
	"bufio"
	"encoding/binary"
	"hidlikewindows"
	"hidlikewindows/slave/devices"
	"hidlikewindows/slave/epp"

	"github.com/sirupsen/logrus"

	evdev "github.com/gvalkov/golang-evdev"
)

type Slave struct {
	keyboard       *devices.KeyboardHID
	keyboardReport *devices.KeyboardReport
	mouse          *devices.MouseHID
	mouseReport    *devices.MouseReport

	serialDevice   *hidlikewindows.SerialDevice
	OutputEventCH  chan *evdev.InputEvent
	NeedReleaseGUI bool

	Config *HLWConfig

	X int32
	Y int32
}

func NewSlave() (obj *Slave, e error) {
	obj = &Slave{}

	obj.OutputEventCH = make(chan *evdev.InputEvent, 3)

	obj.keyboard, e = devices.OpenKeyboardHID("/dev/hidg0")
	if e != nil {
		return
	}
	obj.keyboardReport = devices.NewKeyboardReport()

	obj.mouse, e = devices.OpenMouseHID("/dev/hidg1")
	if e != nil {
		return
	}
	obj.mouseReport = devices.NewMouseReport()

	obj.serialDevice, e = hidlikewindows.OpenSerialPort("")
	if e != nil {
		return
	}

	obj.Config, e = LoadHLWConfig()
	if e != nil {
		return
	}
	logrus.Info(obj.Config)
	return obj, nil

}

const HOST_PKT_LEN int = 6

func (obj *Slave) Run() {

	go obj.StartWebInterface()

	var serialReader = bufio.NewReader(obj.serialDevice.Port)
	go func() {
		tmpBuff := []byte{}
		for {
			bb, err := serialReader.ReadBytes('\n')
			if err != nil {
				logrus.Error(err)
			}

			l1, l2 := len(tmpBuff), len(bb)
			if totalLength := l1 + l2; totalLength < HOST_PKT_LEN {
				tmpBuff = append(tmpBuff, bb...) //如果长度不满足,则存着继续
				// if devices.Debug {
				// 	logrus.WithFields(logrus.Fields{"PktLen": l2, "CacheLen": totalLength}).Warning(tmpBuff, " continue")
				// }
				continue

			} else if totalLength > HOST_PKT_LEN {
				if l1 > 0 {
					logrus.WithFields(logrus.Fields{"L1": l1}).Error(tmpBuff, " dropped")
					tmpBuff = nil
				}
				if l2 != HOST_PKT_LEN {
					logrus.WithFields(logrus.Fields{"L2": l2}).Error(bb, " dropped")
					continue
				}
			}

			tmpBuff = append(tmpBuff, bb...)

			_type := uint8(tmpBuff[0])
			_code := binary.LittleEndian.Uint16(tmpBuff[1:])
			_value := int16(binary.LittleEndian.Uint16(tmpBuff[3:]))

			e := &evdev.InputEvent{Type: uint16(_type), Code: _code, Value: int32(_value)}
			// if devices.Debug {
			// 	logrus.WithFields(logrus.Fields{"[01]Type": e.Type, "[23]Code": e.Code, "[4567]Value": e.Value}).Info(tmpBuff)
			// }
			tmpBuff = nil

			obj.OutputEventCH <- e
		}
	}()

	for {
		select {

		case e := <-obj.OutputEventCH:
			{
				if e.Type == evdev.EV_KEY {
					switch e.Code {
					case evdev.KEY_LEFTCTRL, evdev.KEY_LEFTSHIFT, evdev.KEY_LEFTALT, evdev.KEY_LEFTMETA,
						evdev.KEY_RIGHTCTRL, evdev.KEY_RIGHTSHIFT, evdev.KEY_RIGHTALT, evdev.KEY_RIGHTMETA:
						{
							// if scancode, ok := MODMAP[e.Code]; ok {
							// 	obj.keyboardReport.OnMod(KEY_LEFTMETA, 0)
							// }
							// if scancode, ok := MODMAP[e.Code]; ok {
							// 	if e.Value == 0 && obj.NeedReleaseGUI && (scancode == KEY_LEFTCTRL || scancode == KEY_LEFTALT) {
							// 		obj.keyboardReport.OnMod(KEY_LEFTMETA, 0)
							// 		obj.NeedReleaseGUI = false
							// 	}
							// 	if b := obj.keyboardReport.OnMod(scancode, e.Value); b != nil {
							// 		obj.keyboard.Write(b)
							// 	}
							// } else {
							// 	logrus.Info("not found mod ", e.Code)
							// }
							break
						}
					case evdev.BTN_LEFT, evdev.BTN_RIGHT, evdev.BTN_MIDDLE,
						evdev.BTN_SIDE, evdev.BTN_EXTRA, evdev.BTN_FORWARD, evdev.BTN_BACK:
						{
							if scancode, ok := MOUSEMAP[e.Code]; ok {
								if b := obj.mouseReport.OnButton(scancode, e.Value); b != nil {
									obj.mouse.Write(b)
								}
							}
							break
						}
					default:
						{

							//if e.Code==evdev.KEY_HOME

							if scancode, ok := HIDMAP[e.Code]; ok {
								if e.Value == 1 { // press key
									if obj.keyboardReport.ModStatus(KEY_LEFTCTRL) && scancode != KEY_SPACE {
										obj.keyboardReport.OnMod(KEY_LEFTCTRL, 0)
										obj.keyboardReport.OnMod(KEY_LEFTMETA, 1)
										obj.NeedReleaseGUI = true
									}
									if obj.keyboardReport.ModStatus(KEY_LEFTALT) && scancode == KEY_TAB {
										obj.keyboardReport.OnMod(KEY_LEFTALT, 0)
										obj.keyboardReport.OnMod(KEY_LEFTMETA, 1)
										obj.NeedReleaseGUI = true
									}
								}

								if !obj.NeedReleaseGUI && (scancode == KEY_HOME) { //one key press

									obj.keyboardReport.OnMod(KEY_LEFTMETA, e.Value)
									if b := obj.keyboardReport.OnKey(KEY_LEFT, e.Value); b != nil {
										obj.keyboard.Write(b)
									}
									break
								} else if !obj.NeedReleaseGUI && (scancode == KEY_END) {
									obj.keyboardReport.OnMod(KEY_LEFTMETA, e.Value)
									if b := obj.keyboardReport.OnKey(KEY_RIGHT, e.Value); b != nil {
										obj.keyboard.Write(b)
									}
									break
								} else if b := obj.keyboardReport.OnKey(scancode, e.Value); b != nil {
									obj.keyboard.Write(b)
								}

							}
							break
						}
					}

				} else if e.Type == evdev.EV_REL {
					//mouse move ,scroll
					if e.Code == evdev.REL_X || e.Code == evdev.REL_Y {
						//x,y:=0,0
						if e.Code == evdev.REL_X {
							v, _ := epp.Apply(float64(obj.Config.MouseSpeed), int(e.Value), 0)
							e.Value = int32(v)

						} else if e.Code == evdev.REL_Y {
							_, v := epp.Apply(float64(obj.Config.MouseSpeed), 0, int(e.Value))
							e.Value = int32(v)
						}

						if b := obj.mouseReport.OnMove(e.Code, e.Value); b != nil {
							if e := obj.mouse.Write(b); e != nil {
								logrus.Error("write to mouse hid: ", e)
							}
						}
					} else if e.Code == evdev.REL_WHEEL {
						if b := obj.mouseReport.OnWheel(e.Value); b != nil {
							obj.mouse.Write(b)
						}
					}

				}
			}

		case ret := <-obj.keyboard.ReadCh():
			{
				obj.serialDevice.Port.Write(ret)
			}
			//default:
		}

	}

}
