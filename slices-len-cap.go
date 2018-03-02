package main
import "fmt"
func main(){
	s := []int{2,3,5,7,11,13}
	printSlice(s)//6,6
	s = s[:0]
	printSlice(s)//0,6
	s = s[:4]
	printSlice(s)//4,6
	s = s[3:]
	printSlice(s)//2,4
	
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}