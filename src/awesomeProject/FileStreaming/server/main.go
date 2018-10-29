package main

import (
	pb "awesomeProject/FileStreaming/fileStreampb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	port = ":8080"
	portB = ":9000"
	directory = "FileStreaming/media/"
)

type (
	server struct{}
)

var(
	fileMap []string
	errChan chan error
)

func buildPath(dir string){
	// initialise the file map
	fileMap = []string{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error{
		if strings.HasSuffix(path, ".jpg"){
			fileMap = append(fileMap, info.Name())
		}
		if err != nil {
			log.Fatalf("The error is %v", err)
		}

		return nil
	})
}

func totalChunks(fileSize int64) uint64 {
	var totalParts uint64
	var imageChunk = 1 * (1 << 20)

	totalParts = uint64(math.Ceil(float64(fileSize) / float64(imageChunk)))

	return totalParts
}

func (s *server) ShowFiles(fn *pb.FileName, stream pb.ShareFileService_ShowFilesServer) error {
	
	fn.Dir.Folder = directory
	//build the path for the media files.
	buildPath(fn.Dir.Folder)

	for _, f := range fileMap {
		stream.Send(&pb.FileName{FileName: f})
	}

	return nil
}

func (s *server) ServerUpload(fn *pb.FileName, stream pb.ShareFileService_ServerUploadServer) error {

	fn.Dir.Folder = directory

	image, err := os.Open(fn.Dir.Folder + fn.FileName)
	if err != nil {
		log.Fatalf("the requested file does not exist %v", err)
	}

	errStream := stream.SendHeader(metadata.Pairs("filename", fn.FileName))
	if err != nil{
		log.Fatalf("The SendHeader had an error: %v", errStream)
	}

	fmt.Printf("The stream Header() is: %v\n", fn.FileName)

	imageInfo, _ := image.Stat()

	chunkContent := make([]byte, totalChunks(imageInfo.Size())) //totalChunks(imageInfo.Size()) //1024*1024

	for {

		bytesRead, err := image.Read(chunkContent)
		if err != nil {
			if err != io.EOF{
				return err
			}else {
				break
			}

		}

		chunkMessage := pb.ChunkPackage{
			Batch: &pb.Chunk{Content:chunkContent,},
			Status: &pb.UploadStatus{
				Message: "the package was sent.",
				Code: pb.UploadStatusCode_OK,
			},
			TotalSize: int64(imageInfo.Size()),
			BytesSent: int64(bytesRead),
		}

		err = stream.Send(&chunkMessage)
		if err != nil{
			log.Fatalf("Stream error: %v",err)
		}
	}

	return nil
}

/*
	The Main Program...
 */

func main() {
	go servegRPC()

	go serveHTTP()

	log.Fatal(<- errChan)

}
func serveHTTP() {

	fmt.Printf("The server is also listening to 9000..\n")

	fs := http.FileServer(http.Dir("/Users/pavlos/go/"))
	http.Handle("/", fs)

	errChan <- http.ListenAndServe("127.0.0.1:" + portB, nil)
}


func servegRPC() {
	fmt.Println("The server is listening to 8080..")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShareFileServiceServer(s, &server{})
	reflection.Register(s)
	errChan <- s.Serve(lis)
}

