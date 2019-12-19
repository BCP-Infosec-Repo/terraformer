// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListOrganizationalUnitsForParentInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Use this to limit the number of results you want included per
	// page in the response. If you do not include this parameter, it defaults to
	// a value that is specific to the operation. If additional items exist beyond
	// the maximum you specify, the NextToken response element is present and has
	// a value (is not null). Include that value as the NextToken request parameter
	// in the next call to the operation to get the next part of the results. Note
	// that Organizations might return fewer results than the maximum even when
	// there are more results available. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// Use this parameter if you receive a NextToken response in a previous request
	// that indicates that there is more output available. Set it to the value of
	// the previous call's NextToken response to indicate where the output should
	// continue from.
	NextToken *string `type:"string"`

	// The unique identifier (ID) of the root or OU whose child OUs you want to
	// list.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//    * Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//    letters or digits.
	//
	//    * Organizational unit (OU) - A string that begins with "ou-" followed
	//    by from 4 to 32 lowercase letters or digits (the ID of the root that the
	//    OU is in). This string is followed by a second "-" dash and from 8 to
	//    32 additional lowercase letters or digits.
	//
	// ParentId is a required field
	ParentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListOrganizationalUnitsForParentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOrganizationalUnitsForParentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOrganizationalUnitsForParentInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ParentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListOrganizationalUnitsForParentOutput struct {
	_ struct{} `type:"structure"`

	// If present, this value indicates that there is more output available than
	// is included in the current response. Use this value in the NextToken request
	// parameter in a subsequent call to the operation to get the next part of the
	// output. You should repeat this until the NextToken response element comes
	// back as null.
	NextToken *string `type:"string"`

	// A list of the OUs in the specified root or parent OU.
	OrganizationalUnits []OrganizationalUnit `type:"list"`
}

// String returns the string representation
func (s ListOrganizationalUnitsForParentOutput) String() string {
	return awsutil.Prettify(s)
}

const opListOrganizationalUnitsForParent = "ListOrganizationalUnitsForParent"

// ListOrganizationalUnitsForParentRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists the organizational units (OUs) in a parent organizational unit or root.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using ListOrganizationalUnitsForParentRequest.
//    req := client.ListOrganizationalUnitsForParentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListOrganizationalUnitsForParent
func (c *Client) ListOrganizationalUnitsForParentRequest(input *ListOrganizationalUnitsForParentInput) ListOrganizationalUnitsForParentRequest {
	op := &aws.Operation{
		Name:       opListOrganizationalUnitsForParent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListOrganizationalUnitsForParentInput{}
	}

	req := c.newRequest(op, input, &ListOrganizationalUnitsForParentOutput{})
	return ListOrganizationalUnitsForParentRequest{Request: req, Input: input, Copy: c.ListOrganizationalUnitsForParentRequest}
}

// ListOrganizationalUnitsForParentRequest is the request type for the
// ListOrganizationalUnitsForParent API operation.
type ListOrganizationalUnitsForParentRequest struct {
	*aws.Request
	Input *ListOrganizationalUnitsForParentInput
	Copy  func(*ListOrganizationalUnitsForParentInput) ListOrganizationalUnitsForParentRequest
}

// Send marshals and sends the ListOrganizationalUnitsForParent API request.
func (r ListOrganizationalUnitsForParentRequest) Send(ctx context.Context) (*ListOrganizationalUnitsForParentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOrganizationalUnitsForParentResponse{
		ListOrganizationalUnitsForParentOutput: r.Request.Data.(*ListOrganizationalUnitsForParentOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListOrganizationalUnitsForParentRequestPaginator returns a paginator for ListOrganizationalUnitsForParent.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListOrganizationalUnitsForParentRequest(input)
//   p := organizations.NewListOrganizationalUnitsForParentRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListOrganizationalUnitsForParentPaginator(req ListOrganizationalUnitsForParentRequest) ListOrganizationalUnitsForParentPaginator {
	return ListOrganizationalUnitsForParentPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListOrganizationalUnitsForParentInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListOrganizationalUnitsForParentPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListOrganizationalUnitsForParentPaginator struct {
	aws.Pager
}

func (p *ListOrganizationalUnitsForParentPaginator) CurrentPage() *ListOrganizationalUnitsForParentOutput {
	return p.Pager.CurrentPage().(*ListOrganizationalUnitsForParentOutput)
}

// ListOrganizationalUnitsForParentResponse is the response type for the
// ListOrganizationalUnitsForParent API operation.
type ListOrganizationalUnitsForParentResponse struct {
	*ListOrganizationalUnitsForParentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOrganizationalUnitsForParent request.
func (r *ListOrganizationalUnitsForParentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
