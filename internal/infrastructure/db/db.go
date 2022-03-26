package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type db struct {
	conn *dynamo.DB
}

type DB interface {
	Conn() *dynamo.DB
}

func New() (*db, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("ap-northeast-1"),
			Endpoint:    aws.String("dynamodb:8000"),
			Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
			DisableSSL:  aws.Bool(true),
		},
	)
	if err != nil {
		return nil, err
	}

	return &db{conn: dynamo.New(sess)}, nil
}

func (d *db) Conn() *dynamo.DB {
	return d.conn
}
