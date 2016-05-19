// +build freebsd openbsd netbsd

package termios

import (
	"fmt"
	"syscall"
	"unsafe"
)

func Ptsname(fd uintptr) (string, error) {
	var n uintptr
	err := ioctl(fd, syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/dev/pts/%d", n), nil
}

func grantpt(fd uintptr) error {
	var n uintptr
	return ioctl(fd, syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n)))
}

func unlockpt(fd uintptr) error {
	return nil
}
