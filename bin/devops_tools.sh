#!/bin/bash
echo "## DevOps tools statistics rating"
echo ""
./ghstat -r \
jenkinsci/jenkins,\
concourse/concourse,\
woodpecker-ci/woodpecker,\
go-gitea/gitea,\
gogs/gogs,\
harness/gitness,\
spinnaker/spinnaker,\
hashicorp/terraform,\
opentofu/opentofu,\
pulumi/pulumi,\
gruntwork-io/terragrunt,\
hashicorp/packer,\
ansible/ansible,\
puppetlabs/puppet,\
saltstack/salt,\
chef/chef,\
hashicorp/vault,\
hashicorp/consul,\
hashicorp/nomad,\
traefik/traefik,\
Kong/kong,\
rancher/rancher,\
moby/moby,\
containers/buildah,\
containers/skopeo,\
GoogleContainerTools/kaniko,\
distribution/distribution,\
sonatype/nexus-public,\
aquasecurity/trivy,\
anchore/grype,\
anchore/syft,\
quay/clair,\
aquasecurity/kube-bench,\
bridgecrewio/checkov,\
hadolint/hadolint,\
oxsecurity/megalinter,\
grafana/k6,\
zabbix/zabbix,\
icinga/icinga2,\
netdata/netdata \
-f stats/devops_tools.csv -t tmp
echo "[Detailed DevOps tools statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/devops_tools.csv)"
echo ""
