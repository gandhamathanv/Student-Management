package UI

import (
	"fmt"
	"student/app/components"
	"student/app/database"
)
func signup(choice string){
	switch choice {
	case "staff":
		name:=components.Info("enter your name")
		dept:=components.Info("enter your dept")
		class:=components.Info("enter your class")

		email:=components.Info("enter your email")
		password:=components.Info("enter your password")
		confirmPassword:=components.Info("enter your confirmPassword")
		if(confirmPassword!=password){
			fmt.Println("Password Mismatch")
			signup("staff")
		}else{
			database.SignupDatabase(
				"staff",
				name,dept,class,email,password)
		}

	case "student":
		name:=components.Info("enter your name")
		dept:=components.Info("enter your dept")
		class:=components.Info("enter your class")

		rollnumber:=components.Info("enter your rollnumber")
		password:=components.Info("enter your password")
		confirmPassword:=components.Info("enter your confirmPassword")
		if(confirmPassword!=password){
			fmt.Println("Password Mismatch")
			signup("student")
		}else{
			database.SignupDatabase(
				"student",
				name,dept,class,rollnumber,password)
		}
	}
}
func login(choice string){
	switch choice {
	case "staff":
		email:=components.Info("enter your email")
		password:=components.Info("enter your password")
		status:=database.LoginDatabase("staff",email,password)
		if(status){
			DisplayStaff(email)
		}


	case "student":
		rollnumber:=components.Info("enter your rollnumber")
		password:=components.Info("enter your password")
		status:=database.LoginDatabase("student",rollnumber,password)
		if(status){
			DisplayStudent(rollnumber)
		}
	}	
}