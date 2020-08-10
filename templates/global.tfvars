aws_region="{{ .configyaml.Cloud.Region }}"
aws_access_key_id="{{ .configyaml.Auth.Key }}"
aws_secret_access_key="{{ .configyaml.Auth.Secret }}"
vpc_id="{{ .configyaml.Cloud.VpcId }}"
rhcos_ami="{{ .configyaml.Cluster.AmiId }}"
target_environment="{{ .configyaml.Cluster.Target }}"
cluster_name="{{ .configyaml.Cluster.ClusterName }}"
name_domain="{{ .configyaml.Cluster.BaseDomain }}"
vpc_name="{{ .configyaml.Cluster.VpcName }}"
cluster_domain="{{ .configyaml.Cluster.BaseDomain }}"
private_vpc_cidr="{{ .configyaml.Cloud.VpcId }}"
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


