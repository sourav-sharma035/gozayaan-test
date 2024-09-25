package main

import (
    "encoding/json"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
)

type Response struct {
    Message   string `json:"message"`
    ImageURL  string `json:"image_url"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    bucket := "airport-images-bucket"
    airportId := request.QueryStringParameters["airportId"]
    image := request.Body

    // Create S3 session
    sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
    svc := s3.New(sess)

    key := fmt.Sprintf("images/airport-%s.jpg", airportId)

    _, err := svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   aws.ReadSeekCloser(strings.NewReader(image)),
        ContentType: aws.String("image/jpeg"),
    })

    if err != nil {
        log.Printf("Failed to upload image: %v", err)
        return events.APIGatewayProxyResponse{
            StatusCode: 500,
            Body:       "Failed to upload image",
        }, nil
    }

    imageUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
    response := Response{
        Message:  "Image uploaded successfully",
        ImageURL: imageUrl,
    }
    responseBody, _ := json.Marshal(response)

    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       string(responseBody),
    }, nil
}

func main() {
    lambda.Start(handler)
}
