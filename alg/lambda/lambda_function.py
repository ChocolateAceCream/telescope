import json
import cv2
import numpy as np
import boto3
import urllib.request
import os

# S3 Client
s3_client = boto3.client("s3")

# Model and class names paths
MODEL_PATH = "/opt/models/DenseNet_121.caffemodel"
PROTO_PATH = "/opt/models/DenseNet_121.prototxt"
CLASS_NAMES_PATH = "/opt/models/classification_classes_animals.txt"

def load_class_names():
    """Loads the class names from the text file."""
    try:
        with open(CLASS_NAMES_PATH, "r") as f:
            image_net_names = f.read().split("\n")
        class_names = [name.split(",")[0] for name in image_net_names]  # Get first name if multiple exist
        return class_names
    except Exception as e:
        print(f"Error loading class names: {e}")
        raise


def classify_image(image_url):
    """Download image from S3 URL, classify it, and return results."""

    # Download the image to /tmp
    temp_image_path = "/tmp/input.jpg"
    urllib.request.urlretrieve(image_url, temp_image_path)

    # Load model and class names
    class_names = load_class_names()

    model = cv2.dnn.readNet(MODEL_PATH, PROTO_PATH)

    # Read Image
    image = cv2.imread(temp_image_path)
    blob = cv2.dnn.blobFromImage(image=image, scalefactor=0.017, size=(224, 224), mean=(104, 117, 123))

    model.setInput(blob)
    outputs = model.forward()

    # Process results
    label_id = int(np.argmax(outputs[0]))  # Convert to int for JSON compatibility
    confidence = float(np.max(outputs[0])) * 100.0  # Convert to float for JSON
    class_name = class_names[label_id] if label_id < len(class_names) else "Unknown"

    return {"class_name": class_name, "confidence": confidence}


def lambda_handler(event, context):
    """AWS Lambda Entry Point"""
    try:
        print(f"Received event: {event}")

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
        print(f"Lambda function failed: {e}")
        return {"statusCode": 500, "body": json.dumps({"error": str(e)})}
