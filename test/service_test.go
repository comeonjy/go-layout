package test

import (
	"context"
	"testing"

	"github.com/comeonjy/go-kit/grpc/reloadconfig"
	"google.golang.org/grpc"
)


func TestReloadConfig(t *testing.T) {
	dial, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	client := reloadconfig.NewReloadConfigClient(dial)
	_, err = client.ReloadConfig(context.TODO(), &reloadconfig.Empty{})
	if err != nil {
		t.Error(err)
	}
}
