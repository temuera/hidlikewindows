#!/bin/bash

modprobe dwc2
modprobe libcomposite
# remove all USB Ethernet drivers
modprobe -r g_ether usb_f_ecm u_ether

# insert modules for HID and ECM Ethernet
modprobe usb_f_hid
#modprobe usb_f_ecm
modprobe usb_f_rndis

 

mkdir -p /sys/kernel/config/usb_gadget/hidlikewindows  #just a name
cd /sys/kernel/config/usb_gadget/hidlikewindows


echo 0x1d6b > idVendor # Linux Foundation fixed
echo 0x0104 > idProduct # Multifunction Composite Gadget  fixed
 
echo 0x0100 > bcdDevice # v1.0.0 fixed
echo 0x0200 > bcdUSB # USB2 fixed

#Composite USB devices with multiple functions need to indicate this to Windows by using a special class & protocol code. 
echo 0xEF > bDeviceClass
echo 0x02 > bDeviceSubClass
echo 0x01 > bDeviceProtocol

#OS descriptors 
echo 1       > os_desc/use
echo 0xcd    > os_desc/b_vendor_code
echo MSFT100 > os_desc/qw_sign



#名子随便起,序列号随便填
mkdir -p strings/0x409
echo "fedcba9876543210" > strings/0x409/serialnumber
echo "github.com/dazhoudotnet/hidlikewindows" > strings/0x409/manufacturer
echo "github.com/dazhoudotnet/hidlikewindows" > strings/0x409/product

mkdir -p configs/c.1/strings/0x409
echo "HID" > configs/c.1/strings/0x409/configuration
echo 500 > configs/c.1/MaxPower


#网络 这里泥马浪费了老子至少20小时, rndis设备一定要排在HID的前面
mkdir -p functions/rndis.usb0
echo 1a:55:89:a2:69:31 > functions/rndis.usb0/dev_addr
echo 1a:55:89:a2:69:32 > functions/rndis.usb0/host_addr
echo RNDIS      > functions/rndis.usb0/os_desc/interface.rndis/compatible_id
echo 5162001 > functions/rndis.usb0/os_desc/interface.rndis/sub_compatible_id
ln -s functions/rndis.usb0 configs/c.1/


# 键盘 ,标准104PC
mkdir -p functions/hid.usb0
echo 1 > functions/hid.usb0/protocol
echo 1 > functions/hid.usb0/subclass
echo 8 > functions/hid.usb0/report_length
echo -ne \\x05\\x01\\x09\\x06\\xa1\\x01\\x05\\x07\\x19\\xe0\\x29\\xe7\\x15\\x00\\x25\\x01\\x75\\x01\\x95\\x08\\x81\\x02\\x95\\x01\\x75\\x08\\x81\\x03\\x95\\x05\\x75\\x01\\x05\\x08\\x19\\x01\\x29\\x05\\x91\\x02\\x95\\x01\\x75\\x03\\x91\\x03\\x95\\x06\\x75\\x08\\x15\\x00\\x25\\x65\\x05\\x07\\x19\\x00\\x29\\x65\\x81\\x00\\xc0 > functions/hid.usb0/report_desc
ln -s functions/hid.usb0 configs/c.1/
 

#鼠标,这是5-8键鼠标的配置 .但正常鼠标是 1-5键以鼠标消息发出, 右侧的键则是模拟键盘消息page down up 发出
mkdir -p functions/hid.usb1
echo 1 > functions/hid.usb1/protocol
echo 1 > functions/hid.usb1/subclass
echo 8 > functions/hid.usb1/report_length
echo -ne \\x05\\x01\\x09\\x02\\xa1\\x01\\x09\\x01\\xa1\\x00\\x05\\x09\\x19\\x01\\x29\\x08\\x15\\x00\\x25\\x01\\x95\\x08\\x75\\x01\\x81\\x02\\x05\\x01\\x09\\x30\\x09\\x31\\x16\\x00\\x80\\x26\\xff\\x7f\\x75\\x10\\x95\\x02\\x81\\x06\\x09\\x38\\x15\\x81\\x25\\x7f\\x75\\x08\\x95\\x01\\x81\\x06\\xc0\\xc0 > functions/hid.usb1/report_desc
ln -s functions/hid.usb1 configs/c.1/


#U盘,用来放说明文档

# if [[ ! -e /home/usbdisk.img ]]; then
#   dd if=/dev/zero of=/home/usbdisk.img bs=1M count=32
#   mkdosfs /home/usbdisk.img
# fi
#  FILE=/home/usbdisk.img
#  mkdir -p ${FILE/img/d}
#  mount -o loop,ro, -t vfat $FILE ${FILE/img/d} # FOR IMAGE CREATED WITH DD

#  mkdir -p functions/mass_storage.usb2
#  echo 1 > functions/mass_storage.usb2/stall
#  echo 0 > functions/mass_storage.usb2/lun.0/cdrom
#  echo 0 > functions/mass_storage.usb2/lun.0/ro
#  echo 0 > functions/mass_storage.usb2/lun.0/nofua
#  echo $FILE > functions/mass_storage.usb2/lun.0/file
#  ln -s functions/mass_storage.usb2 configs/c.1/


#ECM好像对 LINUX和MAC更友好?
#mkdir functions/ecm.usb0
#echo 1a:55:89:a2:69:41 > functions/ecm.usb0/dev_addr
#echo 1a:55:89:a2:69:42 > functions/ecm.usb0/host_addr
#ln -s functions/ecm.usb0 configs/c.1/


ln -s configs/c.1 os_desc

udevadm settle -t 5 || :
ls /sys/class/udc > UDC





     