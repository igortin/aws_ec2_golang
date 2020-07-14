# aws_ec2_golang
It is a repository for example, include code for creating EBS snapshots for AWS EC2 instances based on tags. 
  - Discovery EC2 instances based on tags - role:etcd
  - Parse []*ec2.DescribeInstancesOutput and create ec2 objects based on your custom structure Ec2object
  - Create EBS Snapshot and tag it for non-root Volumes without stop EC2 instance

> Note: Ec2 root volumes should be created in EC2 stopped state

### Prerequisites 
GO environment should be configured, check by command below
```sh
$ go env 
```

### How to install
```sh
$ git clone git@github.com:igortin/aws_ec2_golang.git
$ go get -v .
$ cd aws_ec2_golang
$ change Makefile according to your OS (default is Linux)
$ make build
Done 
```

### How to use
Go to AWS Managment Console EC2 Tab. 
Tag EC2 instance: 
- key: role 
- Value: etcd
> Note: you can change in code and replace by os.Getenv("tagValue")
```sh
$ ./goEbsSnap
```
Check in AWS Managment ConsoleSnapshots Tab!
Done
