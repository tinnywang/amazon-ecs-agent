// Code generated by smithy-go-codegen DO NOT EDIT.


package ecs

import (
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

// Deregisters the specified task definition by family and revision. Upon
// deregistration, the task definition is marked as INACTIVE . Existing tasks and
// services that reference an INACTIVE task definition continue to run without
// disruption. Existing services that reference an INACTIVE task definition can
// still scale up or down by modifying the service's desired count. If you want to
// delete a task definition revision, you must first deregister the task definition
// revision.
//
// You can't use an INACTIVE task definition to run new tasks or create new
// services, and you can't update an existing service to reference an INACTIVE
// task definition. However, there may be up to a 10-minute window following
// deregistration where these restrictions have not yet taken effect.
//
// At this time, INACTIVE task definitions remain discoverable in your account
// indefinitely. However, this behavior is subject to change in the future. We
// don't recommend that you rely on INACTIVE task definitions persisting beyond
// the lifecycle of any associated tasks and services.
//
// You must deregister a task definition revision before you delete it. For more
// information, see [DeleteTaskDefinitions].
//
// [DeleteTaskDefinitions]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeleteTaskDefinitions.html
func (c *Client) DeregisterTaskDefinition(ctx context.Context, params *DeregisterTaskDefinitionInput, optFns ...func(*Options)) (*DeregisterTaskDefinitionOutput, error) {
	if params == nil { params = &DeregisterTaskDefinitionInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DeregisterTaskDefinition", params, optFns, c.addOperationDeregisterTaskDefinitionMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DeregisterTaskDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTaskDefinitionInput struct {
	
	// The family and revision ( family:revision ) or full Amazon Resource Name (ARN)
	// of the task definition to deregister. You must specify a revision .
	//
	// This member is required.
	TaskDefinition *string
	
	noSmithyDocumentSerde
}

type DeregisterTaskDefinitionOutput struct {
	
	// The full description of the deregistered task.
	TaskDefinition *types.TaskDefinition
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterTaskDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
	    return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterTaskDefinition{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterTaskDefinition{}, middleware.After)
	if err != nil { return err }
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterTaskDefinition"); err != nil {
	    return fmt.Errorf("add protocol finalizers: %v", err)
	}
	
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
	return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
	return err
	}
	if err = addClientRequestID(stack); err != nil {
	return err
	}
	if err = addComputeContentLength(stack); err != nil {
	return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
	return err
	}
	if err = addRetry(stack, options); err != nil {
	return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
	return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
	return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
	return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
	return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
	return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
	return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
	return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
	return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
	return err
	}
	if err = addOpDeregisterTaskDefinitionValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterTaskDefinition(options.Region, ), middleware.Before); err != nil {
	return err
	}
	if err = addRecursionDetection(stack); err != nil {
	return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
	return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
	return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
	return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
	return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
	return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
	return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
	return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
	return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeregisterTaskDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "DeregisterTaskDefinition",
	}
}