// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an account setting for all users on an account for whom no individual
// account setting has been specified. Account settings are set on a per-Region
// basis.
func (c *Client) PutAccountSettingDefault(ctx context.Context, params *PutAccountSettingDefaultInput, optFns ...func(*Options)) (*PutAccountSettingDefaultOutput, error) {
	if params == nil {
		params = &PutAccountSettingDefaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountSettingDefault", params, optFns, c.addOperationPutAccountSettingDefaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountSettingDefaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountSettingDefaultInput struct {

	// The resource name for which to modify the account setting. If you specify
	// serviceLongArnFormat , the ARN for your Amazon ECS services is affected. If you
	// specify taskLongArnFormat , the ARN and resource ID for your Amazon ECS tasks is
	// affected. If you specify containerInstanceLongArnFormat , the ARN and resource
	// ID for your Amazon ECS container instances is affected. If you specify
	// awsvpcTrunking , the ENI limit for your Amazon ECS container instances is
	// affected. If you specify containerInsights , the default setting for Amazon Web
	// Services CloudWatch Container Insights for your clusters is affected. If you
	// specify tagResourceAuthorization , the opt-in option for tagging resources on
	// creation is affected. For information about the opt-in timeline, see Tagging
	// authorization timeline (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#tag-resources)
	// in the Amazon ECS Developer Guide. If you specify
	// fargateTaskRetirementWaitPeriod , the default wait time to retire a Fargate task
	// due to required maintenance is affected. When you specify fargateFIPSMode for
	// the name and enabled for the value , Fargate uses FIPS-140 compliant
	// cryptographic algorithms on your tasks. For more information about FIPS-140
	// compliance with Fargate, see Amazon Web Services Fargate Federal Information
	// Processing Standard (FIPS) 140-2 compliance (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-fips-compliance.html)
	// in the Amazon Elastic Container Service Developer Guide. When Amazon Web
	// Services determines that a security or infrastructure update is needed for an
	// Amazon ECS task hosted on Fargate, the tasks need to be stopped and new tasks
	// launched to replace them. Use fargateTaskRetirementWaitPeriod to set the wait
	// time to retire a Fargate task to the default. For information about the Fargate
	// tasks maintenance, see Amazon Web Services Fargate task maintenance (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-maintenance.html)
	// in the Amazon ECS Developer Guide.
	//
	// This member is required.
	Name types.SettingName

	// The account setting value for the specified principal ARN. Accepted values are
	// enabled , disabled , on , and off . When you specify
	// fargateTaskRetirementWaitPeriod for the name , the following are the valid
	// values:
	//   - 0 - Amazon Web Services sends the notification, and immediately retires the
	//   affected tasks.
	//   - 7 - Amazon Web Services sends the notification, and waits 7 calendar days to
	//   retire the tasks.
	//   - 14 - Amazon Web Services sends the notification, and waits 14 calendar days
	//   to retire the tasks.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type PutAccountSettingDefaultOutput struct {

	// The current setting for a resource.
	Setting *types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountSettingDefaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountSettingDefault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountSettingDefault{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addPutAccountSettingDefaultResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutAccountSettingDefaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSettingDefault(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutAccountSettingDefault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "PutAccountSettingDefault",
	}
}

type opPutAccountSettingDefaultResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutAccountSettingDefaultResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutAccountSettingDefaultResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ecs"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ecs"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ecs")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addPutAccountSettingDefaultResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutAccountSettingDefaultResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}