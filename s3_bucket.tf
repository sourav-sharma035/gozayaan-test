provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "gozayaan_airport_images" {
  bucket = "airport-images-bucket"
  acl    = "private"
  
  versioning {
    enabled = true
  }

  tags = {
    Name        = "AirportImagesBucket"
    Environment = "dev"
  }
}

output "bucket_name" {
  value = aws_s3_bucket.gozayaan_airport_images.bucket
}
