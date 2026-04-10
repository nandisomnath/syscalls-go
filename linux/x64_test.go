package linux

import (
	"fmt"
	"syscall"
	"testing"
)



func TestRead(t *testing.T)  {
	fd, _ := syscall.Open("/dev/zero", syscall.O_RDONLY, 0)
	

	buf := make([]byte, 16)
	n, err := Read(uint(fd), buf)

	fmt.Println("n:", n)
	fmt.Println("data:", string(buf))
	fmt.Println("err:", err)
}