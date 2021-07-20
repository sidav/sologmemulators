package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type consoleWrapperStdout struct{}

func (c *consoleWrapperStdout) init() {}

func (c *consoleWrapperStdout) flush() {}

func (c *consoleWrapperStdout) closeConsole() {}

func (c *consoleWrapperStdout) clear() {
	fmt.Println("\033[2J") // linux only!
}

func (c *consoleWrapperStdout) print(s string) {
	print(s)
}

func (c *consoleWrapperStdout) println(s string) {
	println(s)
}

func (c *consoleWrapperStdout) read() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, " \n")
	return input
}
