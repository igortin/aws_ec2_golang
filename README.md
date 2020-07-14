# aws_ec2_golang
It is a repository for example, include code for creating EBS snapshots for AWS EC2 instances based on tags. 
  - Discovery EC2 instances based on tags - role:etcd
  - Parse []*ec2.DescribeInstancesOutput and create ec2 objects based on your custom structure Ec2object
  - Create EBS Snapshot and tag it for non-root Volumes without stop EC2 instance

> Note
> Ec2 root volumes should be created in EC2 stopped state

### Prerequisites 
- go environment shoukld configured 

### How to install
> Go to AWS Managment Console EC2 Tab
> tag EC2 instance below: 
> key: role 
> Value: etcd
> (you can change in code and replace by os.Getenv("tagValue"))

```sh
$ git clone git@github.com:igortin/aws_ec2_golang.git
$ go get -v .
$ cd aws_ec2_golang
$ change Makefile according to your OS (default is Linux)
$ make build
$ ./goEbsSnap
Done! 
```
