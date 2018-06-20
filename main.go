package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/imega/teleport-server/health"
	"github.com/imega/teleport-server/mysql"
	"github.com/imega/teleport-server/resolver"
	"github.com/imega/teleport-server/schema"
	"github.com/imega/teleport-server/shutdown"
	"github.com/improbable-eng/go-httpwares/logging/logrus"
	"github.com/improbable-eng/go-httpwares/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	logger *logrus.Entry
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
	})
	logger = logrus.WithField("channel", "graphql-proxy")

	grpcSrv := grpc.NewServer()
	health.RegisterHealthServer(grpcSrv)

	l, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		logger.Errorf("failed to listen on the TCP network address 0.0.0.0:9000, %s", err)
	}

	eDB, err := mysql.NewEntityDB(logger)
	if err != nil {
		logger.Fatalf("failed to create instance db, %s", err)
	}
	health.RegisterHealthCheckFunc(eDB.HealthCheckFunc())
	shutdown.RegisterShutdownFunc(eDB.ShutdownFunc())

	m := http.NewServeMux()

	gqlSchema := graphql.MustParseSchema(schema.String(), &resolver.Resolver{EntityDB: eDB})

	gqlHandler := &relay.Handler{
		Schema: gqlSchema,
	}

	m.Handle("/", gqlContextOwnerIDMiddleware(gqlHandler))
	hm := http_logrus.Middleware(logger, http_logrus.WithRequestFieldExtractor(func(req *http.Request) map[string]interface{} {
		return map[string]interface{}{
			"http.request.x-req-id": "unset",
		}
	}))(m)
	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      hm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
	shutdown.RegisterShutdownFunc(func() {
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		s.Shutdown(ctx)
	})

	go grpcSrv.Serve(l)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			logrus.Errorf("failed to listen on the TCP network address %s and handle requests on incoming connections, %s", s.Addr, err)
		}
	}()

	logger.Info("server is started")
	shutdown.LoopUntilShutdown(15 * time.Second)
	logger.Info("server is stopped")
}

func gqlContextOwnerIDMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			logger = ctxlogrus.Extract(ctx)
		)

		ownerID := r.Header.Get("GRPC-METADATA-X-OWNER-ID")
		if len(ownerID) == 0 {
			logger.Errorf("owner_id not found in headers")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, "owner_id", ownerID)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
