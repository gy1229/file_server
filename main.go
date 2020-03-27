package main

import (
	"context"
	"fmt"
	"github.com/gy1229/file_server/proto_file"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"runtime"
)

const (
	PORT = ":50001"
)

type Server struct {
	proto_file.FileServerServer
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	listen, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto_file.RegisterFileServerServer(s, &Server{})
	log.Println("rpc服务已经开启")
	s.Serve(listen)
}
func WriteWithIoutil(name string, content []byte) {
	name = "static/img/" + name
	if ioutil.WriteFile(name, content, 0644) == nil {
		fmt.Println("写入文件成功:", name)
	}
}
func (s *Server) UploadFile(ctx context.Context, req *proto_file.UploadFileRequsest) (*proto_file.UploadFileResponse, error) {
	fileName := fmt.Sprintf("%d.%s", req.Id, req.FileType)
	WriteWithIoutil(fileName, req.FileContent)
	//if req.FileType == "img" || req.FileType == "png" || req.FileType == "jpeg"{
	//	WriteWithIoutil(fileName, req.FileContent)
	//}
	return &proto_file.UploadFileResponse{
		Status:               "success",
	}, nil
}

func(s *Server) DownloadFile(ctx context.Context, req *proto_file.DownloadFileRequest) (*proto_file.DownloadFileResponse, error) {
	fileName := fmt.Sprintf("%d.%s", req.Id, req.FileType)
	if req.FileType == "png" {
		fileName = "static/img/" + fileName
	}

	// todo 其他类型再确定
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return &proto_file.DownloadFileResponse{
		FileContent:          fileByte,
		Status:               "success",
	}, nil
}