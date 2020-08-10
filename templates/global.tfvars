aws_region="{{ .configyaml.Cloud.Region }}"
aws_access_key_id="{{ .configyaml. }}"
aws_secret_access_key="jlInQW7HYrTgtkzQBs2X5FnL35yPVz1qnd/fvHj5"
vpc_id="vpc-0aef6256b40f30778"
rhcos_ami="ami-e06e5081"
target_environment="govcloud"
cluster_name="i"
name_domain="am.groot"
vpc_name="iamgroot"
cluster_domain="i.am.groot"
private_vpc_cidr="10.0.0.0/24"
subnet_list=["subnet-02bf7c8c69067b993", "subnet-0d75d5033bfc98414", "subnet-058e00cfdb41ca5ce"]

openshift:
  version: 4.5.4

cluster:
  target: govcloud
  vpc-name: iamgroot
  cluster-name: i
  base-domain: am.groot
  cluster-domain: i.am.groot
  ami-id: ami-e06e5081

cloud:
  provider: aws
  region: us-gov-west-1
  vpc-id: vpc-0aef6256b40f30778
  cidr-private: 10.0.0.0/24

subnets:
  private:
    - subnet-02bf7c8c69067b993
    - subnet-0d75d5033bfc98414
    - subnet-058e00cfdb41ca5ce

provider-auth:
  keys: true
  secret: XXXXXXXXXXXXXXXXXXXX
  key: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

redsord:
  enabled: false


