# Use Amazon Linux 2023 with Python 3.12 (AWS Lambda base image)
FROM public.ecr.aws/lambda/python:3.12

# Install system dependencies (including zip)
RUN dnf install -y \
    gcc-c++ make cmake unzip zip \
    python3-devel python3-pip python3-setuptools \
    mesa-libGL libX11 libXext libXrender && \
    dnf clean all

# Install OpenCV and NumPy inside the Lambda layer directory
RUN mkdir -p /opt/python
RUN pip install --upgrade pip && \
    pip install --target=/opt/python opencv-python-headless numpy

# Copy model and class names file into the layer
RUN mkdir -p /opt/models
COPY DenseNet_121.caffemodel /opt/models/DenseNet_121.caffemodel
COPY DenseNet_121.prototxt /opt/models/DenseNet_121.prototxt
COPY classification_classes_animals.txt /opt/models/classification_classes_animals.txt

# Set working directory
WORKDIR /opt

# Zip the layer for AWS Lambda
RUN zip -r /lambda-layer.zip python models

# Override entrypoint (Disable Lambda runtime)
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["cp /lambda-layer.zip /out/lambda-layer.zip"]
