package request

import "github.com/gorilla/mux"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
ResolveClusterGroupRequest
Method:  GET
Route:   /v1/cluster/groups/{groupid}
GroupID: cluster groupid
*/
type ClusterGroupRequest struct {
	GroupID string `json:"groupid"`
}

func ResolveClusterGroupRequest(r *http.Request) (*ClusterGroupRequest, error) {

	vars := mux.Vars(r)
	groupid := strings.TrimSpace(vars["groupid"])
	if groupid == "" {
		return nil, fmt.Errorf("request groupid invalid.")
	}

	return &ClusterGroupRequest{
		GroupID: groupid,
	}, nil
}

/*
ResolveClusterEngineRequest
Method:  GET
Route:   /v1/cluster/engines/{server}
Server: cluster engine host address
*/
type ClusterEngineRequest struct {
	Server string `json:"server"`
}

func ResolveClusterEngineRequest(r *http.Request) (*ClusterEngineRequest, error) {

	vars := mux.Vars(r)
	server := strings.TrimSpace(vars["server"])
	if server == "" {
		return nil, fmt.Errorf("request engine server invalid.")
	}

	return &ClusterEngineRequest{
		Server: server,
	}, nil
}

const (
	GROUP_CREATE_EVENT = "create"
	GROUP_REMOVE_EVENT = "remove"
	GROUP_CHANGE_EVENT = "change"
)

type ClusterGroupEventRequest struct {
	GroupID string `json:"groupid"`
	Event   string `json:"event"`
}

/*
ResolveClusterGroupEventRequest
Method:  POST
Route:   /v1/cluster/groups/event
GroupID: cluster groupid
Event:   event type
*/
func ResolveClusterGroupEventRequest(r *http.Request) (*ClusterGroupEventRequest, error) {

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	request := &ClusterGroupEventRequest{}
	if err := json.NewDecoder(bytes.NewReader(buf)).Decode(request); err != nil {
		return nil, err
	}

	request.Event = strings.ToLower(request.Event)
	request.Event = strings.TrimSpace(request.Event)
	if request.Event != GROUP_CREATE_EVENT &&
		request.Event != GROUP_REMOVE_EVENT &&
		request.Event != GROUP_CHANGE_EVENT {
		return nil, fmt.Errorf("request group event invalid.")
	}
	return request, nil
}
