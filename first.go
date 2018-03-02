package main
import "fmt"
func main(){
	var a int = 100;
	var b int = 200;
	var ret int ;
	ret = max(a,b);
	fmt.Printf("最大值:%d\n",ret);
}
func max(num1 int, num2 int) int {
	var ret int;
	if(num1 > num2){
		ret = num1;
	} else {
		ret = num2;
	}
	return ret;
}