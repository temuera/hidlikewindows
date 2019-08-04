package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	evdev "github.com/gvalkov/golang-evdev"
	"github.com/sirupsen/logrus"
)

func (obj *Host) WebWin(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("draw") == "1" {
		defer func() {
			obj.InputEventCH <- &evdev.InputEvent{Type: evdev.EV_KEY, Code: evdev.BTN_MIDDLE, Value: 1}
			obj.InputEventCH <- &evdev.InputEvent{Type: evdev.EV_KEY, Code: evdev.BTN_MIDDLE, Value: 0}
		}()

		centerX, _ := strconv.Atoi(r.FormValue("centerx"))
		centerY, _ := strconv.Atoi(r.FormValue("centery"))
		if centerX == 0 || centerY == 0 {
			return
		}
		for n := int32(1); n <= 4; n++ {
			for i := 0; i < 4; i++ {
				moved := int32(0)
				movedSteps := []string{}
				for j := int32(n); j <= n*10; j += n {
					moved += j
					movedSteps = append(movedSteps, strconv.Itoa(int(j)))
					switch i {
					case 0:
						if b := obj.mouseReport.OnMove(evdev.REL_X, j); b != nil {
							if e := obj.mouse.Write(b); e != nil {
								logrus.Error("write to mouse hid: ", e)
							}
						}
						break
					case 1:
						if b := obj.mouseReport.OnMove(evdev.REL_Y, j); b != nil {
							if e := obj.mouse.Write(b); e != nil {
								logrus.Error("write to mouse hid: ", e)
							}
						}
						break
					case 2:
						if b := obj.mouseReport.OnMove(evdev.REL_X, -j); b != nil {
							if e := obj.mouse.Write(b); e != nil {
								logrus.Error("write to mouse hid: ", e)
							}
						}
						break
					case 3:
						if b := obj.mouseReport.OnMove(evdev.REL_Y, -j); b != nil {
							if e := obj.mouse.Write(b); e != nil {
								logrus.Error("write to mouse hid: ", e)
							}
						}
						break
					}
					obj.InputEventCH <- &evdev.InputEvent{Type: evdev.EV_KEY, Code: evdev.BTN_LEFT, Value: 1}
					obj.InputEventCH <- &evdev.InputEvent{Type: evdev.EV_KEY, Code: evdev.BTN_LEFT, Value: 0}

					time.Sleep(10 * time.Millisecond)
				}

				logrus.Infof("Win Move to %d -> %s = %d", i, strings.Join(movedSteps, "+"), moved)

			}
			time.Sleep(1000 * time.Millisecond)
		}
		//先右移
		//步子从 1-centerX

		return
	} else if msg := r.FormValue("hid"); msg != "" {
		ss := strings.Split(msg, ":")
		if len(ss) != 3 {
			logrus.Error("ws hid,bad msg: ", msg)
			return
		}

		v, _ := strconv.Atoi(ss[2])
		e := &evdev.InputEvent{}
		e.Value = int32(v)
		if ss[0] == "M" {
			switch ss[1] {
			case "X":
			case "Y":
				{
					e.Type = evdev.EV_REL
					e.Code = evdev.REL_Y
				}
			case "C":
				{
					e.Type = evdev.EV_KEY
					e.Code = evdev.BTN_LEFT
				}
			default:
				{
					logrus.Error("ws hid,bad msg: ", msg)
					return
				}
			}

		}
		//obj.OutputEventCH <- e WIN版直接输出数据给OS.
		{
			if b := obj.mouseReport.OnMove(e.Code, e.Value); b != nil {
				if e := obj.mouse.Write(b); e != nil {
					logrus.Error("write to mouse hid: ", e)
				}
			}
		}

	}

	w.Header().Set("Content-Type", "text/html")
	filename := CurrentDir() + "assets/win.html"
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if b, e := ioutil.ReadFile(filename); e == nil {
			w.Write(b)
			return
		}
	}
	//write from file
	w.Write(content_win)
}

