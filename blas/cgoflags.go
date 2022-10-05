package cublas

// #cgo CFLAGS: -I/usr/local/cuda-11.7/targets/x86_64-linux/include -I/usr/local/cuda-11.7/include
// #cgo LDFLAGS: -lcublas
// #cgo LDFLAGS: -L/usr/local/cuda-11.7/targets/x86_64-linux/lib -L/usr/local/cuda-11.7/lib64 -L/usr/lib/x86_64-linux-gnu
import "C"
