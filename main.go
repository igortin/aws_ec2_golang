package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	// profileaName = os.Getenv("AWS_PROFILE")
	// awsRegion    = os.Getenv("AWS_REGION")
	profileaName        = "private"
	awsRegion           = "us-east-1"
	tagKey              = "tag:Environment"
	tagValue            = "qa"
	instanceStateFilter = "running"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(awsRegion)},
		Profile: profileaName,
	})
	if err != nil {
		log.Println("Error: Can not create new session.")
	}

	// Get data from AWS service
	result, err := ec2GetResponse(sess)
	if err != nil {
		log.Println("Error: Can noot get data")
		return
	}

	// Check  length of array ec2.DescribeInstancesOutput.Reservations
	if len(result.Reservations) == 0 {
		log.Println("Error: ec2.DescribeInstancesOutput.Reservations is empty")
		return
	}

	// Func parse result to []*Ec2object
	ec2Params, err := parseEc2Response(result)
	if err != nil {
		log.Println(err)
		return
	}

	// Show objects parameters
	for _, item := range ec2Params {
		ShowOutput(item)
		// fmt.Println(string(item.ToJSON()))
	}
}

// Func to get data from AWS EC2 service
func ec2GetResponse(sess *session.Session) (*ec2.DescribeInstancesOutput, error) {

	// New client with a session
	ec2svc := ec2.New(sess)

	// Init instance request of structure DescribeInstancesInput
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(tagKey),
				Values: []*string{
					aws.String(tagValue),
				},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String(instanceStateFilter), aws.String("pending")},
			},
		},
	}
	// Return instance structure of DescribeInstancesOutput by Method
	response, err := ec2svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func parseEc2Response(data *ec2.DescribeInstancesOutput) ([]*Ec2object, error) {
	Ec2objectList := []*Ec2object{}
	// iterate []*Reservations
	for _, reservation := range data.Reservations {

		//  iterate []*Instance
		for _, instance := range reservation.Instances {
			//log.Printf("Procced number %d - instnace ID: %v\n", index + 1, *instance.InstanceId)

			object := Ec2object{
				InstanceID:     *instance.InstanceId,
				InstanceState:  *instance.State.Name,
				PrivateDNSName: *instance.PrivateDnsName,
				PublicDNSame:   *instance.PublicDnsName,
			}

			// Init blockDev list for each object
			blockDeviceList := []*BlockDevice{}

			// Create  object's BlockDevicesList

			// []*InstanceBlockDeviceMapping
			// Please mention that there are method for instance blk for example SetDeviceName (can be used)
			for _, blk := range instance.BlockDeviceMappings {
				// Init block
				blockDev := &BlockDevice{
					DeviceName: *blk.DeviceName,
					State:      *blk.Ebs.Status,
					VolumeID:   *blk.Ebs.VolumeId,
				}
				// fullfill array []*BlockDevice{}
				blockDeviceList = append(blockDeviceList, blockDev)
			}

			// fullfill object's prperty by BlockDevicesList
			object.BlockDevicesList = blockDeviceList

			// Append to list of pointers a new pointer to object in memory
			Ec2objectList = append(Ec2objectList, &object)
		}
	}
	return Ec2objectList, nil
}
