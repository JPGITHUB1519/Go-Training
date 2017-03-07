package main

func Fibonnacci(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return Fibonnacci(x-1) + Fibonnacci(x-2)
}
