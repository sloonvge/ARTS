package main

import (
	"testing"
	"fmt"
	"strings"
)
//https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

//
// 与
//
func TestAnd(t *testing.T) {
	var x uint8 = 0xAC
	fmt.Printf("0xAC %b\n", x)
	x = x & 0xF0
	fmt.Printf("0xF0 %b\n", 0xF0)
	fmt.Printf("     %b\n", x)
}

// 1. 判别一个数是奇/偶

func TestOddEven(t *testing.T) {
	OddEven := func(x int) {
		if x & 1 == 1 {
			fmt.Printf("Odd\n")
		} else {
			fmt.Printf("Even\n")
		}
	}

	OddEven(1)
	OddEven(-2)
}

//
// 或
//
func TestOr(t *testing.T) {
	var a uint8 = 0
	a |= 196
	fmt.Printf("%b\n", a)
}

// 或操作，可以用来配置程序

func TestConfigurations(t *testing.T) {
	const (
		UPPER = iota + 1
		LOWER
		CAP
		REV
	)

	prostr := func(str string, conf byte) string{
		rev := func(s string) string{
			runes := []rune(s)
			n := len(runes)
			for i := 0; i < n / 2; i++ {
				runes[i], runes[n - 1 - i] = runes[n - 1 - i], runes[i]
			}
			return string(runes)
		}

		if conf & UPPER != 0 {
			str = strings.ToUpper(str)
		}
		if conf & LOWER != 0 {
			str = strings.ToLower(str)
		}
		if conf & CAP != 0 {
			str = strings.Title(str)
		}
		if conf & REV != 0 {
			str = rev(str)
		}
		return str
	}
	fmt.Printf("%v\n", prostr("HELLO WORLD!", LOWER|REV|CAP))
}

//
// 异或 XOR(a, b) = 1 ; only if a != b
//				 = 0
// a ^ b ^ a = b
func TestXOR(t *testing.T) {
	var a uint16 = 0xCEFF
	a ^= 0xFF00
	a ^= 0xCEFF
	fmt.Printf("%x\n", a)
}

// 异或可以判断啊a, b符号是否相同
func TestXOR1(t *testing.T) {
	a, b := -1, 2
	fmt.Printf("%b %b %v\n", a, b, (a ^ b) > 0)
}

// 可用来求反码
func TestXOR2(t *testing.T) {
	var a uint8 = 0x03
	fmt.Printf("%08b %08b\n", a, ^a)
}

//
// 与非 &^ %^(a, b) = AND(a, NOT(b))
// AND_NOT(a, 1) = 0; AND_NOT(a, 0) = a;

//
// << 左移 >> 右移
// 乘以2^n 除以2^n

func TestLeftShift(t *testing.T) {
	var a int8 = 3
	fmt.Printf("%08b %v\n", a, a)
	fmt.Printf("%08b %v\n", a<<1, a<<1)
	fmt.Printf("%08b %v\n", a<<2, a<<2)
	fmt.Printf("%08b %v\n", a<<3, a<<3)
}

// | << 将某个位置置1
// & << 判断某个位置是否为1
// &^ << 将某个位置置0