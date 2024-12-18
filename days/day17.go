package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	RegisterA int
	RegisterB int
	RegisterC int
	Program   []int
	Output    string
}

func Day17() {
	computer := getDay17Data()
	runComputer(computer)
}

func runComputer(computer Computer) {
	operation := 0
	for {
		if operation > len(computer.Program)-1 {
			break
		}
		log.Println("i, instruction", operation, computer.Program[operation])
		switch computer.Program[operation] {
		case 0:
			log.Println("run reciever function 0")
			operation = computer.adv(operation)
		case 1:
			log.Println("run reciever function 1")
			operation = computer.bxl(operation)
		case 2:
			log.Println("run reciever function 2")
			operation = computer.bst(operation)
		case 3:
			log.Println("run reciever function 3")
			operation = computer.jnx(operation)
		case 4:
			log.Println("run reciever function 4")
			operation = computer.bxc(operation)
		case 5:
			log.Println("run reciever function 5")
			operation = computer.out(operation)
		case 6:
			log.Println("run reciever function 6")
			operation = computer.bdv(operation)
		case 7:
			log.Println("run reciever function 7")
			operation = computer.cdv(operation)
		}
		// log.Panicln(operation, computer)

	}
	log.Println(computer.Output)
}

func (computer Computer) getComboOperand(operation int) int {
	operand := computer.Program[operation+1]
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return computer.RegisterA
	case 5:
		return computer.RegisterB
	case 6:
		return computer.RegisterC
	default:
		log.Panicln("Invalid combo operand 7 encountered")
		return 0
	}
}

func (computer *Computer) adv(operation int) int {
	operand := computer.getComboOperand(operation)
	denominator := int(math.Pow(2, float64(operand)))
	computer.RegisterA = computer.RegisterA / int(denominator)
	return operation + 2
}

func (computer *Computer) bxl(operation int) int {
	operand := computer.Program[operation+1]
	computer.RegisterB = computer.RegisterB ^ operand
	return operation + 2
}

func (computer *Computer) bst(operation int) int {
	operand := computer.getComboOperand(operation)
	computer.RegisterB = operand % 8
	return operation + 2
}

func (computer *Computer) jnx(operation int) int {
	if computer.RegisterA == 0 {
		return operation + 2
	}
	return computer.getComboOperand(operation)
}

func (computer *Computer) bxc(operation int) int {
	computer.RegisterB = computer.RegisterB ^ computer.RegisterC
	return operation + 2
}

func (computer *Computer) out(operation int) int {
	operand := computer.getComboOperand(operation)
	if len(computer.Output) > 0 {
		computer.Output += ","
	}
	computer.Output += fmt.Sprintf("%d", operand%8)
	return operation + 2
}

func (computer *Computer) bdv(operation int) int {
	operand := computer.getComboOperand(operation)
	denominator := int(math.Pow(2, float64(operand)))
	computer.RegisterB = computer.RegisterA / int(denominator)
	return operation + 2
}

func (computer *Computer) cdv(operation int) int {
	operand := computer.getComboOperand(operation)
	denominator := int(math.Pow(2, float64(operand)))
	computer.RegisterC = computer.RegisterA / int(denominator)
	return operation + 2
}

func getDay17Data() Computer {
	file, err := os.Open("inputs/input17.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	computer := Computer{
		Program: []int{},
	}
	registerCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), " ")

		switch len(sections) {
		case 2:
			instructions := strings.Split(sections[1], ",")
			for _, instruction := range instructions {
				operation, err := strconv.Atoi(instruction)
				if err != nil {
					log.Panicln("Error parsing instruction string", err, sections)
				}
				computer.Program = append(computer.Program, operation)
			}

		case 3:
			register, err := strconv.Atoi(sections[2])
			if err != nil {
				log.Panicln("Error parsing program string", err, sections)
			}

			switch registerCount {
			case 0:
				computer.RegisterA = register
			case 1:
				computer.RegisterB = register
			case 2:
				computer.RegisterC = register
			}
			registerCount++
		}

	}

	return computer
}
