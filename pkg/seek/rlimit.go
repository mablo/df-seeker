package seek

import "syscall"

func getRlimit(percent uint) uint64 {
	var rlimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		return 768
	}

	return uint64(float64(rlimit.Cur) * float64(percent) / 100)
}
