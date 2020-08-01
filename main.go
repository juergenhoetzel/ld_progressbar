// #cgo CFLAGS: xx-ldl
package main

import "github.com/schollz/progressbar/v3"
import "syscall"
import "log"
import "fmt"
import "os"

/*
#cgo LDFLAGS: -ldl
#include <unistd.h>
void init_write();
*/
import "C"

var pb *progressbar.ProgressBar

//export ProgressWrite
func ProgressWrite(n C.size_t) {
	if n > 0 {
		pb.Add(int(n))
	} else {
		pb.Finish() // FIXME should send newline?
		fmt.Fprint(os.Stderr, "\n")
	}
}

func init() {
	pb = progressbar.DefaultBytes(-1, "Written")
	C.init_write()
}

func main() {

	for {
		n, err := syscall.Sendfile(syscall.Stdout, syscall.Stdin, nil, 1024)
		if err != nil {
			pb.Add(1)
			log.Fatal(err)
		}
		print(n)
	}

}
