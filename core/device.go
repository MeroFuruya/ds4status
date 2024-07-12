package core

import (
	"fmt"

	"github.com/sstallion/go-hid"
)

func IsDS4Device(info hid.DeviceInfo) bool {
	// print all devices
	// fmt.Printf("name: %s, VendorID: 0x%04x, ProductID: 0x%04x\n", info.ProductStr, info.VendorID, info.ProductID)
	if info.VendorID == 0x054C && info.ProductID == 0x05C4 {
		return true
	}
	if info.VendorID == 0x054C && info.ProductID == 0x09CC {
		return true
	}
	return false
}

func OpenDevice(device hid.DeviceInfo) (deviceHandle *hid.Device, err error) {
	return hid.Open(device.VendorID, device.ProductID, device.SerialNbr)
}

func CloseDevice(device *hid.Device) {
	device.Close()
}

func ReadDeviceState(device *hid.Device) (buf []byte, bytesWritten int, err error) {
	buf = make([]byte, 127)
	bytesWritten, err = device.Read(buf)
	return
}

func PrintDeviceState(device *hid.Device) {
	buf, bytesWritten, err := ReadDeviceState(device)
	if err != nil {
		return
	}

	fmt.Printf("bytesWritten: %d\n", bytesWritten)

	fmt.Printf("LX: %d, LY: %d, RX: %d, RY: %d\n", buf[1], buf[2], buf[3], buf[4])
	fmt.Printf("L2: %d, R2: %d\n", buf[8], buf[9])

	triangle := buf[5]&0x80 != 0
	circle := buf[5]&0x40 != 0
	cross := buf[5]&0x20 != 0
	square := buf[5]&0x10 != 0
	fmt.Printf("Triangle: %t, Circle: %t, Cross: %t, Square: %t\n", triangle, circle, cross, square)

	switch buf[5] & 0x0F {
	case 0x00:
		fmt.Println("DPAD: UP")
	case 0x01:
		fmt.Println("DPAD: UP-RIGHT")
	case 0x02:
		fmt.Println("DPAD: RIGHT")
	case 0x03:
		fmt.Println("DPAD: DOWN-RIGHT")
	case 0x04:
		fmt.Println("DPAD: DOWN")
	case 0x05:
		fmt.Println("DPAD: DOWN-LEFT")
	case 0x06:
		fmt.Println("DPAD: LEFT")
	case 0x07:
		fmt.Println("DPAD: UP-LEFT")
	default:
		fmt.Println("DPAD: CENTER")
	}

	R3 := buf[6]&0x80 != 0
	L3 := buf[6]&0x40 != 0
	Options := buf[6]&0x20 != 0
	Share := buf[6]&0x10 != 0
	R2Btn := buf[6]&0x08 != 0
	L2Btn := buf[6]&0x04 != 0
	R1 := buf[6]&0x02 != 0
	L1 := buf[6]&0x01 != 0
	fmt.Printf("R3: %t, L3: %t, Options: %t, Share: %t, R2: %t, L2: %t, R1: %t, L1: %t\n", R3, L3, Options, Share, R2Btn, L2Btn, R1, L1)

	fmt.Printf("battery: 0x%04x\n", buf[30])

}
