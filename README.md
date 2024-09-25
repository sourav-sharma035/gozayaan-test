# gozayaan-test
For completing my task I will follow step by step. In the below I will define my task doc

Task-1:Provision a Cloud Storage Bucket using Infrastructure as Code (IaC)

Step-1:

I will use terraform for provisioning buckets. Please check the file s3_bucket.tf 

command:

$terraform init

$terraform apply

Task-2: Make an endpoint /update_airport_image to update an airport’s image.

Now I will create an Endpoint /update_airport_image to Update an Airport’s Image. For this I will create API gateway for  /update_airport_image with POST method 
also Integrate it with a Lambda function by GO lang that handles the image upload to the previously created S3 bucket.

Step-1:

Build the s3_lambda.go file

$go build s3_lambda.

$zip s3_lambda.zip s3_lambda

Step-2:

Than upload the s3_lambda.zip file in my already created lambda function and deploy it 

Task-3: Containerize the Go application

Now I will  create a Dockerfile to build an image from my existing project directory /sre_airports_api

Step-1: Go to directory /sre_airports_api and create Dockerfile with requirements

$cd /sre_airports_api 

$vim Dockerfile

Step-2: Build the image

$docker build -t airport-api .

Step-3: Run the container from build image 

$docker run -p 8080:8080 airport-api