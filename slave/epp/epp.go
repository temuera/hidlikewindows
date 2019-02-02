package epp

import (
	"math"
)

var previousMouseRawX int = 0
var previousMouseRawY int = 0
var previousMouseXRemainder float64 = 0.0
var previousMouseYRemainder float64 = 0.0

var previousSegmentIndex int = 0

//var screenResolution int = 96     // DPI
//var screenRefreshRate int = 60    // RefererRate
//var mouseSensitivity float64 = 20 // From registry HKEY_CURRENT_USER\Control Panel\Mouse\MouseSensitivity
var pixelGain float64 = 0.0

var FINDSEGMENT = -1 //插值使用

func smoothMouseGain(deviceSpeed float64, segment *int) float64 {
	/*
	   values of threshold that give the pointer speed in inches/s from
	   the speed of the device in inches/s intermediate values are
	   interpolated
	   http://www.microsoft.com/whdc/archive/pointer-bal.mspx

	   [HKEY_CURRENT_USER\Control Panel\Mouse]
	   "SmoothMouseXCurve"=hex:00,00,00,00,00,00,00,00,	\
	   15,6e,00,00,00,00,00,00,				\
	   00,40,01,00,00,00,00,00,				\
	   29,dc,03,00,00,00,00,00,				\
	   00,00,28,00,00,00,00,00
	   "SmoothMouseYCurve"=hex:00,00,00,00,00,00,00,00,	\
	   b8,5e,01,00,00,00,00,00,				\
	   cd,4c,05,00,00,00,00,00,				\
	   cd,4c,18,00,00,00,00,00,				\
	   00,00,38,02,00,00,00,00
	*/
	if deviceSpeed == 0.0 {
		*segment = 0
		return deviceSpeed
	}
	// smoothX := [5]float64{0.0, 0.43, 1.25, 3.86, 40.0}
	// smoothY := [5]float64{0.0, 1.37, 5.30, 24.30, 568.0}

	smoothX := [5]float64{0.0, 0.43001, 1.25, 3.86, 40.0}
	smoothY := [5]float64{0.0, 1.07027, 4.14063, 18.98438, 443.75}

	var i int
	if *segment == FINDSEGMENT {
		for i = 0; i < 3; i++ {
			if deviceSpeed < smoothX[i+1] {
				break
			}
		}
		*segment = i
	} else {
		i = *segment
	}
	slope := (smoothY[i+1] - smoothY[i]) / (smoothX[i+1] - smoothX[i])
	intercept := smoothY[i] - slope*smoothX[i]
	return slope + intercept/deviceSpeed
}

func Apply(eppFactor float64, mouseRawX, mouseRawY int) (mouseX, mouseY int) {
	mouseMag := float64(math.Max(math.Abs(float64(mouseRawX)), math.Abs(float64(mouseRawY))) +
		math.Min(math.Abs(float64(mouseRawX)), math.Abs(float64(mouseRawY)))/2.0)
	var currentSegmentIndex int = FINDSEGMENT
	pixelGain = eppFactor * smoothMouseGain(mouseMag/3.5, &currentSegmentIndex) / 3.5

	if currentSegmentIndex > previousSegmentIndex {
		pixelGainUsingPreviousSegment := eppFactor * smoothMouseGain(mouseMag/3.5, &previousSegmentIndex) / 3.5
		pixelGain = (pixelGain + pixelGainUsingPreviousSegment) / 2.0
	}
	previousSegmentIndex = currentSegmentIndex

	mouseXplusRemainder := float64(mouseRawX)*pixelGain + previousMouseXRemainder
	mouseYplusRemainder := float64(mouseRawY)*pixelGain + previousMouseYRemainder

	if mouseXplusRemainder >= 0 {
		mouseX = int(math.Floor(mouseXplusRemainder))
	} else {
		mouseX = -int(math.Floor(-mouseXplusRemainder))
	}
	previousMouseXRemainder = mouseXplusRemainder - float64(mouseX)

	if mouseYplusRemainder >= 0 {
		mouseY = int(math.Floor(mouseYplusRemainder))
	} else {
		mouseY = -int(math.Floor(-mouseYplusRemainder))
	}
	previousMouseYRemainder = mouseYplusRemainder - float64(mouseY)
	return
}
