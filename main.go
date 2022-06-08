package main

import (
	"flag"
	"fmt"
	"github.com/Stafford1986/test_manticore/adapter/grpc_serv"
	"github.com/Stafford1986/test_manticore/adapter/manticore_native"
	"github.com/Stafford1986/test_manticore/pb"
	"github.com/Stafford1986/test_manticore/usecase"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 9000, "service port")
}

func main() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("@tcp(%s:%d)/", "127.0.0.1", 9306))
	if err != nil {
		log.Fatal(err)
	}

	resumeRepo := manticore_native.NewResumeRepository(db)
	vacancyRepo := manticore_native.NewVacancyRepository(db)

	resumeUseCase := usecase.NewResumeUseCase(resumeRepo)
	vacancyUseCase := usecase.NewVacancyUseCase(vacancyRepo)

	svc := grpc_serv.New(
		resumeUseCase,
		vacancyUseCase,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to listen", zap.String("error", err.Error()))
	}

	var serverOptions []grpc.ServerOption
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
	serverOptions = append(serverOptions, grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()))

	grpcServer := grpc.NewServer(serverOptions...)
	defer grpcServer.Stop()

	pb.RegisterSearchServiceServer(grpcServer, svc)
	reflection.Register(grpcServer)
	log.Print(fmt.Sprintf("%s running ::%d", "test", port))
	log.Fatal(grpcServer.Serve(lis).Error())
}
