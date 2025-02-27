# coding=utf-8
# Copyright 2024 The TensorFlow Datasets Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Caltech birds dataset."""

import collections
import concurrent.futures
import os
import re

from etils import epath
import numpy as np
from tensorflow_datasets.core.utils.lazy_imports_utils import tensorflow as tf
import tensorflow_datasets.public_api as tfds
import matplotlib.pyplot as plt

# Construct a tf.data.Dataset
ds,info = tfds.load('caltech_birds2011', split='train', as_supervised=True, shuffle_files=True, with_info=True)

# Get class names
label_names = info.features['label'].names
# Print dataset structure
for image, label in ds.take(1):
    print("Image shape:", image.shape)
    print("Label:", label.numpy())

# Plot first 5 images
plt.figure(figsize=(10, 5))
for i, (image, label) in enumerate(ds.take(5)):
    plt.subplot(1, 5, i+1)
    plt.imshow(image.numpy())
    plt.axis("off")
    plt.title(label_names[label.numpy()], fontsize=8)  # Get label name
plt.show()