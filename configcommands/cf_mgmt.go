package configcommands

type CfMgmtConfigCommand struct {
	Version                             VersionCommand                      `command:"version" description:"Print version information and exit"`
	InitConfigurationCommand            InitConfigurationCommand            `command:"init" description:"Initializes folder structure for configuration"`
	OrgConfigurationCommand             OrgConfigurationCommand             `command:"org" description:"Adds/updates specified org to configuration"`
	SpaceConfigurationCommand           SpaceConfigurationCommand           `command:"space" description:"adds/updates space configuration"`
	AddOrgToConfigurationCommand        AddOrgToConfigurationCommand        `command:"add-org" description:"Adds specified org to configuration"`
	AddSpaceToConfigurationCommand      AddSpaceToConfigurationCommand      `command:"add-space" description:"Adds specified space to configuration for org"`
	GenerateConcoursePipelineCommand    GenerateConcoursePipelineCommand    `command:"generate-concourse-pipeline" description:"generates a concourse pipline to be used to drive cf-mgmt"`
	UpdateOrgConfigurationCommand       UpdateOrgConfigurationCommand       `command:"update-org" description:"updates org configuration"`
	UpdateSpaceConfigurationCommand     UpdateSpaceConfigurationCommand     `command:"update-space" description:"updates space configuration"`
	DeleteOrgConfigurationCommand       DeleteOrgConfigurationCommand       `command:"delete-org" description:"deletes org configuration"`
	DeleteSpaceConfigurationCommand     DeleteSpaceConfigurationCommand     `command:"delete-space" description:"deletes space configuration"`
	AddASGToConfigurationCommand        AddASGToConfigurationCommand        `command:"add-asg" description:"add a named asg to configuration"`
	ASGToConfigurationCommand           ASGToConfigurationCommand           `command:"asg" description:"creates/updates a named asg"`
	UpdateOrgsConfigurationCommand      UpdateOrgsConfigurationCommand      `command:"update-orgs" description:"updates orgs.yml"`
	RenameOrgConfigurationCommand       RenameOrgConfigurationCommand       `command:"rename-org" description:"renames an org"`
	RenameSpaceConfigurationCommand     RenameSpaceConfigurationCommand     `command:"rename-space" description:"renames a space for a given org"`
	OrgNamedQuotaConfigurationCommand   OrgNamedQuotaConfigurationCommand   `command:"named-org-quota" description:"creates/updates named org quota"`
	SpaceNamedQuotaConfigurationCommand SpaceNamedQuotaConfigurationCommand `command:"named-space-quota" description:"creates/updates named space quota"`
}

var CfMgmtConfig CfMgmtConfigCommand