var content_win = []byte(`<!doctype html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>WINDOWS Mouse</title>
  <link rel="stylesheet" href="/style.css">
  <script src="/jquery.js"></script>
</head>

<body>

  <div class="box">
    <ul class="nav">
      <li><a href="/">Settings</a></li>
      <li><a href="/win" class="active">Win Mouse</a></li>
      <li><a href="/nonwin">Non-Win Mouse</a></li>
    </ul>
    <div id="content" class="row content"></div>
  </div>
  </div>

  <script>
    var Step = "";
    var canvas = document.createElement('canvas');
    canvas.id = "CursorLayer";
    pW = $('#content').width() - 40, pH = $('#content').height() - 20;
    canvas.width = pW - pW % 100;
    canvas.height = pH - pH % 100;
    canvas.style.border = "1px solid";
    canvas.style.margin = (pH - canvas.height) / 2 + "px " + (pW - canvas.width) / 2 + "px";
    $('#content').append(canvas);
    var ctx = canvas.getContext("2d");
    drawGrid();

    var centerX = canvas.width / 4, centerY = canvas.height / 4;

    var clickCount = 0;
    canvas.onmousedown = function (event) {
      if (Step == "") {
        alert("press enter and leave your hand from mouse !!!")
        Step = "zero"
        return;
      }
      console.log(event.button);
      if (Step == "draw") {
        if (event.button == 1) {
          event.preventDefault();
          Step = "done";
          console.log("done!");
          return;
        }
        clickCount++;
        var x = event.pageX - canvas.offsetLeft, y = event.pageY - canvas.offsetTop;
        console.log("click " + clickCount + ": " + (x - centerX) + "," + (y - centerY));
        drawPoint(x, y, "green");
        drawLine(LastPoint.X, LastPoint.Y, x, y)
        if (clickCount % 10 == 0)
          drawText(x, y);
        LastPoint.X = x;
        LastPoint.Y = y;
      }
    };

    canvas.onmouseout = function (event) { if (Step == "zero") { Step = ""; alert("failed,click the panel to retry.") } }
    canvas.onmousemove = function (event) {
      CurrentPoint = { X: event.pageX - canvas.offsetLeft, Y: event.pageY - canvas.offsetTop };
    }




    function drawPoint(x, y, color) {
      ctx.beginPath();
      ctx.rect(x, y, 2, 2);
      ctx.fillStyle = color;
      ctx.fill();

    }
    function drawText(x, y) {
      ctx.fillText("" + (x - centerX) + "," + (y - centerY), x, y);
    }
    function drawLine(sx, sy, dx, dy) {
      ctx.beginPath();
      ctx.moveTo(sx, sy);
      ctx.lineTo(dx, dy);
      ctx.strokeStyle = '#ff0000';
      ctx.stroke();
    }

    function drawGrid() {
      var grid_size = 10;
      var canvas_width = canvas.width;
      var canvas_height = canvas.height;
      var num_lines_x = Math.floor(canvas_height / grid_size);
      var num_lines_y = Math.floor(canvas_width / grid_size);
      for (var i = 0; i <= num_lines_x; i++) {
        ctx.beginPath();
        ctx.lineWidth = 1;
        ctx.strokeStyle = "#efefef";
        if (i == num_lines_x) {
          ctx.moveTo(0, grid_size * i);
          ctx.lineTo(canvas_width, grid_size * i);
        }
        else {
          ctx.moveTo(0, grid_size * i + 0.5);
          ctx.lineTo(canvas_width, grid_size * i + 0.5);
        }
        ctx.stroke();
      }
      // Draw grid lines along Y-axis
      for (i = 0; i <= num_lines_y; i++) {
        ctx.beginPath();
        ctx.lineWidth = 1;
        ctx.strokeStyle = "#efefef";
        if (i == num_lines_y) {
          ctx.moveTo(grid_size * i, 0);
          ctx.lineTo(grid_size * i, canvas_height);
        }
        else {
          ctx.moveTo(grid_size * i + 0.5, 0);
          ctx.lineTo(grid_size * i + 0.5, canvas_height);
        }
        ctx.stroke();
      }
    }



    var CurrentPoint = { X: -1, Y: -1 };
    var LastPoint = { X: 0, Y: 0 };
    var MoveTimer = setInterval(Frame, 20);
    function Frame() {
      if (Step == "zero") {
        var xOffset = centerX - CurrentPoint.X;
        var yOffset = centerY - CurrentPoint.Y;
        if (xOffset != 0 || yOffset != 0) {
          console.log("zero: ", CurrentPoint.X, "," + CurrentPoint.Y + " -> " + centerX + "," + centerY)
          if (xOffset != 0)
            $.get("/win?hid=M:X:" + (xOffset < 0 ? -1 : 1), function (data) { });
          if (yOffset != 0)
            $.get("/win?hid=M:Y:" + (yOffset < 0 ? -1 : 1), function (data) { });
        } else {
          Step = "draw";
          clearInterval(MoveTimer);
          LastPoint.X = centerX;
          LastPoint.Y = centerY;
          drawPoint(centerX, centerY, "black")
          setTimeout(function () {
            $.get("/win?draw=1&centerx=" + centerX + "&centery=" + centerY, function (data) { });
          }, 1000);

        }
      }
    }



  </script>
</body>

</html>
`)
