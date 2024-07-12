package core

import (
	"github.com/sstallion/go-hid"
)

func GetDevices() (devices []hid.DeviceInfo, err error) {
	err = hid.Enumerate(
		hid.VendorIDAny,
		hid.ProductIDAny,
		func(info *hid.DeviceInfo) error {
			devices = append(devices, *info)
			return nil
		})
	return
}

func GetDS4Devices() (devices []hid.DeviceInfo, err error) {
	allDevices, err := GetDevices()
	if err != nil {
		return
	}

	// filter out DS4 devices+
	for _, device := range allDevices {
		if IsDS4Device(device) {
			devices = append(devices, device)
		}
	}

	return
}
