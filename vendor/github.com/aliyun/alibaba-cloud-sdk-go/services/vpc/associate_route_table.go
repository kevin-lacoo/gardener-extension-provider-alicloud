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

// AssociateRouteTable invokes the vpc.AssociateRouteTable API synchronously
// api document: https://help.aliyun.com/api/vpc/associateroutetable.html
func (client *Client) AssociateRouteTable(request *AssociateRouteTableRequest) (response *AssociateRouteTableResponse, err error) {
	response = CreateAssociateRouteTableResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateRouteTableWithChan invokes the vpc.AssociateRouteTable API asynchronously
// api document: https://help.aliyun.com/api/vpc/associateroutetable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateRouteTableWithChan(request *AssociateRouteTableRequest) (<-chan *AssociateRouteTableResponse, <-chan error) {
	responseChan := make(chan *AssociateRouteTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateRouteTable(request)
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

// AssociateRouteTableWithCallback invokes the vpc.AssociateRouteTable API asynchronously
// api document: https://help.aliyun.com/api/vpc/associateroutetable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateRouteTableWithCallback(request *AssociateRouteTableRequest, callback func(response *AssociateRouteTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateRouteTableResponse
		var err error
		defer close(result)
		response, err = client.AssociateRouteTable(request)
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

// AssociateRouteTableRequest is the request struct for api AssociateRouteTable
type AssociateRouteTableRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	RouteTableId         string           `position:"Query" name:"RouteTableId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
}

// AssociateRouteTableResponse is the response struct for api AssociateRouteTable
type AssociateRouteTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateRouteTableRequest creates a request to invoke AssociateRouteTable API
func CreateAssociateRouteTableRequest() (request *AssociateRouteTableRequest) {
	request = &AssociateRouteTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociateRouteTable", "vpc", "openAPI")
	return
}

// CreateAssociateRouteTableResponse creates a response to parse from AssociateRouteTable response
func CreateAssociateRouteTableResponse() (response *AssociateRouteTableResponse) {
	response = &AssociateRouteTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
