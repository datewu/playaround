package main
import "fmt"

func main() {
	defer lol(100)
	defer dota()
	defer lol(998)
}
func dota() {
	defer lol(1)
	defer lol(2)
	defer lol(3)
	defer lol(4)
	defer lol(5)
}

func lol(i int)  {
	fmt.Println(i)
}
