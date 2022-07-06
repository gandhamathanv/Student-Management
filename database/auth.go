package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginDatabase(user string ,id string ,password string)(bool) {
	fmt.Println(user)
	client := Connect()
	coll := client.Database("studentusers").Collection("user")

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"id", id},{"user",user}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No user was found with the %s\n", id)
		return false
	}
	if err != nil {
		return false
	}
	println(result)
	return true
}
func SignupDatabase(user string ,name string ,dept string,class string,id string,password string)(bool) {
	fmt.Println(user)
	client := Connect()
	coll := client.Database("studentusers").Collection("user")
	doc := bson.D{{"user",user},{"name",name}, {"dept", dept},{"class",class},{"id",id},{"password",password}}

	result, err := coll.InsertOne(context.TODO(), doc)
if err != nil {
    return false
}
	println(result)
	return true
}
func InfoDatabase(id string)([]bson.M) {

		
	client := Connect()
	coll := client.Database("studentusers").Collection("user")
	cursor,err := coll.Find(context.TODO(), bson.M{"id": id})
	if err != nil {
		fmt.Println("err",err)
	}
	var result []bson.M
	if err=cursor.All(context.TODO(),&result); err!=nil{
		fmt.Println("err",err)
	}
	fmt.Println("NAME:",result["name"])
	fmt.Println("id:",result["rollnumber"])
}