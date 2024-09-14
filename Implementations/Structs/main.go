package main


import("fmt")

type ContactInfo struct{
	email string 
	zipCode int 
}

type Person struct{
	firstName string
	lastName string
	contact  ContactInfo
}

// you can declare it like that
// type Person struct{
// 	firstName string
// 	lastName string
// 	ContactInfo
      /*the same as => ContactInfo ContactInfo*/
// }

// func (p Person) print(){
// 	fmt.Printf("%+v",p)
// }

// func (p *Person) updateName(newFirstName string){
// 	(*p).firstName = newFirstName
// }


func updateSlice(s []string){
	s[0] = "Bye"
}

func main(){

// alex := Person{firstName: "Alex", lastName: "Anderson"}
// alex2 := Person{"Alex","Anderson"}
// var alex3 Person

// fmt.Println(alex)
// fmt.Println(alex2)
// fmt.Println(alex3)

// fmt.Printf("%+v",alex)

// jim := Person{
// 	firstName: "Jim",
// 	lastName: "Party",
// 	contact: ContactInfo{
// 		email: "jim@gmail.com",
// 		zipCode: 94000,
// 	},
// }

// fmt.Printf("%+v",jim)


// 	girgis := Person{
// 		firstName: "Girgis",
// 		lastName: "Girgis",
// 		contact: ContactInfo{
// 			email: "girgis@email.com",
// 			zipCode: 94000,
// 		},

// 	}


//      var girgisPointer *Person = &girgis
//      girgisPointer.updateName("Girgis2")
//      girgis.print()
// 	// {firstName:Girgis2 lastName:Girgis contact:{email:girgis@email.com zipCode:94000}}

//       girgis.updateName("Girgis3")
//       girgis.print()
// 	// {firstName:Girgis3 lastName:Girgis contact:{email:girgis@email.com zipCode:94000}}

mySlice := []string{"Hi","There","How","Are","You"}
updateSlice(mySlice)
fmt.Println(mySlice)



}