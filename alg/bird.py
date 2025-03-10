# coding=utf-8
import tensorflow as tf
import tensorflow_datasets as tfds
import matplotlib.pyplot as plt
import numpy as np
from tensorflow.keras.applications import DenseNet121
from tensorflow.keras import layers, models

# ✅ Load the Caltech Birds dataset
ds, info = tfds.load('caltech_birds2011', split='train', as_supervised=True, shuffle_files=True, with_info=True)

# ✅ Extract class names
label_names = info.features["label"].names
num_classes = len(label_names)  # 200 bird species

# ✅ Data Augmentation function
def augment(image, label):
    """Apply data augmentation transformations."""
    image = tf.image.resize(image, (224, 224))  # Resize for DenseNet
    image = tf.image.random_flip_left_right(image)  # Random horizontal flip
    image = tf.image.random_brightness(image, max_delta=0.2)  # Adjust brightness
    image = tf.image.random_contrast(image, lower=0.8, upper=1.2)  # Adjust contrast
    image = tf.image.random_saturation(image, lower=0.8, upper=1.2)  # Adjust saturation
    image = tf.cast(image, tf.float32) / 255.0  # Normalize (0-1)

    return image, label

# ✅ Apply augmentation & batch the dataset
batch_size = 32
train_ds = ds.map(augment, num_parallel_calls=tf.data.AUTOTUNE)
train_ds = train_ds.batch(batch_size).shuffle(1000).prefetch(tf.data.AUTOTUNE)

# ✅ Load Pretrained DenseNet121 Model
base_model = DenseNet121(weights="imagenet", include_top=False, input_shape=(224, 224, 3))
base_model.trainable = False  # Freeze base model

# ✅ Build Model
model = models.Sequential([
    base_model,
    layers.GlobalAveragePooling2D(),  # Pool features
    layers.Dense(512, activation="relu"),
    layers.Dropout(0.5),  # Prevent overfitting
    layers.Dense(num_classes, activation="softmax")  # 200 bird classes
])

# ✅ Compile Model
model.compile(optimizer="adam",
              loss="sparse_categorical_crossentropy",
              metrics=["accuracy"])

# ✅ Train the Model
epochs = 200
history = model.fit(train_ds, epochs=epochs)

# ✅ Plot Training History
# plt.figure(figsize=(12, 5))

# # Accuracy
# plt.subplot(1, 2, 1)
# plt.plot(history.history["accuracy"], label="Train Accuracy")
# plt.xlabel("Epoch")
# plt.ylabel("Accuracy")
# plt.legend()
# plt.title("Training Accuracy")

# # Loss
# plt.subplot(1, 2, 2)
# plt.plot(history.history["loss"], label="Train Loss")
# plt.xlabel("Epoch")
# plt.ylabel("Loss")
# plt.legend()
# plt.title("Training Loss")

# plt.show()

# ✅ Save Model
model.save("caltech_birds_densenet.keras")
print("✅ Model saved as caltech_birds_densenet.keras")
