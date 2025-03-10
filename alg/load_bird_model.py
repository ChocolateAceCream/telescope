# ✅ Import Required Libraries
import keras
import numpy as np
import tensorflow as tf
import matplotlib.pyplot as plt
from tensorflow.keras.preprocessing import image
import tensorflow_datasets as tfds

# ✅ Step 1: Load the Trained Model
print("🔄 Loading Model...")
# model = keras.models.load_model("caltech_birds_densenet.keras", compile=False)
model = keras.models.load_model("cnn_model_v3.h5", compile=False)
print("✅ Model Loaded Successfully!")

# ✅ Step 2: Load Class Labels from Dataset
ds_info = tfds.builder("caltech_birds2011").info
label_names = ds_info.features["label"].names  # List of 200 bird species names

# ✅ Step 3: Image Preprocessing Function
def preprocess_image(img_path):
    """
    Load and preprocess an image for model prediction.
    - Resizes the image to 224x224 (DenseNet input size).
    - Converts to NumPy array and normalizes pixel values (0-1).
    - Adds batch dimension.
    """
    img = image.load_img(img_path, target_size=(224, 224))  # Resize
    img_array = image.img_to_array(img)  # Convert to NumPy array
    img_array = np.expand_dims(img_array, axis=0)  # Add batch dimension
    img_array /= 255.0  # Normalize pixel values (same as training)
    return img_array, img  # Return both preprocessed array & original image

# ✅ Step 4: Prediction Function
def predict_bird(img_path):
    """
    Predict the bird species from an image using the trained model.
    - Preprocesses the image.
    - Uses the model to predict the class.
    - Displays the image with the predicted bird species.
    """
    img_array, img = preprocess_image(img_path)

    # 🔍 Predict the class
    predictions = model.predict(img_array)
    predicted_class = np.argmax(predictions, axis=1)[0]  # Get highest probability class index
    predicted_label = label_names[predicted_class]  # Convert index to bird name

    # 🎨 Display the image with prediction
    plt.figure(figsize=(6, 6))
    plt.imshow(img)
    plt.axis("off")
    plt.title(f"Predicted: {predicted_label}", fontsize=14)
    plt.show()

    print(f"🔍 Predicted Bird Species: {predicted_label}")

# ✅ Step 5: Test with an Image
img_path = "cow.png"  # Replace with your own bird image path
predict_bird(img_path)
