package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/sminamot/cloudtasks-example/task"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	ct *CloudTasksClient
	task.UnimplementedTaskAPIServer
}

var _ task.TaskAPIServer = &Server{}

func (s *Server) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*emptypb.Empty, error) {
	m := &task.HandleTaskRequest{
		Name: req.Name,
	}
	if err := s.ct.CreateTasks(ctx, m); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) HandleTask(ctx context.Context, req *task.HandleTaskRequest) (*emptypb.Empty, error) {
	var taskName, retryCount string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		taskName = strings.Join(md.Get("X-CloudTasks-TaskName"), ",")
		if cs := md.Get("X-CloudTasks-TaskRetryCount"); len(cs) > 0 {
			retryCount = cs[0]
		}
	}
	fmt.Printf("task_name: %s, name: %s\n", taskName, req.GetName())
	fmt.Printf("retry_count: %s\n", retryCount)

	return &emptypb.Empty{}, nil
}

const defaultPort = 13333

func main() {
	ctx := context.Background()

	credentials, err := google.FindDefaultCredentials(ctx)
	_, _ = credentials, err
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get projectID:", err)
		os.Exit(1)
	}
	projectID := credentials.ProjectID

	ctHandleEndpoint, ok := os.LookupEnv("CLOUD_TASKS_HANDLE_ENDPOINT")
	if !ok {
		fmt.Fprintln(os.Stderr, "CLOUD_TASKS_HANDLE_ENDPOINT is not set")
		os.Exit(1)
	}
	ct, err := NewCloudTasksClient(ctx, projectID, "asia-northeast1", "test-queue2", ctHandleEndpoint)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create CloudTasksClient:", err)
		os.Exit(1)
	}

	serv := grpc.NewServer()
	reflection.Register(serv)

	task.RegisterTaskAPIServer(serv, &Server{ct: ct})

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = fmt.Sprint(defaultPort)
	}
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to listen:", err)
		os.Exit(1)
	}

	if err := serv.Serve(l); err != nil {
		fmt.Fprintln(os.Stderr, "failed to serve:", err)
		os.Exit(1)
	}
}
