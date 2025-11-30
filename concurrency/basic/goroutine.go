package main

import (
	"fmt"
	"time"
)

func selfIntroduction(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, i+1)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go selfIntroduction("I am Kotaro Ikeda")
	go selfIntroduction("I am a backend engineer")
	go selfIntroduction("I am a software engineer")
	go selfIntroduction("I am a web developer")
	go selfIntroduction("I am a mobile developer")
	go selfIntroduction("I am a fullstack developer")
	go selfIntroduction("I am a frontend developer")
	go selfIntroduction("I am a backend developer")
	go selfIntroduction("I am a software developer")
	go selfIntroduction("I am a web developer")
	selfIntroduction("I am a mobile developer")
}
