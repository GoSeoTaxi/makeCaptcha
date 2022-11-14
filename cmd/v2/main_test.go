package main

import (
	"log"
	"testing"
)

//func BenchmarkFib1(b *testing.B) { benchmarkFib(b) }

func BenchmarkCrC2(b *testing.B) {
	for n := 0; n < b.N; n++ {

		_, _, err := CreateCaptcha2()
		if err != nil {
			log.Fatal(err)
		}

		//	fmt.Println(n)
		//time.Sleep(1 * time.Second)
	}
}
