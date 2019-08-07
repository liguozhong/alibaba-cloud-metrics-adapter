package scdn

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

// DescribeScdnRefreshQuota invokes the scdn.DescribeScdnRefreshQuota API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdnrefreshquota.html
func (client *Client) DescribeScdnRefreshQuota(request *DescribeScdnRefreshQuotaRequest) (response *DescribeScdnRefreshQuotaResponse, err error) {
	response = CreateDescribeScdnRefreshQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnRefreshQuotaWithChan invokes the scdn.DescribeScdnRefreshQuota API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnrefreshquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnRefreshQuotaWithChan(request *DescribeScdnRefreshQuotaRequest) (<-chan *DescribeScdnRefreshQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnRefreshQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnRefreshQuota(request)
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

// DescribeScdnRefreshQuotaWithCallback invokes the scdn.DescribeScdnRefreshQuota API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnrefreshquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnRefreshQuotaWithCallback(request *DescribeScdnRefreshQuotaRequest, callback func(response *DescribeScdnRefreshQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnRefreshQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnRefreshQuota(request)
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

// DescribeScdnRefreshQuotaRequest is the request struct for api DescribeScdnRefreshQuota
type DescribeScdnRefreshQuotaRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnRefreshQuotaResponse is the response struct for api DescribeScdnRefreshQuota
type DescribeScdnRefreshQuotaResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	UrlQuota      string `json:"UrlQuota" xml:"UrlQuota"`
	DirQuota      string `json:"DirQuota" xml:"DirQuota"`
	UrlRemain     string `json:"UrlRemain" xml:"UrlRemain"`
	DirRemain     string `json:"DirRemain" xml:"DirRemain"`
	PreloadQuota  string `json:"PreloadQuota" xml:"PreloadQuota"`
	BlockQuota    string `json:"BlockQuota" xml:"BlockQuota"`
	PreloadRemain string `json:"PreloadRemain" xml:"PreloadRemain"`
	BlockRemain   string `json:"blockRemain" xml:"blockRemain"`
}

// CreateDescribeScdnRefreshQuotaRequest creates a request to invoke DescribeScdnRefreshQuota API
func CreateDescribeScdnRefreshQuotaRequest() (request *DescribeScdnRefreshQuotaRequest) {
	request = &DescribeScdnRefreshQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnRefreshQuota", "scdn", "openAPI")
	return
}

// CreateDescribeScdnRefreshQuotaResponse creates a response to parse from DescribeScdnRefreshQuota response
func CreateDescribeScdnRefreshQuotaResponse() (response *DescribeScdnRefreshQuotaResponse) {
	response = &DescribeScdnRefreshQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}