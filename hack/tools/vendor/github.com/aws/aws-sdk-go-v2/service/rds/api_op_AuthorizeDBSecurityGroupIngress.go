// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables ingress to a DBSecurityGroup using one of two forms of authorization.
// First, EC2 or VPC security groups can be added to the DBSecurityGroup if the
// application using the database is running on EC2 or VPC instances. Second, IP
// ranges are available if the application accessing your database is running on
// the internet. Required parameters for this API are one of CIDR range,
// EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId for non-VPC). You can't authorize
// ingress from an EC2 security group in one Amazon Web Services Region to an
// Amazon RDS DB instance in another. You can't authorize ingress from a VPC
// security group in one VPC to an Amazon RDS DB instance in another. For an
// overview of CIDR ranges, go to the Wikipedia Tutorial (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
// . EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see Migrate from EC2-Classic to a VPC (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html)
// in the Amazon EC2 User Guide, the blog EC2-Classic Networking is Retiring –
// Here’s How to Prepare (http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/)
// , and Moving a DB instance not in a VPC into a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html)
// in the Amazon RDS User Guide.
func (c *Client) AuthorizeDBSecurityGroupIngress(ctx context.Context, params *AuthorizeDBSecurityGroupIngressInput, optFns ...func(*Options)) (*AuthorizeDBSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &AuthorizeDBSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeDBSecurityGroupIngress", params, optFns, c.addOperationAuthorizeDBSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeDBSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeDBSecurityGroupIngressInput struct {

	// The name of the DB security group to add authorization to.
	//
	// This member is required.
	DBSecurityGroupName *string

	// The IP range to authorize.
	CIDRIP *string

	// Id of the EC2 security group to authorize. For VPC DB security groups,
	// EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId and
	// either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupId *string

	// Name of the EC2 security group to authorize. For VPC DB security groups,
	// EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId and
	// either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupName *string

	// Amazon Web Services account number of the owner of the EC2 security group
	// specified in the EC2SecurityGroupName parameter. The Amazon Web Services access
	// key ID isn't an acceptable value. For VPC DB security groups, EC2SecurityGroupId
	// must be provided. Otherwise, EC2SecurityGroupOwnerId and either
	// EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupOwnerId *string

	noSmithyDocumentSerde
}

type AuthorizeDBSecurityGroupIngressOutput struct {

	// Contains the details for an Amazon RDS DB security group. This data type is
	// used as a response element in the DescribeDBSecurityGroups action.
	DBSecurityGroup *types.DBSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAuthorizeDBSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAuthorizeDBSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAuthorizeDBSecurityGroupIngress{}, middleware.After)
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
	if err = addAuthorizeDBSecurityGroupIngressResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAuthorizeDBSecurityGroupIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeDBSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeDBSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "AuthorizeDBSecurityGroupIngress",
	}
}

type opAuthorizeDBSecurityGroupIngressResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opAuthorizeDBSecurityGroupIngressResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opAuthorizeDBSecurityGroupIngressResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "rds"
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
				signingName = "rds"
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
				v4aScheme.SigningName = aws.String("rds")
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

func addAuthorizeDBSecurityGroupIngressResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opAuthorizeDBSecurityGroupIngressResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}