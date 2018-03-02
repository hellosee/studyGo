package main
import(
	"fmt"
)
func swap(x,y string)(string,string){
	return y,x
}
func main(){
	//a,b := swap("aaa","bbb")
	//fmt.Println(a,b)
	var a = 100
	var b = 200
	//swapInt(a,b)
	//swapIntX(&a,&b)
	swap := func(x *int,y *int){
		var temp int
		temp = *x
		*x = *y
		*y = temp
	}
	swap(&a,&b)
	fmt.Println(a,b)
}

func swapInt(x,y int) {
	var temp int 
	temp = x
	x = y
	y=temp
}

func swapIntX(x *int,y *int){
	var temp int 
	temp = *x
	*x = *y
	*y = temp
}

