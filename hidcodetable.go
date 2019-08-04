package main

import (
	"github.com/sirupsen/logrus"
)

const KEY_A byte = 4
const KEY_B byte = 5
const KEY_C byte = 6
const KEY_D byte = 7
const KEY_E byte = 8
const KEY_F byte = 9
const KEY_G byte = 10
const KEY_H byte = 11
const KEY_I byte = 12
const KEY_J byte = 13
const KEY_K byte = 14
const KEY_L byte = 15
const KEY_M byte = 16
const KEY_N byte = 17
const KEY_O byte = 18
const KEY_P byte = 19
const KEY_Q byte = 20
const KEY_R byte = 21
const KEY_S byte = 22
const KEY_T byte = 23
const KEY_U byte = 24
const KEY_V byte = 25
const KEY_W byte = 26
const KEY_X byte = 27
const KEY_Y byte = 28
const KEY_Z byte = 29
const KEY_1 byte = 30
const KEY_2 byte = 31
const KEY_3 byte = 32
const KEY_4 byte = 33
const KEY_5 byte = 34
const KEY_6 byte = 35
const KEY_7 byte = 36
const KEY_8 byte = 37
const KEY_9 byte = 38
const KEY_0 byte = 39
const KEY_ENTER byte = 40
const KEY_ESC byte = 41
const KEY_BACKSPACE byte = 42
const KEY_TAB byte = 43
const KEY_SPACE byte = 44
const KEY_MINUS byte = 45
const KEY_EQUAL byte = 46
const KEY_LEFTBRACE byte = 47
const KEY_RIGHTBRACE byte = 48
const KEY_BACKSLASH byte = 49
const KEY_SEMICOLON byte = 51
const KEY_APOSTROPHE byte = 52
const KEY_GRAVE byte = 53
const KEY_COMMA byte = 54
const KEY_DOT byte = 55
const KEY_SLASH byte = 56
const KEY_CAPSLOCK byte = 57
const KEY_F1 byte = 58
const KEY_F2 byte = 59
const KEY_F3 byte = 60
const KEY_F4 byte = 61
const KEY_F5 byte = 62
const KEY_F6 byte = 63
const KEY_F7 byte = 64
const KEY_F8 byte = 65
const KEY_F9 byte = 66
const KEY_F10 byte = 67
const KEY_F11 byte = 68
const KEY_F12 byte = 69
const KEY_SYSRQ byte = 70
const KEY_SCROLLLOCK byte = 71
const KEY_PAUSE byte = 72
const KEY_INSERT byte = 73
const KEY_HOME byte = 74
const KEY_PAGEUP byte = 75
const KEY_DELETE byte = 76
const KEY_END byte = 77
const KEY_PAGEDOWN byte = 78
const KEY_RIGHT byte = 79
const KEY_LEFT byte = 80
const KEY_DOWN byte = 81
const KEY_UP byte = 82
const KEY_NUMLOCK byte = 83
const KEY_KPSLASH byte = 84
const KEY_KPASTERISK byte = 85
const KEY_KPMINUS byte = 86
const KEY_KPPLUS byte = 87
const KEY_KPENTER byte = 88
const KEY_KP1 byte = 89
const KEY_KP2 byte = 90
const KEY_KP3 byte = 91
const KEY_KP4 byte = 92
const KEY_KP5 byte = 93
const KEY_KP6 byte = 94
const KEY_KP7 byte = 95
const KEY_KP8 byte = 96
const KEY_KP9 byte = 97
const KEY_KP0 byte = 98
const KEY_KPDOT byte = 99
const KEY_102ND byte = 100
const KEY_COMPOSE byte = 101
const KEY_POWER byte = 102
const KEY_KPEQUAL byte = 103
const KEY_F13 byte = 104
const KEY_F14 byte = 105
const KEY_F15 byte = 106
const KEY_F16 byte = 107
const KEY_F17 byte = 108
const KEY_F18 byte = 109
const KEY_F19 byte = 110
const KEY_F20 byte = 111
const KEY_F21 byte = 112
const KEY_F22 byte = 113
const KEY_F23 byte = 114
const KEY_F24 byte = 115
const KEY_OPEN byte = 116
const KEY_HELP byte = 117
const KEY_PROPS byte = 118
const KEY_FRONT byte = 119
const KEY_STOP byte = 120
const KEY_AGAIN byte = 121
const KEY_UNDO byte = 122
const KEY_CUT byte = 123
const KEY_COPY byte = 124
const KEY_PASTE byte = 125
const KEY_FIND byte = 126
const KEY_MUTE byte = 127
const KEY_VOLUMEUP byte = 128
const KEY_VOLUMEDOWN byte = 129
const KEY_KPCOMMA byte = 133
const KEY_RO byte = 135
const KEY_KATAKANAHIRAGANA byte = 136
const KEY_YEN byte = 137
const KEY_HENKAN byte = 138
const KEY_MUHENKAN byte = 139
const KEY_KPJPCOMMA byte = 140
const KEY_HANJA byte = 145
const KEY_KATAKANA byte = 146
const KEY_HIRAGANA byte = 147
const KEY_ZENKAKUHANKAKU byte = 148
const KEY_KPLEFTPAREN byte = 182
const KEY_KPRIGHTPAREN byte = 183

