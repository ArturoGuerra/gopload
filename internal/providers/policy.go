package providers
import (
    "fmt"
    "github.com/arturoguerra/goimgupload/internal/config"
)

var cfg = config.Load()

const BucketExistsError = "Your previous request to create the named bucket succeeded and you already own it."
var policy = fmt.Sprintf(`
{
 "Statement": [
  {
   "Action": [
    "s3:GetBucketLocation",
    "s3:ListBucket"
   ],
   "Effect": "Allow",
   "Principal": {
    "AWS": [
     "*"
    ]
   },
   "Resource": [
    "arn:aws:s3:::%s"
   ]
  },
  {
   "Action": [
    "s3:GetObject"
   ],
   "Effect": "Allow",
   "Principal": {
    "AWS": [
     "*"
    ]
   },
   "Resource": [
    "arn:aws:s3:::%s/*"
   ]
  }
 ],
 "Version": "2012-10-17"
}`, cfg.Bucket, cfg.Bucket)
