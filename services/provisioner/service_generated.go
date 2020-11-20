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
 * Provisioner
 *
 * With the Provisioner service in Splunk Cloud Services, you can provision and manage tenants.
 *
 * API version: v1beta1.4 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package provisioner

import (
	"net/http"

	"github.com/splunk/go-dependencies/services"
	"github.com/splunk/go-dependencies/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new provisioner service client from the given Config
func NewService(iClient services.IClient) *Service {
	return &Service{Client: iClient}
}

/*
	CreateInvite - provisioner service endpoint
	Creates an invitation for a person to join the tenant using their email address.
	Parameters:
		inviteBody
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateInvite(inviteBody InviteBody, resp ...*http.Response) (*InviteInfo, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/provisioner/v1beta1/invites`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: inviteBody})
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
	var rb InviteInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateProvisionJob - provisioner service endpoint
	Creates a new job that provisions a new tenant and subscribes apps to the tenant.
	Parameters:
		createProvisionJobBody
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateProvisionJob(createProvisionJobBody CreateProvisionJobBody, resp ...*http.Response) (*ProvisionJobInfo, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/provisioner/v1beta1/jobs/tenants/provision`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: createProvisionJobBody})
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
	var rb ProvisionJobInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteInvite - provisioner service endpoint
	Removes an invitation in the given tenant.
	Parameters:
		inviteId
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteInvite(inviteId string, resp ...*http.Response) error {
	pp := struct {
		InviteId string
	}{
		InviteId: inviteId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/provisioner/v1beta1/invites/{{.InviteId}}`, pp)
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
	GetInvite - provisioner service endpoint
	Returns an invitation in the given tenant.
	Parameters:
		inviteId
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetInvite(inviteId string, resp ...*http.Response) (*InviteInfo, error) {
	pp := struct {
		InviteId string
	}{
		InviteId: inviteId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/provisioner/v1beta1/invites/{{.InviteId}}`, pp)
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
	var rb InviteInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetProvisionJob - provisioner service endpoint
	Returns details of a specific provision job.
	Parameters:
		jobId
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetProvisionJob(jobId string, resp ...*http.Response) (*ProvisionJobInfo, error) {
	pp := struct {
		JobId string
	}{
		JobId: jobId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/provisioner/v1beta1/jobs/tenants/provision/{{.JobId}}`, pp)
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
	var rb ProvisionJobInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetTenant - provisioner service endpoint
	Returns a specific tenant.
	Parameters:
		tenantName
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetTenant(tenantName string, resp ...*http.Response) (*TenantInfo, error) {
	pp := struct {
		TenantName string
	}{
		TenantName: tenantName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/provisioner/v1beta1/tenants/{{.TenantName}}`, pp)
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
	var rb TenantInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListInvites - provisioner service endpoint
	Returns a list of invitations in a given tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListInvites(resp ...*http.Response) (*Invites, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/provisioner/v1beta1/invites`, nil)
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
	var rb Invites
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListProvisionJobs - provisioner service endpoint
	Returns a list of all provision jobs created by the user.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListProvisionJobs(resp ...*http.Response) (*ProvisionJobs, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/provisioner/v1beta1/jobs/tenants/provision`, nil)
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
	var rb ProvisionJobs
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListTenants - provisioner service endpoint
	Returns all tenants that the user can read.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListTenants(resp ...*http.Response) (*Tenants, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/provisioner/v1beta1/tenants`, nil)
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
	var rb Tenants
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	UpdateInvite - provisioner service endpoint
	Modifies an invitation in the given tenant.
	Parameters:
		inviteId
		updateInviteBody
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UpdateInvite(inviteId string, updateInviteBody UpdateInviteBody, resp ...*http.Response) (*InviteInfo, error) {
	pp := struct {
		InviteId string
	}{
		InviteId: inviteId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/provisioner/v1beta1/invites/{{.InviteId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: updateInviteBody})
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
	var rb InviteInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}
