package main

import (
	"hidlikewindows/devices"
	"hidlikewindows/epp"

	//"bufio"

	"regexp"
	"sync"
	"time"

	evdev "github.com/gvalkov/golang-evdev"
	"github.com/radovskyb/watcher"
	"github.com/sirupsen/logrus"
)

type Host struct {
	Debug        bool
	DevicesMap   sync.Map
	InputEventCH chan *evdev.InputEvent

	keyboard       *devices.KeyboardHID
	keyboardReport *devices.KeyboardReport
	mouse          *devices.MouseHID
	mouseReport    *devices.MouseReport
	//OutputEventCH  chan *evdev.InputEvent
	NeedReleaseGUI bool
	Config         *HLWConfig
	X              int32
	Y              int32
}

func NewHost() (obj *Host, e error) {
	obj = &Host{}
	obj.InputEventCH = make(chan *evdev.InputEvent, 100)
	if e != nil {
		return
	}

	//obj.OutputEventCH = make(chan *evdev.InputEvent, 3)
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

	obj.Config, e = LoadHLWConfig()
	if e != nil {
		return
	}
	logrus.Info(obj.Config)
	return obj, nil
}

func (obj *Host) Run() error {
	w := watcher.New()
	defer w.Close()
	//w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Remove)
	w.AddFilterHook(watcher.RegexFilterHook(regexp.MustCompile("/dev/input/event.*"), true))
	if err := w.Add("/dev/input/"); err != nil {
		return err
	}
	go func() {
		for {
			select {
			case event := <-w.Event:
				if event.Op == watcher.Remove {
					obj.DevicesMap.Delete(event.Path)
				}
				if event.Op == watcher.Create {
					if device, e := evdev.Open(event.Path); e == nil {
						device.Grab()
						obj.DevicesMap.Store(event.Path, OpenInputDevice(device, obj.InputEventCH))
					}

				}
			case <-w.Closed:
				return
			}
		}
	}()
	inputDevices, _ := evdev.ListInputDevices()
	for _, device := range inputDevices {
		device.Grab()
		obj.DevicesMap.Store(device.Fn, OpenInputDevice(device, obj.InputEventCH))
	}
	err := w.Start(time.Millisecond * 100)
	return err
}

func (obj *Host) ProcessEvents() {

	for {

		select {

		case e := <-obj.InputEventCH:
			{
				//logrus.Info(e)

				if e.Type == evdev.EV_KEY {
					if e.Value == 2 { //repeat key
						continue
					}

					switch e.Code {
					//mouse button
					case evdev.BTN_LEFT, evdev.BTN_RIGHT, evdev.BTN_MIDDLE,
						evdev.BTN_SIDE, evdev.BTN_EXTRA, evdev.BTN_FORWARD, evdev.BTN_BACK:
						{
							if scancode, ok := MOUSEMAP[e.Code]; ok {
								if obj.ProcessMouseMap(scancode, e.Value) {
									break
								}
								if b := obj.mouseReport.OnButton(scancode, e.Value); b != nil {
									obj.mouse.Write(b)
								}
							}
							break
						}

					case evdev.KEY_LEFTCTRL, evdev.KEY_LEFTSHIFT, evdev.KEY_LEFTALT, evdev.KEY_LEFTMETA,
						evdev.KEY_RIGHTCTRL, evdev.KEY_RIGHTSHIFT, evdev.KEY_RIGHTALT, evdev.KEY_RIGHTMETA:
						{
							if scancode, ok := MODMAP[e.Code]; ok {
								if b := obj.keyboardReport.OnMod(scancode, e.Value); b != nil {
									obj.keyboard.Write(b)
								}
							}
							break
						}

					default:
						{
							logrus.Info(e)
							if scancode, ok := HIDMAP[e.Code]; ok {
								if obj.ProcessKeyboardMap(scancode, e.Value) {
									break
								}
								if b := obj.keyboardReport.OnKey(scancode, e.Value); b != nil {
									obj.keyboard.Write(b)
								}
							}
							break
						}
					}

				} else if e.Type == evdev.EV_REL {
					//mouse move ,scroll
					if e.Code == evdev.REL_X || e.Code == evdev.REL_Y {
						if e.Code == evdev.REL_X {
							v, _ := epp.Apply(obj.Config.EPPFactor, int(e.Value), 0)
							e.Value = int32(v)

						} else if e.Code == evdev.REL_Y {
							_, v := epp.Apply(obj.Config.EPPFactor, 0, int(e.Value))
							e.Value = int32(v)
						}
						if b := obj.mouseReport.OnMove(e.Code, e.Value); b != nil {
							if e := obj.mouse.Write(b); e != nil {
								logrus.Error("write to mouse hid: ", e)
							}
						}
					} else if e.Code == evdev.REL_WHEEL {
						for index := 0; index < obj.Config.ScrollSpeed; index++ {
							if b := obj.mouseReport.OnWheel(e.Value); b != nil {
								obj.mouse.Write(b)
							}
						}
					}

				}
			}

		case ret := <-obj.keyboard.ReadCh():
			{
				_ = ret
				//obj.serialDevice.Port.Write(ret)
			}
			//default:
		}

	}

	logrus.Info("HLW Existed!")

}
