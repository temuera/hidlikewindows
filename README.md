# hidlikewindows
use pi zero make mac/linux mouse keyboard like windows


![settings](https://github.com/dazhoudotnet/hidlikewindows/blob/master/settings.png?raw=true)

![keyboard remap](https://github.com/dazhoudotnet/hidlikewindows/blob/master/keyboard_remap.png?raw=true)

![mouse button remap](https://github.com/dazhoudotnet/hidlikewindows/blob/master/mouse_remap.png?raw=true)

![mouse speed test](https://github.com/dazhoudotnet/hidlikewindows/blob/master/mousespeedtest.png?raw=true)


 **Host PI Zero Config**

`config.txt`
dtoverlay=pi3-disable-bt
enable_uart=1
init_uart_clock=64000000
init_uart_baud=4000000

`cmdline.txt`
dwc_otg.lpm_enable=0 console=tty1 root=PARTUUID=a34c07ff-02 rootfstype=ext4 elevator=deadline fsck.repair=yes usbhid.mousepoll=2  rootwait


 **Slave PI Zero Config**
`config.txt`
dtoverlay=dwc2 #Open otg mode
dr_mode=peripheral #unknow
dtoverlay=pi3-disable-bt # disable bt and use hardware serialport.
enable_uart=1
init_uart_clock=64000000
init_uart_baud=4000000

`cmdline.txt`
dwc_otg.lpm_enable=0 console=tty1 root=PARTUUID=18d2a8c5-02 rootfstype=ext4 elevator=deadline fsck.repair=yes rootwait

