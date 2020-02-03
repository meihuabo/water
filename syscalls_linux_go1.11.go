// +build linux,go1.11

package water

import (
	"os"
	"runtime"
	"syscall"
)

func openDev(config Config) (ifce *Interface, err error) {
	var fdInt int
	fdLoc := "/dev/net/tun"
	if runtime.GOARCH == "arm" {
		fdLoc := "/dev/tun"
	}

	if fdInt, err = syscall.Open(
		fdLoc, os.O_RDWR|syscall.O_NONBLOCK, 0); err != nil {
		return nil, err
	}

	name, err := setupFd(config, uintptr(fdInt))
	if err != nil {
		return nil, err
	}

	return &Interface{
		isTAP:           config.DeviceType == TAP,
		ReadWriteCloser: os.NewFile(uintptr(fdInt), "tun"),
		name:            name,
	}, nil
}
