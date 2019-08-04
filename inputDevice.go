package main

import (
	evdev "github.com/gvalkov/golang-evdev"
	"github.com/sirupsen/logrus"
)

type InputDevice struct {
	Path string
	Name string
	Dev  *evdev.InputDevice

	outCH chan *evdev.InputEvent
}

func OpenInputDevice(dev *evdev.InputDevice, outch chan *evdev.InputEvent) (obj *InputDevice) {
	obj = &InputDevice{}
	obj.Path = dev.Fn
	obj.Name = dev.Name
	obj.outCH = outch
	obj.Dev = dev
	go obj.readLoop()
	logrus.Info("Input Plugin:", obj.Path, ",", obj.Name)
	return obj
}

func (obj *InputDevice) readLoop() {
	defer func() {
		defer obj.Dev.Release()
		logrus.Info("Input Removed:", obj.Path, ",", obj.Name)
	}()
	for {
		n, err := obj.Dev.ReadOne()
		if err != nil {
			return
		}
		obj.outCH <- n
	}
}
