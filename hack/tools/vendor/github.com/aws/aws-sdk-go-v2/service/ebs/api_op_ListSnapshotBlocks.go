// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the blocks in an Amazon Elastic Block Store snapshot.
// You should always retry requests that receive server ( 5xx ) error responses,
// and ThrottlingException and RequestThrottledException client error responses.
// For more information see Error retries (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ListSnapshotBlocks(ctx context.Context, params *ListSnapshotBlocksInput, optFns ...func(*Options)) (*ListSnapshotBlocksOutput, error) {
	if params == nil {
		params = &ListSnapshotBlocksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSnapshotBlocks", params, optFns, c.addOperationListSnapshotBlocksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSnapshotBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSnapshotBlocksInput struct {

	// The ID of the snapshot from which to get block indexes and block tokens.
	//
	// This member is required.
	SnapshotId *string

	// The maximum number of blocks to be returned by the request. Even if additional
	// blocks can be retrieved from the snapshot, the request can return less blocks
	// than MaxResults or an empty array of blocks. To retrieve the next set of blocks
	// from the snapshot, make another request with the returned NextToken value. The
	// value of NextToken is null when there are no more blocks to return.
	MaxResults *int32

	// The token to request the next page of results. If you specify NextToken, then
	// StartingBlockIndex is ignored.
	NextToken *string

	// The block index from which the list should start. The list in the response will
	// start from this block index or the next valid block index in the snapshot. If
	// you specify NextToken, then StartingBlockIndex is ignored.
	StartingBlockIndex *int32

	noSmithyDocumentSerde
}

type ListSnapshotBlocksOutput struct {

	// The size of the blocks in the snapshot, in bytes.
	BlockSize *int32

	// An array of objects containing information about the blocks.
	Blocks []types.Block

	// The time when the BlockToken expires.
	ExpiryTime *time.Time

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The size of the volume in GB.
	VolumeSize *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSnapshotBlocksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSnapshotBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSnapshotBlocks{}, middleware.After)
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
	if err = addListSnapshotBlocksResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListSnapshotBlocksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSnapshotBlocks(options.Region), middleware.Before); err != nil {
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

// ListSnapshotBlocksAPIClient is a client that implements the ListSnapshotBlocks
// operation.
type ListSnapshotBlocksAPIClient interface {
	ListSnapshotBlocks(context.Context, *ListSnapshotBlocksInput, ...func(*Options)) (*ListSnapshotBlocksOutput, error)
}

var _ ListSnapshotBlocksAPIClient = (*Client)(nil)

// ListSnapshotBlocksPaginatorOptions is the paginator options for
// ListSnapshotBlocks
type ListSnapshotBlocksPaginatorOptions struct {
	// The maximum number of blocks to be returned by the request. Even if additional
	// blocks can be retrieved from the snapshot, the request can return less blocks
	// than MaxResults or an empty array of blocks. To retrieve the next set of blocks
	// from the snapshot, make another request with the returned NextToken value. The
	// value of NextToken is null when there are no more blocks to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSnapshotBlocksPaginator is a paginator for ListSnapshotBlocks
type ListSnapshotBlocksPaginator struct {
	options   ListSnapshotBlocksPaginatorOptions
	client    ListSnapshotBlocksAPIClient
	params    *ListSnapshotBlocksInput
	nextToken *string
	firstPage bool
}

// NewListSnapshotBlocksPaginator returns a new ListSnapshotBlocksPaginator
func NewListSnapshotBlocksPaginator(client ListSnapshotBlocksAPIClient, params *ListSnapshotBlocksInput, optFns ...func(*ListSnapshotBlocksPaginatorOptions)) *ListSnapshotBlocksPaginator {
	if params == nil {
		params = &ListSnapshotBlocksInput{}
	}

	options := ListSnapshotBlocksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSnapshotBlocksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSnapshotBlocksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSnapshotBlocks page.
func (p *ListSnapshotBlocksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSnapshotBlocksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListSnapshotBlocks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListSnapshotBlocks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "ListSnapshotBlocks",
	}
}

type opListSnapshotBlocksResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListSnapshotBlocksResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListSnapshotBlocksResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ebs"
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
				signingName = "ebs"
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
				v4aScheme.SigningName = aws.String("ebs")
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

func addListSnapshotBlocksResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListSnapshotBlocksResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}