package ravello

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Application are applications properties
type Application struct {
	ID                          uint64        `json:"id,omitempty"`
	Name                        string        `json:"name,omitempty"`
	Owner                       string        `json:"owner,omitempty"`
	Description                 string        `json:"description,omitempty"`
	BaseBlueprintID             int           `json:"baseBlueprintId,omitempty"`
	BlueprintName               string        `json:"blueprintName,omitempty"`
	DesignDiffersFromDeployment bool          `json:"designDiffersFromDeployment,omitempty"`
	Published                   bool          `json:"published,omitempty"`
	CreationTime                uint8         `json:"creationTime,omitempty"`
	Version                     uint64        `json:"version,omitempty"`
	NextStopTask                *NextStopTask `json:"nextStopTask,omitempty"`
	CostBucket                  *CostBucket   `json:"costBucket,omitempty"`
	Deployment                  *Deployment   `json:"deployment,omitempty"`
}

// Deployment specifics for an application
type Deployment struct {
	TotalActiveVms int `json:"totalActiveVms,omitempty"`
	TotalErrorVms  int `json:"totalErrorVms,omitempty"`
}

//CostBucket tracking struct
type CostBucket struct {
	ID uint64 `json:"id,omitempty"`
}

// NextStopTask returns ID and time of task
type NextStopTask struct {
	ID   uint64    `json:"id,omitempty"`
	Time time.Time `json:"time,omitempty"`
}

// VM are vm properties from a given application
type VM struct {
	ID                 uint64              `json:"id,omitempty"`
	Name               string              `json:"name,omitempty"`
	SuppliedServices   []SuppliedService   `json:"suppliedServices,omitempty"`
	NetworkConnections []NetworkConnection `json:"networkConnections,omitempty"`
}

// SuppliedService are exposed service for a VM
// They are part of VM struct
type SuppliedService struct {
	Name         string `json:"name,omitempty"`
	PortRange    string `json:"portRange,omitempty"`
	External     bool   `json:"external,omitempty"`
	IPConfigLuid uint64 `json:"ipConfigLuid,omitempty"`
}

// NetworkConnection properties for a VM
type NetworkConnection struct {
	IPConfig *IPConfig `json:"ipConfig,omitempty"`
}

// IPConfig for a VM NetworkConnection
type IPConfig struct {
	ID   uint64 `json:"id,omitempty"`
	Fqdn string `json:"fqdn,omitempty"`
}

// ActionResult are response object for application acitons
type ActionResult struct {
	CompletedSuccessfuly string             `json:"completedSuccessfuly,omitempty"`
	OperationMessages    *OperationMessages `json:"operationMessages,omitempty"`
}

// OperationMessages are part of ActionResult response
type OperationMessages struct {
	VMID       string `json:"vmId,omitempty"`
	ErrorLevel string `json:"errorLevel,omitempty"`
	Message    string `json:"message,omitempty"`
}

// ExpirationTime are SetApplicationExpirationTime properties
type ExpirationTime struct {
	ExpirationTimeFromNowSeconds int `json:"expirationFromNowSeconds,omitempty"`
}

// Preference for publishing applications
type Preference struct {
	PreferredRegion   string `json:"preferredRegion,omitempty"`
	OptimizationLevel string `json:"optimizationLevel,omitempty"`
	StartAllVms       string `json:"startAllVms,omitempty"`
}

// FQDN returns fqdn for a VM
type FQDN struct {
	Value string `json:"value,omitempty"`
}

// GetApplicationsList returns list of applications
func GetApplicationsList() (apps []Application, err error) {
	r, err := handler("GET", "/applications", nil)
	json.Unmarshal(r, &apps)
	return apps, err
}

// GetApplication returns application details
func GetApplication(id uint64) (a Application, err error) {
	r, err := handler("GET", "/applications/"+strconv.Itoa(int(id)), nil)
	json.Unmarshal(r, &a)
	return
}

// CreateApplication creates a new application
func CreateApplication(data []byte) (a Application, err error) {
	r, err := handler("POST", "/applications/", data)
	json.Unmarshal(r, &a)
	return
}

// PublishApplication publishes the application
func PublishApplication(id uint64, data []byte) (err error) {
	r, err := handler("POST", "/applications/"+strconv.Itoa(int(id))+"/publish", data)
	fmt.Println(string(r))
	return
}

// ExecuteApplicatonAction with possible action :
// start,stop,restart,resetDisks
func ExecuteApplicatonAction(id uint64, action string) (ar ActionResult, err error) {
	r, err := handler("POST", "/applications/"+strconv.Itoa(int(id))+"/"+action, nil)
	json.Unmarshal(r, &ar)
	return
}

// GetVMSList returns the vms for a given application id
// Hardcoded deployment endpoint
func GetVMSList(id uint64) (vms []VM, err error) {
	r, err := handler("GET", "/applications/"+strconv.Itoa(int(id))+"/vms;deployment", nil)
	json.Unmarshal(r, &vms)
	return
}

// GetVM returns a single vm properties for a given application id
// Hardcoded deployment endpoint
func GetVM(id uint64, vmID uint64) (vm VM, err error) {
	r, err := handler("GET", "/applications/"+strconv.Itoa(int(id))+"/vms/"+strconv.Itoa(int(vmID))+";deployment", nil)
	json.Unmarshal(r, &vm)
	return
}

// SetApplicationExpirationTime sets the expiration time for a given application id
func SetApplicationExpirationTime(id uint64, time int) (err error) {
	j, _ := json.Marshal(ExpirationTime{ExpirationTimeFromNowSeconds: time})
	_, err = handler("POST", "/applications/"+strconv.Itoa(int(id))+"/setExpiration", j)
	return
}

// GetVMVncURL returns the VNC url for a VM in a given application
// URL lifetime is 1 minute
func GetVMVncURL(id uint64, vmID uint64) (url string, err error) {
	r, err := handler("GET", "/applications/"+strconv.Itoa(int(id))+"/vms/"+strconv.Itoa(int(vmID))+"/vncUrl", nil)
	url = string(r)
	return
}

// GetFQDN returns the fqdn for a VM in a given application
func GetFQDN(id uint64, vmID uint64) (f FQDN, err error) {
	r, err := handler("GET", "/applications/"+strconv.Itoa(int(id))+"/vms/"+strconv.Itoa(int(vmID))+"/fqdn;deployment", nil)
	json.Unmarshal(r, &f)
	return
}

// DeleteApplication deletes a given application
func DeleteApplication(id uint64) (err error) {
	_, err = handler("DELETE", "/applications/"+strconv.Itoa(int(id)), nil)
	return
}
