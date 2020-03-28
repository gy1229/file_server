package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gy1229/file_server/proto_file"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
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
	go HttpServer()
	s.Serve(listen)
}

func HttpServer() {
	r := gin.Default()
	r.GET("/img/:name", DownLoadImage)
	r.Run(":50002")
}

func DownLoadImage(c *gin.Context) {
	name := c.Param("name")
	fileName := fmt.Sprintf("static/img/%s", name)
	file, err := os.Open(fileName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "cant find file",
		})
		return
	}
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "cant find file",
		})
		return
	}
	c.Writer.Write(fileByte)
}