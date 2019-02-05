package main

import (
	"encoding/hex"
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

		for _, v := range r.Form["mod_in"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.InputMod |= byte(n) //append(km.InputMod, byte(n))
			}
		}
		if n, e := strconv.Atoi(r.FormValue("key_in")); e == nil && n > 0 {
			km.InputKey = byte(n) //append(km.InputKey, byte(n))
		}

		for _, v := range r.Form["mod_out"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.OutputMod |= byte(n) //= append(km.OutputMod, byte(n))
			}
		}

		if n, e := strconv.Atoi(r.FormValue("key_out")); e == nil && n > 0 {
			km.OutputKey = byte(n) //append(km.OutputKey, byte(n))
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

	if action == "savemousemap" {

		km := KeyMap{}

		for _, v := range r.Form["mod_in"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.InputMod |= byte(n) //append(km.InputMod, byte(n))
			}
		}
		if n, e := strconv.Atoi(r.FormValue("mouse_in")); e == nil && n > 0 {
			km.InputKey = byte(n) //append(km.InputKey, byte(n))
		}

		for _, v := range r.Form["mod_out"] {
			if n, e := strconv.Atoi(v); e == nil && n > 0 {
				km.OutputMod |= byte(n) //= append(km.OutputMod, byte(n))
			}
		}
		if n, e := strconv.Atoi(r.FormValue("key_out")); e == nil && n > 0 {
			km.OutputKey = byte(n) //append(km.OutputKey, byte(n))
		}

		w.Header().Set("Content-Type", "application/json")

		err := obj.Config.InsertMouseMap(km)
		if err != nil {
			w.Write([]byte(`{"result":false,"msg":"` + err.Error() + `"}`))
		} else {
			w.Write([]byte(`{"result":true,"input":"` + km.InputString() + `","output":"` + km.OutputString() + `"}`))
		}

		return
	}

	if action == "removekeymap" {
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		if b, e := hex.DecodeString(id); e == nil && len(b) == 2 {
			obj.Config.RemoveKeyMap([2]byte{b[0], b[1]})
			w.Write([]byte(`{"result":true,"id":"` + id + `"}`))
		} else {
			w.Write([]byte(`{"result":false,"id":"` + id + `"}`))
		}
		return
	}

	if action == "removemousemap" {
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		if b, e := hex.DecodeString(id); e == nil && len(b) == 2 {
			obj.Config.RemoveMouseMap([2]byte{b[0], b[1]})
			w.Write([]byte(`{"result":true,"id":"` + id + `"}`))
		} else {
			w.Write([]byte(`{"result":false,"id":"` + id + `"}`))
		}
		return
	}

	if action == "setmousespeed" {
		speed, _ := strconv.Atoi(r.FormValue("speed"))
		if speed != obj.Config.MouseSpeed && speed > 0 && speed <= 60 {
			obj.Config.UpdateMouseSpeed(speed)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":true}`))
		return
	}
	if action == "setscrollspeed" {
		speed, _ := strconv.Atoi(r.FormValue("speed"))
		if speed != obj.Config.ScrollSpeed && speed > 0 && speed <= 60 {
			obj.Config.UpdateScrollSpeed(speed)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":true}`))
		return
	}
	w.Header().Set("Content-Type", "text/html")

	t := template.New("settings.html")

	filename := CurrentDir() + "assets/settings.html"
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if b, e := ioutil.ReadFile(filename); e == nil {
			//w.Write(b)
			t.Parse(string(b))

		}
	} else {
		t.Parse(string(content_settings))
	}
	t.Execute(w, obj.Config)

}

