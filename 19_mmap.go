package main

import (
	"fmt"
	"github.com/pkg/profile"
	"os"
	"strings"
	"syscall"
)
/*
go build concat
./concat -> profile.out
go tool pprof --pdf concat <profile.out> > file.pdf
*/

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	file, _ := os.OpenFile("./test.txt", syscall.O_RDWR|syscall.O_APPEND|syscall.O_CREAT, 0600)
	ptr := (int)(file.Fd())
	stat, _ := file.Stat()
	size := int(stat.Size())
	fmt.Println(size)
	m, _ :=syscall.Mmap(ptr, 0, size, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED )

	//grow file+memorymap
	file.Write(make([]byte, size))
	syscall.Munmap(m)
	m, _ =syscall.Mmap(ptr, 0, size*2, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED )
	defer syscall.Munmap(m)

	str := fmt.Sprintf("%s", string(m))
	fmt.Println(strings.ToLower(str))
	m[size] = '#'
	fmt.Println(string(m))
}