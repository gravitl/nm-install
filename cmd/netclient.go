package cmd

import (
	"os"
	"runtime"
	"strings"

	"github.com/bitfield/script"
	"github.com/pterm/pterm"
)

func installNetclient() {
	newInstall := false
	existingVersion, _ := script.Exec("netclient version").String()
	pterm.Println("checking for installed netclient ", "existing version", existingVersion, "latest", latest)
	if strings.TrimSpace(existingVersion) == latest {
		pterm.Println("latest version of netclient already installed, skipping...")
		return
	} else if strings.Contains(existingVersion, "file not found") {
		newInstall = true
	}
	pterm.Println("netclient new install:", newInstall)
	pterm.Println("retrieving the latest version of netclient")
	arch := runtime.GOARCH
	baseURL := "https://github.com/gravitl/netclient/releases/download/" + latest + "/netclient-linux-" + arch
	getFile(baseURL, "", "/tmp/netclient")
	os.Chmod("/tmp/netclient", 0700)
	if _, err := script.Exec("/tmp/netclient install").Stdout(); err != nil {
		panic(err)
	}
	if newInstall {
		pterm.Println("latest version of netclient installed")
	} else {
		pterm.Println("netclient updated to latest version")
		// since this was an update don't contine
		return
	}
	pterm.Println("joining network with token", token)
	if token != "" {
		if _, err := script.Exec("netclient join -t " + token).Stdout(); err != nil {
			panic(err)
		}
	} else {
		pterm.Println("enrollment key not defined")
	}
	nodeID, err := script.Exec("/tmp/nmctl node list --output json").JQ(".[]").JQ(".id").String()
	if err != nil {
		pterm.Println("error reading netclient config:", err)
		return
	}
	hostID, err := script.Exec("/tmp/nmctl node list --output json").JQ(".[]").JQ(".hostid").String()
	if err != nil {
		pterm.Println("error reading netclient config:", err)
		return
	}
	pterm.Println("setting default node (nmctl host update " + hostID + " --default)")
	script.Exec("/tmp/nmctl host update " + hostID + " --default").Stdout()
	pterm.Println("creating ingress gateway (nmctl node create_ingress netmaker " + nodeID + ")")
	if hostID == "" {
		pterm.Println("something went wrong joining network, unable to create ingress")
		return
	}
	script.Exec("/tmp/nmctl node create_ingress netmaker " + nodeID).Stdout()
}
