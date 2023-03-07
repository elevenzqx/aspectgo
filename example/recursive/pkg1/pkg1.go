package pkg1

import (
	"fmt"

	"github.com/elevenzqx/aspectgo/example/recursive/pkg1/pkg2"
)

func SayHello1() {
	fmt.Println("hello from pkg1")
	pkg2.SayHello2()
}
