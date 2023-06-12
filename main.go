package main

import (
	"AccManager/src"
)

func main() {
	src.InitDb()
	src.InitBot().Idle()
}
