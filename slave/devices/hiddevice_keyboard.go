package devices

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type IHIDDevice interface {
}

type KeyboardHID struct {
	Path      string
	File      *os.File
	readSetup sync.Once
	readErr   error
	readCh    chan []byte
}

func OpenKeyboardHID(path string) (obj *KeyboardHID, e error) {
	obj = &KeyboardHID{}
	obj.Path = path

	logrus.Info("Open keyboard hid: " + obj.Path + " for read write data.")
	obj.File, e = os.OpenFile(obj.Path, os.O_APPEND|os.O_RDWR, 0) //os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		logrus.Error("Can not open "+obj.Path+" with error :", e)
		return nil, e
	}
	logrus.Info("Keyboard HID file ready...")
	return obj, nil
}

func (obj *KeyboardHID) Close() {
	defer obj.File.Close()
}

func (obj *KeyboardHID) ReadCh() <-chan []byte {
	obj.readSetup.Do(func() {
		obj.readCh = make(chan []byte, 64)
		go obj.readThread()
	})
	return obj.readCh
}
func (obj *KeyboardHID) readThread() {
	defer close(obj.readCh)

	for {
		buf := make([]byte, 64)
		n, err := obj.File.Read(buf)
		if err != nil {
			obj.readErr = err
			return
		}
		obj.readCh <- buf[:n]
	}

}
func (obj *KeyboardHID) Write(data []byte) error {
	_, err := obj.File.Write(data)
	return err

}
