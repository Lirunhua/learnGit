package main
import(
	"fmt"
	"math"
)

type I interface{
	M()
}

type T struct {
	S string
}

type F float64

//  *T implement interface I
// but T didn't implement interface I
func (t *T) M(){
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	// if : i = T{"hello"}, the wrong msg will be : type T didn't implement interface I
	i = &T{"hello"} 
}
