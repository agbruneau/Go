package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "Main")
	f := fibonacci()
	for i := 0; i < 1000; i++ {
		fmt.Println(f())
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/*
Solution 1: */
func fibonacci() func() uint {
	var v1, v2 uint = 0, 1
	return func() uint {
		v2 = v2 + v1
		v1 = v2 - v1
		return v2 - v1
	}
}

/*
Solution 2:
func fibonacci() func() int {
	v1, v2 := 0, 1

	return func() int {
		defer func() {
			v1, v2 = v2, v1+v2
		}()

		return v1
	}
}
*/

/*
Solution 3:
func fibonacci() func() int {
	v1, v2 := 0, 1
	return func() int {
		f := v1
		v1, v2 = v2, v2+f
		return f
	}
}
*/

/*
Solution 4:
func fibonacci() func() int {
    a, b := 1, 0
    return func() int {
        a, b = b, a+b
        return a
    }
}
*/

/*
Solution 5:
func fibonacci() func(int) int {
	first := 0
	second := 1
	return func(i int) int {
		if i == 0 || i == 1 {
			return i
		}
		sum := first + second
		first = second
		second = sum
		return sum
	}
}
*/

/*
Solution 6:
func fibonacci() func() int {
	previous := 0
	last := 0
	index := 0

	return func() int {
		var result int
		if index == 1 {
			result = 1
		} else {
			result = last + previous
		}

		previous = last;
		last = result;
		index++

		return int(result)
	}
}
*/

/*
Solution 7:
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a // Store the result value before translations
		c := a + b // Fibonacci
		a = b // Translation
		b = c // Translation
		return result
	}
}
*/

/*
Solution 8:
func fibonacci() func() int {
    a: = 0
    b: = 1
    sum: = a + b
    return func() int {
        if a == 0 {
            a = b
            return 0
        }
        a = b
        b = sum
        sum = a + b
        return a
    }
}
*/

/*
Solution 9:
func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b // two assignment happens parallely, but if you enforce it sequentially, it won't work
		return a
	}
}
*/

/*
Solution 10, with Slice:
func fibonacci() func() int {
	var i []int
	return func() int {
		if len(i) == 0 {
			i = append(i, 0)
		} else if len(i) == 1 {
			i = append(i, 1)
		} else {
			i = append(i, i[len(i)-2]+i[len(i)-1])
		}
		return i[len(i)-1]
	}
}
*/
