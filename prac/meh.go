package main

import "fmt"


type Avatar struct{
	username string
	age 	 int
	skills 	 []string
}

func (a *Avatar)getData(){
	fmt.Printf("username: %v\nage: %d\nskills: %v\n",a.username, a.age, a.skills)
}



func main(){
	av1 := &Avatar{username: "noornee",age: 0, skills: []string{"golang", "python"}}
	av1.getData()

}