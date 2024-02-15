// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
)

// This operation establishes an HTTP/2 connection between the consumer you
// specify in the ConsumerARN parameter and the shard you specify in the ShardId
// parameter. After the connection is successfully established, Kinesis Data
// Streams pushes records from the shard to the consumer over this connection.
// Before you call this operation, call RegisterStreamConsumer to register the
// consumer with Kinesis Data Streams. When the SubscribeToShard call succeeds,
// your consumer starts receiving events of type SubscribeToShardEvent over the
// HTTP/2 connection for up to 5 minutes, after which time you need to call
// SubscribeToShard again to renew the subscription if you want to continue to
// receive records. You can make one call to SubscribeToShard per second per
// registered consumer per shard. For example, if you have a 4000 shard stream and
// two registered stream consumers, you can make one SubscribeToShard request per
// second for each combination of shard and registered consumer, allowing you to
// subscribe both consumers to all 4000 shards in one second. If you call
// SubscribeToShard again with the same ConsumerARN and ShardId within 5 seconds
// of a successful call, you'll get a ResourceInUseException . If you call
// SubscribeToShard 5 seconds or more after a successful call, the second call
// takes over the subscription and the previous connection expires or fails with a
// ResourceInUseException . For an example of how to use this operations, see
// Enhanced Fan-Out Using the Kinesis Data Streams API .
func (c *Client) SubscribeToShard(ctx context.Context, params *SubscribeToShardInput, optFns ...func(*Options)) (*SubscribeToShardOutput, error) {
	if params == nil {
		params = &SubscribeToShardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubscribeToShard", params, optFns, c.addOperationSubscribeToShardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubscribeToShardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubscribeToShardInput struct {

	// For this parameter, use the value you obtained when you called
	// RegisterStreamConsumer .
	//
	// This member is required.
	ConsumerARN *string

	// The ID of the shard you want to subscribe to. To see a list of all the shards
	// for a given stream, use ListShards .
	//
	// This member is required.
	ShardId *string

	// The starting position in the data stream from which to start streaming.
	//
	// This member is required.
	StartingPosition *types.StartingPosition

	noSmithyDocumentSerde
}

type SubscribeToShardOutput struct {
	eventStream *SubscribeToShardEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *SubscribeToShardOutput) GetStream() *SubscribeToShardEventStream {
	return o.eventStream
}

func (c *Client) addOperationSubscribeToShardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSubscribeToShard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubscribeToShard{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamSubscribeToShardMiddleware(stack, options); err != nil {
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
	if err = addSubscribeToShardResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSubscribeToShardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubscribeToShard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubscribeToShard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "SubscribeToShard",
	}
}

// SubscribeToShardEventStream provides the event stream handling for the SubscribeToShard operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewSubscribeToShardEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type SubscribeToShardEventStream struct {
	// SubscribeToShardEventStreamReader is the EventStream reader for the
	// SubscribeToShardEventStream events. This value is automatically set by the SDK
	// when the API call is made Use this member when unit testing your code with the
	// SDK to mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader SubscribeToShardEventStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewSubscribeToShardEventStream initializes an SubscribeToShardEventStream.
// This function should only be used for testing and mocking the SubscribeToShardEventStream
// stream within your application.
//
// The Reader member must be set before reading events from the stream.
func NewSubscribeToShardEventStream(optFns ...func(*SubscribeToShardEventStream)) *SubscribeToShardEventStream {
	es := &SubscribeToShardEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Events returns a channel to read events from.
func (es *SubscribeToShardEventStream) Events() <-chan types.SubscribeToShardEventStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *SubscribeToShardEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *SubscribeToShardEventStream) safeClose() {
	close(es.done)

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *SubscribeToShardEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *SubscribeToShardEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}

type opSubscribeToShardResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSubscribeToShardResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSubscribeToShardResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*SubscribeToShardInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.ConsumerARN = input.ConsumerARN

	params.OperationType = ptr.String("data")

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
			signingName := "kinesis"
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
				signingName = "kinesis"
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
				v4aScheme.SigningName = aws.String("kinesis")
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

func addSubscribeToShardResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSubscribeToShardResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}