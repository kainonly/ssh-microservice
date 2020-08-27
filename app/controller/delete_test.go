package controller

import (
	"context"
	pb "ssh-microservice/router"
	"testing"
)

func TestController_Delete(t *testing.T) {
	response, err := client.Delete(
		context.Background(),
		&pb.DeleteParameter{
			Identity: "debug-1",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
	response, err = client.Delete(
		context.Background(),
		&pb.DeleteParameter{
			Identity: "debug-2",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
}
