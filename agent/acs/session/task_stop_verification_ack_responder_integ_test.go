//go:build integration
// +build integration

package session_test

import (
	"context"
	"testing"

	"github.com/aws/amazon-ecs-agent/agent/acs/session"
	"github.com/aws/amazon-ecs-agent/agent/data"
	"github.com/aws/amazon-ecs-agent/agent/engine"
	"github.com/aws/amazon-ecs-agent/ecs-agent/acs/model/ecsacs"
	acssession "github.com/aws/amazon-ecs-agent/ecs-agent/acs/session"
	"github.com/aws/amazon-ecs-agent/ecs-agent/acs/session/testconst"
	apicontainerstatus "github.com/aws/amazon-ecs-agent/ecs-agent/api/container/status"
	"github.com/aws/amazon-ecs-agent/ecs-agent/metrics"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/require"
)

// Tests that a task, its containers, and its resources are all stopped when a task stop verification ACK message is received.
func TestTaskStopVerificationACKResponder(t *testing.T) {
	taskEngine, done, _ := engine.Setup(engine.DefaultTestConfigIntegTest(), nil, t)
	defer done()

	task := engine.CreateTestTask("test_task")
	go taskEngine.AddTask(task)

	engine.VerifyContainerManifestPulledStateChange(t, taskEngine)
	engine.VerifyTaskManifestPulledStateChange(t, taskEngine)
	engine.VerifyContainerRunningStateChange(t, taskEngine)
	engine.VerifyTaskRunningStateChange(t, taskEngine)

	manifestMessageIDAccessor := session.NewManifestMessageIDAccessor()
	manifestMessageIDAccessor.SetMessageID("manifest-message-id")

	taskStopper := session.NewTaskStopper(taskEngine, data.NewNoopClient())
	responder := acssession.NewTaskStopVerificationACKResponder(taskStopper, manifestMessageIDAccessor, metrics.NewNopEntryFactory())

	handler := responder.HandlerFunc().(func(*ecsacs.TaskStopVerificationAck))
	handler(&ecsacs.TaskStopVerificationAck{
		GeneratedAt: aws.Int64(testconst.DummyInt),
		MessageId:   aws.String(manifestMessageIDAccessor.GetMessageID()),
		StopTasks: []*ecsacs.TaskIdentifier{
			{
				TaskArn: aws.String(task.Arn),
			},
		},
	})

	engine.VerifyContainerStoppedStateChange(t, taskEngine)
	engine.VerifyTaskStoppedStateChange(t, taskEngine)

	dockerClient := taskEngine.(*DockerTaskEngine).client
	status, container := dockerClient.DescribeContainer(context.Background(), task.Containers[0].RuntimeID)
	require.Equal(t, apicontainerstatus.ContainerStopped, status)
	require.NoError(t, container.Error)
}
