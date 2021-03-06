package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/infrakit/pkg/cli"
	plugin_base "github.com/docker/infrakit/pkg/plugin"
	instance_rpc "github.com/docker/infrakit/pkg/rpc/instance"
	"github.com/docker/infrakit/pkg/run"
	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Use:   os.Args[0],
		Short: "VMware vSphere instance plugin",
	}

	// This will hold the configuration that is used to communicate with VMware vCenter or vSphere
	var newVCenter vCenter

	name := cmd.Flags().String("name", "instance-vsphere", "Plugin name to advertise for discovery")
	logLevel := cmd.Flags().Int("log", cli.DefaultLogLevel, "Logging level. 0 is least verbose. Max is 5")

	// Attributes of the VMware vCenter Server to connect to
	newVCenter.vCenterURL = cmd.Flags().String("url", os.Getenv("VCURL"), "URL of VMware vCenter in the format of https://username:password@VCaddress/sdk")
	newVCenter.dcName = cmd.Flags().String("datacenter", os.Getenv("VCDATACENTER"), "The name of a Datacenter within vCenter")
	newVCenter.dsName = cmd.Flags().String("datastore", os.Getenv("VCDATASTORE"), "The name of the DataStore to host the VM")
	newVCenter.networkName = cmd.Flags().String("network", os.Getenv("VCNETWORK"), "The network label the VM will use")
	newVCenter.vSphereHost = cmd.Flags().String("hostname", os.Getenv("VCHOST"), "The server that will run the VM")

	// Testing flag to ensure VMs are never deleted
	ignoreOnDestroy := cmd.Flags().Bool("ignoreOnDestroy", false, "When set to true, InfraKit will only Mark VMs as deleted")

	cmd.Run = func(c *cobra.Command, args []string) {
		cli.SetLogLevel(*logLevel)
		run.Plugin(plugin_base.DefaultTransport(*name),
			instance_rpc.PluginServer(NewVSphereInstancePlugin(&newVCenter, *ignoreOnDestroy)))
	}

	cmd.AddCommand(cli.VersionCommand())

	err := cmd.Execute()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
