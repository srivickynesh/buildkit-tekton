// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or edits tags on a customer managed key
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk).
// Tagging or untagging a KMS key can allow or deny permission to the KMS key. For
// details, see ABAC for KMS
// (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the Key
// Management Service Developer Guide. Each tag consists of a tag key and a tag
// value, both of which are case-sensitive strings. The tag value can be an empty
// (null) string. To add a tag, specify a new tag key and a tag value. To edit a
// tag, specify an existing tag key and a new tag value. You can use this operation
// to tag a customer managed key
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk),
// but you cannot tag an Amazon Web Services managed key
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk),
// an Amazon Web Services owned key
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk),
// a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#keystore-concept),
// or an alias
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#alias-concept).
// You can also add tags to a KMS key while creating it (CreateKey) or replicating
// it (ReplicateKey). For information about using tags in KMS, see Tagging keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html). For
// general information about tags, including the format and syntax, see Tagging
// Amazon Web Services resources
// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon
// Web Services General Reference. The KMS key that you use for this operation must
// be in a compatible key state. For details, see Key states of KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions: kms:TagResource
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations
//
// * CreateKey
//
// * ListResourceTags
//
// *
// ReplicateKey
//
// * UntagResource
func (c *Client) TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) {
	if params == nil {
		params = &TagResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagResource", params, optFns, c.addOperationTagResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagResourceInput struct {

	// Identifies a customer managed key in the account and Region. Specify the key ID
	// or key ARN of the KMS key. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// One or more tags. Each tag consists of a tag key and a tag value. The tag value
	// can be an empty (null) string. You cannot have more than one tag on a KMS key
	// with the same tag key. If you specify an existing tag key with a different tag
	// value, KMS replaces the current tag value with the specified one.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagResource{}, middleware.After)
	if err != nil {
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpTagResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagResource(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opTagResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "TagResource",
	}
}
