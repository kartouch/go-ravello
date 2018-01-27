# Go-Ravello

Limited api binding to ravello.


## FUNCTIONS

~~~
func DeleteApplication(id uint64) (err error)
    DeleteApplication deletes a given application

func GetApplicationsList() (apps []Application, err error)
    GetApplicationsList returns list of applications

func GetVMSList(id uint64) (vms []VM, err error)
    GetVMSList returns the vms for a given application id Hardcoded
    deployment endpoint

func GetVMVncURL(id uint64, vmID uint64) (url string, err error)
    GetVMVncURL returns the VNC url for a VM in a given application URL
    lifetime is 1 minute

func PublishApplication(id uint64, prefs []byte) (err error)
    PublishApplication publishes the application

func SetApplicationExpirationTime(id uint64, time int) (err error)
    SetApplicationExpirationTime sets the expiration time for a given
    application id
~~~