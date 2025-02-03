import json
import cv2
import numpy as np
import boto3
import tempfile
import urllib.request

# S3 Client
s3_client = boto3.client("s3")

# Model paths in /tmp (Lambda only allows write access here)
MODEL_PATH = "/tmp/DenseNet_121.caffemodel"
PROTO_PATH = "/tmp/DenseNet_121.prototxt"

# S3 Bucket where model files are stored
MODEL_BUCKET = "telescope-develop"
CAFFEMODEL_KEY = "DenseNet_121.caffemodel"
PROTOTXT_KEY = "DenseNet_121.prototxt"


def download_model():
    """Download model files from S3 to /tmp if not already present."""
    for key, path in [(CAFFEMODEL_KEY, MODEL_PATH), (PROTOTXT_KEY, PROTO_PATH)]:
        try:
            s3_client.download_file(MODEL_BUCKET, key, path)
        except Exception as e:
            print(f"Error downloading model: {e}")
            raise


def classify_image(image_url):
    """Download image from S3 URL, classify it, and return results."""

    # Download the image to /tmp
    temp_image_path = "/tmp/input.jpg"
    urllib.request.urlretrieve(image_url, temp_image_path)

    # Load model
    download_model()
    model = cv2.dnn.readNet(MODEL_PATH, PROTO_PATH)

    # Read Image
    image = cv2.imread(temp_image_path)
    blob = cv2.dnn.blobFromImage(image, scalefactor=0.017, size=(224, 224), mean=(104, 117, 123))

    model.setInput(blob)
    outputs = model.forward()

    # Process results
    label_id = int(np.argmax(outputs[0]))  # Convert to int for JSON compatibility
    confidence = float(np.max(outputs[0])) * 100.0  # Convert to float for JSON

    return {"class_id": label_id, "confidence": confidence}


def lambda_handler(event, context):
    """AWS Lambda Entry Point"""
    try:
        # Get the S3 image URL from the input event
        body = json.loads(event["body"])
        image_url = body.get("image_url")

        if not image_url:
            return {"statusCode": 400, "body": json.dumps({"error": "Missing image_url"})}

        # Run classification
        result = classify_image(image_url)

        return {
            "statusCode": 200,
            "body": json.dumps(result)
        }

    except Exception as e:
        return {"statusCode": 500, "body": json.dumps({"error": str(e)})}
