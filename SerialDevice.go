package hidlikewindows

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type SerialDevice struct {
	Path string
	Port *serial.Port

	readSetup sync.Once
	readErr   error
	readCh    chan []byte
}

func OpenSerialPort(path string) (obj *SerialDevice, e error) {
	obj = &SerialDevice{}
	if path == "" {
		path = "/dev/ttyAMA0"
	}
	obj.Path = path
	Baud := 4000000
	//logrus.Info("Open serial port: " + obj.Path + ".")
	obj.Port, e = serial.OpenPort(&serial.Config{Name: obj.Path, Baud: Baud})
	if e != nil {
		logrus.Error("Can not open serial port with error :", e)
		return nil, e
	}
	logrus.WithFields(logrus.Fields{"Baud": Baud, "Port": path}).Info("Serialport ready:")
	return obj, nil
}
