package main

import "phase1/concurrency"

func main() {
	// fmt.Println("Hello, World!")

	// // Factorial + Fibonacchi
	// num := 7
	// fmt.Printf("Factorial of %d = %d", num, factFibo.CalFactorials(num))
	// fmt.Printf("\nFibonacci %d = %d\n", num, factFibo.CalFibonacci(num))

	// // Calculator
	// num1 := 3.0
	// num2 := 2.0
	// fmt.Print("num1 = ")
	// fmt.Scanf("%f ", &num1)
	// mode := "plus"
	// fmt.Print("mode = ")
	// fmt.Scanf("%s ", &mode)
	// fmt.Print("num2 = ")
	// fmt.Scanf("%f ", &num2)
	// num3 := calculator.Calculator(mode, num1, num2)
	// fmt.Printf("Calculator : %s : %.2f,%.2f = %.2f", mode, num1, num2, num3)

	// // Concurrency
	// concurrency.NormalPrint()
	// concurrency.GoRoutinesPrint()
	// concurrency.ChannelPrint()
	// concurrency.ChannelWaitGroupPrint()
	// concurrency.NormalLoop()
	// concurrency.GoRoutinesLoop()
	// concurrency.ChannelLoop()
	concurrency.Worker()

	// done := make(chan bool)
	// go worker(done)
	// <-done
	// fmt.Println("Finished!")
}

// func worker(done chan bool) {
// 	fmt.Println("Working...")
// 	done <- true
// }
