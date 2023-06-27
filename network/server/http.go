package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zutim/ztm-demo/network/ngin/handler"
	"github.com/zutim/ztm-demo/network/ngin/router"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"os"
	"path/filepath"
	"reflect"
)

type http struct {
	Engine *gin.Engine
}

func NewHttp() *http {
	r := gin.Default()
	router.LoadRouter(r)

	return &http{
		Engine: r,
	}
}

func (h *http) ReadRoot() {
	currentDir, _ := os.Getwd()
	//// 打印当前目录
	fmt.Println("当前目录:", currentDir)
	// 指定要遍历的 proto 文件所在的目录路径
	protoDir := "./proto"

	targetPath := filepath.Join(currentDir, protoDir)

	// 打印目标路径
	fmt.Println("目标路径:", targetPath)

	// 遍历指定目录下的 .proto 文件
	err := filepath.Walk(protoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// 判断是否为 .proto 文件
		if !info.IsDir() && filepath.Ext(path) == ".pb" {
			// 加载并解析 .proto 文件
			descriptorSet := &descriptorpb.FileDescriptorSet{}
			fmt.Println(path)
			fileBytes, err := os.ReadFile(path)
			if err != nil {
				fmt.Println(1, err)
				return nil
			}
			if err := proto.Unmarshal(fileBytes, descriptorSet); err != nil {
				fmt.Println(2, err)
				return nil
			}

			// 注册 .proto 文件中定义的服务路由
			h.registerProtoRoutes(descriptorSet)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h *http) registerProtoRoutes(set *descriptorpb.FileDescriptorSet) {
	for _, fileDescriptorProto := range set.GetFile() {
		for _, serviceDescriptorProto := range fileDescriptorProto.GetService() {
			for _, methodDescriptorProto := range serviceDescriptorProto.GetMethod() {
				fullMethodName := fmt.Sprintf("/%s/%s",
					serviceDescriptorProto.GetName(),
					methodDescriptorProto.GetName(),
				)

				HandlerName := fmt.Sprintf("%s%s",
					serviceDescriptorProto.GetName(),
					methodDescriptorProto.GetName(),
				)

				//in := methodDescriptorProto.GetInputType()
				//out := methodDescriptorProto.GetOutputType()
				//
				//in = strings.Replace(in, ".proto", "v1", -1)
				//out = strings.Replace(out, ".proto", "v1", -1)

				//fmt.Println(in)
				//fmt.Println(out)

				fmt.Printf("Registering route: %s\n", fullMethodName)

				hand := handler.NewHandler()

				handlerValue := reflect.ValueOf(hand)

				methodValue := handlerValue.MethodByName(HandlerName)

				funcValue := methodValue.Interface().(func(*gin.Context))

				h.Engine.POST(fullMethodName, func(ctx *gin.Context) {
					funcValue(ctx)
				})
			}
		}
	}
}

func (h *http) Run() {
	h.Engine.Run(viper.GetString("http"))
}
