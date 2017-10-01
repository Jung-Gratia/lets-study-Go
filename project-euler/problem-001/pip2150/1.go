package main
import "fmt"

func main(){ 
	n:=100
	sum:=0

	for i:=0;i<n;i++ {
		if (i % 5) == 0 {
			sum += i
		} else if (i % 3) == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}