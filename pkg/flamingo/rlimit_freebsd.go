// +build freebsd

package flamingo

import (
	"syscall"
)

// IncreaseFileLimit tries to increase our available file limits to the maximum possible
func IncreaseFileLimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return
	}

	limits := []int{999999, 99999, 49999, 32766, 9999, 4999, 2048}

	for _, l := range limits {
		rLimit.Max = int64(l)
		rLimit.Cur = int64(l)
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err == nil {
			return
		}
	}
	return
}
