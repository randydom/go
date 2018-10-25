# For this Dockerfile I will make an image from scratch
# One of the objectives are to compare the image from scratch with the golang.alpine image
# Check the actual image as well
FROM scratch

# I add the server and the protobuf file to the image.
ADD main /

EXPOSE 8080/tcp
EXPOSE 9000/tcp

ENTRYPOINT ["/main"]


