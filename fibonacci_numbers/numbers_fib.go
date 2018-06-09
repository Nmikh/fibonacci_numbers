package main

import (
"time"
"fmt"
"encoding/json"
)

type Message struct {
	Position int
	Number   int
}

func main() {
	var number int = 0
	var previousNumber int;
	var position int = 0;
	var trueAnswers int = 0;
	var falseAnswers int = 0;
	for true {
		var answer int
		ch := make(chan bool, 1)
		timeout := make(chan bool, 1)
		defer close(ch)
		defer close(timeout)

		go func() {
			time.Sleep(10 * time.Second)
			timeout <- true
		}()

		go func() {
			fmt.Print("\nEnter Number: ")
			_, _ = fmt.Scanf("%d", &answer)
			ch <- true
		}()

		select {
		case <-ch:
			number, position, previousNumber, trueAnswers, falseAnswers =
				checkAnswer(number, position, previousNumber, answer, trueAnswers, falseAnswers)
		case <-timeout:
			number, position, previousNumber, trueAnswers, falseAnswers =
				checkAnswer(number, position, previousNumber, -1, trueAnswers, falseAnswers)
		}

		if trueAnswers >= 10 {
			fmt.Print("\nYou win")
			break
		} else if falseAnswers >= 3 {
			fmt.Print("\nYou lose")
			break
		}
	}
}

func checkAnswer(number int, position int, previousNumber int, answer int, trueAnswers int, falseAnswers int) (int, int, int, int, int) {
	if number == 0 && position == 0 {
		previousNumber = 0;
		position++;
		if answer == 0 {
			trueAnswers ++;
		} else {
			trueAnswers, falseAnswers = wrongAnswer(trueAnswers, falseAnswers, position, number)
		}
		number++;
	} else if number == 1 && position == 1 {
		position++;
		if answer == 1 {
			trueAnswers ++;
		} else {
			trueAnswers, falseAnswers = wrongAnswer(trueAnswers, falseAnswers, position, number)
		}
	} else {
		var a int = number;
		number = number + previousNumber;
		previousNumber = a;
		position ++;
		if answer == number {
			trueAnswers ++;
		} else {
			trueAnswers, falseAnswers = wrongAnswer(trueAnswers, falseAnswers, position, number)
		}
	}
	return number, position, previousNumber, trueAnswers, falseAnswers
}

func wrongAnswer(trueAnswers int, falseAnswers int, position int, number int) (int, int) {
	trueAnswers = 0
	falseAnswers++
	view := Message{
		Position: position,
		Number:   number,
	}
	b, _ := json.Marshal(view)
	fmt.Println("\nWrong\n")
	s := string(b)
	fmt.Println(s)
	return trueAnswers, falseAnswers
}