var content_settings = []byte(`<!doctype html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>HID Like Windows</title>
    <link rel="stylesheet" href="/style.css">
    <script src="/jquery.js"></script>
</head>

<body>
    <div class="box">
        <ul class="nav">
            <li><a href="/" class="active">Settings</a></li>
            <li><a href="/win">Win Mouse</a></li>
            <li><a href="/nonwin">Non-Win Mouse</a></li>
        </ul>
        <div id="content" class="row content">
            <div style="margin: 20px 40px; background: #fff;padding:10px 50px;min-height: 500px">
                <h1 class="left">🖱️ Mouse Settings</h1>
                <div class="left">
                    <ul class="f">
                        <li class="l ">
                            Mouse Speed
                        </li>
                        <li class="r">
                            <input type="radio" name="mousespeed" value="10" {{if eq .MouseSpeed 10}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="12" {{if eq .MouseSpeed 12}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="14" {{if eq .MouseSpeed 14}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="16" {{if eq .MouseSpeed 16}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="18" {{if eq .MouseSpeed 18}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="20" {{if eq .MouseSpeed 20}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="22" {{if eq .MouseSpeed 22}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="24" {{if eq .MouseSpeed 24}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="26" {{if eq .MouseSpeed 26}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="28" {{if eq .MouseSpeed 28}} checked {{end}} />
                            <input type="radio" name="mousespeed" value="30" {{if eq .MouseSpeed 30}} checked {{end}} />
                        </li>
                        <li class="l ">
                            Scroll Speed
                        </li>
                        <li class="r">
                            <input type="radio" name="scrollspeed" value="1" {{if eq .ScrollSpeed 1}} checked {{end}} />
                            <input type="radio" name="scrollspeed" value="2" {{if eq .ScrollSpeed 2}} checked {{end}} />
                            <input type="radio" name="scrollspeed" value="3" {{if eq .ScrollSpeed 3}} checked {{end}} />
                            <input type="radio" name="scrollspeed" value="4" {{if eq .ScrollSpeed 4}} checked {{end}} />
                            <input type="radio" name="scrollspeed" value="5" {{if eq .ScrollSpeed 5}} checked {{end}} />
                            <input type="radio" name="scrollspeed" value="6" {{if eq .ScrollSpeed 6}} checked {{end}} />
                        </li>
                    </ul>
                    <table class="table" style="margin:0 auto;">
                        <tr>
                            <th style="width:20px"></th>
                            <th class="left" style="width:30%">Mouse Button</th>
                            <th class="left">Output</th>
                        </tr>
                        <tr>
                            <td colspan="3" class="left"><a id="btnmouseeditor" class='inline btn' href="javascript:void(0)">+</a></td>
                        </tr>
                        {{ range $key, $km := .MouseMaps }}
                        <tr id="{{$km.IdString}}">
                            <td><button class="btnRemoveMouseMap" objid="{{$km.IdString}}">-</button></td>
                            <td>{{$km.InputString}}</td>
                            <td>{{$km.OutputString}}</td>

                        </tr>
                        {{ end }}
                    </table>
                </div>



                <h1 class="left">⌨️ Keyboard Settings</h1>

                <table class="table" style="margin:0 auto;">
                    <tr>
                        <th style="width:20px"></th>
                        <th class="left" style="width:30%">Keyboard</th>
                        <th class="left">Output</th>
                    </tr>
                    <tr>
                        <td colspan="3" class="left"><a id="btnkeyeditor" href="javascript:void(0)" class='inline btn'>+</a></td>
                    </tr>
                    {{ range $key, $km := .KeyMaps }}
                    <tr id="{{$km.IdString}}">
                        <td><a href="javascript:void(0)" class="btnRemoveKeyMap btn" objid="{{$km.IdString}}">-</a></td>
                        <td>{{$km.InputString}}</td>
                        <td>{{$km.OutputString}}</td>

                    </tr>
                    {{ end }}
                </table>

            </div>
        </div>
    </div>
    </div>



    <div style='display:none;'>

        <div id='mouseeditor' style='padding:10px; background:#fff;'>
            <form id="frmMouse">
                <ul style="width:auto">
                    <li class="sectionleft">
                        <div style="width:435px;">
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mouse_in" value="1"><span>Left Button</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mouse_in" value="2"><span>Right Button</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mouse_in" value="4"><span>Middle Button</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mouse_in" value="8"><span>Side Button</span>
                                    </label>
                                </li>

                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mouse_in" value="16"><span>Extra Button</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mouse_in" value="32"><span>Forward Button</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mouse_in" value="64"><span>Back Button</span>
                                    </label>
                                </li>



                            </ul>


                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="1"><span>CTRL</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="8"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="4"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key space">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="1" disabled><span>SPACE</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="64"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="128"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" disabled><span>FN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="16"><span>CTRL</span>
                                    </label>
                                </li>

                            </ul>
                        </div>
                    </li>

                </ul>


                <div class="center">⇓⇓⇓⇓</div>



                <ul style="width:auto">
                    <li class="sectionleft">
                        <div style="width:435px;">
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="41"><span>ESC</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="58"><span>F1</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="59"><span>F2</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="60"><span>F3</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="61"><span>F4</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="62"><span>F5</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="63"><span>F6</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="64"><span>F7</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="65"><span>F8</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="66"><span>F9</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="67"><span>F10</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="68"><span>F11</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="69"><span>F12</span>
                                    </label>
                                </li>
                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="53"><span>~</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="30"><span>1</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="31"><span>2</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="32"><span>3</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="33"><span>4</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="34"><span>5</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="35"><span>6</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="36"><span>7</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="37"><span>8</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="38"><span>9</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="39"><span>0</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="45"><span>-</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="46"><span>=</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="42"><span>BACKSPACE</span>
                                    </label>
                                </li>

                            </ul>
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="43"><span>TAB</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="20"><span>Q</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="26"><span>W</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="8"><span>E</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="21"><span>R</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="23"><span>T</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="28"><span>Y</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="24"><span>U</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="12"><span>I</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="18"><span>O</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="19"><span>P</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="47"><span>[</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="48"><span>]</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="49"><span>\</span>
                                    </label>
                                </li>
                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="0" disabled><span>CAPS LOCK</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="4"><span>A</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="22"><span>S</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="7"><span>D</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="9"><span>F</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="10"><span>G</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="11"><span>H</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="13"><span>J</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="14"><span>K</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="15"><span>L</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="51"><span>;</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="52"><span>'</span>
                                    </label>
                                </li>
                                <li class="key enter">
                                    <label>
                                        <input type="checkbox" name="key_out" value="40"><span>ENTER</span>
                                    </label>
                                </li>
                            </ul>


                            <ul>
                                <li class="key shift">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="2"><span>SHIFT</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="29"><span>Z</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="27"><span>X</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="6"><span>C</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="25"><span>V</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="5"><span>B</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="17"><span>N</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="16"><span>M</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="54"><span>,</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="55"><span>.</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="56"><span>/</span>
                                    </label>
                                </li>
                                <li class="key shift">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="32"><span>SHIFT</span>
                                    </label>
                                </li>

                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="1"><span>CTRL</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="8"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="4"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key space">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="1"><span>SPACE</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="64"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="128"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" disabled><span>FN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="16"><span>CTRL</span>
                                    </label>
                                </li>

                            </ul>
                        </div>
                    </li>
                    <li class="sectionright">
                        <div style="width:96px;">
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="73"><span>INS</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="74"><span>HOME</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="75"><span>PG UP</span>
                                    </label>
                                </li>
                            </ul>
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="76"><span>DEL</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="77"><span>END</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="78"><span>PG DN</span>
                                    </label>
                                </li>
                            </ul>

                            <ul style="padding-top:33px">
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="82"><span>↑</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                            </ul>
                            <ul>
                                <li class="key">
                                    <label><input type="checkbox" name="key_out" value="80"><span>←</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="81"><span>↓</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label><input type="checkbox" name="key_out" value="79"><span>→</span>
                                    </label>
                                </li>
                            </ul>
                        </div>
                    </li>
                </ul>



            </form>
            <div style="padding-top:10px">
                <button id="btnSaveMouse">Save</button>
                <button class="btnClear">Clear</button>
                <button class="btnClose fright">Close</button>
                <label id="errMsgMouse" style="color:red"></label></div>
        </div>

        <div id='keyeditor' style='padding:10px; background:#fff;'>
            <form id="frmKey">
                <ul style="width:auto">
                    <li class="sectionleft">
                        <div style="width:435px;">
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="41"><span>ESC</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="58"><span>F1</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="59"><span>F2</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="60"><span>F3</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="61"><span>F4</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="62"><span>F5</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="63"><span>F6</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="64"><span>F7</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="65"><span>F8</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="66"><span>F9</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="67"><span>F10</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="68"><span>F11</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="69"><span>F12</span>
                                    </label>
                                </li>
                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="53"><span>~</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="30"><span>1</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="31"><span>2</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="32"><span>3</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="33"><span>4</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="34"><span>5</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="35"><span>6</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="36"><span>7</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="37"><span>8</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="38"><span>9</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="39"><span>0</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="45"><span>-</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="46"><span>=</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="42"><span>BACKSPACE</span>
                                    </label>
                                </li>

                            </ul>
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="43"><span>TAB</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="20"><span>Q</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="26"><span>W</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="8"><span>E</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="21"><span>R</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="23"><span>T</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="28"><span>Y</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="24"><span>U</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="12"><span>I</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="18"><span>O</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="19"><span>P</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="47"><span>[</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="48"><span>]</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="49"><span>\</span>
                                    </label>
                                </li>
                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="0" disabled><span>CAPS LOCK</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="4"><span>A</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="22"><span>S</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="7"><span>D</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="9"><span>F</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="10"><span>G</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="11"><span>H</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="13"><span>J</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="14"><span>K</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="15"><span>L</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="51"><span>;</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="52"><span>'</span>
                                    </label>
                                </li>
                                <li class="key enter">
                                    <label>
                                        <input type="checkbox" name="key_in" value="40"><span>ENTER</span>
                                    </label>
                                </li>
                            </ul>


                            <ul>
                                <li class="key shift">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="2"><span>SHIFT</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="29"><span>Z</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="27"><span>X</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="6"><span>C</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="25"><span>V</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="5"><span>B</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="17"><span>N</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="16"><span>M</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="54"><span>,</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="55"><span>.</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="56"><span>/</span>
                                    </label>
                                </li>
                                <li class="key shift">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="32"><span>SHIFT</span>
                                    </label>
                                </li>

                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="1"><span>CTRL</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="8"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="4"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key space">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="1"><span>SPACE</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="64"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="128"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" disabled><span>FN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_in" value="16"><span>CTRL</span>
                                    </label>
                                </li>

                            </ul>
                        </div>
                    </li>
                    <li class="sectionright">
                        <div style="width:96px;">
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="73"><span>INS</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="74"><span>HOME</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="75"><span>PG UP</span>
                                    </label>
                                </li>
                            </ul>
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="76"><span>DEL</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="77"><span>END</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="78"><span>PG DN</span>
                                    </label>
                                </li>
                            </ul>

                            <ul style="padding-top:33px">
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="82"><span>↑</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                            </ul>
                            <ul>
                                <li class="key">
                                    <label><input type="checkbox" name="key_in" value="80"><span>←</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_in" value="81"><span>↓</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label><input type="checkbox" name="key_in" value="79"><span>→</span>
                                    </label>
                                </li>
                            </ul>
                        </div>
                    </li>
                </ul>


                <div class="center">⇓⇓⇓⇓</div>



                <ul style="width:auto">
                    <li class="sectionleft">
                        <div style="width:435px;">
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="41"><span>ESC</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="58"><span>F1</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="59"><span>F2</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="60"><span>F3</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="61"><span>F4</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="62"><span>F5</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="63"><span>F6</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="64"><span>F7</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="65"><span>F8</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="66"><span>F9</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="67"><span>F10</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="68"><span>F11</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="69"><span>F12</span>
                                    </label>
                                </li>
                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="53"><span>~</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="30"><span>1</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="31"><span>2</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="32"><span>3</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="33"><span>4</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="34"><span>5</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="35"><span>6</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="36"><span>7</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="37"><span>8</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="38"><span>9</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="39"><span>0</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="45"><span>-</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="46"><span>=</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="42"><span>BACKSPACE</span>
                                    </label>
                                </li>

                            </ul>
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="43"><span>TAB</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="20"><span>Q</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="26"><span>W</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="8"><span>E</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="21"><span>R</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="23"><span>T</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="28"><span>Y</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="24"><span>U</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="12"><span>I</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="18"><span>O</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="19"><span>P</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="47"><span>[</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="48"><span>]</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="49"><span>\</span>
                                    </label>
                                </li>
                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="0" disabled><span>CAPS LOCK</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="4"><span>A</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="22"><span>S</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="7"><span>D</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="9"><span>F</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="10"><span>G</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="11"><span>H</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="13"><span>J</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="14"><span>K</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="15"><span>L</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="51"><span>;</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="52"><span>'</span>
                                    </label>
                                </li>
                                <li class="key enter">
                                    <label>
                                        <input type="checkbox" name="key_out" value="40"><span>ENTER</span>
                                    </label>
                                </li>
                            </ul>


                            <ul>
                                <li class="key shift">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="2"><span>SHIFT</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="29"><span>Z</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="27"><span>X</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="6"><span>C</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="25"><span>V</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="5"><span>B</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="17"><span>N</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="16"><span>M</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="54"><span>,</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="55"><span>.</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="56"><span>/</span>
                                    </label>
                                </li>
                                <li class="key shift">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="32"><span>SHIFT</span>
                                    </label>
                                </li>

                            </ul>

                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="1"><span>CTRL</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="8"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="4"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key space">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="1"><span>SPACE</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="64"><span>ALT</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="128"><span>WIN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" disabled><span>FN</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="mod_out" value="16"><span>CTRL</span>
                                    </label>
                                </li>

                            </ul>
                        </div>
                    </li>
                    <li class="sectionright">
                        <div style="width:96px;">
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="73"><span>INS</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="74"><span>HOME</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="75"><span>PG UP</span>
                                    </label>
                                </li>
                            </ul>
                            <ul>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="76"><span>DEL</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="77"><span>END</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="78"><span>PG DN</span>
                                    </label>
                                </li>
                            </ul>

                            <ul style="padding-top:33px">
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="82"><span>↑</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                    </label>
                                </li>
                            </ul>
                            <ul>
                                <li class="key">
                                    <label><input type="checkbox" name="key_out" value="80"><span>←</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label>
                                        <input type="checkbox" name="key_out" value="81"><span>↓</span>
                                    </label>
                                </li>
                                <li class="key">
                                    <label><input type="checkbox" name="key_out" value="79"><span>→</span>
                                    </label>
                                </li>
                            </ul>
                        </div>
                    </li>
                </ul>



            </form>
            <div style="padding-top:10px">
                <button id="btnSaveKeyboard">Save</button>
                <button class="btnClear">Clear</button>
                <button class="btnClose fright">Close</button>
                <label id="errMsgKey" style="color:red"></label>
            </div>
        </div>
    </div>


    <!-- <script>
        var charfield = document.getElementById("btnSave");
        charfield.onkeypress = function (e) {
            e = e || window.event;
            var charCode = (typeof e.which == "number") ? e.which : e.keyCode;
            if (charCode > 0) {
                alert("Typed character: " + String.fromCharCode(charCode));
            }
        };
    </script> -->


    <script>
        $(document).ready(function () {
            $("#btnkeyeditor").click(function () {
                $('#frmKey').trigger("reset");
                $('#errMsg').text('');
                $("#keyeditor").modal({
                    escapeClose: false,
                    clickClose: false,
                    showClose: false
                });
            })
            $("#btnmouseeditor").click(function () {
                $('#frmMouse').trigger("reset");
                $('#errMsg').text('');
                $("#mouseeditor").modal({
                    escapeClose: false,
                    clickClose: false,
                    showClose: false
                });
            })
            $(".btnClose").click(function () {
                $.modal.close();
            })

            $('input[type=radio][name=mousespeed]').change(function () {
                $.get("/?action=setmousespeed&speed=" + $(this).val(), function (r) {
                    console.debug("new mouse speed " + r);
                });
            });
            $('input[type=radio][name=scrollspeed]').change(function () {
                $.get("/?action=setscrollspeed&speed=" + $(this).val(), function (r) {
                    console.debug("new mouse scroll speed " + r);
                });
            });

            $(".btnClear").click(function () {
                $('#frmKey,#frmMouse').trigger("reset");
                $('#errMsg').text('');
            });

            $("#btnSaveKeyboard").click(function () {
                var input_modkeys = [];
                $('#frmKey input[name="mod_in"]:checked').each(function () {
                    input_modkeys.push($(this).val());
                });
                console.debug("mod_in " + input_modkeys)

                var input_keys = [];
                $('#frmKey input[name="key_in"]:checked').each(function () {
                    input_keys.push($(this).val());
                });
                console.debug("key_in " + input_keys)

                var output_modkeys = [];
                $('#frmKey input[name="mod_out"]:checked').each(function () {
                    output_modkeys.push($(this).val());
                });
                console.debug("mod_out " + output_modkeys)

                var output_keys = [];
                $('#frmKey input[name="key_out"]:checked').each(function () {
                    output_keys.push($(this).val());
                });
                console.debug("key_out " + output_keys)

                if (output_keys.length > 6) {
                    $('#errMsgKey').text("max output 6 keys");
                    return;
                }

                $.get("/?action=savekeymap", $("#frmKey").serialize(), function (r) {
                    if (r.result) {
                        $.modal.close();
                        window.location.reload(true);
                    }
                    else
                        $('#errMsgKey').text(r.msg);
                });
            });

            $(".btnRemoveKeyMap").click(function () {
                $.get("/?action=removekeymap&id=" + $(this).attr("objid"), function (r) {
                    window.location.reload(true);
                });
            });


            $("#btnSaveMouse").click(function () {
                var input_modkeys = [];
                $('#frmMouse input[name="mod_in"]:checked').each(function () {
                    input_modkeys.push($(this).val());
                });
                console.debug("mod_in " + input_modkeys)

                var input_keys = [];
                $('#frmMouse input[name="mouse_in"]:checked').each(function () {
                    input_keys.push($(this).val());
                });
                console.debug("mouse_in " + input_keys)

                var output_modkeys = [];
                $('#frmMouse input[name="mod_out"]:checked').each(function () {
                    output_modkeys.push($(this).val());
                });
                console.debug("mod_out " + output_modkeys)

                var output_keys = [];
                $('#frmMouse input[name="key_out"]:checked').each(function () {
                    output_keys.push($(this).val());
                });
                console.debug("key_out " + output_keys)

                if (output_keys.length > 6) {
                    $('#errMsgMouse').text("max output 6 keys");
                    return;
                }

                $.get("/?action=savemousemap", $("#frmMouse").serialize(), function (r) {
                    if (r.result) {
                        $.modal.close();
                        window.location.reload(true);
                    }
                    else
                        $('#errMsgMouse').text(r.msg);
                });
            });

            $(".btnRemoveMouseMap").click(function () {
                $.get("/?action=removemousemap&id=" + $(this).attr("objid"), function (r) {
                    window.location.reload(true);
                });
            });


        })


    </script>
</body>

</html>

`)
