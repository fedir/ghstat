#!/bin/bash
echo "## CNCF cloud native projects statistics rating"
echo ""
./ghstat -r \
kubernetes/kubernetes,\
argoproj/argo-cd,\
fluxcd/flux2,\
helm/helm,\
prometheus/prometheus,\
grafana/grafana,\
envoyproxy/envoy,\
istio/istio,\
containerd/containerd,\
open-telemetry/opentelemetry-collector,\
jaegertracing/jaeger,\
fluent/fluentd,\
coredns/coredns,\
etcd-io/etcd,\
thanos-io/thanos,\
cortexproject/cortex,\
cri-o/cri-o,\
goharbor/harbor,\
linkerd/linkerd2,\
open-policy-agent/opa,\
rook/rook,\
cert-manager/cert-manager,\
crossplane/crossplane,\
cilium/cilium,\
falcosecurity/falco,\
kedacore/keda,\
dapr/dapr,\
knative/serving,\
kubeedge/kubeedge,\
kyverno/kyverno,\
spiffe/spire,\
vitessio/vitess,\
tikv/tikv,\
backstage/backstage,\
nats-io/nats-server,\
grafana/loki,\
grafana/mimir,\
longhorn/longhorn,\
strimzi/strimzi-kafka-operator,\
chaos-mesh/chaos-mesh,\
kubevirt/kubevirt,\
operator-framework/operator-sdk,\
VictoriaMetrics/VictoriaMetrics,\
VictoriaMetrics/VictoriaLogs,\
kubeflow/kubeflow,\
containers/podman,\
tektoncd/pipeline,\
buildpacks/pack,\
emissary-ingress/emissary,\
artifacthub/hub,\
openkruise/kruise,\
wasmCloud/wasmCloud \
-f stats/cncf_projects.csv -t tmp
echo "[Detailed CNCF projects statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/cncf_projects.csv)"
echo ""
