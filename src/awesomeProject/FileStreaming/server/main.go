package main

import (
	pb "awesomeProject/FileStreaming/fileStreampb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

const (
	port = ":8023"
)

type (
	server struct{}
)

var(
	fileMap []string
)

func buildPath(dir string){

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error{
		if strings.HasSuffix(path, ".jpg"){
			fileMap = append(fileMap, path)
		}
		if err != nil {
			log.Fatalf("The error is %v", err)
		}

		return nil
	})
}

//func totalChunks(fileSize int64) uint64 {
//	var totalParts uint64
//	var imageChunk = 1 * (1 << 20)
//
//	totalParts = uint64(math.Ceil(float64(fileSize) / float64(imageChunk)))
//
//	return totalParts
//}

func (s *server) ShowFiles(dir *pb.Folder, stream pb.ShareFileService_ShowFilesServer) error {
	
	buildPath(dir.Folder)

	for _, f := range fileMap {
		stream.Send(&pb.FileName{FileName: f})
	}

	return nil
}

func (s *server) ServerUpload(fn *pb.FileName, stream pb.ShareFileService_ServerUploadServer) error {

	image, err := os.Open(fn.FileName)
	if err != nil {
		log.Fatalf("the requested file does not exist %v", err)
	}

	stream.SendHeader(metadata.Pairs("filename: ", fn.FileName))
	fmt.Printf("The stream Header() is: %v", fn.FileName)

	//imageInfo, _ := image.Stat()

	chunkContent := make([]byte, 2048*2048) //totalChunks(imageInfo.Size())

	for {
		_, err := image.Read(chunkContent)
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
		}

		fmt.Printf("chunkMessage %v", chunkMessage)

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
	fmt.Println("The server is listening..")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterShareFileServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
