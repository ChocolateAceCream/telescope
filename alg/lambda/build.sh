docker build -t lambda-layer-opencv .
docker run --rm -v "$PWD/out":/out lambda-layer-opencv
