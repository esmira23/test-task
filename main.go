package main

import (
	"fmt"
	"log"
	"strconv"
)

type Rule struct {
	divisor int
	value   string
}

type RuleSet []Rule

type FizzBuzzProblem interface {
	FizzBuzz(n int) []string
}

func (rules RuleSet) FizzBuzz(end int) []string {
	var ans []string

	for i := 0; i <= end; i++ {
		output := ""
		for _, rule := range rules {
			if i%rule.divisor == 0 {
				output += rule.value
			}
		}
		if output == "" {
			output = strconv.Itoa(i)
		}
		ans = append(ans, output)
	}

	return ans
}

func main() {

	rules := RuleSet{
		{divisor: 3, value: "Fizz"},
		{divisor: 5, value: "Buzz"},
	}

	var f float64
	fmt.Print("Enter an integer value\n>> ")

	_, err := fmt.Scanf("%f", &f)
	if err != nil {
		log.Fatalln(err)
	}

	i := int(f)
	for key, value := range rules.FizzBuzz(i) {
		fmt.Printf("%d: %s\n", key, value)
	}

	return
}
