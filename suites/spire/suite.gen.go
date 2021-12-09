// Code generated by gotestmd DO NOT EDIT.
package spire

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
)

type Suite struct {
	base.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/spire")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete crd spiffeids.spiffeid.spiffe.io` + "\n" + `kubectl delete ns spire`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/spire?ref=7dc4670c03f4db100fee6f9759b69fb29c8b21a4`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-agent`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-server`)
}
func (s *Suite) Test() {}
