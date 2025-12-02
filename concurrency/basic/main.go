package main

import "concurrency/basic/examples"

func main() {
	examples.Goroutine()
	examples.ChannelBasic()
	examples.ChannelBuffered()
	examples.ChannelRangeClose()
	examples.ChannelSelect()
	examples.Mutex()
}
