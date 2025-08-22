package main

import "github.com/alfianyulianto/go-room-managament/halpers"

func main() {
	server := NewInitializedServer()
	err := server.Listen(":3000")
	halpers.IfPanicError(err)
}
