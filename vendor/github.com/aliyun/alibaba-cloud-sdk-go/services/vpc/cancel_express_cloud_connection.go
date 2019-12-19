package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CancelExpressCloudConnection invokes the vpc.CancelExpressCloudConnection API synchronously
// api document: https://help.aliyun.com/api/vpc/cancelexpresscloudconnection.html
func (client *Client) CancelExpressCloudConnection(request *CancelExpressCloudConnectionRequest) (response *CancelExpressCloudConnectionResponse, err error) {
	response = CreateCancelExpressCloudConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// CancelExpressCloudConnectionWithChan invokes the vpc.CancelExpressCloudConnection API asynchronously
// api document: https://help.aliyun.com/api/vpc/cancelexpresscloudconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelExpressCloudConnectionWithChan(request *CancelExpressCloudConnectionRequest) (<-chan *CancelExpressCloudConnectionResponse, <-chan error) {
	responseChan := make(chan *CancelExpressCloudConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelExpressCloudConnection(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CancelExpressCloudConnectionWithCallback invokes the vpc.CancelExpressCloudConnection API asynchronously
// api document: https://help.aliyun.com/api/vpc/cancelexpresscloudconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelExpressCloudConnectionWithCallback(request *CancelExpressCloudConnectionRequest, callback func(response *CancelExpressCloudConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelExpressCloudConnectionResponse
		var err error
		defer close(result)
		response, err = client.CancelExpressCloudConnection(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CancelExpressCloudConnectionRequest is the request struct for api CancelExpressCloudConnection
type CancelExpressCloudConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EccId                string           `position:"Query" name:"EccId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CancelExpressCloudConnectionResponse is the response struct for api CancelExpressCloudConnection
type CancelExpressCloudConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelExpressCloudConnectionRequest creates a request to invoke CancelExpressCloudConnection API
func CreateCancelExpressCloudConnectionRequest() (request *CancelExpressCloudConnectionRequest) {
	request = &CancelExpressCloudConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CancelExpressCloudConnection", "vpc", "openAPI")
	return
}

// CreateCancelExpressCloudConnectionResponse creates a response to parse from CancelExpressCloudConnection response
func CreateCancelExpressCloudConnectionResponse() (response *CancelExpressCloudConnectionResponse) {
	response = &CancelExpressCloudConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
