package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// 创建容器
	boyd, err := cli.ContainerCreate(context.TODO(), &container.Config{
		Tty:       true,
		OpenStdin: true,
		Image:     "nginx:1.16.1", // 该容器使用的镜像
	}, &container.HostConfig{
		// 容器端口 80 tcp 协议，暴露至宿主机的 8081，容器名为 testnginx
		PortBindings: nat.PortMap{nat.Port("80/tcp"): []nat.PortBinding{{"0.0.0.0", "8082"}}},
	}, nil, nil, "testnginx2")

	if err != nil {
		panic(err)
	}
	fmt.Println(boyd, err)

	containerID := boyd.ID
	err = cli.ContainerStart(context.TODO(), containerID, types.ContainerStartOptions{})
	fmt.Println(err)

}
