# poc-go-dynamodb

PoC Golang reference for using [AWS DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html).

## Usage

```
go run *.go
``` 

## References

* [Docs](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html)
  * [Preventing Item Overwrite](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html#Expressions.ConditionExpressions.PreventingOverwrites) (`attribute_not_exists(Id)`)
* [Examples](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html)
* [Golang SDK API](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/)
  * [PutItem](http://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.PutItem)
  * [GetItem](http://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.GetItem)
  * [UpdateItem](http://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.UpdateItem)
  * [DeleteItem](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.DeleteItem)
  * [Query](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.Query)
  * [Scan](http://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.Scan)
  * [Marshal/Unmarshal](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbattribute/)
  * [Expression Builders](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/expression/)
* https://github.com/aws/aws-sdk-go : AWS SDK for the Go programming language
* Alternative libs
  * https://github.com/guregu/dynamo : Expressive DynamoDB library for Go
  * https://github.com/underarmour/dynago : A DynamoDB client for Go
