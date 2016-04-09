package main

import "fmt"

func main() {
	fmt.Print("Ingrese una palabra: ")
  	var input string
  	fmt.Scanf("%s", &input)
  	escribir(input)
  	escribirAlReves(input)
}

func escribir(s string) {
	for i := 1; i <= len(s); i++ {
		fmt.Println(s[0:i])
	}
}

func escribirAlReves(s string) {
	for i := len(s); i >= 0; i-- {
		fmt.Println(s[0:i])
	}
}