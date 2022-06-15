package main

import "fmt"

func main() {
	myPc := otherPc{ram:16, brand:"msi", disk: 1024}
	fmt.Println(myPc)
}

type otherPc struct{
	ram int
	brand string
	disk int
}

func (myPc otherPc) String() string {
	return fmt.Sprintf("Tengo %d GB de ram, %d GB de disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}
