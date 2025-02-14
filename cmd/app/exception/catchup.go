package exception

import "fmt"

func CatchUp() {
	if err := recover(); err != nil {
		fmt.Println(err.(error))
	}
}
