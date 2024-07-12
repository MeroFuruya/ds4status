package core

import (
	"fmt"

	"github.com/MeroFuruya/ds4status/util"
	"github.com/sstallion/go-hid"
)

func Test() {
	devices, err := GetDS4Devices()
	if err != nil {
		fmt.Println(err)
		return
	}

	var deviceHandle *hid.Device
	var buf []byte
	var bytesWritten int

	for _, device := range devices {
		fmt.Println(device)

		deviceHandle, err = OpenDevice(device)
		if err != nil {
			fmt.Println(err)
			continue
		}

		buf, bytesWritten, err = ReadDeviceState(deviceHandle)
		if err != nil {
			fmt.Println(err)
			continue
		}

		util.PrintBytes(buf[:bytesWritten])

		PrintDeviceState(deviceHandle)

		CloseDevice(deviceHandle)
	}
}
