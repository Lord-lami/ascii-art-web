#!/bin/bash

# Build the image
imagename="sheeplami/ascii-art-web-docker:1.0"
echo "BUILDING THE $imagename IMAGE"
docker image build -f Dockerfile --label "env=dev" -t $imagename .

# Run the container
containername="ascii-art-web-container"
echo "STARTING THE $containername CONTAINER OFF $imagename IMAGE"
docker container run -p 8080:8080 --label "env=dev" --detach --name $containername $imagename