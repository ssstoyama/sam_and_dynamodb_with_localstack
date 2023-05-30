package main

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestMain(m *testing.M) {
	os.Setenv("AWS_REGION", "ap-northeast-1")
	os.Setenv("AWS_ENDPOINT", "http://localhost:4566")

	os.Exit(m.Run())
}

func TestHandler(t *testing.T) {
	const expected = `[{"id":"id1"},{"id":"id2"},{"id":"id3"}]`

	ctx := context.Background()

	res, err := handler(ctx, &events.APIGatewayProxyRequest{})

	if err != nil {
		t.Fatal(err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatal(res.StatusCode)
	}
	if expected != res.Body {
		t.Fatal(res.Body)
	}
}
