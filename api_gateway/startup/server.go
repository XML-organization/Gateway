package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	cfg "api_gateway/startup/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	autentificationGw "github.com/XML-organization/common/proto/autentification_service"
	userGw "github.com/XML-organization/common/proto/user_service"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
	autentificationEmdpoint := fmt.Sprintf("%s:%s", server.config.AutentificationHost, server.config.AutentificationPort)
	err = autentificationGw.RegisterAutentificationServiceHandlerFromEndpoint(context.TODO(), server.mux, autentificationEmdpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
