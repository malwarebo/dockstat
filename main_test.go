package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
)

type MockDockerClient struct {
	containers []types.Container
	err        error
}

func (m *MockDockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return m.containers, m.err
}

func (m *MockDockerClient) ContainerKill(ctx context.Context, containerID, signal string) error {
	return m.err
}

func (m *MockDockerClient) ContainerLogs(ctx context.Context, containerID string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	if m.err != nil {
		return nil, m.err
	}
	return io.NopCloser(bytes.NewBufferString("test log output")), nil
}

func (m *MockDockerClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	return m.err
}

func TestFormat(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"test", 10, "test      "},
		{"verylongstring", 5, "veryl"},
		{"", 3, "   "},
	}

	for _, tc := range tests {
		result := format(tc.input, tc.length)
		if result != tc.expected {
			t.Errorf("format(%q, %d) = %q; expected %q", tc.input, tc.length, result, tc.expected)
		}
	}
}

func TestListContainers(t *testing.T) {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	mockContainers := []types.Container{
		{
			ID:     "container1234567890",
			Names:  []string{"/test-container"},
			State:  "running",
			Status: "Up 2 hours",
		},
	}

	mockClient := &MockDockerClient{containers: mockContainers}
	ctx := context.Background()

	listContainers(mockClient, ctx)

	w.Close()
	os.Stdout = originalStdout
	var output bytes.Buffer
	io.Copy(&output, r)

	t.Logf("Full output:\n%s", output.String())

	outputStr := output.String()

	expectedID := "container1"
	if !strings.Contains(outputStr, expectedID) {
		t.Errorf("Expected output to contain container ID '%s', got: %s", expectedID, outputStr)
	}

	if !strings.Contains(outputStr, "test-container") {
		t.Errorf("Expected output to contain container name, got: %s", outputStr)
	}

	if len(outputStr) < 10 {
		t.Errorf("Output is too short, expected more content, got: %s", outputStr)
	}
}

func TestNoContainers(t *testing.T) {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	mockClient := &MockDockerClient{containers: []types.Container{}}
	ctx := context.Background()

	listContainers(mockClient, ctx)

	w.Close()
	os.Stdout = originalStdout
	var output bytes.Buffer
	io.Copy(&output, r)

	outputStr := output.String()
	if !strings.Contains(outputStr, "No running containers found") {
		t.Errorf("Expected 'No running containers found' message, got: %s", outputStr)
	}
}

func TestRunContainer(t *testing.T) {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	mockClient := &MockDockerClient{}
	ctx := context.Background()
	containerID := "test-container"

	runContainer(mockClient, ctx, containerID)

	w.Close()
	os.Stdout = originalStdout
	var output bytes.Buffer
	io.Copy(&output, r)

	outputStr := output.String()
	expectedMsg := "Container test-container started"
	if !strings.Contains(outputStr, expectedMsg) {
		t.Errorf("Expected '%s', got: '%s'", expectedMsg, outputStr)
	}
}
