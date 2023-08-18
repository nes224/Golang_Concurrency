package main 


import "fmt"

func zeroValue(ivalue int){
	ivalue = 0
}

func zeropointer(ipointer *int){
	*ipointer = 0
} 

func main(){
	i := 1 
	fmt.Println("i =",i)

	zeroValue(i)
	fmt.Println("i from function zerovalue =",i)

	zeropointer(&i) // การเข้าถึง Address เราจะต้องใช้เครื่องหมาย &
	fmt.Println("i from function zerovalue =",i)

}