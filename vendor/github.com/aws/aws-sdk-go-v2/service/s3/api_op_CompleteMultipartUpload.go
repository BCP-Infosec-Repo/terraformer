// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CompleteMultipartUploadInput struct {
	_ struct{} `type:"structure" payload:"MultipartUpload"`

	// Name of the bucket to which the multipart upload was initiated.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Object key for which the multipart upload was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// The container for the multipart upload request information.
	MultipartUpload *CompletedMultipartUpload `locationName:"CompleteMultipartUpload" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// For information about downloading objects from Requester Pays buckets, see
	// Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/http:/docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// ID for the initiated multipart upload.
	//
	// UploadId is a required field
	UploadId *string `location:"querystring" locationName:"uploadId" type:"string" required:"true"`
}

// String returns the string representation
func (s CompleteMultipartUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CompleteMultipartUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CompleteMultipartUploadInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *CompleteMultipartUploadInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CompleteMultipartUploadInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.MultipartUpload != nil {
		v := s.MultipartUpload

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "CompleteMultipartUpload", v, metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "uploadId", protocol.StringValue(v), metadata)
	}
	return nil
}

type CompleteMultipartUploadOutput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket that contains the newly created object.
	Bucket *string `type:"string"`

	// Entity tag that identifies the newly created object's data. Objects with
	// different object data will have different entity tags. The entity tag is
	// an opaque string. The entity tag may or may not be an MD5 digest of the object
	// data. If the entity tag is not an MD5 digest of the object data, it will
	// contain one or more nonhexadecimal characters and/or will consist of less
	// than 32 or more than 32 hexadecimal digits.
	ETag *string `type:"string"`

	// If the object expiration is configured, this will contain the expiration
	// date (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// The object key of the newly created object.
	Key *string `min:"1" type:"string"`

	// The URI that identifies the newly created object.
	Location *string `type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// customer master key (CMK) that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// If you specified server-side encryption either with an Amazon S3-managed
	// encryption key or an AWS KMS customer master key (CMK) in your initiate multipart
	// upload request, the response includes this header. It confirms the encryption
	// algorithm that Amazon S3 used to encrypt the object.
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// Version ID of the newly created object, in case the bucket has versioning
	// turned on.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s CompleteMultipartUploadOutput) String() string {
	return awsutil.Prettify(s)
}

func (s *CompleteMultipartUploadOutput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CompleteMultipartUploadOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Location", protocol.StringValue(v), metadata)
	}
	if s.Expiration != nil {
		v := *s.Expiration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-expiration", protocol.StringValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCompleteMultipartUpload = "CompleteMultipartUpload"

// CompleteMultipartUploadRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Completes a multipart upload by assembling previously uploaded parts.
//
// You first initiate the multipart upload and then upload all parts using the
// UploadPart operation. After successfully uploading all relevant parts of
// an upload, you call this operation to complete the upload. Upon receiving
// this request, Amazon S3 concatenates all the parts in ascending order by
// part number to create a new object. In the Complete Multipart Upload request,
// you must provide the parts list. You must ensure that the parts list is complete.
// This operation concatenates the parts that you provide in the list. For each
// part in the list, you must provide the part number and the ETag value, returned
// after that part was uploaded.
//
// Processing of a Complete Multipart Upload request could take several minutes
// to complete. After Amazon S3 begins processing the request, it sends an HTTP
// response header that specifies a 200 OK response. While processing is in
// progress, Amazon S3 periodically sends white space characters to keep the
// connection from timing out. Because a request could fail after the initial
// 200 OK response has been sent, it is important that you check the response
// body to determine whether the request succeeded.
//
// Note that if CompleteMultipartUpload fails, applications should be prepared
// to retry the failed requests. For more information, see Amazon S3 Error Best
// Practices (https://docs.aws.amazon.com/AmazonS3/latest/dev/ErrorBestPractices.html).
//
// For more information about multipart uploads, see Uploading Objects Using
// Multipart Upload (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html).
//
// For information about permissions required to use the multipart upload API,
// see Multipart Upload API and Permissions (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html).
//
// GetBucketLifecycle has the following special errors:
//
//    * Error code: EntityTooSmall Description: Your proposed upload is smaller
//    than the minimum allowed object size. Each part must be at least 5 MB
//    in size, except the last part. 400 Bad Request
//
//    * Error code: InvalidPart Description: One or more of the specified parts
//    could not be found. The part might not have been uploaded, or the specified
//    entity tag might not have matched the part's entity tag. 400 Bad Request
//
//    * Error code: InvalidPartOrder Description: The list of parts was not
//    in ascending order. The parts list must be specified in order by part
//    number. 400 Bad Request
//
//    * Error code: NoSuchUpload Description: The specified multipart upload
//    does not exist. The upload ID might be invalid, or the multipart upload
//    might have been aborted or completed. 404 Not Found
//
// The following operations are related to DeleteBucketMetricsConfiguration:
//
//    * CreateMultipartUpload
//
//    * UploadPart
//
//    * AbortMultipartUpload
//
//    * ListParts
//
//    * ListMultipartUploads
//
//    // Example sending a request using CompleteMultipartUploadRequest.
//    req := client.CompleteMultipartUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CompleteMultipartUpload
func (c *Client) CompleteMultipartUploadRequest(input *CompleteMultipartUploadInput) CompleteMultipartUploadRequest {
	op := &aws.Operation{
		Name:       opCompleteMultipartUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &CompleteMultipartUploadInput{}
	}

	req := c.newRequest(op, input, &CompleteMultipartUploadOutput{})
	return CompleteMultipartUploadRequest{Request: req, Input: input, Copy: c.CompleteMultipartUploadRequest}
}

// CompleteMultipartUploadRequest is the request type for the
// CompleteMultipartUpload API operation.
type CompleteMultipartUploadRequest struct {
	*aws.Request
	Input *CompleteMultipartUploadInput
	Copy  func(*CompleteMultipartUploadInput) CompleteMultipartUploadRequest
}

// Send marshals and sends the CompleteMultipartUpload API request.
func (r CompleteMultipartUploadRequest) Send(ctx context.Context) (*CompleteMultipartUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CompleteMultipartUploadResponse{
		CompleteMultipartUploadOutput: r.Request.Data.(*CompleteMultipartUploadOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CompleteMultipartUploadResponse is the response type for the
// CompleteMultipartUpload API operation.
type CompleteMultipartUploadResponse struct {
	*CompleteMultipartUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CompleteMultipartUpload request.
func (r *CompleteMultipartUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
