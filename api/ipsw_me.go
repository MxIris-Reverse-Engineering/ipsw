package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ipswMeAPI        = "https://api.ipsw.me/v4/"
	canijailbreakAPI = "https://canijailbreak.com/jailbreaks.json"
)

// Device struct
type Device struct {
	Name        string `json:"name,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
	BoardConfig string `json:"boardconfig,omitempty"`
	Platform    string `json:"platform,omitempty"`
	CpID        int    `json:"cpid,omitempty"`
	BdID        int    `json:"bdid,omitempty"`
	Firmwares   []IPSW `json:"firmwares,omitempty"`
}

// IPSW struct
type IPSW struct {
	Identifier  string    `json:"identifier,omitempty"`
	Version     string    `json:"version,omitempty"`
	BuildID     string    `json:"buildid,omitempty"`
	SHA1        string    `json:"sha1sum,omitempty"`
	MD5         string    `json:"md5sum,omitempty"`
	Filesize    int       `json:"filesize,omitempty"`
	URL         string    `json:"url,omitempty"`
	ReleaseDate time.Time `json:"releasedate,omitempty"`
	UploadDate  time.Time `json:"uploaddate,omitempty"`
	Signed      bool      `json:"signed,omitempty"`
}

// GetAllDevices returns a list of all devices
func GetAllDevices() ([]Device, error) {
	devices := []Device{}

	res, err := http.Get(ipswMeAPI + "devices")
	if err != nil {
		return devices, err
	}
	if res.StatusCode != http.StatusOK {
		return devices, fmt.Errorf("api returned status: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return devices, err
	}
	res.Body.Close()

	err = json.Unmarshal(body, &devices)
	if err != nil {
		return devices, err
	}

	return devices, nil
}

// GetDevice returns a device from it's identifier
func GetDevice(identifier string) (Device, error) {
	d := Device{}

	res, err := http.Get(ipswMeAPI + "device" + "/" + identifier)
	if err != nil {
		return d, err
	}
	if res.StatusCode != http.StatusOK {
		return d, fmt.Errorf("api returned status: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return d, err
	}
	res.Body.Close()

	err = json.Unmarshal(body, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

// GetAllIPSW finds all IPSW files for a given iOS version
func GetAllIPSW(version string) ([]IPSW, error) {
	ipsws := []IPSW{}

	res, err := http.Get(ipswMeAPI + "ipsw/" + version)
	if err != nil {
		return ipsws, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned status: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ipsws, err
	}
	res.Body.Close()

	err = json.Unmarshal(body, &ipsws)
	if err != nil {
		return ipsws, err
	}

	return ipsws, nil
}

// GetIPSW will get an IPSW when supplied an identifier and build ID
func GetIPSW(identifier, buildID string) (IPSW, error) {
	i := IPSW{}

	res, err := http.Get(ipswMeAPI + "ipsw/" + identifier + "/" + buildID)
	if err != nil {
		return i, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return i, fmt.Errorf("api returned status: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return i, err
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		return i, err
	}

	return i, nil
}

// https://api.ipsw.me/v4/releases
// func GetReleases() []Release {}
