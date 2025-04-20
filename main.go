package main

// func add(a int, b int) int {
// 	return a + b
// }

// func randomValues() (int, int) {
// 	return 133, 122
// }

// func multiAdd(nums ...int) int {
// 	sum := 0
// 	for _, n := range nums {
// 		sum = sum + n
// 	}

// 	return sum
// }

// func intSeq() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// func printRange(id string) {
// 	for i:= range 4 {
// 		fmt.Println(id, i)
// 	}
// }

// func someTask(done chan bool) {
// 	fmt.Println("Task complete")
// 	done <- true
// }

// func worker (id int) {
// 	fmt.Printf("Worker %d Starting\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d Executed\n", id)
// }

func main() {
	/* Print */
	// fmt.Print("Hello World!");

	/* Values */
	// var a int = 4;
	// var b int = 5;

	// fmt.Print("Sum is ", a  + b)

	// a := 4
	// fmt.Print("Value is ", a)

	/* For loop */
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for j := 0; j < 5; j++ {
	// 	fmt.Println(j)
	// }

	// for k := range 7 {
	// 	fmt.Print(k)
	// }

	// if 7 % 2 == 0 || 1.0 / 8.0 == 0 {
	// 	fmt.Print("Success")
	// } else {
	// 	fmt.Print("Failure")
	// }

	/* Arrays */
	// var arr [5]int = [5]int{1, 2, 3, 4, 5};
	// fmt.Println(arr)

	// arr[3] = 256;
	// fmt.Println(arr)
	// fmt.Println(len(arr))

	// twoD := [2][3]string{
	// 	{"a1", "a2", "a3"},
	// 	{"b1", "b2", "b3"},
	// }
	// fmt.Println(twoD)


	/* Slices */
	// var str []string
	// fmt.Println(str, str == nil, len(str) == 0)

	// var str2 = make([]string, 3)
	// fmt.Println(str2, len(str2), cap(str2))

	// str2[0] = "yash"
	// fmt.Println(str2)

	// newStr := append(str2, "solanki")
	// fmt.Println(newStr[0])
	// fmt.Println(newStr[1])
	// fmt.Println(newStr[2])
	// fmt.Println(newStr[3])

	// str3 := make([]string, 1)
	// copy(str3, str2)
	// fmt.Println(str3)

	// var str4 []string
	// str4 = append(str4, "a")
	// str4 = append(str4, "b")
	// str4 = append(str4, "c")
	// str4 = append(str4, "d")
	// fmt.Println(str4)

	// str5, str6, str7 := str4[0:3], str4[2:], str4[:3]
	// fmt.Println(str5, str6, str7)

	// if slices.Equal(str5, str7) {
	// 	fmt.Println("Equal slices")
	// }

	/* Maps */
	// var map1 = make(map[string]int)
	// map1["a"] = 3
	// map1["b"] = 35
	// fmt.Println(map1)

	// a := map1["a"]
	// fmt.Println(a)

	// l := len(map1)
	// fmt.Println(l)

	// delete(map1, "a")
	// fmt.Println(map1)

	// clear(map1)
	// fmt.Println(map1)

	// newMap := map[string]int{"a": 122, "b": 876}
	// fmt.Println(newMap)

	/* Functions */
	// fmt.Println("Sum is", add(1, 8))

	// a, b := randomValues()
	// fmt.Println("Random values", a, b)

	// fmt.Println(multiAdd(1, 2, 3))

	// nums := []int{1, 2, 4, 5, 6}
	// fmt.Println(multiAdd(nums...))

	/* Closures */
	// nextInt := intSeq()
	// fmt.Println(nextInt(), nextInt(), nextInt())

	// newInt := intSeq()
	// fmt.Println(newInt())

	/* Goroutine */

	// printRange("direct")

	// go printRange("inside goroutine")

	// fmt.Println("After goroutine")

	// time.Sleep(time.Second)

	// fmt.Println("End")

	/* WaitGroup */

	// var wg sync.WaitGroup

	// for i := range 5 {
	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()
	// 		worker(i)
	// 	}()
	// }

	// wg.Wait()


	/* Channels */
	// messages := make(chan string)

	// go func() { messages <- "ping"}()
	// msg := <-messages
	// fmt.Println(msg)

	// go func() { messages <- "pong"}()
	// msg = <-messages
	// fmt.Println(msg)

	/* Buffered Channels */
	// bufferedMessages := make(chan string, 2)
	// bufferedMessages <- "msg 1"
	// bufferedMessages <- "msg 2"

	// fmt.Println(<- bufferedMessages)
	// fmt.Println(<- bufferedMessages)

	/* Channel sync */

	// fmt.Println("Starting")
	// done := make(chan bool)

	// go someTask(done)

	// <- done
}