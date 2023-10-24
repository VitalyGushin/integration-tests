// Code generated by gotestmd DO NOT EDIT.
package interdomain

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/dns"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/loadbalancer"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/nsm"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/spiffe_federation"
	"github.com/networkservicemesh/integration-tests/suites/spire/cluster1"
	"github.com/networkservicemesh/integration-tests/suites/spire/cluster2"
)

type Suite struct {
	base.Suite
	loadbalancerSuite      loadbalancer.Suite
	dnsSuite               dns.Suite
	cluster1Suite          cluster1.Suite
	cluster2Suite          cluster2.Suite
	spiffe_federationSuite spiffe_federation.Suite
	nsmSuite               nsm.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.loadbalancerSuite, &s.dnsSuite, &s.cluster1Suite, &s.cluster2Suite, &s.spiffe_federationSuite, &s.nsmSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
}
func (s *Suite) TestNsm_consul() {
	r := s.Runner("../deployments-k8s/examples/interdomain/nsm_consul")
	s.T().Cleanup(func() {
		r.Run(`pkill -f "port-forward"`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/server/counting_nsm.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/client/dashboard.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_consul/nse-auto-scale-client?ref=ecedf527284cbc3f57a32697fc0027c050db0547` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_consul/nse-auto-scale-server?ref=ecedf527284cbc3f57a32697fc0027c050db0547` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/service.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/server/counting_service.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/netsvc.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete pods --all`)
		r.Run(`consul-k8s uninstall --kubeconfig=$KUBECONFIG2 -auto-approve=true -wipe-data=true`)
	})
	r.Run(`curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -`)
	r.Run(`sudo apt-add-repository -y "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"`)
	r.Run(`sudo apt-get update && sudo apt-get install -y consul-k8s=0.48.0-1`)
	r.Run(`consul-k8s version`)
	r.Run(`consul-k8s install -config-file=helm-consul-values.yaml -set global.image=hashicorp/consul:1.12.0 -auto-approve --kubeconfig=$KUBECONFIG2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/server/counting_service.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/server/counting.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/netsvc.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_consul/nse-auto-scale-client?ref=ecedf527284cbc3f57a32697fc0027c050db0547` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_consul/nse-auto-scale-server?ref=ecedf527284cbc3f57a32697fc0027c050db0547`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/service.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/client/dashboard.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --timeout=10m --for=condition=ready pod -l app=dashboard-nsc`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pod/dashboard-nsc -c cmd-nsc -- apk add curl`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pod/dashboard-nsc -c cmd-nsc -- curl counting:9001`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 port-forward pod/dashboard-nsc 9002:9002 &`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete deploy counting`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_consul/server/counting_nsm.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --timeout=5m --for=condition=ready pod -l app=counting`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pod/dashboard-nsc -c cmd-nsc -- curl counting:9001`)
}
func (s *Suite) TestNsm_consul_vl3() {
	r := s.Runner("../deployments-k8s/examples/interdomain/nsm_consul_vl3")
	s.T().Cleanup(func() {
		r.Run(`pkill -f "port-forward"` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 delete -n ns-nsm-consul-vl3 -k ./cluster1` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -n ns-nsm-consul-vl3 -k ./cluster2`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k ./cluster1` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 apply -k ./cluster2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=nse-vl3-vpp -n ns-nsm-consul-vl3` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=vl3-ipam -n ns-nsm-consul-vl3` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l name=control-plane -n ns-nsm-consul-vl3` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod counting -n ns-nsm-consul-vl3` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=5m pod dashboard -n ns-nsm-consul-vl3`)
	r.Run(`export CP=$(kubectl --kubeconfig=$KUBECONFIG1 get pods -n ns-nsm-consul-vl3 -l name=control-plane --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`ENCRYPTION_KEY=$(kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec ${CP} -c ubuntu -- /bin/sh -c 'consul keygen')`)
	r.Run(`CP_IP_VL3_ADDRESS=$(kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec ${CP} -c ubuntu -- ifconfig nsm-1 | grep -Eo 'inet addr:[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 11-)`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec ${CP} -c ubuntu -- consul tls ca create`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec ${CP} -c ubuntu -- consul tls cert create -server -dc dc1`)
	r.Run(`cat > consul.hcl <<EOF` + "\n" + `data_dir = "/opt/consul"` + "\n" + `datacenter = "dc1"` + "\n" + `encrypt = "${ENCRYPTION_KEY}"` + "\n" + `tls {` + "\n" + `  defaults {` + "\n" + `    ca_file = "consul-agent-ca.pem"` + "\n" + `    cert_file = "dc1-server-consul-0.pem"` + "\n" + `    key_file = "dc1-server-consul-0-key.pem"` + "\n" + `    verify_incoming = true` + "\n" + `    verify_outgoing = true` + "\n" + `  }` + "\n" + `  internal_rpc {` + "\n" + `    verify_server_hostname = true` + "\n" + `  }` + "\n" + `}` + "\n" + `auto_encrypt {` + "\n" + `  allow_tls = true` + "\n" + `}` + "\n" + `acl {` + "\n" + `  enabled = true` + "\n" + `  default_policy = "allow"` + "\n" + `  enable_token_persistence = true` + "\n" + `}` + "\n" + `EOF`)
	r.Run(`cat > server.hcl <<EOF` + "\n" + `server = true` + "\n" + `bootstrap_expect = 1` + "\n" + `bind_addr = "${CP_IP_VL3_ADDRESS}"` + "\n" + `connect {` + "\n" + `  enabled = true` + "\n" + `}` + "\n" + `` + "\n" + `addresses {` + "\n" + `  grpc = "127.0.0.1"` + "\n" + `}` + "\n" + `ports {` + "\n" + `  grpc  = 8502` + "\n" + `}` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 cp consul.hcl ns-nsm-consul-vl3/${CP}:/consul/config/` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 cp server.hcl ns-nsm-consul-vl3/${CP}:/consul/config/`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec ${CP} -c ubuntu -- consul validate /consul/config/`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec ${CP} -c ubuntu -- /bin/sh -c 'consul agent -config-dir=/consul/config/  1>/dev/null 2>&1 &'`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec ${CP} -c ubuntu -- consul members`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- /bin/bash -c  'apt update & apt upgrade -y'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- apt-get install curl gnupg sudo lsb-release net-tools iproute2 apt-utils systemctl -y`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- /bin/bash -c 'curl --fail --silent --show-error --location https://apt.releases.hashicorp.com/gpg | \` + "\n" + `      gpg --dearmor | \` + "\n" + `      sudo dd of=/usr/share/keyrings/hashicorp-archive-keyring.gpg '`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- /bin/bash -c 'echo "deb [arch=amd64 signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \` + "\n" + ` sudo tee -a /etc/apt/sources.list.d/hashicorp.list'`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo apt-get update`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo apt-get install consul=1.12.0-1`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- consul version`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- /bin/bash -c 'curl -L https://func-e.io/install.sh | bash -s -- -b /usr/local/bin'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- /bin/bash -c 'export FUNC_E_PLATFORM=linux/amd64'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- /bin/bash -c 'func-e use 1.22.2'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- /bin/bash -c 'sudo cp ~/.func-e/versions/1.22.2/bin/envoy /usr/bin/'`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- envoy --version`)
	r.Run(`COUNTING_IP_VL3_ADDRESS=$(kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- ifconfig nsm-1 | grep -Eo 'inet [0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 6-)`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 cp  ns-nsm-consul-vl3/${CP}:consul-agent-ca.pem consul-agent-ca.pem` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 cp  ns-nsm-consul-vl3/${CP}:consul-agent-ca-key.pem consul-agent-ca-key.pem` + "\n" + `` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 cp consul-agent-ca.pem ns-nsm-consul-vl3/counting:/etc/consul.d` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 cp consul-agent-ca-key.pem ns-nsm-consul-vl3/counting:/etc/consul.d`)
	r.Run(`cat > consul-counting.hcl <<EOF` + "\n" + `data_dir = "/opt/consul"` + "\n" + `encrypt = "${ENCRYPTION_KEY}"` + "\n" + `tls {` + "\n" + `  defaults {` + "\n" + `    ca_file = "/etc/consul.d/consul-agent-ca.pem"` + "\n" + `    verify_incoming = false` + "\n" + `    verify_outgoing = true` + "\n" + `  }` + "\n" + `  internal_rpc {` + "\n" + `    verify_server_hostname = true` + "\n" + `  }` + "\n" + `}` + "\n" + `auto_encrypt {` + "\n" + `  tls = true` + "\n" + `}` + "\n" + `acl {` + "\n" + `  enabled = true` + "\n" + `  default_policy = "allow"` + "\n" + `  enable_token_persistence = true` + "\n" + `}` + "\n" + `bind_addr = "${COUNTING_IP_VL3_ADDRESS}"` + "\n" + `connect {` + "\n" + `  enabled = true` + "\n" + `}` + "\n" + `addresses {` + "\n" + `  grpc = "127.0.0.1"` + "\n" + `}` + "\n" + `ports {` + "\n" + `  grpc  = 8502` + "\n" + `}` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 cp consul-counting.hcl ns-nsm-consul-vl3/counting:/etc/consul.d/`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo consul validate /etc/consul.d/`)
	r.Run(`cat > consul.service <<EOF` + "\n" + `[Unit]` + "\n" + `Description=Consul` + "\n" + `Documentation=https://www.consul.io/` + "\n" + `` + "\n" + `[Service]` + "\n" + `ExecStart=/usr/bin/consul agent -join ${CP_IP_VL3_ADDRESS} -config-dir=/etc/consul.d/ ` + "\n" + `ExecReload=/bin/kill -HUP $MAINPID` + "\n" + `LimitNOFILE=65536` + "\n" + `` + "\n" + `[Install]` + "\n" + `WantedBy=multi-user.target` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 cp consul.service ns-nsm-consul-vl3/counting:/etc/systemd/system/consul.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo systemctl daemon-reload`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo systemctl start consul.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo systemctl enable consul.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- consul members`)
	r.Run(`cat > counting.hcl <<EOF` + "\n" + `service {` + "\n" + `  name = "counting"` + "\n" + `  id = "counting-1"` + "\n" + `  port = 9001` + "\n" + `` + "\n" + `  connect {` + "\n" + `    sidecar_service {}` + "\n" + `  }` + "\n" + `` + "\n" + `  check {` + "\n" + `    id       = "counting-check"` + "\n" + `    http     = "http://localhost:9001/health"` + "\n" + `    method   = "GET"` + "\n" + `    interval = "1s"` + "\n" + `    timeout  = "1s"` + "\n" + `  }` + "\n" + `}` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- mkdir service` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 cp counting.hcl ns-nsm-consul-vl3/counting:/service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- consul services register /service/counting.hcl`)
	r.Run(`cat > consul-envoy.service <<EOF` + "\n" + `[Unit]` + "\n" + `Description=Consul` + "\n" + `Documentation=https://www.consul.io/` + "\n" + `` + "\n" + `[Service]` + "\n" + `ExecStart=/usr/bin/consul connect envoy -sidecar-for counting-1 -admin-bind localhost:19001 ` + "\n" + `ExecReload=/bin/kill -HUP $MAINPID` + "\n" + `LimitNOFILE=65536` + "\n" + `` + "\n" + `[Install]` + "\n" + `WantedBy=multi-user.target` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 cp consul-envoy.service ns-nsm-consul-vl3/counting:/etc/systemd/system/consul-envoy.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo systemctl daemon-reload ` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo systemctl start consul-envoy.service ` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n ns-nsm-consul-vl3 exec counting -c ubuntu -- sudo systemctl enable consul-envoy.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c  'apt update & apt upgrade -y'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- apt-get install curl gnupg sudo lsb-release net-tools iproute2 apt-utils systemctl -y`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c 'curl --fail --silent --show-error --location https://apt.releases.hashicorp.com/gpg | \` + "\n" + `      gpg --dearmor | \` + "\n" + `      sudo dd of=/usr/share/keyrings/hashicorp-archive-keyring.gpg '`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c 'echo "deb [arch=amd64 signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \` + "\n" + ` sudo tee -a /etc/apt/sources.list.d/hashicorp.list'`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c 'sudo apt-get update & sudo apt-get install consul=1.12.0-1'`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- consul version`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c 'curl -L https://func-e.io/install.sh | bash -s -- -b /usr/local/bin'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c 'export FUNC_E_PLATFORM=linux/amd64'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c 'func-e use 1.22.2'` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- /bin/bash -c 'sudo cp ~/.func-e/versions/1.22.2/bin/envoy /usr/bin/'`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- envoy --version`)
	r.Run(`DASHBOARD_IP_VL3_ADDRESS=$(kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- ifconfig nsm-1 | grep -Eo 'inet [0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 6-)`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 cp consul-agent-ca.pem ns-nsm-consul-vl3/dashboard:/etc/consul.d` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 cp consul-agent-ca-key.pem ns-nsm-consul-vl3/dashboard:/etc/consul.d`)
	r.Run(`cat > consul-dashboard.hcl <<EOF` + "\n" + `encrypt = "${ENCRYPTION_KEY}"` + "\n" + `data_dir = "/opt/consul"` + "\n" + `tls {` + "\n" + `  defaults {` + "\n" + `    ca_file = "/etc/consul.d/consul-agent-ca.pem"` + "\n" + `    verify_incoming = false` + "\n" + `    verify_outgoing = true` + "\n" + `  }` + "\n" + `  internal_rpc {` + "\n" + `    verify_server_hostname = true` + "\n" + `  }` + "\n" + `}` + "\n" + `datacenter = "dc1"` + "\n" + `auto_encrypt {` + "\n" + `  tls = true` + "\n" + `}` + "\n" + `acl {` + "\n" + `  enabled = true` + "\n" + `  default_policy = "allow"` + "\n" + `  enable_token_persistence = true` + "\n" + `}` + "\n" + `bind_addr = "${DASHBOARD_IP_VL3_ADDRESS}"` + "\n" + `connect {` + "\n" + `  enabled = true` + "\n" + `}` + "\n" + `addresses {` + "\n" + `  grpc = "127.0.0.1"` + "\n" + `}` + "\n" + `ports {` + "\n" + `  grpc  = 8502` + "\n" + `}` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 cp consul-dashboard.hcl ns-nsm-consul-vl3/dashboard:/etc/consul.d/`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- sudo consul validate /etc/consul.d/`)
	r.Run(`cat > consul.service <<EOF` + "\n" + `[Unit]` + "\n" + `Description=Consul` + "\n" + `Documentation=https://www.consul.io/` + "\n" + `` + "\n" + `[Service]` + "\n" + `ExecStart=/usr/bin/consul agent -join ${CP_IP_VL3_ADDRESS} -config-dir=/etc/consul.d/` + "\n" + `ExecReload=/bin/kill -HUP $MAINPID` + "\n" + `LimitNOFILE=65536` + "\n" + `` + "\n" + `[Install]` + "\n" + `WantedBy=multi-user.target` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 cp consul.service ns-nsm-consul-vl3/dashboard:/etc/systemd/system/consul.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- sudo systemctl daemon-reload ` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- sudo systemctl start consul.service ` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- sudo systemctl enable consul.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- consul members`)
	r.Run(`cat > dashboard.hcl <<EOF` + "\n" + `service {` + "\n" + `  name = "dashboard"` + "\n" + `  port = 9002` + "\n" + `` + "\n" + `  connect {` + "\n" + `    sidecar_service {` + "\n" + `      proxy {` + "\n" + `        upstreams = [` + "\n" + `          {` + "\n" + `            destination_name = "counting"` + "\n" + `            local_bind_port  = 5000` + "\n" + `          }` + "\n" + `        ]` + "\n" + `      }` + "\n" + `    }` + "\n" + `  }` + "\n" + `` + "\n" + `  check {` + "\n" + `    id       = "dashboard-check"` + "\n" + `    http     = "http://localhost:9002/health"` + "\n" + `    method   = "GET"` + "\n" + `    interval = "1s"` + "\n" + `    timeout  = "1s"` + "\n" + `  }` + "\n" + `}` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- mkdir service` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 cp dashboard.hcl ns-nsm-consul-vl3/dashboard:/service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- consul services register /service/dashboard.hcl`)
	r.Run(`cat > consul-envoy.service <<EOF` + "\n" + `[Unit]` + "\n" + `Description=Consul` + "\n" + `Documentation=https://www.consul.io/` + "\n" + `` + "\n" + `[Service]` + "\n" + `ExecStart=/usr/bin/consul connect envoy -sidecar-for dashboard` + "\n" + `ExecReload=/bin/kill -HUP $MAINPID` + "\n" + `LimitNOFILE=65536` + "\n" + `` + "\n" + `[Install]` + "\n" + `WantedBy=multi-user.target` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 cp consul-envoy.service ns-nsm-consul-vl3/dashboard:/etc/systemd/system/consul-envoy.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- sudo systemctl daemon-reload ` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- sudo systemctl start consul-envoy.service ` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -c ubuntu -- sudo systemctl enable consul-envoy.service`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 port-forward dashboard 9002:9002 &`)
	r.Run(`result=$(kubectl --kubeconfig=$KUBECONFIG2 -n ns-nsm-consul-vl3 exec dashboard -- curl localhost:5000)` + "\n" + `echo ${result} | grep  -o '\"count\":[1-9]\d*'`)
}
func (s *Suite) TestNsm_istio() {
	r := s.Runner("../deployments-k8s/examples/interdomain/nsm_istio")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_istio/greeting/server.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_istio/nse-auto-scale?ref=ecedf527284cbc3f57a32697fc0027c050db0547` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_istio/greeting/client.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_istio/netsvc.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete ns istio-system` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 label namespace default istio-injection-` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete pods --all`)
	})
	r.Run(`curl -sL https://istio.io/downloadIstioctl | sh -` + "\n" + `export PATH=$PATH:$HOME/.istioctl/bin` + "\n" + `istioctl install --readiness-timeout 10m0s --set profile=minimal -y --kubeconfig=$KUBECONFIG2` + "\n" + `istioctl --kubeconfig=$KUBECONFIG2 proxy-status`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_istio/netsvc.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_istio/greeting/client.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_istio/nse-auto-scale?ref=ecedf527284cbc3f57a32697fc0027c050db0547`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 label namespace default istio-injection=enabled` + "\n" + `` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_istio/greeting/server.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --timeout=5m --for=condition=ready pod -l app=alpine`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec deploy/alpine -c cmd-nsc -- apk add curl`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec deploy/alpine -c cmd-nsc -- curl -s greeting.default:9080 | grep -o "hello world from istio"`)
}
func (s *Suite) TestNsm_kuma_universal_vl3() {
	r := s.Runner("../deployments-k8s/examples/interdomain/nsm_kuma_universal_vl3")
	s.T().Cleanup(func() {
		r.Run(`pkill -f "port-forward"` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 delete ns kuma-system kuma-demo ns-dns-vl3` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete ns kuma-demo` + "\n" + `rm tls.crt tls.key ca.crt kustomization.yaml control-plane.yaml` + "\n" + `rm -rf kuma-1.7.0`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_kuma_universal_vl3/vl3-dns?ref=ecedf527284cbc3f57a32697fc0027c050db0547` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n ns-dns-vl3 wait --for=condition=ready --timeout=5m pod -l app=vl3-ipam`)
	r.Run(`curl -L https://kuma.io/installer.sh | VERSION=1.7.0 ARCH=amd64 bash -` + "\n" + `export PATH=$PWD/kuma-1.7.0/bin:$PATH`)
	r.Run(`kumactl generate tls-certificate --hostname=control-plane-kuma.my-vl3-network --hostname=kuma-control-plane.kuma-system.svc --type=server --key-file=./tls.key --cert-file=./tls.crt`)
	r.Run(`cp ./tls.crt ./ca.crt`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_kuma_universal_vl3/namespace.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 create secret generic general-tls-certs --namespace=kuma-system --from-file=./tls.key --from-file=./tls.crt --from-file=./ca.crt`)
	r.Run(`kumactl install control-plane --tls-general-secret=general-tls-certs --tls-general-ca-bundle=$(cat ./ca.crt | base64) > control-plane.yaml`)
	r.Run(`cat > kustomization.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: kustomize.config.k8s.io/v1beta1` + "\n" + `kind: Kustomization` + "\n" + `` + "\n" + `resources:` + "\n" + `- control-plane.yaml` + "\n" + `` + "\n" + `patchesStrategicMerge:` + "\n" + `- https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_kuma_universal_vl3/patch-control-plane.yaml` + "\n" + `EOF`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k .`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_kuma_universal_vl3/demo-redis.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 -n kuma-demo wait --for=condition=ready --timeout=5m pod -l app=redis`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ecedf527284cbc3f57a32697fc0027c050db0547/examples/interdomain/nsm_kuma_universal_vl3/demo-app.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 -n kuma-demo wait --for=condition=ready --timeout=5m pod -l app=demo-app`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 port-forward svc/demo-app -n kuma-demo 8081:5000 &`)
	r.Run(`response=$(curl -X POST localhost:8081/increment)` + "\n" + `echo $response | grep '"err":null'`)
}
