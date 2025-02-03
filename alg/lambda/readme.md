# Usage
1. run build.sh
2. create a layer using lambda-layer.zip from /out
3. create lambda function using lambda_function.py
4. upload image to s3 and copy the image url
5. test using event JSON
```json
{
  "body": "{\"image_url\": \"https://telescope-develop.s3.us-east-1.amazonaws.com/pexels-photo-306036.jpeg\"}",
  "httpMethod": "POST",
  "headers": {
    "Content-Type": "application/json"
  },
  "isBase64Encoded": false
}

```