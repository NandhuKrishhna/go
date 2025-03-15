package main

import "fmt"

func main() {
	fmt.Println("Hello,World")
}

//*No arguments: Unlike some languages (e.g., C’s int main(int argc, char *argv[])), Go’s main() takes no arguments and returns nothing. It’s simple by desig
//*If your package is main, you must have a main() function, or Go will complain when you try to build it.
