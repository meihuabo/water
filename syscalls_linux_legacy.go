// +build linux,!go1.11

package water

import (
	"os"
	"runtime"
)

func openDev(config Config) (ifce *Interface, err error) {
	var file *os.File
	fdLoc := "/dev/net/tun"
	if runtime.GOARCH == "arm" {
		fdLoc = "/dev/tun"
	}

	if file, err = os.OpenFile(
		fdLoc, os.O_RDWR, 0); err != nil {
		return nil, err
	}

	name, err := setupFd(config, file.Fd())
	if err != nil {
		return nil, err
	}

	return &Interface{
		isTAP:           config.DeviceType == TAP,
		ReadWriteCloser: file,
		name:            name,
	}, nil
}
