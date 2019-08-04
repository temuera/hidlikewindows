package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type KeyMap struct {
	Id        [2]byte `json:"id"`
	Type      int     //0==keyboard ,1 =mouse
	InputMod  byte    `json:"inputmod"`
	InputKey  byte    `json:"inputkey"`
	OutputMod byte    `json:"outputmod"`
	OutputKey byte    `json:"outputkey"`
}

func (obj KeyMap) IdString() string {
	return hex.EncodeToString(obj.Id[0:])
}

func (obj KeyMap) InputString() string {
	ms := []string{}
	for i := 1; i <= 256; i *= 2 {
		if obj.InputMod&byte(i) == byte(i) {
			if n, existed := MODNAME[i]; existed {
				ms = append(ms, n)
			}
		}
	}
	ks := []string{}
	if obj.Type == 0 { //keyboard
		if n, existed := HIDNAME[int(obj.InputKey)]; existed {
			ks = append(ks, n)
		}
	} else {
		if n, existed := MOUSENAME[int(obj.InputKey)]; existed {
			ks = append(ks, n)
		}
	}
	if len(ms) > 0 {
		return strings.Join(ms, ",") + "+" + strings.Join(ks, ",")
	}
	return strings.Join(ks, ",")
}
func (obj KeyMap) OutputString() string {
	ms := []string{}
	for i := 1; i <= 256; i *= 2 {
		if obj.OutputMod&byte(i) == byte(i) {
			if n, existed := MODNAME[i]; existed {
				ms = append(ms, n)
			}
		}
	}
	ks := []string{}
	if n, existed := HIDNAME[int(obj.OutputKey)]; existed {
		ks = append(ks, n)
	}
	if len(ms) > 0 {
		return strings.Join(ms, ",") + "+" + strings.Join(ks, ",")
	}
	return strings.Join(ks, ",")
}

type HLWConfig struct {
	MouseSpeed  int                `json:"mousespeed"`
	ScrollSpeed int                `json:"scrollspeed"`
	KeyMaps     map[[2]byte]KeyMap `json:"keymaps"`
	MouseMaps   map[[2]byte]KeyMap `json:"mousemaps"`

	EPPFactor float64 `json:"-"`
}

var lock sync.Mutex

var configPath = CurrentDir() + "./hlwconfig.gob"

var Marshal = func(v interface{}) (io.Reader, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}
var Unmarshal = func(r io.Reader, v interface{}) error {
	dec := gob.NewDecoder(r)
	return dec.Decode(v)

}

func LoadHLWConfig() (cfg *HLWConfig, err error) {
	cfg = &HLWConfig{MouseSpeed: 20, ScrollSpeed: 3, KeyMaps: make(map[[2]byte]KeyMap), MouseMaps: make(map[[2]byte]KeyMap)}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		cfg.Save()
	}

	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = Unmarshal(f, cfg)
	if err == nil {
		cfg.EPPFactor = (96 / 150.0) * (float64(cfg.MouseSpeed) / 16.0)

		if cfg.KeyMaps == nil {
			cfg.KeyMaps = make(map[[2]byte]KeyMap)
		}
		if cfg.MouseMaps == nil {
			cfg.MouseMaps = make(map[[2]byte]KeyMap)
		}

	}
	return cfg, err
}
func (obj *HLWConfig) Save() error {

	obj.EPPFactor = (96 / 150.0) * (float64(obj.MouseSpeed) / 16.0)

	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(obj)
	if err != nil {
		logrus.Error(err)
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func (obj *HLWConfig) UpdateMouseSpeed(newSpeed int) {
	lock.Lock()
	defer lock.Unlock()
	obj.MouseSpeed = newSpeed
	obj.Save()
}
func (obj *HLWConfig) UpdateScrollSpeed(newSpeed int) {
	lock.Lock()
	defer lock.Unlock()
	obj.ScrollSpeed = newSpeed
	obj.Save()
}
func (obj *HLWConfig) InsertKeyMap(km KeyMap) (err error) {
	lock.Lock()
	defer lock.Unlock()
	km.Type = 0
	copy(km.Id[:], []byte{km.InputMod, km.InputKey})
	if _, ok := obj.KeyMaps[km.Id]; ok {
		err = errors.New("Key existed:" + km.InputString())
		return
	}
	obj.KeyMaps[km.Id] = km
	obj.Save()
	return nil
}
func (obj *HLWConfig) RemoveKeyMap(id [2]byte) {
	lock.Lock()
	defer lock.Unlock()
	delete(obj.KeyMaps, id)

	obj.Save()
}

func (obj *HLWConfig) InsertMouseMap(km KeyMap) (err error) {
	lock.Lock()
	defer lock.Unlock()
	km.Type = 1
	copy(km.Id[:], []byte{km.InputMod, km.InputKey})
	if _, ok := obj.MouseMaps[km.Id]; ok {
		err = errors.New("Key existed:" + km.InputString())
		return
	}
	obj.MouseMaps[km.Id] = km
	obj.Save()
	return nil
}
func (obj *HLWConfig) RemoveMouseMap(id [2]byte) {
	lock.Lock()
	defer lock.Unlock()
	delete(obj.MouseMaps, id)
	obj.Save()
}
