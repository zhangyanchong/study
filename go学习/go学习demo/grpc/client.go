package main
import (
	"log"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	User "grpc/pei"
)


func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := User.NewSimpleClient(conn)

	// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}

	// 调用gRPC接口
	tr, err := t.GetList(context.Background(), &User.SimpleRequest{Data: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应code: %d,服务端响应val:%s", tr.Code,tr.Value)
}
