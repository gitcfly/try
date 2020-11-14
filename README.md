# tryct
golang try_catch_finally 实现

使用教程
go get -u git push -u origin master

```
import (
	"fmt"
	"github.com/gitcfly/try"
)

func main() {

	try.Try(func() {
  
		fmt.Println("start excute some code ")
    
		panic("panic by myself")
    
		fmt.Println("end excute some code ")
    
	}).Catch(func(err interface{}) {
  
		fmt.Println(err)
    
	}).Final(func() {
  
		fmt.Println("finally do something")
    
	})
  
}

//输出

start excute some code 
panic by myself
finally do something

```