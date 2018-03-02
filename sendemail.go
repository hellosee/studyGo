package main
import (
	"fmt"
	"net/smtp"
	"strings"
)
func main(){
	sendList := make(map[string]string)
	sendList["1335244575@qq.com"] = "这封邮件是用Go脚本发出来的!2018年狗年，加油！！";
	sendMail(sendList)
	fmt.Print("按下回车结束")
}

func sendToMail(user,password,host,to,subject,body,mailtype string) error{
	auth := smtp.PlainAuth("",user,password,strings.Split(host,":")[0])
	msg  := []byte("To:" + to + "\r\nFrom" + user + "\r\nSubject:" + subject + "\r\n" + "Content-Type:text/" + mailtype + ";charset=UTF-8" + "\r\n\r\n" + body)
	sendto := strings.Split(to,";")
	err := smtp.SendMail(host,auth,user,sendto,msg)
	return err
}

func sendMail(sendList map[string] string){
	fmt.Printf("共需要发送%d封邮件\n",len(sendList))
	index := 1;
	for mail,content := range sendList{
		fmt.Printf("发送第%d封",index);
		if err := sendToMail("lishoujie08@126.com",
		"password",
		"smtp.126.com:25",
		mail,
		"这封邮件是用Go脚本发出来的",
		content,
		"html");
		err != nil {
			fmt.Printf("... 发送错误(X) %s %s \n",mail,err.Error())
		} else {
			fmt.Printf("... 发送成功(V) %s \n",mail);
		}
		index++
	}
}












