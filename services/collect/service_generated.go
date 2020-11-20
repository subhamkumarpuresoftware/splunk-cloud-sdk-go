/*
 * Copyright © 2020 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * Collect Service
 *
 * With the Collect service in Splunk Cloud Services, you can manage how data collection jobs ingest event and metric data.
 *
 * API version: v1beta1.8 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package collect

import (
	"net/http"

	"github.com/splunk/go-dependencies/services"
	"github.com/splunk/go-dependencies/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new collect service client from the given Config
func NewService(iClient services.IClient) *Service {
	return &Service{Client: iClient}
}

/*
	CreateExecution - Creates an execution for a scheduled job based on the job ID.
	Parameters:
		jobId: The job ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateExecution(jobId string, resp ...*http.Response) (*SingleExecutionResponse, error) {
	pp := struct {
		JobId string
	}{
		JobId: jobId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs/{{.JobId}}/executions`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SingleExecutionResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateJob - Creates a job.
	This API returns &#x60;403&#x60; if the number of collect workers is over a certain limit.
	Parameters:
		job: The API request schema for the job.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateJob(job Job, resp ...*http.Response) (*SingleJobResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: job})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SingleJobResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteJob - Removes a job based on the job ID.
	Parameters:
		jobId: The job ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteJob(jobId string, resp ...*http.Response) error {
	pp := struct {
		JobId string
	}{
		JobId: jobId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs/{{.JobId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeleteJobs - Removes all jobs on a tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteJobs(resp ...*http.Response) (*DeleteJobsResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb DeleteJobsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetExecution - Returns the execution details based on the execution ID and job ID.
	Parameters:
		jobId: The job ID.
		executionUid: The execution UID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetExecution(jobId string, executionUid string, resp ...*http.Response) (*SingleExecutionResponse, error) {
	pp := struct {
		JobId        string
		ExecutionUid string
	}{
		JobId:        jobId,
		ExecutionUid: executionUid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs/{{.JobId}}/executions/{{.ExecutionUid}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SingleExecutionResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetJob - Returns a job based on the job ID.
	Parameters:
		jobId: The job ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetJob(jobId string, resp ...*http.Response) (*SingleJobResponse, error) {
	pp := struct {
		JobId string
	}{
		JobId: jobId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs/{{.JobId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SingleJobResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListJobs - Returns a list of all jobs that belong to a tenant.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListJobs(query *ListJobsQueryParams, resp ...*http.Response) (*ListJobsResponse, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/collect/v1beta1/jobs`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ListJobsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PatchExecution - Modifies an execution based on the job ID.
	Parameters:
		jobId: The job ID.
		executionUid: The execution UID.
		executionPatch: The API request schema for patching an execution.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PatchExecution(jobId string, executionUid string, executionPatch ExecutionPatch, resp ...*http.Response) error {
	pp := struct {
		JobId        string
		ExecutionUid string
	}{
		JobId:        jobId,
		ExecutionUid: executionUid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs/{{.JobId}}/executions/{{.ExecutionUid}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: executionPatch})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	PatchJob - Modifies a job based on the job ID.
	This API returns &#x60;403&#x60; if the number of collect workers is over a certain limit.
	Parameters:
		jobId: The job ID.
		jobPatch: The API request schema for patching a job.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PatchJob(jobId string, jobPatch JobPatch, resp ...*http.Response) (*SingleJobResponse, error) {
	pp := struct {
		JobId string
	}{
		JobId: jobId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/collect/v1beta1/jobs/{{.JobId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: jobPatch})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SingleJobResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PatchJobs - Finds all jobs that match the query and modifies the with the changes specified in the request.
	This is a non-atomic operation and the results are returned as a list with each job patch result as its element. This API returns &#x60;200 OK&#x60; regardless of how many jobs were successfully patched. You must read the response body to find out if all jobs are patched. When the API is called, the &#x60;jobIDs&#x60; or &#x60;connectorID&#x60; must be specified. Do not specify more than one of them at the same time. This API returns &#x60;403&#x60; if the number of collect workers is over a certain limit.
	Parameters:
		jobsPatch: The API request schema for patching jobs.
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PatchJobs(jobsPatch JobsPatch, query *PatchJobsQueryParams, resp ...*http.Response) (*PatchJobsResponse, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/collect/v1beta1/jobs`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: jobsPatch})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PatchJobsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}
