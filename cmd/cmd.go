package cmd

import (
	"fmt"
	"log"

	"zayyid-go/delivery/container"
	"zayyid-go/delivery/http"
)

func Execute() {
	container := container.SetupContainer()

	http := http.ServeHttp(container)
	if err := http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port)); err != nil {
		// Let it panic when thing goes wrong when running the server
		log.Fatalf("Error starting HTTP server: %s\n", err)
	}
}
