// Code generated by gotestmd DO NOT EDIT.
package remote_nsmgr_restart

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/basic"
)

type Suite struct {
	base.Suite
	basicSuite basic.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.basicSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/heal/remote-nsmgr-restart")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ${NAMESPACE}`)
	})
	r.Run(`NAMESPACE=($(kubectl create -f ../namespace.yaml)[0])` + "\n" + `NAMESPACE=${NAMESPACE:10}`)
	r.Run(`kubectl exec -n spire spire-server-0 -- \` + "\n" + `/opt/spire/bin/spire-server entry create \` + "\n" + `-spiffeID spiffe://example.org/ns/${NAMESPACE}/sa/default \` + "\n" + `-parentID spiffe://example.org/ns/spire/sa/spire-agent \` + "\n" + `-selector k8s:ns:${NAMESPACE} \` + "\n" + `-selector k8s:sa:default`)
	r.Run(`NODES=($(kubectl get nodes -o go-template='{{range .items}}{{ if not .spec.taints  }}{{index .metadata.labels "kubernetes.io/hostname"}} {{end}}{{end}}'))`)
	r.Run(`cat > kustomization.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: kustomize.config.k8s.io/v1beta1` + "\n" + `kind: Kustomization` + "\n" + `` + "\n" + `namespace: ${NAMESPACE}` + "\n" + `` + "\n" + `bases:` + "\n" + `- ../../../apps/nsc-kernel` + "\n" + `- ../../../apps/nse-kernel` + "\n" + `` + "\n" + `patchesStrategicMerge:` + "\n" + `- patch-nsc.yaml` + "\n" + `- patch-nse.yaml` + "\n" + `EOF`)
	r.Run(`cat > patch-nsc.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: nsc-kernel` + "\n" + `spec:` + "\n" + `  template:` + "\n" + `    spec:` + "\n" + `      containers:` + "\n" + `        - name: nsc` + "\n" + `          env:` + "\n" + `            - name: NSM_NETWORK_SERVICES` + "\n" + `              value: kernel://icmp-responder/nsm-1` + "\n" + `      nodeSelector:` + "\n" + `        kubernetes.io/hostname: ${NODES[0]}` + "\n" + `EOF`)
	r.Run(`cat > patch-nse.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: nse-kernel` + "\n" + `spec:` + "\n" + `  template:` + "\n" + `    spec:` + "\n" + `      containers:` + "\n" + `        - name: nse` + "\n" + `          env:` + "\n" + `            - name: NSE_CIDR_PREFIX` + "\n" + `              value: 172.16.1.100/30` + "\n" + `      nodeSelector:` + "\n" + `        kubernetes.io/hostname: ${NODES[1]}` + "\n" + `EOF`)
	r.Run(`kubectl apply -k .`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-kernel -n ${NAMESPACE}`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ${NAMESPACE}`)
	r.Run(`NSC=$(kubectl get pods -l app=nsc-kernel -n ${NAMESPACE} --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ${NAMESPACE} --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl exec ${NSC} -n ${NAMESPACE} -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec ${NSE} -n ${NAMESPACE} -- ping -c 4 172.16.1.101`)
	r.Run(`NSMGR=$(kubectl get pods -l app=nsmgr --field-selector spec.nodeName==${NODES[1]} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod ${NSMGR} -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsmgr --field-selector spec.nodeName==${NODES[1]} -n nsm-system`)
	r.Run(`kubectl exec ${NSC} -n ${NAMESPACE} -- ping -c 4 172.16.1.102`)
	r.Run(`kubectl exec ${NSE} -n ${NAMESPACE} -- ping -c 4 172.16.1.103`)
}
func (s *Suite) Test() {}
