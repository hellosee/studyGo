package main
import("fmt")
func main(){
	//寻找100以内所有的素数
	var count int
	var flag int
	count = 1
	for count < 100 {
		count++
		flag = prime(count)
		if flag == 0{
			fmt.Println(count,"素数")
		}
	}
	
}

func prime(n int) int{
	for tmp:=2;tmp<n;tmp++{
		if n % tmp == 0 {
			return n
		} 
	}
	return 0
}