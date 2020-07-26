package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func count_sheep(num string, iteration int) {
	var num_int, cur_num_int int
	num_list := make(map[rune]rune)
	i := 1

	num_int, err := strconv.Atoi(num)
	if err != nil {
		panic("Error converting num to int")
	}

	if num_int == 0 {
		fmt.Printf("case #%v: INSOMNIA\n", iteration)
		return
	}

	for {
		for _, char := range num {
			num_list[char] = char
		}

		if len(num_list) == 10 {
			fmt.Printf("case #%v: %s\n", iteration, num)
			return
		}

		cur_num_int = num_int * i
		i++

		num = strconv.Itoa(cur_num_int)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	i := 1

	// Discard first line
	sc.Scan()

	for sc.Scan() {
		count_sheep(sc.Text(), i)
		i++
	}
}
