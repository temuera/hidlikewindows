package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func (obj *Slave) WebSettings(w http.ResponseWriter, r *http.Request) {

	action := r.FormValue("action")
	if action == "savekeymap" {

		km := KeyMap{}

		for _, v := range r.Form["modkey"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.InputMod = append(km.InputMod, n)
			}
		}
		for _, v := range r.Form["key"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.InputKey = append(km.InputKey, n)
			}
		}

		for _, v := range r.Form["modkey_out"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.OutputMod = append(km.OutputMod, n)
			}
		}
		for _, v := range r.Form["key_out"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.OutputKey = append(km.OutputKey, n)
			}
		}

		w.Header().Set("Content-Type", "application/json")

		err := obj.Config.InsertKeyMap(km)
		if err != nil {
			w.Write([]byte(`{"result":false,"msg":"` + err.Error() + `"}`))
		} else {
			w.Write([]byte(`{"result":true,"input":"` + km.InputString() + `","output":"` + km.OutputString() + `"}`))
		}

		return
	}
	if action == "removekeymap" {
		id := r.FormValue("id")
		obj.Config.RemoveKeyMap(id)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":true,"id":"` + id + `"}`))
		return
	}
	if action == "setspeed" {
		speed, _ := strconv.Atoi(r.FormValue("speed"))
		if speed != obj.Config.MouseSpeed && speed > 0 && speed <= 60 {
			obj.Config.UpdateMouseSpeed(speed)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":true}`))
		return
	}

	w.Header().Set("Content-Type", "text/html")

	t := template.New("settings.html")

	filename := "assets/settings.html"
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if b, e := ioutil.ReadFile(filename); e == nil {
			//w.Write(b)
			t.Parse(string(b))

		}
	} else {
		t.Parse(string(content_settings))
	}
	//write from file

	//w.Write(content_settings)
	t.Execute(w, obj.Config)

}

var content_settings = []byte(`
`)
