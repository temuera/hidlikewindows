# hidlikewindows
use pi zero make mac/linux mouse keyboard like windows

 **Host PI Zero Config**

`config.txt`
dtoverlay=pi3-disable-bt
enable_uart=1
init_uart_clock=64000000
init_uart_baud=4000000

`cmdline.txt`
dwc_otg.lpm_enable=0 console=tty1 root=PARTUUID=a34c07ff-02 rootfstype=ext4 elevator=deadline fsck.repair=yes usbhid.mousepoll=2  rootwait
