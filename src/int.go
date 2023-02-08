package main

import "fmt"

type FullName struct {
	firstName string
	lastName  string
}

//func (fullName *FullName) FirstName() string {
//	return fullName.firstName
//}
//

func (fullName *FullName) LastName() string {
	fullName.lastName = "aaa"
	return fullName.lastName
}

func NewFullName(firstName string, lastName string) *FullName {
	return &FullName{firstName: firstName, lastName: lastName}
}

func main() {
	userName := NewFullName("Jack", "Cooper")
	fmt.Println(userName.LastName())
}
