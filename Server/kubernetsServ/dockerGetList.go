package kubernetsServ

// import (
// 	"context"
// 	"fmt"

// 	"github.com/docker/docker/api/types"
// 	"github.com/docker/docker/client"
// 	"github.com/moby/moby/api/types"
// )

// func GetDocker() {

// 	cli, err := ConnectDocker()
// 	if err != "" {
// 		return
// 	}

// 	fmt.Println("Connect to docker success.")
// 	GetContainers(cli)
// }

// func ConnectDocker() (cli *client.Client, error string) {
// 	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation(), client.WithHost("tcp://192.168.1.211:2375"))
// 	if err != nil {
// 		return
// 	}
// 	return cli
// }

// func GetContainers(cli *client.Client) error {
// 	containers, err := cli.ContainerList(context.Background(), types.ContinerListOptions{All: true})
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, container := range containers {
// 		fmt.Println("%s %s \n", container.ID[:10], container.Image)
// 	}
// 	return nil
// }
