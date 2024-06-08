# awic - AWS instance basic hardware viewer and calculator

Outputs HW information about indicated AWS instance type

Allows to quickly calculate CPU and Memory resources for the `N` number instances of the same type.

Another useful tool https://github.com/aws/amazon-ec2-instance-selector

```txt
awic --help
Usage of awic:
  -i string
        Instance type. Required. Example: c5.4xlarge
  -m int
        Instances multiplier (default 1)
  -p string
        AWS profile name
  -r string
        AWS region name
```

Examples:
```sh
# default information about instance type
awic -p nonprd -i r5a.16xlarge
r5a.16xlarge - vCPUs=32, Mem GiB=512, Gbps=12/12, Pods=737

# Quickly output resulting CPU and Memory for X instances of the same type
awic -i r5a.16xlarge -m 2
r5a.16xlarge - vCPUs=32, Mem GiB=512, Gbps=12/12, Pods=737
x2 - vCPUs=64, Mem=1024
```
