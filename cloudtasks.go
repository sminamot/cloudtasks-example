package main

import (
	"context"
	"fmt"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type CloudTasksClient struct {
	client     *cloudtasks.Client
	projectID  string
	locationID string
	queueID    string
	url        string
}

func NewCloudTasksClient(ctx context.Context, projectID, locationID, queueID, url string) (*CloudTasksClient, error) {
	cli, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &CloudTasksClient{
		client:     cli,
		projectID:  projectID,
		locationID: locationID,
		queueID:    queueID,
		url:        url,
	}, nil
}

func (c *CloudTasksClient) CreateTasks(ctx context.Context, data proto.Message) error {
	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", c.projectID, c.locationID, c.queueID)
	v, err := protojson.Marshal(data)
	if err != nil {
		return err
	}

	req := &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2#HttpRequest
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        c.url,
					Body:       v,
				},
			},
		},
	}

	createdTask, err := c.client.CreateTask(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudtasks.CreateTask: %v", err)
	}

	//fmt.Println(createdTask)
	fmt.Printf("created task name: %s\n", createdTask.GetName())
	return nil
}
