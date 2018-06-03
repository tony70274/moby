package daemon

import (
//	"context"
//	"fmt"
	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/api/types/container"
//	"github.com/docker/docker/errdefs"
//	"github.com/pkg/errors"
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
    println("In the func of ContainerFDS from Daemon.")
    println("Policy is ",policy.Policy)

	return nil
}
