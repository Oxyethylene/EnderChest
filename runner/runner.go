package runner

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

func Run(router http.Handler) {
	httpHandler := router
	const ADDRESS = "0.0.0.0"
	const PORT = 8080
	addr := fmt.Sprintf("%s:%d", ADDRESS, PORT)
	zap.S().Infow("Started Listening for plain HTTP connection",
		"addr", ADDRESS,
		"port", PORT,
	)
	server := &http.Server{Addr: addr, Handler: httpHandler}

	zap.S().Fatal(server.Serve(startListening(addr, 10)))
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
