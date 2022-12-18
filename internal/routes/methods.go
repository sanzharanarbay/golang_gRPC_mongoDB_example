package routes

import (
	pb "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/pb"
	service_servers "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/service-servers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitMethods(port string) {
	appPort := ":" + port
	srvCategory := service_servers.CategoryServiceServer{}
	srvCity := service_servers.CityServiceServer{}
	srvVacancy := service_servers.VacancyServiceServer{}
	lis, err := net.Listen("tcp", appPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, &srvCategory)
	pb.RegisterCityServiceServer(grpcServer, &srvCity)
	pb.RegisterVacancyServiceServer(grpcServer, &srvVacancy)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
