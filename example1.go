package main

func doCalc1() {
	// 2 seconds spent here
}

func doCalc2() {
	// 2 seconds spent here
}

func doCalc3() {
	// 2 seconds spent here
}

func main() {
	// ...
	// insert long program here...
	// ...

	// We are executing three functions serially, like usual.
	// Time = 6 seconds.
	doCalc1()
	doCalc2()
	doCalc3()

	// ...
	// insert rest of the program here...
	// ...
}
