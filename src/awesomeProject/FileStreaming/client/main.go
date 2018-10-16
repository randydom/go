package main

import (
	pb "awesomeProject/FileStreaming/fileStreampb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"math"
	"os"
)

const (
	address = "localhost:8023"
	directory = "FileStreaming/media/"
)

var (
	images = map[int]string{}
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}

	defer conn.Close()

	c := pb.NewShareFileServiceClient(conn)

	for {
		getImages(c)

		imageFile,isStreamAlive := pickImage()

		if isStreamAlive{
			fmt.Println("Good bye then.")
			break
		}

		streamImage(*imageFile, c)
	}



}
func getImages(c pb.ShareFileServiceClient){
	// This part is the choice of the specific image and the display to a browser
	fmt.Println("\nThe server files are:")

	fName := pb.FileName{
		FileName: "",
		Dir: &pb.Folder{Folder: directory,},
	}

	stream, err := c.ShowFiles(context.Background(), &fName)
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

func pickImage() (*pb.FileName, bool){

	var keyNumber int
	var exitMsg string
	var exitChat bool
	var fp pb.FileName

	fmt.Println("Do you want to exit the chat? Y/N")
	fmt.Scanf("%s", &exitMsg)

	if exitMsg == "Y" || exitMsg == "y"{
		exitChat = true

		fp = pb.FileName{
			FileName: "",
			Dir: &pb.Folder{ Folder: directory,},
		}
	}else
	{
		fmt.Println("Choose a number to stream a file:")
		fmt.Scanf("%d", &keyNumber)
		filePath := images[keyNumber]

		fp = pb.FileName{
			FileName: filePath,
			Dir: &pb.Folder{ Folder: directory,},
		}

		exitChat = false
	}

	return &fp, exitChat
}

func streamImage(fp pb.FileName, c pb.ShareFileServiceClient) {

	totalSent := int64(0)

	imageStream, err := c.ServerUpload(context.Background(), &fp)
	if err != nil {
		log.Fatalf("Client. The error is: %v", err)
	}

	meta, err := imageStream.Header()

	if err != nil || meta["filename"] == nil || meta["filename"][0] == "" {
		log.Fatalf("The file does not exist anywhere %v", err)
	}

	outfile, err := os.Create("/Users/pavlos/go/"+meta["filename"][0])
	fmt.Println(meta["filename"])
	if err != nil {
		log.Fatalf("Create os faced an error: %v", err)
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

		totalSent += chunkPart.BytesSent
		fmt.Print("\r", math.Round(float64(totalSent)/float64(chunkPart.TotalSize)*100), "%")

	}
}


