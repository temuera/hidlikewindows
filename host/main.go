package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/kardianos/service"

	"github.com/sirupsen/logrus"
)

const APP_NAME = "hlw_host"
const APP_DisplayName = "HLW HOST"
const APP_Description = "github.com/dazhoudotnet/hidlikewindows"

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	svcConfig := &service.Config{
		Name:        APP_NAME,
		DisplayName: APP_DisplayName,
		Description: APP_Description,
	}
	prg := &program{}
	s, _ := service.New(prg, svcConfig)

	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err = s.Install()
			if err != nil {
				logrus.Errorln("Installation failed:", err)
				logrus.Exit(1)
				return
			}
			logrus.Infoln(s.String(), "Successful installation.")
		case "uninstall":
		case "remove":
			err = s.Uninstall()
			if err != nil {
				logrus.Errorln("Uninstall failed:", err)
				logrus.Exit(1)
				return
			}
			logrus.Infoln(s.String(), "Successfully removed.")
		case "start":
			err = s.Start()
			if err != nil {
				logrus.Errorln(s.String(), "Startup failed:", err)
				return
			}
			logrus.Infoln(s.String(), "Successful startup. pid=", os.Getpid())
		case "stop":
			err = s.Stop()
			if err != nil {
				logrus.Errorln(s.String(), "Stop failed:", err)
				return
			}
			logrus.Infoln(s.String(), "Successfully stopped.")
			return
		}
	} else {
		err := s.Run()
		if err != nil {
			logrus.Errorln("Failed to run:", s.String())
			logrus.Exit(1)
			return
		}
	}

}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {
	doWork()
}
func (p *program) Stop(s service.Service) error {
	return nil
}

func doWork() {
	defer func() {
		if e := recover(); e != nil {
			logrus.Errorln(APP_NAME, "WTF!!!")
			logrus.Errorln(e)
			logrus.Exit(1)
			return
		}
	}()
	runtime.GOMAXPROCS(runtime.NumCPU())

	{
		//systemctl disable getty@tty1.service
		cmd := exec.Command("/usr/sbin/service", "getty@tty1", "stop")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		if e := cmd.Run(); e != nil {
			logrus.Error(out.String())
			logrus.Error(e)
			logrus.Exit(1)
		}
	}

	{ //remap keyboard ignore ctrl+alt+del (reboot)
		cmd := exec.Command("/bin/loadkeys", CurrentDir()+"disablekeyboard.kmap")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		if e := cmd.Run(); e != nil {
			logrus.Error(out.String())
			logrus.Error(e)
			logrus.Exit(1)
		}
	}

	host, e := NewHost()
	if e != nil {
		logrus.Fatal(e)
	}
	go host.Run()
	logrus.Info("HLW Host started.")
	select {}
}

func CurrentDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir + string(os.PathSeparator)
}
