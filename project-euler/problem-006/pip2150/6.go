package main
import "fmt"

func main(){ 
	
	n := 100

	sum:=0
	sum_2:=0
	for i:=1; i<=n; i++{
		sum+=i*i
		sum_2 += i
	}
	sum_2 *= sum_2

	fmt.Println(sum_2-sum)
}