const KEY_LEFTCTRL byte = 1
const KEY_LEFTSHIFT byte = 2
const KEY_LEFTALT byte = 4
const KEY_LEFTMETA byte = 8
const KEY_RIGHTCTRL byte = 16
const KEY_RIGHTSHIFT byte = 32
const KEY_RIGHTALT byte = 64
const KEY_RIGHTMETA byte = 128

var HIDNAME map[int]string

var HIDMAP map[uint16]byte

var MODNAME map[int]string

var MOUSENAME map[int]string

var MODMAP map[uint16]byte
var MOUSEMAP map[uint16]byte

func init() {

	logrus.Info("Keytable inited.");

	MOUSEMAP = make(map[uint16]byte)
	/*	BTN_LEFT                     = 0x110	272
		BTN_RIGHT                    = 0x111	273
		BTN_MIDDLE                   = 0x112	274
		BTN_SIDE                     = 0x113	275
		BTN_EXTRA                    = 0x114	276
		BTN_FORWARD                  = 0x115	277
		BTN_BACK                     = 0x116	278
		BTN_TASK                     = 0x117	279

	*/
	MOUSEMAP[272] = 0x01
	MOUSEMAP[273] = 0x02
	MOUSEMAP[274] = 0x04

	MOUSEMAP[275] = 0x08
	MOUSEMAP[276] = 0x10

	MOUSEMAP[277] = 0x20
	MOUSEMAP[278] = 0x40

	MOUSEMAP[279] = 0x80

	MOUSENAME = make(map[int]string)
	MOUSENAME[1] = "M.LEFT"
	MOUSENAME[2] = "M.RIGHT"
	MOUSENAME[4] = "M.MIDDLE"
	MOUSENAME[8] = "M.SIDE"
	MOUSENAME[16] = "M.EXTRA"
	MOUSENAME[32] = "M.FORWARD"
	MOUSENAME[64] = "M.BACK"
	MOUSENAME[128] = "M.TASK"

	MODMAP = make(map[uint16]byte)
	MODMAP[29] = KEY_LEFTCTRL
	MODMAP[42] = KEY_LEFTSHIFT
	MODMAP[56] = KEY_LEFTALT
	MODMAP[125] = KEY_LEFTMETA

	MODMAP[97] = KEY_RIGHTCTRL
	MODMAP[54] = KEY_RIGHTSHIFT
	MODMAP[100] = KEY_RIGHTALT
	MODMAP[126] = KEY_RIGHTMETA

	MODNAME = make(map[int]string)
	MODNAME[int(KEY_LEFTCTRL)] = "L.CTRL"
	MODNAME[int(KEY_LEFTSHIFT)] = "L.SHIFT"
	MODNAME[int(KEY_LEFTALT)] = "L.ALT"
	MODNAME[int(KEY_LEFTMETA)] = "L.META"
	MODNAME[int(KEY_RIGHTCTRL)] = "R.CTRL"
	MODNAME[int(KEY_RIGHTSHIFT)] = "R.SHIFT"
	MODNAME[int(KEY_RIGHTALT)] = "R.ALT"
	MODNAME[int(KEY_RIGHTMETA)] = "R.META"

	HIDMAP = make(map[uint16]byte)
	HIDMAP[1] = KEY_ESC               //41
	HIDMAP[2] = KEY_1                 //30
	HIDMAP[3] = KEY_2                 //31
	HIDMAP[4] = KEY_3                 //32
	HIDMAP[5] = KEY_4                 //33
	HIDMAP[6] = KEY_5                 //34
	HIDMAP[7] = KEY_6                 //35
	HIDMAP[8] = KEY_7                 //36
	HIDMAP[9] = KEY_8                 //37
	HIDMAP[10] = KEY_9                //38
	HIDMAP[11] = KEY_0                //39
	HIDMAP[12] = KEY_MINUS            //45
	HIDMAP[13] = KEY_EQUAL            //46
	HIDMAP[14] = KEY_BACKSPACE        //42
	HIDMAP[15] = KEY_TAB              //43
	HIDMAP[16] = KEY_Q                //20
	HIDMAP[17] = KEY_W                //26
	HIDMAP[18] = KEY_E                //8
	HIDMAP[19] = KEY_R                //21
	HIDMAP[20] = KEY_T                //23
	HIDMAP[21] = KEY_Y                //28
	HIDMAP[22] = KEY_U                //24
	HIDMAP[23] = KEY_I                //12
	HIDMAP[24] = KEY_O                //18
	HIDMAP[25] = KEY_P                //19
	HIDMAP[26] = KEY_LEFTBRACE        //47
	HIDMAP[27] = KEY_RIGHTBRACE       //48
	HIDMAP[28] = KEY_ENTER            //40
	HIDMAP[30] = KEY_A                //4
	HIDMAP[31] = KEY_S                //22
	HIDMAP[32] = KEY_D                //7
	HIDMAP[33] = KEY_F                //9
	HIDMAP[34] = KEY_G                //10
	HIDMAP[35] = KEY_H                //11
	HIDMAP[36] = KEY_J                //13
	HIDMAP[37] = KEY_K                //14
	HIDMAP[38] = KEY_L                //15
	HIDMAP[39] = KEY_SEMICOLON        //51
	HIDMAP[40] = KEY_APOSTROPHE       //52
	HIDMAP[41] = KEY_GRAVE            //53
	HIDMAP[43] = KEY_BACKSLASH        //49
	HIDMAP[44] = KEY_Z                //29
	HIDMAP[45] = KEY_X                //27
	HIDMAP[46] = KEY_C                //6
	HIDMAP[47] = KEY_V                //25
	HIDMAP[48] = KEY_B                //5
	HIDMAP[49] = KEY_N                //17
	HIDMAP[50] = KEY_M                //16
	HIDMAP[51] = KEY_COMMA            //54
	HIDMAP[52] = KEY_DOT              //55
	HIDMAP[53] = KEY_SLASH            //56
	HIDMAP[55] = KEY_KPASTERISK       //85
	HIDMAP[57] = KEY_SPACE            //44
	HIDMAP[58] = KEY_CAPSLOCK         //57
	HIDMAP[59] = KEY_F1               //58
	HIDMAP[60] = KEY_F2               //59
	HIDMAP[61] = KEY_F3               //60
	HIDMAP[62] = KEY_F4               //61
	HIDMAP[63] = KEY_F5               //62
	HIDMAP[64] = KEY_F6               //63
	HIDMAP[65] = KEY_F7               //64
	HIDMAP[66] = KEY_F8               //65
	HIDMAP[67] = KEY_F9               //66
	HIDMAP[68] = KEY_F10              //67
	HIDMAP[69] = KEY_NUMLOCK          //83
	HIDMAP[70] = KEY_SCROLLLOCK       //71
	HIDMAP[71] = KEY_KP7              //95
	HIDMAP[72] = KEY_KP8              //96
	HIDMAP[73] = KEY_KP9              //97
	HIDMAP[74] = KEY_KPMINUS          //86
	HIDMAP[75] = KEY_KP4              //92
	HIDMAP[76] = KEY_KP5              //93
	HIDMAP[77] = KEY_KP6              //94
	HIDMAP[78] = KEY_KPPLUS           //87
	HIDMAP[79] = KEY_KP1              //89
	HIDMAP[80] = KEY_KP2              //90
	HIDMAP[81] = KEY_KP3              //91
	HIDMAP[82] = KEY_KP0              //98
	HIDMAP[83] = KEY_KPDOT            //99
	HIDMAP[85] = KEY_ZENKAKUHANKAKU   //148
	HIDMAP[86] = KEY_102ND            //100
	HIDMAP[87] = KEY_F11              //68
	HIDMAP[88] = KEY_F12              //69
	HIDMAP[89] = KEY_RO               //135
	HIDMAP[90] = KEY_KATAKANA         //146
	HIDMAP[91] = KEY_HIRAGANA         //147
	HIDMAP[92] = KEY_HENKAN           //138
	HIDMAP[93] = KEY_KATAKANAHIRAGANA //136
	HIDMAP[94] = KEY_MUHENKAN         //139
	HIDMAP[95] = KEY_KPJPCOMMA        //140
	HIDMAP[96] = KEY_KPENTER          //88
	HIDMAP[98] = KEY_KPSLASH          //84
	HIDMAP[99] = KEY_SYSRQ            //70
	HIDMAP[102] = KEY_HOME            //74
	HIDMAP[103] = KEY_UP              //82
	HIDMAP[104] = KEY_PAGEUP          //75
	HIDMAP[105] = KEY_LEFT            //80
	HIDMAP[106] = KEY_RIGHT           //79
	HIDMAP[107] = KEY_END             //77
	HIDMAP[108] = KEY_DOWN            //81
	HIDMAP[109] = KEY_PAGEDOWN        //78
	HIDMAP[110] = KEY_INSERT          //73
	HIDMAP[111] = KEY_DELETE          //76
	HIDMAP[113] = KEY_MUTE            //127
	HIDMAP[114] = KEY_VOLUMEDOWN      //129
	HIDMAP[115] = KEY_VOLUMEUP        //128
	HIDMAP[116] = KEY_POWER           //102
	HIDMAP[117] = KEY_KPEQUAL         //103
	HIDMAP[119] = KEY_PAUSE           //72
	HIDMAP[121] = KEY_KPCOMMA         //133
	//HIDMAP[122] = KEY_HANGEUL         //144
	HIDMAP[123] = KEY_HANJA        //145
	HIDMAP[124] = KEY_YEN          //137
	HIDMAP[127] = KEY_COMPOSE      //101
	HIDMAP[128] = KEY_STOP         //120
	HIDMAP[129] = KEY_AGAIN        //121
	HIDMAP[130] = KEY_PROPS        //118
	HIDMAP[131] = KEY_UNDO         //122
	HIDMAP[132] = KEY_FRONT        //119
	HIDMAP[133] = KEY_COPY         //124
	HIDMAP[134] = KEY_OPEN         //116
	HIDMAP[135] = KEY_PASTE        //125
	HIDMAP[136] = KEY_FIND         //126
	HIDMAP[137] = KEY_CUT          //123
	HIDMAP[138] = KEY_HELP         //117
	HIDMAP[179] = KEY_KPLEFTPAREN  //182
	HIDMAP[180] = KEY_KPRIGHTPAREN //183
	HIDMAP[183] = KEY_F13          //104
	HIDMAP[184] = KEY_F14          //105
	HIDMAP[185] = KEY_F15          //106
	HIDMAP[186] = KEY_F16          //107
	HIDMAP[187] = KEY_F17          //108
	HIDMAP[188] = KEY_F18          //109
	HIDMAP[189] = KEY_F19          //110
	HIDMAP[190] = KEY_F20          //111
	HIDMAP[191] = KEY_F21          //112
	HIDMAP[192] = KEY_F22          //113
	HIDMAP[193] = KEY_F23          //114
	HIDMAP[194] = KEY_F24          //115

	HIDNAME = make(map[int]string)
	HIDNAME[int(KEY_ESC)] = "ESC"                           //41
	HIDNAME[int(KEY_1)] = "1"                               //30
	HIDNAME[int(KEY_2)] = "2"                               //31
	HIDNAME[int(KEY_3)] = "3"                               //32
	HIDNAME[int(KEY_4)] = "4"                               //33
	HIDNAME[int(KEY_5)] = "5"                               //34
	HIDNAME[int(KEY_6)] = "6"                               //35
	HIDNAME[int(KEY_7)] = "7"                               //36
	HIDNAME[int(KEY_8)] = "8"                               //37
	HIDNAME[int(KEY_9)] = "9"                               //38
	HIDNAME[int(KEY_0)] = "0"                               //39
	HIDNAME[int(KEY_MINUS)] = "MINUS"                       //45
	HIDNAME[int(KEY_EQUAL)] = "EQUAL"                       //46
	HIDNAME[int(KEY_BACKSPACE)] = "BACKSPACE"               //42
	HIDNAME[int(KEY_TAB)] = "TAB"                           //43
	HIDNAME[int(KEY_Q)] = "Q"                               //20
	HIDNAME[int(KEY_W)] = "W"                               //26
	HIDNAME[int(KEY_E)] = "E"                               //8
	HIDNAME[int(KEY_R)] = "R"                               //21
	HIDNAME[int(KEY_T)] = "T"                               //23
	HIDNAME[int(KEY_Y)] = "Y"                               //28
	HIDNAME[int(KEY_U)] = "U"                               //24
	HIDNAME[int(KEY_I)] = "I"                               //12
	HIDNAME[int(KEY_O)] = "O"                               //18
	HIDNAME[int(KEY_P)] = "P"                               //19
	HIDNAME[int(KEY_LEFTBRACE)] = "LEFTBRACE"               //47
	HIDNAME[int(KEY_RIGHTBRACE)] = "RIGHTBRACE"             //48
	HIDNAME[int(KEY_ENTER)] = "ENTER"                       //40
	HIDNAME[int(KEY_A)] = "A"                               //4
	HIDNAME[int(KEY_S)] = "S"                               //22
	HIDNAME[int(KEY_D)] = "D"                               //7
	HIDNAME[int(KEY_F)] = "F"                               //9
	HIDNAME[int(KEY_G)] = "G"                               //10
	HIDNAME[int(KEY_H)] = "H"                               //11
	HIDNAME[int(KEY_J)] = "J"                               //13
	HIDNAME[int(KEY_K)] = "K"                               //14
	HIDNAME[int(KEY_L)] = "L"                               //15
	HIDNAME[int(KEY_SEMICOLON)] = "SEMICOLON"               //51
	HIDNAME[int(KEY_APOSTROPHE)] = "APOSTROPHE"             //52
	HIDNAME[int(KEY_GRAVE)] = "GRAVE"                       //53
	HIDNAME[int(KEY_BACKSLASH)] = "BACKSLASH"               //49
	HIDNAME[int(KEY_Z)] = "Z"                               //29
	HIDNAME[int(KEY_X)] = "X"                               //27
	HIDNAME[int(KEY_C)] = "C"                               //6
	HIDNAME[int(KEY_V)] = "V"                               //25
	HIDNAME[int(KEY_B)] = "B"                               //5
	HIDNAME[int(KEY_N)] = "N"                               //17
	HIDNAME[int(KEY_M)] = "M"                               //16
	HIDNAME[int(KEY_COMMA)] = "COMMA"                       //54
	HIDNAME[int(KEY_DOT)] = "DOT"                           //55
	HIDNAME[int(KEY_SLASH)] = "SLASH"                       //56
	HIDNAME[int(KEY_KPASTERISK)] = "KPASTERISK"             //85
	HIDNAME[int(KEY_SPACE)] = "SPACE"                       //44
	HIDNAME[int(KEY_CAPSLOCK)] = "CAPSLOCK"                 //57
	HIDNAME[int(KEY_F1)] = "F1"                             //58
	HIDNAME[int(KEY_F2)] = "F2"                             //59
	HIDNAME[int(KEY_F3)] = "F3"                             //60
	HIDNAME[int(KEY_F4)] = "F4"                             //61
	HIDNAME[int(KEY_F5)] = "F5"                             //62
	HIDNAME[int(KEY_F6)] = "F6"                             //63
	HIDNAME[int(KEY_F7)] = "F7"                             //64
	HIDNAME[int(KEY_F8)] = "F8"                             //65
	HIDNAME[int(KEY_F9)] = "F9"                             //66
	HIDNAME[int(KEY_F10)] = "F10"                           //67
	HIDNAME[int(KEY_NUMLOCK)] = "NUMLOCK"                   //83
	HIDNAME[int(KEY_SCROLLLOCK)] = "SCROLLLOCK"             //71
	HIDNAME[int(KEY_KP7)] = "KP7"                           //95
	HIDNAME[int(KEY_KP8)] = "KP8"                           //96
	HIDNAME[int(KEY_KP9)] = "KP9"                           //97
	HIDNAME[int(KEY_KPMINUS)] = "KPMINUS"                   //86
	HIDNAME[int(KEY_KP4)] = "KP4"                           //92
	HIDNAME[int(KEY_KP5)] = "KP5"                           //93
	HIDNAME[int(KEY_KP6)] = "KP6"                           //94
	HIDNAME[int(KEY_KPPLUS)] = "KPPLUS"                     //87
	HIDNAME[int(KEY_KP1)] = "KP1"                           //89
	HIDNAME[int(KEY_KP2)] = "KP2"                           //90
	HIDNAME[int(KEY_KP3)] = "KP3"                           //91
	HIDNAME[int(KEY_KP0)] = "KP0"                           //98
	HIDNAME[int(KEY_KPDOT)] = "KPDOT"                       //99
	HIDNAME[int(KEY_ZENKAKUHANKAKU)] = "ZENKAKUHANKAKU"     //148
	HIDNAME[int(KEY_102ND)] = "102ND"                       //100
	HIDNAME[int(KEY_F11)] = "F11"                           //68
	HIDNAME[int(KEY_F12)] = "F12"                           //69
	HIDNAME[int(KEY_RO)] = "RO"                             //135
	HIDNAME[int(KEY_KATAKANA)] = "KATAKANA"                 //146
	HIDNAME[int(KEY_HIRAGANA)] = "HIRAGANA"                 //147
	HIDNAME[int(KEY_HENKAN)] = "HENKAN"                     //138
	HIDNAME[int(KEY_KATAKANAHIRAGANA)] = "KATAKANAHIRAGANA" //136
	HIDNAME[int(KEY_MUHENKAN)] = "MUHENKAN"                 //139
	HIDNAME[int(KEY_KPJPCOMMA)] = "KPJPCOMMA"               //140
	HIDNAME[int(KEY_KPENTER)] = "KPENTER"                   //88
	HIDNAME[int(KEY_KPSLASH)] = "KPSLASH"                   //84
	HIDNAME[int(KEY_SYSRQ)] = "SYSRQ"                       //70
	HIDNAME[int(KEY_HOME)] = "HOME"                         //74
	HIDNAME[int(KEY_UP)] = "UP"                             //82
	HIDNAME[int(KEY_PAGEUP)] = "PAGEUP"                     //75
	HIDNAME[int(KEY_LEFT)] = "LEFT"                         //80
	HIDNAME[int(KEY_RIGHT)] = "RIGHT"                       //79
	HIDNAME[int(KEY_END)] = "END"                           //77
	HIDNAME[int(KEY_DOWN)] = "DOWN"                         //81
	HIDNAME[int(KEY_PAGEDOWN)] = "PAGEDOWN"                 //78
	HIDNAME[int(KEY_INSERT)] = "INSERT"                     //73
	HIDNAME[int(KEY_DELETE)] = "DELETE"                     //76
	HIDNAME[int(KEY_MUTE)] = "MUTE"                         //127
	HIDNAME[int(KEY_VOLUMEDOWN)] = "VOLUMEDOWN"             //129
	HIDNAME[int(KEY_VOLUMEUP)] = "VOLUMEUP"                 //128
	HIDNAME[int(KEY_POWER)] = "POWER"                       //102
	HIDNAME[int(KEY_KPEQUAL)] = "KPEQUAL"                   //103
	HIDNAME[int(KEY_PAUSE)] = "PAUSE"                       //72
	HIDNAME[int(KEY_KPCOMMA)] = "KPCOMMA"                   //133
	//HIDNAME[int(KEY_HANGEUL)] = "HANGEUL"         //144
	HIDNAME[int(KEY_HANJA)] = "HANJA"               //145
	HIDNAME[int(KEY_YEN)] = "YEN"                   //137
	HIDNAME[int(KEY_COMPOSE)] = "COMPOSE"           //101
	HIDNAME[int(KEY_STOP)] = "STOP"                 //120
	HIDNAME[int(KEY_AGAIN)] = "AGAIN"               //121
	HIDNAME[int(KEY_PROPS)] = "PROPS"               //118
	HIDNAME[int(KEY_UNDO)] = "UNDO"                 //122
	HIDNAME[int(KEY_FRONT)] = "FRONT"               //119
	HIDNAME[int(KEY_COPY)] = "COPY"                 //124
	HIDNAME[int(KEY_OPEN)] = "OPEN"                 //116
	HIDNAME[int(KEY_PASTE)] = "PASTE"               //125
	HIDNAME[int(KEY_FIND)] = "FIND"                 //126
	HIDNAME[int(KEY_CUT)] = "CUT"                   //123
	HIDNAME[int(KEY_HELP)] = "HELP"                 //117
	HIDNAME[int(KEY_KPLEFTPAREN)] = "KPLEFTPAREN"   //182
	HIDNAME[int(KEY_KPRIGHTPAREN)] = "KPRIGHTPAREN" //183
	HIDNAME[int(KEY_F13)] = "F13"                   //104
	HIDNAME[int(KEY_F14)] = "F14"                   //105
	HIDNAME[int(KEY_F15)] = "F15"                   //106
	HIDNAME[int(KEY_F16)] = "F16"                   //107
	HIDNAME[int(KEY_F17)] = "F17"                   //108
	HIDNAME[int(KEY_F18)] = "F18"                   //109
	HIDNAME[int(KEY_F19)] = "F19"                   //110
	HIDNAME[int(KEY_F20)] = "F20"                   //111
	HIDNAME[int(KEY_F21)] = "F21"                   //112
	HIDNAME[int(KEY_F22)] = "F22"                   //113
	HIDNAME[int(KEY_F23)] = "F23"                   //114
	HIDNAME[int(KEY_F24)] = "F24"                   //115
}
