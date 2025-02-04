package packetfabric

import (
	"fmt"
)

const lagURI = "/v2/lags"
const lagPortCircuitURI = "/v2/lags/%s"
const lagInterfacesMemberURI = "/v2/lags/%s/members"
const lagInterfacesCircuitMemberURI = "/v2/lags/%s/members/%s"
const lagCircuitIDEnableURI = "/v2/lags/%s/enable"
const lagCircuitIDDisableURI = "/v2/lags/%s/disable"

type LinkAggregationGroup struct {
	Description string   `json:"description,omitempty"`
	Interval    string   `json:"interval,omitempty"`
	Pop         string   `json:"pop,omitempty"`
	Members     []string `json:"members,omitempty"`
}

type LinkAggregationGroupCreateResp struct {
	State         string `json:"state,omitempty"`
	PortCircuitID string `json:"port_circuit_id,omitempty"`
	Description   string `json:"description,omitempty"`
	Number        int    `json:"number,omitempty"`
	TimeCreated   string `json:"time_created,omitempty"`
	Accepted      bool   `json:"accepted,omitempty"`
}

type LinkAggregationGroupWorkflowResp struct {
	WorkflowName string `json:"workflow_name"`
}

func (c *PFClient) CreateLinkAggregationGroup(lag LinkAggregationGroup) (*LinkAggregationGroupCreateResp, error) {
	expectedResp := &LinkAggregationGroupCreateResp{}
	_, err := c.sendRequest(lagURI, postMethod, lag, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) UpdateLinkAggregationGroup(portCircuitID, description, interval string) (*LinkAggregationGroupCreateResp, error) {
	expectedResp := &LinkAggregationGroupCreateResp{}
	type UpdateLag struct {
		Description string `json:"description"`
		Interval    string `json:"interval"`
	}
	payload := UpdateLag{
		Description: description,
		Interval:    interval,
	}
	_, err := c.sendRequest(lagPortCircuitURI, putMethod, payload, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) EnableLinkAggregationGroup(portCircuitID string) (*LinkAggregationGroupWorkflowResp, error) {
	return c.changeLinkAggregationGroupState(portCircuitID, true)
}

func (c *PFClient) DisableLinkAggregationGroup(portCircuitID string) (*LinkAggregationGroupWorkflowResp, error) {
	return c.changeLinkAggregationGroupState(portCircuitID, false)
}

func (c *PFClient) changeLinkAggregationGroupState(portCircuitID string, enable bool) (*LinkAggregationGroupWorkflowResp, error) {
	expectedResp := &LinkAggregationGroupWorkflowResp{}
	var formatedURI string
	if enable {
		formatedURI = fmt.Sprintf(lagCircuitIDEnableURI, portCircuitID)
	} else {
		formatedURI = fmt.Sprintf(lagCircuitIDDisableURI, portCircuitID)
	}
	_, err := c.sendRequest(formatedURI, postMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) GetLAGInterfaces(lagPortCircuitURI string) (*[]InterfaceReadResp, error) {
	formatedURI := fmt.Sprintf(lagInterfacesMemberURI, lagPortCircuitURI)
	expectedResp := make([]InterfaceReadResp, 0)
	_, err := c.sendRequest(formatedURI, getMethod, nil, &expectedResp)
	if err != nil {
		return nil, err
	}
	return &expectedResp, nil
}

func (c *PFClient) DeleteLinkAggregationGroup(portCircuitID string) (*LinkAggregationGroupWorkflowResp, error) {
	formatedURI := fmt.Sprintf(lagPortCircuitURI, portCircuitID)
	expectedResp := &LinkAggregationGroupWorkflowResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) DeleteLinkAggregationGroupMember(lagPortCircuitID, memberPortCircuitID string) (*LinkAggregationGroupWorkflowResp, error) {
	formatedURI := fmt.Sprintf(lagInterfacesCircuitMemberURI, lagPortCircuitID, memberPortCircuitID)
	expectedResp := &LinkAggregationGroupWorkflowResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}
