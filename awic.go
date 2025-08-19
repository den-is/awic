package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	insttype := flag.String("i", "", "Instance type. Required. Example: c5.4xlarge")
	multiplier := flag.Int("x", 1, "Instances multiplier")
	profile := flag.String("p", "", "AWS profile name")
	region := flag.String("r", "", "AWS region name")
	flag.Parse()

	if *insttype == "" {
		fmt.Println("Instance type is required")
		os.Exit(1)
	}

	if *multiplier < 1 {
		fmt.Println("Multiplier should be a positive integer")
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(*profile),
		config.WithRegion(*region),
	)
	if err != nil {
		log.Fatalf("Failed to connect to AWS API: %v", err)
	}

	svc := ec2.NewFromConfig(cfg)

	req := &ec2.DescribeInstanceTypesInput{
		InstanceTypes: []types.InstanceType{
			types.InstanceType(*insttype),
		},
	}

	result, err := svc.DescribeInstanceTypes(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	instMemMiB := *result.InstanceTypes[0].MemoryInfo.SizeInMiB
	instvCPUs := *result.InstanceTypes[0].VCpuInfo.DefaultCores
	baseGbps := *result.InstanceTypes[0].NetworkInfo.NetworkCards[0].BaselineBandwidthInGbps
	peakGbps := *result.InstanceTypes[0].NetworkInfo.NetworkCards[0].PeakBandwidthInGbps
	maxIfs := *result.InstanceTypes[0].NetworkInfo.MaximumNetworkInterfaces
	ipv4PerIf := *result.InstanceTypes[0].NetworkInfo.Ipv4AddressesPerInterface

	// pods per instance
	// https://github.com/awslabs/amazon-eks-ami/blob/master/files/eni-max-pods.txt
	// # of ENI * (# of IPv4 per ENI - 1) + 2
	podsPerInst := maxIfs*(ipv4PerIf-1) + 2

	fmt.Printf(
		"%s - vCPUs=%d, Mem GiB=%d, Gbps=%d/%d, Pods=%d\n",
		*insttype,
		instvCPUs,
		instMemMiB/1024,
		int(baseGbps),
		int(peakGbps),
		podsPerInst,
	)

	if *multiplier > 1 {
		fmt.Printf(
			"x%d - vCPUs=%d, Mem=%d\n",
			*multiplier,
			instvCPUs*int32(*multiplier),
			(instMemMiB/1024)*int64(*multiplier),
		)
	}
}
