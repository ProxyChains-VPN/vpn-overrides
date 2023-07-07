package overrides

import "syscall"

func Read(fd int, buf []byte) (n int, err error) {
	if conns[fd] == nil {
		return syscall.Read(fd, buf)
	}

	return (*conns[fd]).Read(buf)
}
