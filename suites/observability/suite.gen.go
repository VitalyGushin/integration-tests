// Code generated by gotestmd DO NOT EDIT.
package observability

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/spire"
)

type Suite struct {
	base.Suite
	spireSuite spire.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.spireSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
}
func (s *Suite) TestJaeger_and_prometheus() {
	r := s.Runner("../deployments-k8s/examples/observability/jaeger-and-prometheus")
	s.T().Cleanup(func() {
		r.Run(`rm -r example` + "\n" + `kubectl delete ns ${NAMESPACE}`)
		r.Run(`kubectl delete mutatingwebhookconfiguration --all` + "\n" + `kubectl delete ns nsm-system`)
		r.Run(`kubectl delete ns observability` + "\n" + `pkill -f "port-forward"`)
	})
	r.Run(`kubectl apply -k .`)
	r.Run(`kubectl wait -n observability --timeout=1m --for=condition=ready pod -l app=opentelemetry`)
	r.Run(`kubectl create ns nsm-system`)
	r.Run(`kubectl apply -k nsm-system`)
	r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
	r.Run(`NAMESPACE=($(kubectl create -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/1b21796131dec8aad9013703d14425a74dcb504b/examples/use-cases/namespace.yaml)[0])` + "\n" + `NAMESPACE=${NAMESPACE:10}`)
	r.Run(`NODE=($(kubectl get nodes -o go-template='{{range .items}}{{ if not .spec.taints  }}{{index .metadata.labels "kubernetes.io/hostname"}} {{end}}{{end}}')[0])`)
	r.Run(`mkdir example`)
	r.Run(`cat > example/kustomization.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: kustomize.config.k8s.io/v1beta1` + "\n" + `kind: Kustomization` + "\n" + `` + "\n" + `namespace: ${NAMESPACE}` + "\n" + `` + "\n" + `resources: ` + "\n" + `- client.yaml` + "\n" + `bases:` + "\n" + `- https://github.com/networkservicemesh/deployments-k8s/apps/nse-kernel?ref=1b21796131dec8aad9013703d14425a74dcb504b` + "\n" + `` + "\n" + `patchesStrategicMerge:` + "\n" + `- patch-nse.yaml` + "\n" + `EOF`)
	r.Run(`cat > example/client.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: v1` + "\n" + `kind: Pod` + "\n" + `metadata:` + "\n" + `  name: alpine` + "\n" + `  labels:` + "\n" + `    app: alpine    ` + "\n" + `  annotations:` + "\n" + `    networkservicemesh.io: kernel://icmp-responder/nsm-1` + "\n" + `spec:` + "\n" + `  containers:` + "\n" + `  - name: alpine` + "\n" + `    image: alpine:3.15.0` + "\n" + `    imagePullPolicy: IfNotPresent` + "\n" + `    stdin: true` + "\n" + `    tty: true` + "\n" + `  nodeSelector:` + "\n" + `    kubernetes.io/hostname: ${NODE}` + "\n" + `EOF`)
	r.Run(`cat > example/patch-nse.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: nse-kernel` + "\n" + `spec:` + "\n" + `  template:` + "\n" + `    spec:` + "\n" + `      containers:` + "\n" + `        - name: nse` + "\n" + `          env:` + "\n" + `            - name: NSM_CIDR_PREFIX` + "\n" + `              value: 172.16.1.100/31` + "\n" + `            - name: TELEMETRY` + "\n" + `              value: "true"` + "\n" + `      nodeSelector:` + "\n" + `        kubernetes.io/hostname: ${NODE}` + "\n" + `EOF`)
	r.Run(`kubectl apply -k example`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ${NAMESPACE}`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ${NAMESPACE}`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ${NAMESPACE} --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ${NAMESPACE} --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl exec ${NSC} -n ${NAMESPACE} -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec ${NSE} -n ${NAMESPACE} -- ping -c 4 172.16.1.101`)
	r.Run(`NSMGR=$(kubectl get pods -l app=nsmgr -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `NSMGR=${NSMGR:0:11}`)
	r.Run(`kubectl port-forward service/jaeger -n observability 16686:16686&` + "\n" + `kubectl port-forward service/prometheus -n observability 9090:9090&`)
	r.Run(`result=$(curl -X GET localhost:16686/api/traces?service=${NSMGR}&lookback=5m&limit=1)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -q "nsmgr"`)
	r.Run(`result=$(curl -X GET localhost:9090/api/v1/query?query=up)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -q "up"`)
}
