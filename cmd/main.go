package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hizzuu-dev/plate-api/internal/infrastructure/db"
)

type User struct {
	UserID string `dynamo:"UserID,hash"`
	Name   string `dynamo:"Name,range"`
	Age    int    `dynamo:"Age"`
	Text   string `dynamo:"Text"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	fmt.Println("start main.go")
	db, err := db.New()
	if err != nil {
		panic(err)
	}

	if err := db.Conn().Table("UserTable").DeleteTable().Run(); err != nil {
		panic(err)
	}

	if err := db.Conn().CreateTable("UserTable", User{}).Run(); err != nil {
		panic(err)
	}

	tb2 := db.Conn().Table("UserTable")

	fmt.Println(tb2.Name(), "check dynamoDB2")
	lambda.Start(handler)
}
