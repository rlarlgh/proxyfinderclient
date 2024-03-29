package proxyfinderclient

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/rlarlgh/proxyfinderclient/grpc/proxyfinder"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type ProxyFinderClient struct {
	Host string
	Port int
	// client pb.ProxyFinderClient
}

func NewProxyFinderClient(host string, port int) *ProxyFinderClient {
	return &ProxyFinderClient{
		Host: host,
		Port: port,
	}
}

func (c *ProxyFinderClient) FindProxyServer(targetUrl string) (string, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	conn, err := c.tryConnect(c.Host, c.Port)
	if err != nil {
		return "", 0, err
	}
	defer conn.Close()
	client := pb.NewProxyFinderClient(conn)
	resp, err := client.FindProxyServer(ctx, &pb.FindProxyServerRequest{TargetUrl: targetUrl})
	if err != nil {
		return "", 0, err
	}
	// fmt.Printf("Found proxy server: %s:%d\n", resp.Host, resp.Port)
	return resp.Host, int(resp.Port), nil
}

// func (c *ProxyFinderClient) startConn(host string, port int) {
// 	conn, err := c.tryConnect(host, port)
// 	if err != nil {
// 		log.Fatalf("Failed to connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c.client = pb.NewProxyFinderClient(conn)
// }

func (c *ProxyFinderClient) tryConnect(host string, port int) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	var conn *grpc.ClientConn
	var err error

	// 지수 백오프 설정
	backoffConfig := backoff.Config{
		BaseDelay:  1.0 * time.Second,  // 최소 지연 시간
		Multiplier: 1.5,                // 지연 시간 증가 인자
		MaxDelay:   60.0 * time.Second, // 최대 지연 시간
	}

	// 연결 옵션 설정
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), // TLS 설정
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: backoffConfig}),
		grpc.WithBlock(), // 초기 연결 시 블록킹 모드
	}

	// 연결 시도
	conn, err = grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	// 연결 상태 모니터링
	go func() {
		for {
			state := conn.GetState()
			log.Printf("Connection state: %v\n", state)
			if state == connectivity.TransientFailure {
				// 연결 상태가 TransientFailure이면 재연결 시도
				conn.ResetConnectBackoff()
			}

			// 연결 상태가 변할 때까지 대기
			if !conn.WaitForStateChange(context.Background(), state) {
				// 연결이 종료되면 루프 종료
				break
			}
		}
	}()

	return conn, nil
}
