package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/20-CleanArch/configurations"
	"github.com/devfullcycle/20-CleanArch/internal/event/handler"
	"github.com/devfullcycle/20-CleanArch/internal/infra/graph"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/service"
	"github.com/devfullcycle/20-CleanArch/internal/infra/web/webserver"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conf, err := configurations.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(conf.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// migrations
	cmd := exec.Command(
		"migrate",
		"-path=sql/migrations",
		"-database=mysql://"+conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+":"+conf.DBPort+")/orders",
		"-verbose",
		"up",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar as migrations: %v. Output: %s", err, string(output))
	}
	log.Println("Migrations executadas com sucesso.")

	rabbitMQChannel := getRabbitMQChannel(conf.RabbitMQHost, conf.RabbitMQPort, conf.RabbitMQUser, conf.RabbitMQPassword)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	eventDispatcher.Register("OrderGet", &handler.OrderGetHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	getOrdersUSeCase := NewGetOrdersUseCase(db, eventDispatcher)

	httpServer := webserver.NewWebServer(conf.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	getOrderHandler := NewGetOrdersHandler(db, eventDispatcher)
	httpServer.AddHandler("/orders", webOrderHandler.Create)
	httpServer.AddHandler("/orders/get", getOrderHandler.GetByID)
	httpServer.AddHandler("/orders/all", getOrderHandler.GetAll)
	fmt.Println("Starting web server on port", conf.WebServerPort)
	go httpServer.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *getOrdersUSeCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", conf.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		GetOrderUseCase:    *getOrdersUSeCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", conf.GraphQLServerPort)
	http.ListenAndServe(":"+conf.GraphQLServerPort, nil)
}

func getRabbitMQChannel(host, port, user, password string) *amqp.Channel {
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
