package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Stafford1986/test_manticore/adapter/grpc_serv"
	"github.com/Stafford1986/test_manticore/adapter/manticore_native"
	"github.com/Stafford1986/test_manticore/adapter/mdb"
	"github.com/Stafford1986/test_manticore/pb"
	"github.com/Stafford1986/test_manticore/usecase"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
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
	start := time.Now()

	ctx := context.Background()

	db.SetMaxOpenConns(50)

	resumeRepo := manticore_native.NewResumeRepository(db)
	vacancyRepo := manticore_native.NewVacancyRepository(db)

	resumeUseCase := usecase.NewResumeUseCase(resumeRepo)
	vacancyUseCase := usecase.NewVacancyUseCase(vacancyRepo)

	clientOptions := options.Client().ApplyURI("mongodb://" + "127.0.0.1:27017").SetAuth(credential)

	mongoCli, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("failed create db connection", zap.String("error", err.Error()))
	}

	err = mongoCli.Connect(ctx)
	if err != nil {
		log.Fatal("failed db connect", zap.String("error", err.Error()))
	}

	vacColl := mongoCli.Database("iwrk").Collection("vacancies")

	vacAdapt := mdb.NewRepository(vacColl)

	vacs, err := vacAdapt.GetVacancies(ctx, 100000, 0)

	if err != nil {
		log.Fatal("failed get vac", zap.String("error", err.Error()))
	}

	fmt.Println(len(vacs))
	fmt.Println(vacs[0].IndustryGroups)

	var vetVac []*pb.VacancyEntity
	insertLimit := 3000
	appended := 0
	for _, v := range vacs {
		vv, err := v.Vet()
		if err != nil {
			log.Fatal("failed vet vac", zap.String("error", err.Error()))
		}
		vetVac = append(vetVac, vv)
		appended++
		if appended == insertLimit {
			fmt.Println(appended)
			fmt.Println(len(vetVac))
			vacList := &pb.VacancyList{
				Items: vetVac,
			}

			err = vacancyRepo.BulkInsert(ctx, vacList)
			if err != nil {
				log.Fatal("failed insert vac into indx", zap.String("error", err.Error()))
			}
			vetVac = vetVac[:0]
			appended = 0
			fmt.Println(len(vetVac))
		}
	}

	if appended != 0 {
		fmt.Println("LAST", appended)
		vacList := &pb.VacancyList{
			Items: vetVac,
		}
		err = vacancyRepo.BulkInsert(ctx, vacList)
		if err != nil {
			log.Fatal("failed insert vac into indx", zap.String("error", err.Error()))
		}
	}

	fmt.Println(time.Since(start))

	defer func() {
		if err = mongoCli.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

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
