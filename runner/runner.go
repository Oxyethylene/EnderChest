package runner

import (
	"context"
	"fmt"
	"github.com/Oxyethylene/EnderChest/config"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

func Run(router http.Handler) {
	httpHandler := router
	address := config.ServerConfig.Addr
	port := config.ServerConfig.Port
	addr := fmt.Sprintf("%s:%d", address, port)
	server := &http.Server{Addr: addr, Handler: httpHandler}
	listener := startListening(addr, 10)
	zap.S().Infow("Started Listening for plain HTTP connection",
		"addr", listener.Addr().String(),
	)

	zap.S().Fatal(server.Serve(listener))
}

func startListening(addr string, keepAlive int) net.Listener {
	lc := net.ListenConfig{KeepAlive: time.Duration(keepAlive) * time.Second}
	conn, err := lc.Listen(context.Background(), "tcp", addr)
	if err != nil {
		zap.S().Fatalw("listen failed",
			"addr", addr,
			zap.Error(err),
		)
	}
	return conn
}
