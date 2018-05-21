package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/satori/go.uuid"
)

// Note: You can also use `dynamodbav` here. See 'dynamodbattribute/encode.go'
type Foo struct {
	ID  string `json:"id"`
	Foo string `json:"foo"`
}

func main() {
	os.Setenv("AWS_PROFILE", "devalias") // One method
	region := "ap-southeast-2"

	cfg := aws.NewConfig().
		WithRegion(region)
		//WithCredentials(credentials.NewSharedCredentials("", "devalias")) // Alternative method

	sess, err := session.NewSession(cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	db := dynamodb.New(sess)

	uuid := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", uuid)

	ListTables(db)

	CreateItem(db, uuid)
	GetItem(db, uuid)
	UpdateItem(db, uuid)
	DeleteItem(db, uuid)
}

// ListTables
// Ref:
//   https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-list-tables.html
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.ListTables
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#ListTablesInput
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#ListTablesOutput
func ListTables(db *dynamodb.DynamoDB) {
	result, err := db.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalf("ListTables error: %v", err)
	}

	fmt.Println("Tables:")
	for _, n := range result.TableNames {
		fmt.Println(*n)
	}
	fmt.Println("")
}

// CreateItem
// Ref:
//   https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-create-table-item.html
//   https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-load-table-items-from-json.html
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.PutItem
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#PutItemInput
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#PutItemOutput
//   http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html
//   https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.OperatorsAndFunctions.html
func CreateItem(db *dynamodb.DynamoDB, uuid uuid.UUID) {
	item := Foo{
		ID:  uuid.String(),
		Foo: "Bar",
	}

	av, err := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String("test"),
		ConditionExpression: aws.String("attribute_not_exists(id)"), // Don't overwrite existing
	}

	result, err := db.PutItem(input)
	if err != nil {
		switch err.Error() {
		//case dynamodb.ErrCodeConditionalCheckFailedException:
		//	// TODO: if `dynamodb.ErrCodeConditionalCheckFailedException`, try again with new UUID as it probably clashed
		default:
			log.Fatalf("PutItem error: %v", err)
		}
	}

	fmt.Printf("PutItem result: %#v\n\n", result)
}

// GetItem
// Ref:
//   https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-read-table-item.html
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.GetItem
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#GetItemInput
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#GetItemOutput
//   http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html
func GetItem(db *dynamodb.DynamoDB, uuid uuid.UUID) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("test"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(uuid.String())},
		},
		//ProjectionExpression: aws.String("foo"), // Can use this to only return part of the data
	}

	result, err := db.GetItem(input)
	if err != nil {
		log.Fatalf("GetItem error: %v", err)
	}

	fmt.Printf("GetItem result: %#v\n\n", result)
}

// UpdateItem
// Ref:
//   https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-update-table-item.html
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.UpdateItem
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#UpdateItemInput
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#UpdateItemOutput
func UpdateItem(db *dynamodb.DynamoDB, uuid uuid.UUID) {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("test"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(uuid.String())},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {S: aws.String("BARBARBAR")},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set foo = :r"),
		ConditionExpression: aws.String("attribute_exists(id)"), // Don't update if doesn't exist
	}

	result, err := db.UpdateItem(input)
	if err != nil {
		log.Fatalf("UpdateItem error: %v", err)
	}

	fmt.Printf("UpdateItem result: %#v\n\n", result)
}

// DeleteItem
// Ref:
//   https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-delete-table-item.html
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.DeleteItem
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DeleteItemInput
//   https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DeleteItemOutput
func DeleteItem(db *dynamodb.DynamoDB, uuid uuid.UUID) {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("test"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(uuid.String())},
		},
	}

	result, err := db.DeleteItem(input)

	if err != nil {
		log.Fatalf("DeleteItem error: %v", err)
	}

	fmt.Printf("DeleteItem result: %#v\n\n", result)
}

// TODO: Scan example? https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-scan-table-item.html
