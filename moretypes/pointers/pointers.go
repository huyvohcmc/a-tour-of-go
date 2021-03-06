package main

import "fmt"

func main() {
	i, j := 44, 45

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 1          // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 5    // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
