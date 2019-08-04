package devices

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type MouseHID struct {
	Path      string
	File      *os.File
	readSetup sync.Once
	readErr   error
	readCh    chan []byte
}

func OpenMouseHID(path string) (obj *MouseHID, e error) {
	obj = &MouseHID{}
	obj.Path = path

	logrus.Info("Open mouse hid: " + obj.Path + " for read write data.")
	obj.File, e = os.OpenFile(obj.Path, os.O_APPEND|os.O_RDWR, 0) //os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		logrus.Error("Can not open "+obj.Path+" with error :", e)
		return nil, e
	}
	logrus.Info("Mouse HID file ready...")
	return obj, nil
}

func (obj *MouseHID) Close() {
	defer obj.File.Close()
}

func (obj *MouseHID) ReadCh() <-chan []byte {
	obj.readSetup.Do(func() {
		obj.readCh = make(chan []byte, 64)
		go obj.readThread()
	})
	return obj.readCh
}
func (obj *MouseHID) readThread() {
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
func (obj *MouseHID) Write(data []byte) error {
	_, err := obj.File.Write(data)
	return err

	//return nil
}
