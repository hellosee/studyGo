package main
import(
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=1;i<=10;i++{
		z -= (z*z -x) / (2*z)
	}
	return z
}

func main(){
	//用牛顿法实现平方根函数
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}