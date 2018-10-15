package main

import (
	pb "awesomeProject/FileStreaming/fileStreampb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

const (
	address = "localhost:8023"
	directory = "FileStreaming/media"
)

var (
	images map[int]string
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}

	defer conn.Close()

	c := pb.NewShareFileServiceClient(conn)
	getImages(c)

	streamImage(c)

}

func streamImage( c pb.ShareFileServiceClient) {

	var keyNumber int
	fmt.Println("Choose a number to stream a file:")
	fmt.Scanf("%d", &keyNumber)
	filePath := images[keyNumber]

	fp := pb.FileName{
		FileName: filePath,
	}

	imageStream, err := c.ServerUpload(context.Background(), &fp)
	if err != nil {
		log.Fatalf("Client. The error is: %v", err)
	}

	meta, err := imageStream.Header()
	if err != nil || meta["filename"] == nil || meta["filename"][0] == "" {
		log.Printf("The file does not exist anywhere %v", err)
	}

	outfile, err := os.Create(meta["filename"][0])
	if err != nil {
		log.Fatal(err)
	}

	for {
		chunkPart, err := imageStream.Recv()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("The stream reported an error %v", err)
			} else {
				break
			}

		}
		outfile.Write(chunkPart.Batch.Content)

	}
}

func getImages(c pb.ShareFileServiceClient){
	// This part is the choice of the specific image and the display to a browser
	fmt.Println("The server files are:")

	stream, err := c.ShowFiles(context.Background(), &pb.Folder{Folder: directory,})
	if err != nil {
		log.Fatalf("could not find the files: %v", err)
	}
	// initialise the counter that will be used as a key to the file map
	counter := 1
	images = make(map[int]string)
	for {
		feature, err := stream.Recv()

		if err == io.EOF {
			// end of files
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		// store the file name in a map with an integer key
		images[counter] = feature.FileName

		fmt.Printf("%d. %s\n", counter, feature.FileName)
		// increment the counter for each new file
		counter++
	}

}
