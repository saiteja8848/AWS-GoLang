package category

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidCategoryData     = "invalid Category data"
	ErrorInvalidCategoryId       = "invalid categoryId"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
	ErrorCategoryAlreadyExists   = "Category already exists"
	ErrorCategoryDoesNotExists   = "Category does not exist"
)

type Category struct {
	CategoryId   string `json:"categoryId"`
	Categoryname string `json:"categoryName"`
	Date         string `json:"date"`
	Score        int    `json:"score"`
}

func FetchCategory(categoryId, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*Category, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"categoryId": {
				S: aws.String(categoryId),
			},
		},
		TableName: aws.String(tableName),
	}
	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new(Category)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchCategorys(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]Category, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]Category)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	return item, nil
}

func CreateCategory(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Category,
	error,
) {
	var category Category
	if err := json.Unmarshal([]byte(req.Body), &category); err != nil {
		return nil, errors.New(ErrorInvalidCategoryData)
	}
	// Check if category exists
	currentCategory, _ := FetchCategory(category.CategoryId, tableName, dynaClient)
	if currentCategory != nil && len(currentCategory.CategoryId) != 0 {
		return nil, errors.New(ErrorCategoryAlreadyExists)
	}
	// Save Category
	av, err := dynamodbattribute.MarshalMap(category)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &category, nil
}

func UpdateCategory(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Category,
	error,
) {
	var category Category
	if err := json.Unmarshal([]byte(req.Body), &category); err != nil {
		return nil, errors.New(ErrorInvalidCategoryId)
	}

	currentUser, _ := FetchCategory(category.CategoryId, tableName, dynaClient)
	if currentUser != nil && len(currentUser.CategoryId) == 0 {
		return nil, errors.New(ErrorCategoryDoesNotExists)
	}

	av, err := dynamodbattribute.MarshalMap(category)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &category, nil
}

func DeleteCategory(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {
	categoryId := req.QueryStringParameters["categoryId"]
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"categoryId": {
				S: aws.String(categoryId),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}
	return nil
}
