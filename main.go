package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	Red     = "\033[1;31m"
	Green   = "\033[1;32m"
	Yellow  = "\033[1;33m"
	Purple  = "\033[1;95m"
	Magenta = "\033[1;35m"
	Cyan    = "\033[1;36m"
	White   = "\033[1;37m"
	Blue    = "\033[34m"
	Orange  = "\033[38;5;208m"
	reset   = "\033[0m"
)

func main() {
	Args1 := os.Args[1]
	content := strings.Split(Args1, "\\n")
	var s []string
	if len(os.Args) == 3 {
		if strings.Contains(os.Args[2], "--color=") {
			rep := strings.ReplaceAll(os.Args[2], "--color=", "")

			file, err := os.Open("ascii-art.txt")
			if err != nil {
				fmt.Println("Le fichier n'existe pas.")
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				s = append(s, scanner.Text())
			}

			colored, allColored := true, true
			if rep == "cyan" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Cyan)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Cyan)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "magenta" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Magenta)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Magenta)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "red" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Red)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Red)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "green" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Green)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Green)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "white" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(White)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(White)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "yellow" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Yellow)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Yellow)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "blue" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Blue)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Blue)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "orange" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Orange)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Orange)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else if rep == "purple" {
				if !strings.Contains(os.Args[1], "¤") {
					allColored = false
					fmt.Print(Purple)
				}
				for _, element := range content {
					if len(element) > 0 {

						line := []rune(element)
						for a := 0; a < 8; a++ {
							for i := 0; i < len(line); i++ {

								if !colored {
									fmt.Print(Purple)
								} else if colored && allColored {
									fmt.Print(reset)
								}

								if line[i] == '¤' {
									colored = !colored
									continue
								}

								group := (int(line[i]) - 32) * 9
								adress := group + a + 1
								fmt.Print(s[adress])
							}

							fmt.Print(string(rune('\n')))
						}
					} else {
						fmt.Println("Your os.Args[1] is empty")
					}
				}
			} else {
				fmt.Println("Please enter one of there colors:")
				fmt.Println("blue, magenta, orange, green, cyan, green, white, purple, red")
			}
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Println("EX: go run . something standard")
	}
}
