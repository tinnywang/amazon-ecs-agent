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

// Modifies the parameters for a capacity provider.
func (c *Client) UpdateCapacityProvider(ctx context.Context, params *UpdateCapacityProviderInput, optFns ...func(*Options)) (*UpdateCapacityProviderOutput, error) {
	if params == nil { params = &UpdateCapacityProviderInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateCapacityProvider", params, optFns, c.addOperationUpdateCapacityProviderMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateCapacityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCapacityProviderInput struct {
	
	// An object that represent the parameters to update for the Auto Scaling group
	// capacity provider.
	//
	// This member is required.
	AutoScalingGroupProvider *types.AutoScalingGroupProviderUpdate
	
	// The name of the capacity provider to update.
	//
	// This member is required.
	Name *string
	
	noSmithyDocumentSerde
}

type UpdateCapacityProviderOutput struct {
	
	// Details about the capacity provider.
	CapacityProvider *types.CapacityProvider
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCapacityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
	    return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCapacityProvider{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCapacityProvider{}, middleware.After)
	if err != nil { return err }
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCapacityProvider"); err != nil {
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
	if err = addOpUpdateCapacityProviderValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCapacityProvider(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCapacityProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "UpdateCapacityProvider",
	}
}