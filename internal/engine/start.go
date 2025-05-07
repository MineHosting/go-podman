// go-podman/internal/engine/start.go
package engine

import (
	"log"

	"github.com/MineHosting/go-podman/internal/network"
	"github.com/MineHosting/go-podman/internal/socket"
)

func start_rootless_client(sc socket.SocketClientInterface) (*PodmanClient, error) {
	rootlessClient := NewPodmanClient(socket.Rootless, sc)

	err := rootlessClient.ChangeApiVersion()
	if err != nil {
		return nil, err
	}

	return rootlessClient, nil
}

func start_rootfull_client(sc socket.SocketClientInterface) (*PodmanClient, error) {
	rootfullClient := NewPodmanClient(socket.Rootfull, sc)

	err := rootfullClient.ChangeApiVersion()
	if err != nil {
		return nil, err
	}

	return rootfullClient, nil
}

func Start(is_rootless bool) (Client, error) {
	var client *PodmanClient
	var err error

	socketClient := socket.NewSocketClient(&network.RealPayloadSerializer{}, &network.RealHTTPRequestBuilder{}, &network.RealResponseReader{}, &network.RealResponseValidator{}, &network.RealTransportCreator{})

	if is_rootless {
		client, err = start_rootless_client(socketClient)
	} else {
		client, err = start_rootfull_client(socketClient)
	}

	if err != nil {
		log.Println(err)
	}

	err = Ping(client)
	if err != nil {
		log.Println(err)
	}

	return client, err
}
