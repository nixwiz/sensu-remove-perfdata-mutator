[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/nixwiz/sensu-remove-perfdata-mutator)
![Go Test](https://github.com/nixwiz/sensu-remove-perfdata-mutator/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/nixwiz/sensu-remove-perfdata-mutator/workflows/goreleaser/badge.svg)

# Sensu Go Mutator to remove Nagios Performance Data from Check Output

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Mutator definition](#mutator-definition)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

sensu-remove-perfdata-mutator is a [Sensu Mutator][2] that removes
[Nagios Performance Data][3] from an event before passing it to a handler.
This is done in an attempt to make the output presented to the handler more
readable.

## Usage examples

```
Sensu Go Mutator to remove Nagios Performance Data from Check Output

Usage:
  sensu-remove-perfdata-mutator [flags]
  sensu-remove-perfdata-mutator [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -h, --help      help for sensu-remove-perfdata-mutator
  -v, --verbose   Provide verbose output

Use "sensu-remove-perfdata-mutator [command] --help" for more information about a command.
```

Using this mutator will strip off the Nagios Performance Data included in the check output so that
the handler only has access to the more readable version and can choose to surface it as necessary.

Before:
```
CheckCPU OK: user=0.63% system=0.25% iowait=0.25% other=0.38% idle=98.50% | cpu_user=0.63%;80;90 cpu_system=0.25%;80;90 cpu_iowait=0.25%;80;90 cpu_other=0.38%;80;90 cpu_idle=98.50%
```

After:
```
CheckCPU OK: user=0.63% system=0.25% iowait=0.25% other=0.38% idle=98.50%
```

## Configuration

### Asset registration

[Sensu Assets][5] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add nixwiz/sensu-remove-perfdata-mutator
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][4].

### Mutator definition

```yml
---
type: Mutator
api_version: core/v2
metadata:
  name: sensu-remove-perfdata-mutator
  namespace: default
spec:
  command: sensu-remove-perfdata-mutator
  runtime_assets:
  - nixwiz/sensu-remove-perfdata-mutator
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the sensu-remove-perfdata-mutator repository:

```
go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://docs.sensu.io/sensu-go/latest/reference/mutators/
[3]: https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/perfdata.html
[4]: https://bonsai.sensu.io/assets/nixwiz/sensu-remove-perfdata-mutator
[5]: https://docs.sensu.io/sensu-go/latest/reference/assets/
