provider "aws" {
  region = "us-west-1"
}

resource "aws_s3_bucket" "quickread_pdf_upload" {
  bucket = "quickread-pdf-upload"

  tags = {
    Name = "quickread-pdf-upload"
  }

}

resource "aws_s3_access_point" "quickread_pdf_upload" {
  name       = "quickread-pdf-upload"
  bucket     = aws_s3_bucket.quickread_pdf_upload.bucket
  depends_on = [aws_s3_bucket.quickread_pdf_upload]
}

output "bucket_name" {
  value = aws_s3_bucket.quickread_pdf_upload.id
}

output "access_point_alias" {
  value = aws_s3_access_point.quickread_pdf_upload.alias
}

