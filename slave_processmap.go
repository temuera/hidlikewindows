package main

import "github.com/sirupsen/logrus"

var CurrentMouseMap *KeyMap

func (obj *Host) ProcessMouseMap(scancode byte, downup int32) bool {

	keyboardModStatus := obj.keyboardReport.GetModStatus()
	mouseButtonStatus := obj.mouseReport.GetButtonStatus()

	mouseButtonStatus |= scancode
	if km, existed := obj.Config.MouseMaps[[2]byte{keyboardModStatus, mouseButtonStatus}]; existed {
		if downup == 1 { //button down
			logrus.Info(keyboardModStatus, "\t", mouseButtonStatus, "\tMouse Down:", km.OutputString())

			//clean mod
			obj.keyboardReport.OnMod(km.InputMod, 0)
			//down new mod and key
			obj.keyboardReport.OnMod(km.OutputMod, 1)
			if b := obj.keyboardReport.OnKey(km.OutputKey, 1); b != nil {
				obj.keyboard.Write(b)
			}
			//up new mod and key
			obj.keyboardReport.OnMod(km.OutputMod, 0)
			if b := obj.keyboardReport.OnKey(km.OutputKey, 0); b != nil {
				obj.keyboard.Write(b)
			}
			//restore mod
			obj.keyboardReport.OnMod(km.InputMod, 1)

		} else {
			logrus.Info(keyboardModStatus, "\t", mouseButtonStatus, "\tMouse Up:", km.OutputString())
		}
		return true
	}

	return false
}

func (obj *Host) ProcessKeyboardMap(scancode byte, downup int32) bool {
	keyboardModStatus := obj.keyboardReport.GetModStatus()
	if km, existed := obj.Config.KeyMaps[[2]byte{keyboardModStatus, scancode}]; existed {
		if downup == 1 {
			logrus.Info(keyboardModStatus, "\t", scancode, "\tKey Down:", km.OutputString())
			//clean mod
			obj.keyboardReport.OnMod(km.InputMod, 0)
			//down new mod and key
			obj.keyboardReport.OnMod(km.OutputMod, 1)
			if b := obj.keyboardReport.OnKey(km.OutputKey, 1); b != nil {
				obj.keyboard.Write(b)
			}
			//up new mod and key
			obj.keyboardReport.OnMod(km.OutputMod, 0)
			if b := obj.keyboardReport.OnKey(km.OutputKey, 0); b != nil {
				obj.keyboard.Write(b)
			}
			//restore mod
			obj.keyboardReport.OnMod(km.InputMod, 1)

		} else {
			logrus.Info(keyboardModStatus, "\t", scancode, "\tKey Up:", km.OutputString())
		}
		return true
	}
	return false
}
