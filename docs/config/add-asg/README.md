&larr; [back to Commands](../README.md)

# `cf-mgmt-config add-asg`

*** Deprecated *** - Use `asg` command instead

`add-asg` will add a <asg-name>.json file in the asgs folder of the configuration

## Command Usage

```
Usage:
  cf-mgmt-config [OPTIONS] add-asg [add-asg-OPTIONS]

Help Options:
  -h, --help            Show this help message

[add-asg command options]
          --config-dir= Name of the config directory (default: config) [$CONFIG_DIR]
          --asg=        ASG name
          --path=       path to asg definition
          --override    override current definition
          --type=[space|default] Space asg or default asg (default: space)

```
