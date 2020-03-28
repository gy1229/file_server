package main

import (
	"context"
	"fmt"
	"github.com/gy1229/file_server/proto_file"
	"io/ioutil"
	"os"
)

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
