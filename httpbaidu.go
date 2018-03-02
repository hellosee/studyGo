package main
import(
	"fmt"
	"os"
	"log"
	"net/http"
)
func main(){
	resp,err := http.Get("https://www.baidu.com");
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK{
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
	
	buf := make([]byte,1024)
	//create file
	f,errl := os.OpenFile("baidu.html",os.O_RDWR|os.O_CREATE|os.O_APPEND,os.ModePerm)
	if errl != nil{
		panic(errl)
		return
	}
	defer f.Close()
	for{
		n,_ := resp.Body.Read(buf)
		if 0 == n{break}
		f.WriteString(string(buf[:n]))
	}
}