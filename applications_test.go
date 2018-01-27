package ravello

import (
	"encoding/json"
	"regexp"
	"testing"
)

var appID uint64
var vmID uint64

func TestGetApplicationsList(t *testing.T) {
	apps, err := GetApplicationsList()
	if err != nil {
		t.Fatalf("could not get applications, got: %v", err)
	}

	// Let's assume we have at least 1 app in the library
	if len(apps) < 0 {
		t.Fatalf("could not get applications, got: %v", err)
	}
}

func TestCreateApplication(t *testing.T) {
	j, _ := json.Marshal(Application{BaseBlueprintID: 86312957, Name: "ravello-go-test", Description: "ravello test", CostBucket: &CostBucket{ID: 95059969}})
	app, err := CreateApplication(j)
	if err != nil {
		t.Fatalf("could not create application, got: %v", err)
	}

	appID = app.ID
	if app.Name != "ravello-go-test" {
		t.Fatalf("could not create application. Got: %v , expected: ravello-go-test", app.Name)
	}
}

func TestPublishApplication(t *testing.T) {
	prefs, _ := json.Marshal(Preference{PreferredRegion: "us-central-1", OptimizationLevel: "PERFORMANCE_OPTIMIZED", StartAllVms: "true"})
	err := PublishApplication(appID, prefs)
	if err != nil {
		t.Fatalf("could not publish application, got: %v", err)
	}
}

func TestGetApplication(t *testing.T) {
	app, err := GetApplication(appID)
	if err != nil {
		t.Fatalf("could not get application, got: %v", err)
	}

	if app.ID != appID {
		t.Fatalf("could not get application. Got: %v, expected: %v", app.ID, appID)
	}
}

func TestGetVMSList(t *testing.T) {
	vms, err := GetVMSList(appID)
	if err != nil {
		t.Fatalf("could not get vms, got: %v", err)
	}

	// Let's assume there is at least 1 vm
	// Continue with first VM found from Blueprint for rest of tests
	if len(vms) != 0 {
		for _, vm := range vms {
			vmID = vm.ID
		}
	} else {
		t.Fatalf("could not get vms. Got: %v, expected: %v", 0, len(vms))
	}
}

func TestGetVM(t *testing.T) {
	vm, err := GetVM(appID, vmID)
	if err != nil {
		t.Fatalf("could not get vm: %v", err)
	}

	if vm.Name != "noonelovesme.example.com" {
		t.Fatalf("could not get vm. Go %v, expected: noonelovesme.example.com", vm.Name)
	}
}

func TestSetExpirationTime(t *testing.T) {
	err := SetApplicationExpirationTime(appID, 300)
	if err != nil {
		t.Fatalf("could not set expiration time: %v", err)
	}
}

func TestGetVMVncURL(t *testing.T) {
	url, err := GetVMVncURL(appID, vmID)
	if err != nil {
		t.Fatalf("could not get VNC url: %v", err)
	}

	matched, err := regexp.MatchString(".*ravellosystems.com/vnc/.*", url)
	if err != nil {
		t.Fatalf("could not regexp VNC url: %v", err)
	}
	if !matched {
		t.Fatalf("could not get VNC url (regexp not matching): %v", matched)
	}
}

func TestGetFQDN(t *testing.T) {
	fqdn, err := GetFQDN(appID, vmID)
	if err != nil {
		t.Fatalf("could not get vm FQDN: %v", err)
	}

	matched, err := regexp.MatchString(".*srv.ravcloud.com", fqdn.Value)
	if err != nil {
		t.Fatalf("could not regexp VNC url: %v", err)
	}
	if !matched {
		t.Fatalf("could not get fqdn (regexp not matching): %v", matched)
	}

}

func TestExecuteApplicatonAction(t *testing.T) {
	action, err := ExecuteApplicatonAction(appID, "stop")
	if err != nil {
		t.Fatalf("could not request action: %v", err)
	}

	if action.CompletedSuccessfuly != "true" {
		t.Fatal("could not complete request action")
	}
}

func TestDeleteApplication(t *testing.T) {
	err := DeleteApplication(appID)
	if err != nil {
		t.Fatalf("could not delete application, got: %v", err)
	}

	app, err := GetApplication(appID)
	if err != nil {
		t.Fatalf("could not get application, got: %v", err)
	}

	if app.ID == appID {
		t.Fatal("application was not deleted")
	}

}
