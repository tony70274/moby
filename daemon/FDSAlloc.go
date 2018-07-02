package daemon

import (
//	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/container"
	"text/tabwriter"
	//"time"
//	"github.com/docker/docker/api/types/container"
//	"github.com/docker/docker/errdefs"
//	"github.com/pkg/errors"
	"os/exec"
	"os"
	//"time"
	"time"
)

func (daemon *Daemon) ContainerFDS(policy *types.ContainerFDSOptions) error {
/*
	var warnings []string

	c, err := daemon.GetContainer(name)
	if err != nil {
		return container.ContainerUpdateOKBody{Warnings: warnings}, err
	}

	warnings, err = daemon.verifyContainerSettings(c.OS, hostConfig, nil, true)
	if err != nil {
		return container.ContainerUpdateOKBody{Warnings: warnings}, errdefs.InvalidParameter(err)
	}

	if err := daemon.update(name, hostConfig); err != nil {
		return container.ContainerUpdateOKBody{Warnings: warnings}, err
	}
*/
	var containerInfo []*container.Container
	containers ,err := daemon.Containers(&types.ContainerListOptions{})
	if err != nil{
		println("Get ContainerList error")
		return nil
	}else{
		println("Get ContainerList")
	}
	for _, container := range containers{
		tmp, err := daemon.GetContainer(container.ID)
		//println(&tmp)
		println(tmp.HostConfig.Resources.CPUPeriod)
		if err != nil{
			println("Get ContainerInfo error")
			return nil
		}else{
			println("Get ContainerInfo")
		}
		containerInfo = append(containerInfo,tmp)
		println("Info: ",&containerInfo)

	}
	println("Info",&containerInfo)
	//show(containerInfo)
	//sshowInfo(containerInfo)
	for range time.Tick(time.Millisecond * 1000){
		//cleanScreen()
		showInfo(containerInfo)
	}
    println("In the func of ContainerFDS from Daemon.")
    println("Policy is ",policy.Policy)

	return nil
}
func showInfo(cs []*container.Container){
	println(&cs)
	w := tabwriter.NewWriter(os.Stdout, 12, 1 ,3 ,' ',0 )
	//fmt.Fprint(w, "Name\tCPU%\tAVG\tQuota\tPeriod\tMax_CPU\n")
	fmt.Fprint(w, "Name\tPeriod\tQuota\n")
	for _, c := range cs{
		fmt.Fprintf(w, "%s\n",
			c.Name,
			c.HostConfig.CPUPeriod,
			c.HostConfig.CPUQuota,
		)
	}



}
func show(cs []*container.Container){
	println(&cs)
}
func cleanScreen(){
	cmd := exec.Command("Clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
/*
func (daemon *Daemon) getContainerResource() ([]*container.Container,error) {
	var(
	containerInfo = []*container.Container{}
	)
	containers ,err := daemon.Containers(&types.ContainerListOptions{})
	if err != nil{
		println("Get ContainerList error")
		return nil,err
	}else{
		println("Get ContainerList")
	}
	for _, container := range containers{
		tmp, err := daemon.GetContainer(container.ID)
	//println(&tmp)
		println(tmp.HostConfig.Resources.CPUPeriod)
		if err != nil{
			println("Get ContainerInfo error")
			return nil,nil
		}else{
			println("Get ContainerInfo")
		}
	containerInfo = append(containerInfo,tmp)
	println("Info: ",&containerInfo)

	}


}
*/
