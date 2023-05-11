package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cfg "api_gateway/startup/config"

	accomodationGw "github.com/XML-organization/common/proto/accomodation_service"
	autentificationGw "github.com/XML-organization/common/proto/autentification_service"
	bookingGw "github.com/XML-organization/common/proto/booking_service"
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
	bookingEmdpoint := fmt.Sprintf("%s:%s", server.config.BookingHost, server.config.BookingPort)
	err = bookingGw.RegisterBookingServiceHandlerFromEndpoint(context.TODO(), server.mux, bookingEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	accomodationEmdpoint := fmt.Sprintf("%s:%s", server.config.AccomodationHost, server.config.AccomodationPort)
	err = accomodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accomodationEmdpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), cors(server.mux)))
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
