package cudnn

// #cgo LDFLAGS:-lcuda
// #cgo LDFLAGS:-lcudnn
//
// // default locs:
// #cgo LDFLAGS:-L/usr/local/cuda-11.7/lib64 -L/usr/local/cuda-11.7/lib
//
// // Include locations for cudnn.
// #cgo CFLAGS: -I/usr/local/cuda-10.2/targets/x86_64-linux/include
// #cgo CFLAGS: -I/usr/local/cuda-10.1/targets/x86_64-linux/include
// #cgo CFLAGS: -I/usr/include/x86_64-linux-gnu
// #cgo CFLAGS: -I/usr/local/cuda-11.7/include
import "C"
