package try

import (
	"fmt"
)

func main() {
	Try(func() {
		fmt.Println("start excute some code ")
		panic("panic by myself")
		fmt.Println("end excute some code ")
	}).Catch(func(err interface{}) {
		fmt.Println(err)
	}).Final(func() {
		fmt.Println("finally do something")
	})
}
