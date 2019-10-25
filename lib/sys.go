package lib

import (
	"bytes"
	"fmt"
	util "github.com/eagle7410/go_util/libs"
	"io/ioutil"
	"os"
	"text/template"
	"time"
)

const ServicePath = "/lib/systemd/system/testApp.service"
func UninstallService() {
	out, err := SystemctlStopService()

	if err != nil {
		util.LogFatalf("Error on stop service: %s", err)
	}

	fmt.Printf("Out stop service \n %v \n", string(out))

	out, err = SystemctlDisableService()

	if err != nil {
		util.LogFatalf("Error on disable service: %s", err)
	}

	fmt.Printf("Out disable service \n %v \n", string(out))

	if err = os.Remove(ServicePath); err != nil {
		util.LogFatalf("Error on remove file service: %s", err)
	}
}

func InstallAsService() {

	tpl := template.New("service")

	if _, err := tpl.Parse(ServiceTpl); err != nil {
		util.LogFatalf("Error on parse template service : %s", err)
	}

	var buffer bytes.Buffer

	renderData := struct {
		Workdir string
	}{
		Workdir: ENV.WorkDir,
	}

	if err := tpl.Execute(&buffer, renderData); err != nil {
		util.LogFatalf("Error on render service file : %s", err)
	}

	if err := ioutil.WriteFile(ServicePath, buffer.Bytes(), 0644); err != nil {
		util.LogFatalf("Error on write service file : %s", err)
	}

	if _, err := SystemctlEnableService(); err != nil {
		util.LogFatalf("Error on enable service: %s", err)
	}

	fmt.Printf("Enable service Ok")

	if _, err := SystemctlRunService(); err != nil {
		util.LogFatalf("Error on run service: %s", err)
	}

	fmt.Printf("Start service Ok")
}

func SystemctlStopService() ([]byte, error) {
	return util.ExecCommandWithTimeLimit(time.Second*10, ENV.WorkDir, "systemctl", "stop", "testApp")
}
func SystemctlDisableService() ([]byte, error) {
	return util.ExecCommandWithTimeLimit(time.Second*10, ENV.WorkDir, "systemctl", "disable", "testApp")
}
func SystemctlEnableService() ([]byte, error) {
	return util.ExecCommandWithTimeLimit(time.Second*10, ENV.WorkDir, "systemctl", "enable", "testApp")
}
func SystemctlRunService() ([]byte, error) {
	return util.ExecCommandWithTimeLimit(time.Second*10, ENV.WorkDir, "systemctl", "start", "testApp")
}

/**
systemctl daemon-reload
systemctl restart testApp
systemctl status testApp
*/
const ServiceTpl = `[Unit]
After=network.target
Description=Test app
[Service]
Restart=always
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=testApp
WorkingDirectory={{.Workdir}}
ExecStart={{.Workdir}}/serve
ExecReload=/bin/kill -SIGINT $MAINPID
Timeout=0
Type=simplegs

[Install]
WantedBy=multi-user.target
`
