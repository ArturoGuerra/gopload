package providers


const BucketExistsError = "Your previous request to create the named bucket succeeded and you already own it."
const policy = `
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
    "arn:aws:s3:::images"
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
    "arn:aws:s3:::images/*"
   ]
  }
 ],
 "Version": "2012-10-17"
}`