package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type KeyMap struct {
	Id       string `json:"id"`
	InputMod []int  `json:"inputmod"`
	InputKey []int  `json:"inputkey"`

	OutputMod []int `json:"outputmod"`
	OutputKey []int `json:"outputkey"`
}

func (obj KeyMap) InputString() string {
	ms := []string{}
	for _, c := range obj.InputMod {
		if n, existed := MODNAME[c]; existed {
			ms = append(ms, n)
		}
	}
	ks := []string{}
	for _, c := range obj.InputKey {
		if n, existed := HIDNAME[c]; existed {
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
	for _, c := range obj.OutputMod {
		if n, existed := MODNAME[c]; existed {
			ms = append(ms, n)
		}
	}
	ks := []string{}
	for _, c := range obj.OutputKey {
		if n, existed := HIDNAME[c]; existed {
			ks = append(ks, n)
		}
	}
	if len(ms) > 0 {
		return strings.Join(ms, ",") + "+" + strings.Join(ks, ",")
	}
	return strings.Join(ks, ",")
}

type HLWConfig struct {
	MouseSpeed  int               `json:"mousespeed"`
	ScrollSpeed int               `json:scrollspeed`
	KeyMaps     map[string]KeyMap `json:"keymaps"`
}

var lock sync.Mutex

const configPath = "./hlwconfig.json"

var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func LoadHLWConfig() (cfg *HLWConfig, err error) {
	cfg = &HLWConfig{MouseSpeed: 20, ScrollSpeed: 3, KeyMaps: make(map[string]KeyMap)}
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
	return cfg, err
}
func (obj *HLWConfig) Save() error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(obj)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func (obj *HLWConfig) UpdateMouseSpeed(newSpeed int) {
	obj.MouseSpeed = newSpeed
	obj.Save()
}
func (obj *HLWConfig) UpdateScrollSpeed(newSpeed int) {
	obj.ScrollSpeed = newSpeed
	obj.Save()
}
func (obj *HLWConfig) InsertKeyMap(km KeyMap) (err error) {

	km.Id = fmt.Sprintf("%x", md5.Sum([]byte(km.InputString())))
	if _, ok := obj.KeyMaps[km.Id]; ok {
		err = errors.New("Key existed")
		return
	}
	obj.KeyMaps[km.Id] = km
	obj.Save()
	return nil
}

func (obj *HLWConfig) RemoveKeyMap(id string) {
	delete(obj.KeyMaps, id)
	obj.Save()
}
