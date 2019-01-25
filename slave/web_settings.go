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

<!doctype html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>HID Like Windows</title>
  <link rel="stylesheet" href="/style.css">
  <script src="/jquery.js"></script>
</head>

<body>
  <div class="box">
    <ul>
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
        </div>


        
        <h1 class="left">⌨️ Keyboard Settings</h1>

        <table class="table" style="margin:0 auto;">
          <tr>
            <th class="left">Input</th>
            <th class="left">Output</th>
            <th style="width:20px"></th>
          </tr>
          <tr>
            <td colspan="3" class="right"><button class='inline' href="#keyeditor">+</button></td>
          </tr>

          {{ range $key, $km := .KeyMaps }}
          <tr id="{{$key}}">
            <td>{{$km.InputString}}</td>
            <td>{{$km.OutputString}}</td>
            <td><button class="btnRemove" objid="{{$key}}">-</button></td>
          </tr>
          {{ end }}

        </table>

      </div>
    </div>
  </div>
  </div>



  <div style='display:none'>
    <div id='keyeditor' style='padding:10px; background:#fff;'>

 


      <form id="frmEditor">
        <div>
          <input type="checkbox" id="KEY_LEFTCTRL" name="modkey" value="1"><label for="KEY_LEFTCTRL">L.CTRL</label>
          <input type="checkbox" id="KEY_LEFTSHIFT" name="modkey" value="2"><label for="KEY_LEFTSHIFT">L.SHIFT</label>
          <input type="checkbox" id="KEY_LEFTALT" name="modkey" value="4"><label for="KEY_LEFTALT">L.ALT</label>
          <input type="checkbox" id="KEY_LEFTMETA" name="modkey" value="8"><label for="KEY_LEFTMETA">L.META(CMD,WIN)</label>
          <input type="checkbox" id="KEY_RIGHTCTRL" name="modkey" value="16"><label for="KEY_RIGHTCTRL">R.CTRL</label>
          <input type="checkbox" id="KEY_RIGHTSHIFT" name="modkey" value="32"><label for="KEY_RIGHTSHIFT">R.SHIFT</label>
          <input type="checkbox" id="KEY_RIGHTALT" name="modkey" value="64"><label for="KEY_RIGHTALT">R.ALT</label>
          <input type="checkbox" id="KEY_RIGHTMETA" name="modkey" value="128"><label for="KEY_RIGHTMETA">R.META(CMD.WIN)</label>
        </div>
        <div>
          <input type="checkbox" id="KEY_A" name="key" value="4"><label for="KEY_A">A</label>
          <input type="checkbox" id="KEY_B" name="key" value="5"><label for="KEY_B">B</label>
          <input type="checkbox" id="KEY_C" name="key" value="6"><label for="KEY_C">C</label>
          <input type="checkbox" id="KEY_D" name="key" value="7"><label for="KEY_D">D</label>
          <input type="checkbox" id="KEY_E" name="key" value="8"><label for="KEY_E">E</label>
          <input type="checkbox" id="KEY_F" name="key" value="9"><label for="KEY_F">F</label>
          <input type="checkbox" id="KEY_G" name="key" value="10"><label for="KEY_G">G</label>
          <input type="checkbox" id="KEY_H" name="key" value="11"><label for="KEY_H">H</label>
          <input type="checkbox" id="KEY_I" name="key" value="12"><label for="KEY_I">I</label>
          <input type="checkbox" id="KEY_J" name="key" value="13"><label for="KEY_J">J</label>
          <input type="checkbox" id="KEY_K" name="key" value="14"><label for="KEY_K">K</label>
          <input type="checkbox" id="KEY_L" name="key" value="15"><label for="KEY_L">L</label>
          <input type="checkbox" id="KEY_M" name="key" value="16"><label for="KEY_M">M</label>
          <input type="checkbox" id="KEY_N" name="key" value="17"><label for="KEY_N">N</label>
          <input type="checkbox" id="KEY_O" name="key" value="18"><label for="KEY_O">O</label>
          <input type="checkbox" id="KEY_P" name="key" value="19"><label for="KEY_P">P</label>
          <input type="checkbox" id="KEY_Q" name="key" value="20"><label for="KEY_Q">Q</label>
          <input type="checkbox" id="KEY_R" name="key" value="21"><label for="KEY_R">R</label>
          <input type="checkbox" id="KEY_S" name="key" value="22"><label for="KEY_S">S</label>
          <input type="checkbox" id="KEY_T" name="key" value="23"><label for="KEY_T">T</label>
          <input type="checkbox" id="KEY_U" name="key" value="24"><label for="KEY_U">U</label>
          <input type="checkbox" id="KEY_V" name="key" value="25"><label for="KEY_V">V</label>
          <input type="checkbox" id="KEY_W" name="key" value="26"><label for="KEY_W">W</label>
          <input type="checkbox" id="KEY_X" name="key" value="27"><label for="KEY_X">X</label>
          <input type="checkbox" id="KEY_Y" name="key" value="28"><label for="KEY_Y">Y</label>
          <input type="checkbox" id="KEY_Z" name="key" value="29"><label for="KEY_Z">Z</label>
        </div>
        <div>
          <input type="checkbox" id="KEY_1" name="key" value="30"><label for="KEY_1">1</label>
          <input type="checkbox" id="KEY_2" name="key" value="31"><label for="KEY_2">2</label>
          <input type="checkbox" id="KEY_3" name="key" value="32"><label for="KEY_3">3</label>
          <input type="checkbox" id="KEY_4" name="key" value="33"><label for="KEY_4">4</label>
          <input type="checkbox" id="KEY_5" name="key" value="34"><label for="KEY_5">5</label>
          <input type="checkbox" id="KEY_6" name="key" value="35"><label for="KEY_6">6</label>
          <input type="checkbox" id="KEY_7" name="key" value="36"><label for="KEY_7">7</label>
          <input type="checkbox" id="KEY_8" name="key" value="37"><label for="KEY_8">8</label>
          <input type="checkbox" id="KEY_9" name="key" value="38"><label for="KEY_9">9</label>
          <input type="checkbox" id="KEY_0" name="key" value="39"><label for="KEY_0">0</label>
        </div>
        <div>
          <input type="checkbox" id="KEY_ENTER" name="key" value="40"><label for="KEY_ENTER">ENTER</label>
          <input type="checkbox" id="KEY_ESC" name="key" value="41"><label for="KEY_ESC">ESC</label>
          <input type="checkbox" id="KEY_BACKSPACE" name="key" value="42"><label for="KEY_BACKSPACE">BACKSPACE</label>
          <input type="checkbox" id="KEY_TAB" name="key" value="43"><label for="KEY_TAB">TAB</label>
          <input type="checkbox" id="KEY_SPACE" name="key" value="44"><label for="KEY_SPACE">SPACE</label>

          <input type="checkbox" id="KEY_MINUS" name="key" value="45"><label for="KEY_MINUS">-</label>
          <input type="checkbox" id="KEY_EQUAL" name="key" value="46"><label for="KEY_EQUAL">=</label>
          <input type="checkbox" id="KEY_LEFTBRACE" name="key" value="47"><label for="KEY_LEFTBRACE">[</label>
          <input type="checkbox" id="KEY_RIGHTBRACE" name="key" value="48"><label for="KEY_RIGHTBRACE">]</label>
          <input type="checkbox" id="KEY_BACKSLASH" name="key" value="49"><label for="KEY_BACKSLASH">\</label>
          <input type="checkbox" id="KEY_SEMICOLON" name="key" value="51"><label for="KEY_SEMICOLON">;</label>
          <input type="checkbox" id="KEY_APOSTROPHE" name="key" value="52"><label for="KEY_APOSTROPHE">'</label>
          <input type="checkbox" id="KEY_GRAVE" name="key" value="53"><label for="KEY_GRAVE">&grave;</label>
          <input type="checkbox" id="KEY_COMMA" name="key" value="54"><label for="KEY_COMMA">,</label>
          <input type="checkbox" id="KEY_DOT" name="key" value="55"><label for="KEY_DOT">.</label>
          <input type="checkbox" id="KEY_SLASH" name="key" value="56"><label for="KEY_SLASH">/</label>


        </div>
        <div>
          <input type="checkbox" id="KEY_F1" name="key" value="58"><label for="KEY_F1">F1</label>
          <input type="checkbox" id="KEY_F2" name="key" value="59"><label for="KEY_F2">F2</label>
          <input type="checkbox" id="KEY_F3" name="key" value="60"><label for="KEY_F3">F3</label>
          <input type="checkbox" id="KEY_F4" name="key" value="61"><label for="KEY_F4">F4</label>
          <input type="checkbox" id="KEY_F5" name="key" value="62"><label for="KEY_F5">F5</label>
          <input type="checkbox" id="KEY_F6" name="key" value="63"><label for="KEY_F6">F6</label>
          <input type="checkbox" id="KEY_F7" name="key" value="64"><label for="KEY_F7">F7</label>
          <input type="checkbox" id="KEY_F8" name="key" value="65"><label for="KEY_F8">F8</label>
          <input type="checkbox" id="KEY_F9" name="key" value="66"><label for="KEY_F9">F9</label>
          <input type="checkbox" id="KEY_F10" name="key" value="67"><label for="KEY_F10">F10</label>
          <input type="checkbox" id="KEY_F11" name="key" value="68"><label for="KEY_F11">F11</label>
          <input type="checkbox" id="KEY_F12" name="key" value="69"><label for="KEY_F12">F12</label>
        </div>
        <div><input type="checkbox" id="KEY_INSERT" name="key" value="73"><label for="KEY_INSERT">INSERT</label>
          <input type="checkbox" id="KEY_HOME" name="key" value="74"><label for="KEY_HOME">HOME</label>
          <input type="checkbox" id="KEY_PAGEUP" name="key" value="75"><label for="KEY_PAGEUP">PAGEUP</label>
          <input type="checkbox" id="KEY_DELETE" name="key" value="76"><label for="KEY_DELETE">DELETE</label>
          <input type="checkbox" id="KEY_END" name="key" value="77"><label for="KEY_END">END</label>
          <input type="checkbox" id="KEY_PAGEDOWN" name="key" value="78"><label for="KEY_PAGEDOWN">PAGEDOWN</label>
          <input type="checkbox" id="KEY_RIGHT" name="key" value="79"><label for="KEY_RIGHT">RIGHT</label>
          <input type="checkbox" id="KEY_LEFT" name="key" value="80"><label for="KEY_LEFT">LEFT</label>
          <input type="checkbox" id="KEY_DOWN" name="key" value="81"><label for="KEY_DOWN">DOWN</label>
          <input type="checkbox" id="KEY_UP" name="key" value="82"><label for="KEY_UP">UP</label>
        </div>
        <hr />
        <div>
          <input type="checkbox" id="KEY_LEFTCTRL_OUT" name="modkey_out" value="1"><label for="KEY_LEFTCTRL_OUT">L.CTRL</label>
          <input type="checkbox" id="KEY_LEFTSHIFT_OUT" name="modkey_out" value="2"><label for="KEY_LEFTSHIFT_OUT">L.SHIFT</label>
          <input type="checkbox" id="KEY_LEFTALT_OUT" name="modkey_out" value="4"><label for="KEY_LEFTALT_OUT">L.ALT</label>
          <input type="checkbox" id="KEY_LEFTMETA_OUT" name="modkey_out" value="8"><label for="KEY_LEFTMETA_OUT">L.META(CMD,WIN)</label>
          <input type="checkbox" id="KEY_RIGHTCTRL_OUT" name="modkey_out" value="16"><label for="KEY_RIGHTCTRL_OUT">R.CTRL</label>
          <input type="checkbox" id="KEY_RIGHTSHIFT_OUT" name="modkey_out" value="32"><label for="KEY_RIGHTSHIFT_OUT">R.SHIFT</label>
          <input type="checkbox" id="KEY_RIGHTALT_OUT" name="modkey_out" value="64"><label for="KEY_RIGHTALT_OUT">R.ALT</label>
          <input type="checkbox" id="KEY_RIGHTMETA_OUT" name="modkey_out" value="128"><label for="KEY_RIGHTMETA_OUT">R.META(CMD.WIN)</label>
        </div>
        <div>
          <input type="checkbox" id="KEY_A_OUT" name="key_out" value="4"><label for="KEY_A_OUT">A</label>
          <input type="checkbox" id="KEY_B_OUT" name="key_out" value="5"><label for="KEY_B_OUT">B</label>
          <input type="checkbox" id="KEY_C_OUT" name="key_out" value="6"><label for="KEY_C_OUT">C</label>
          <input type="checkbox" id="KEY_D_OUT" name="key_out" value="7"><label for="KEY_D_OUT">D</label>
          <input type="checkbox" id="KEY_E_OUT" name="key_out" value="8"><label for="KEY_E_OUT">E</label>
          <input type="checkbox" id="KEY_F_OUT" name="key_out" value="9"><label for="KEY_F_OUT">F</label>
          <input type="checkbox" id="KEY_G_OUT" name="key_out" value="10"><label for="KEY_G_OUT">G</label>
          <input type="checkbox" id="KEY_H_OUT" name="key_out" value="11"><label for="KEY_H_OUT">H</label>
          <input type="checkbox" id="KEY_I_OUT" name="key_out" value="12"><label for="KEY_I_OUT">I</label>
          <input type="checkbox" id="KEY_J_OUT" name="key_out" value="13"><label for="KEY_J_OUT">J</label>
          <input type="checkbox" id="KEY_K_OUT" name="key_out" value="14"><label for="KEY_K_OUT">K</label>
          <input type="checkbox" id="KEY_L_OUT" name="key_out" value="15"><label for="KEY_L_OUT">L</label>
          <input type="checkbox" id="KEY_M_OUT" name="key_out" value="16"><label for="KEY_M_OUT">M</label>
          <input type="checkbox" id="KEY_N_OUT" name="key_out" value="17"><label for="KEY_N_OUT">N</label>
          <input type="checkbox" id="KEY_O_OUT" name="key_out" value="18"><label for="KEY_O_OUT">O</label>
          <input type="checkbox" id="KEY_P_OUT" name="key_out" value="19"><label for="KEY_P_OUT">P</label>
          <input type="checkbox" id="KEY_Q_OUT" name="key_out" value="20"><label for="KEY_Q_OUT">Q</label>
          <input type="checkbox" id="KEY_R_OUT" name="key_out" value="21"><label for="KEY_R_OUT">R</label>
          <input type="checkbox" id="KEY_S_OUT" name="key_out" value="22"><label for="KEY_S_OUT">S</label>
          <input type="checkbox" id="KEY_T_OUT" name="key_out" value="23"><label for="KEY_T_OUT">T</label>
          <input type="checkbox" id="KEY_U_OUT" name="key_out" value="24"><label for="KEY_U_OUT">U</label>
          <input type="checkbox" id="KEY_V_OUT" name="key_out" value="25"><label for="KEY_V_OUT">V</label>
          <input type="checkbox" id="KEY_W_OUT" name="key_out" value="26"><label for="KEY_W_OUT">W</label>
          <input type="checkbox" id="KEY_X_OUT" name="key_out" value="27"><label for="KEY_X_OUT">X</label>
          <input type="checkbox" id="KEY_Y_OUT" name="key_out" value="28"><label for="KEY_Y_OUT">Y</label>
          <input type="checkbox" id="KEY_Z_OUT" name="key_out" value="29"><label for="KEY_Z_OUT">Z</label>
        </div>
        <div>
          <input type="checkbox" id="KEY_1_OUT" name="key_out" value="30"><label for="KEY_1_OUT">1</label>
          <input type="checkbox" id="KEY_2_OUT" name="key_out" value="31"><label for="KEY_2_OUT">2</label>
          <input type="checkbox" id="KEY_3_OUT" name="key_out" value="32"><label for="KEY_3_OUT">3</label>
          <input type="checkbox" id="KEY_4_OUT" name="key_out" value="33"><label for="KEY_4_OUT">4</label>
          <input type="checkbox" id="KEY_5_OUT" name="key_out" value="34"><label for="KEY_5_OUT">5</label>
          <input type="checkbox" id="KEY_6_OUT" name="key_out" value="35"><label for="KEY_6_OUT">6</label>
          <input type="checkbox" id="KEY_7_OUT" name="key_out" value="36"><label for="KEY_7_OUT">7</label>
          <input type="checkbox" id="KEY_8_OUT" name="key_out" value="37"><label for="KEY_8_OUT">8</label>
          <input type="checkbox" id="KEY_9_OUT" name="key_out" value="38"><label for="KEY_9_OUT">9</label>
          <input type="checkbox" id="KEY_0_OUT" name="key_out" value="39"><label for="KEY_0_OUT">0</label>
        </div>
        <div>
          <input type="checkbox" id="KEY_ENTER_OUT" name="key_out" value="40"><label for="KEY_ENTER_OUT">ENTER</label>
          <input type="checkbox" id="KEY_ESC_OUT" name="key_out" value="41"><label for="KEY_ESC_OUT">ESC</label>
          <input type="checkbox" id="KEY_BACKSPACE_OUT" name="key_out" value="42"><label for="KEY_BACKSPACE_OUT">BACKSPACE</label>
          <input type="checkbox" id="KEY_TAB_OUT" name="key_out" value="43"><label for="KEY_TAB_OUT">TAB</label>
          <input type="checkbox" id="KEY_SPACE_OUT" name="key_out" value="44"><label for="KEY_SPACE_OUT">SPACE</label>

          <input type="checkbox" id="KEY_MINUS_OUT" name="key_out" value="45"><label for="KEY_MINUS_OUT">-</label>
          <input type="checkbox" id="KEY_EQUAL_OUT" name="key_out" value="46"><label for="KEY_EQUAL_OUT">=</label>
          <input type="checkbox" id="KEY_LEFTBRACE_OUT" name="key_out" value="47"><label for="KEY_LEFTBRACE_OUT">[</label>
          <input type="checkbox" id="KEY_RIGHTBRACE_OUT" name="key_out" value="48"><label for="KEY_RIGHTBRACE_OUT">]</label>
          <input type="checkbox" id="KEY_BACKSLASH_OUT" name="key_out" value="49"><label for="KEY_BACKSLASH_OUT">\</label>
          <input type="checkbox" id="KEY_SEMICOLON_OUT" name="key_out" value="51"><label for="KEY_SEMICOLON_OUT">;</label>
          <input type="checkbox" id="KEY_APOSTROPHE_OUT" name="key_out" value="52"><label for="KEY_APOSTROPHE_OUT">'</label>
          <input type="checkbox" id="KEY_GRAVE_OUT" name="key_out" value="53"><label for="KEY_GRAVE_OUT">&grave;</label>
          <input type="checkbox" id="KEY_COMMA_OUT" name="key_out" value="54"><label for="KEY_COMMA_OUT">,</label>
          <input type="checkbox" id="KEY_DOT_OUT" name="key_out" value="55"><label for="KEY_DOT_OUT">.</label>
          <input type="checkbox" id="KEY_SLASH_OUT" name="key_out" value="56"><label for="KEY_SLASH_OUT">/</label>
        </div>
        <div>
          <input type="checkbox" id="KEY_F1_OUT" name="key_out" value="58"><label for="KEY_F1_OUT">F1</label>
          <input type="checkbox" id="KEY_F2_OUT" name="key_out" value="59"><label for="KEY_F2_OUT">F2</label>
          <input type="checkbox" id="KEY_F3_OUT" name="key_out" value="60"><label for="KEY_F3_OUT">F3</label>
          <input type="checkbox" id="KEY_F4_OUT" name="key_out" value="61"><label for="KEY_F4_OUT">F4</label>
          <input type="checkbox" id="KEY_F5_OUT" name="key_out" value="62"><label for="KEY_F5_OUT">F5</label>
          <input type="checkbox" id="KEY_F6_OUT" name="key_out" value="63"><label for="KEY_F6_OUT">F6</label>
          <input type="checkbox" id="KEY_F7_OUT" name="key_out" value="64"><label for="KEY_F7_OUT">F7</label>
          <input type="checkbox" id="KEY_F8_OUT" name="key_out" value="65"><label for="KEY_F8_OUT">F8</label>
          <input type="checkbox" id="KEY_F9_OUT" name="key_out" value="66"><label for="KEY_F9_OUT">F9</label>
          <input type="checkbox" id="KEY_F10_OUT" name="key_out" value="67"><label for="KEY_F10_OUT">F10</label>
          <input type="checkbox" id="KEY_F11_OUT" name="key_out" value="68"><label for="KEY_F11_OUT">F11</label>
          <input type="checkbox" id="KEY_F12_OUT" name="key_out" value="69"><label for="KEY_F12_OUT">F12</label>
        </div>
        <div><input type="checkbox" id="KEY_INSERT_OUT" name="key_out" value="73"><label for="KEY_INSERT_OUT">INSERT</label>
          <input type="checkbox" id="KEY_HOME_OUT" name="key_out" value="74"><label for="KEY_HOME_OUT">HOME</label>
          <input type="checkbox" id="KEY_PAGEUP_OUT" name="key_out" value="75"><label for="KEY_PAGEUP_OUT">PAGEUP</label>
          <input type="checkbox" id="KEY_DELETE_OUT" name="key_out" value="76"><label for="KEY_DELETE_OUT">DELETE</label>
          <input type="checkbox" id="KEY_END_OUT" name="key_out" value="77"><label for="KEY_END_OUT">END</label>
          <input type="checkbox" id="KEY_PAGEDOWN_OUT" name="key_out" value="78"><label for="KEY_PAGEDOWN_OUT">PAGEDOWN</label>
          <input type="checkbox" id="KEY_RIGHT_OUT" name="key_out" value="79"><label for="KEY_RIGHT_OUT">RIGHT</label>
          <input type="checkbox" id="KEY_LEFT_OUT" name="key_out" value="80"><label for="KEY_LEFT_OUT">LEFT</label>
          <input type="checkbox" id="KEY_DOWN_OUT" name="key_out" value="81"><label for="KEY_DOWN_OUT">DOWN</label>
          <input type="checkbox" id="KEY_UP_OUT" name="key_out" value="82"><label for="KEY_UP_OUT">UP</label>
        </div>
      </form>
      <hr />
      <button id="btnSave">Save</button>
      <button id="btnClear">Clear</button>
      <label id="errMsg" style="color:red"></label>
    </div>
  </div>


  <script>
      var charfield = document.getElementById("btnSave");
      charfield.onkeypress = function (e) {
        e = e || window.event;
        var charCode = (typeof e.which == "number") ? e.which : e.keyCode;
        if (charCode > 0) {
          alert("Typed character: " + String.fromCharCode(charCode));
        }
      };
    </script>


  <script>
    $(document).ready(function () {

      $('input[type=radio][name=mousespeed]').change(function () {
        $.get("/?action=setmousespeed&speed=" + $(this).val(), function (r) {
          console.debug(r);
        });
      });
      $('input[type=radio][name=scrollspeed]').change(function () {
        $.get("/?action=setscrollspeed&speed=" + $(this).val(), function (r) {
          console.debug(r);
        });
      });


      $(".inline").colorbox({ inline: true, width: "50%" });
      $("#btnClear").click(function () {
        $('#frmEditor').trigger("reset");
        $('#errMsg').text('');
      });
      $(document).bind('cbox_open', function () {
        $('#frmEditor').trigger("reset");
        $('#errMsg').text('');
      });
      $("#btnSave").click(function () {
        var input_modkeys = [];
        $('input[name="modkey"]:checked').each(function () {
          input_modkeys.push($(this).val());
        });
        console.debug(input_modkeys)

        var input_keys = [];
        $('input[name="key"]:checked').each(function () {
          input_keys.push($(this).val());
        });
        console.debug(input_keys)

        var output_modkeys = [];
        $('input[name="modkey_out"]:checked').each(function () {
          output_modkeys.push($(this).val());
        });
        console.debug(output_modkeys)

        var output_keys = [];
        $('input[name="key_out"]:checked').each(function () {
          output_keys.push($(this).val());
        });
        console.debug(output_keys)

        if (output_keys.length > 6) {
          $('#errMsg').text("max output 6 keys");
          return;
        }

        $.get("/?action=savekeymap", $("#frmEditor").serialize(), function (r) {
          if (r.result) {
            $.colorbox.close();
            window.location.reload(true);
          }
          else
            $('#errMsg').text(r.msg);
        });


      });

      $(".btnRemove").click(function () {
        $.get("/?action=removekeymap&id=" + $(this).attr("objid"), function (r) {
          window.location.reload(true);
        });
      });

    })


  </script>
</body>

</html>

`)
