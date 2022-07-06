package UI

import (
	"fmt"
	"student/app/components"
	"student/app/database"
)

func DisplayHome() { 
	fmt.Println("Hello welcome to KEC COLLEGE")
	viewer:=components.Choice("You Are \n1.Staff\n2.Student")
	switch viewer {
	case 1,2:
		snap:=components.Choice("1.LogIn\n2.SignUp")
		switch snap {
		case 1:
			if(viewer==1){
				login("staff")
			}else{
				login("student")

			}
		case 2:
			if(viewer==1){
				signup("staff")
			}else{
				signup("student")

			}
		}
		
	default:
		fmt.Println("invlid")

	}
}

func DisplayStudent(rollnumber string) { 
	fmt.Println("Hello,",rollnumber)
	choice:=components.Choice("You Are \n1.info\n2.change password")
	switch choice {
	case 1:
		fmt.Println("See Info")
		
		database.InfoDatabase(rollnumber)
		
	case 2:
		fmt.Println("Change Password")
	default:
		fmt.Println("invlid")

	}
}
func DisplayStaff(email string) { 
	fmt.Println("Hello,",email)
	choice:=components.Choice("You Are \n1.info\n2.change password")
	switch choice {
	case 1:
		fmt.Println("See Info")
		database.InfoDatabase(email)
	case 2:
		fmt.Println("Change Password")
	default:
		fmt.Println("invalid")

	}
}