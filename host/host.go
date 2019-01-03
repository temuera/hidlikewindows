package main

import (
	"bufio"
	"hidlikewindows"
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
	serialDevice *hidlikewindows.SerialDevice
	InputEventCH chan *evdev.InputEvent
}

func NewHost() (obj *Host, e error) {
	obj = &Host{}
	obj.InputEventCH = make(chan *evdev.InputEvent, 10)
	obj.serialDevice, e = hidlikewindows.OpenSerialPort("")
	if e != nil {
		return
	}

	return obj, nil
}

func (obj *Host) Run() error {
	go obj.readserial()
	go obj.send2serial()

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
		obj.DevicesMap.Store(device.Fn, OpenInputDevice(device, obj.InputEventCH))
	}

	err := w.Start(time.Millisecond * 100)
	return err
}

func (obj *Host) send2serial() {
	var sendBuff = []byte{0, 0, 0, 0, 0, 0}
	for {
		e := <-obj.InputEventCH

		if e.Type == evdev.EV_KEY {
			if e.Value == 2 { //repeat key
				continue
			}
			//keyboard event
			sendBuff = []byte{0, 0, 0, 0, 0, 0}
			sendBuff[0] = uint8(e.Type)
			sendBuff[1] = byte(uint16(e.Code))
			sendBuff[2] = byte(uint16(e.Code) >> 8)
			sendBuff[3] = byte(uint16(e.Value))
			sendBuff[4] = byte(uint16(e.Value) >> 8)
			sendBuff[5] = byte('\n')

			_, err := obj.serialDevice.Port.Write(sendBuff)
			if err != nil {
				logrus.Error(err)
			}
			if obj.Debug {
				name := ""
				if v, ok := evdev.KEY[int(e.Code)]; ok {
					name = v
				}
				logrus.WithFields(logrus.Fields{"[01]Type": e.Type, "[23]Code": e.Code, "[4567]" + name: e.Value}).Info(sendBuff)
			}

		} else if e.Type == evdev.EV_REL { //mouse event
			sendBuff = []byte{0, 0, 0, 0, 0, 0}
			sendBuff[0] = uint8(e.Type)
			sendBuff[1] = byte(uint16(e.Code))
			sendBuff[2] = byte(uint16(e.Code) >> 8)
			sendBuff[3] = byte(uint16(e.Value))
			sendBuff[4] = byte(uint16(e.Value) >> 8)
			sendBuff[5] = byte('\n')
			_, err := obj.serialDevice.Port.Write(sendBuff)
			if err != nil {
				logrus.Error(err)
			}
			if obj.Debug {
				name := ""
				if v, ok := evdev.REL[int(e.Code)]; ok {
					name = v
				}
				logrus.WithFields(logrus.Fields{"[01]Type": e.Type, "[23]Code": e.Code, "[4567]" + name: e.Value}).Info(sendBuff)
			}
		} //else {

		// 	if e.Type == 4 {
		// 		if devices.Debug {
		// 			//logrus.Infof("- MSC : %+v \t %s", e, evdev.MSC[int(e.Code)])
		// 		}
		// 	} else if e.Type == 0 {
		// 		//logrus.Infof("- SYN: %+v \t %s", e, evdev.SYN[int(e.Code)])
		// 	} else {
		// 		logrus.Infof("??? : %+v \t", e)
		// 	}
		// }

	}
}

func (obj *Host) readserial() {
	var serialReader = bufio.NewReader(obj.serialDevice.Port)
	for {
		d, e := serialReader.ReadByte()
		if e == nil {
			logrus.Infof("LED Change : %x", d)
		}
	}

}
