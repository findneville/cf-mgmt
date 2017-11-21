package configcommands

import (
	"errors"
	"fmt"

	"github.com/pivotalservices/cf-mgmt/config"
)

type AddASGToConfigurationCommand struct {
	ConfigManager config.Manager
	BaseConfigCommand
	ASGName  string `long:"asg" description:"ASG name" required:"true"`
	FilePath string `long:"path" description:"path to asg definition"`
	Override bool   `long:"override" description:"override current definition"`
}

//Execute - adds a named asg to the configuration
func (c *AddASGToConfigurationCommand) Execute([]string) error {
	c.initConfig()

	errorString := ""

	if errorString != "" {
		return errors.New(errorString)
	}

	if !c.Override {
		asgConfigs, err := c.ConfigManager.GetASGConfigs()
		if err != nil {
			return err
		}
		for _, asgConfig := range asgConfigs {
			if c.ASGName == asgConfig.Name {
				return errors.New(fmt.Sprintf("asg [%s] already exists if wanting to update use --override flag", c.ASGName))
			}
		}
	}

	securityGroupsBytes := []byte("[]")
	if c.FilePath != "" {
		bytes, err := config.LoadFileBytes(c.FilePath)
		if err != nil {
			return err
		}
		securityGroupsBytes = bytes
	}
	if err := config.NewManager(c.ConfigDirectory).AddSecurityGroup(c.ASGName, securityGroupsBytes); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("The asg [%s] has been updated", c.ASGName))
	return nil
}

func (c *AddASGToConfigurationCommand) initConfig() {
	if c.ConfigManager == nil {
		c.ConfigManager = config.NewManager(c.ConfigDirectory)
	}
}