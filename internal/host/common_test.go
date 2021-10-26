package host

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	"github.com/openshift/assisted-service/models"
)

type hostNetProfile struct {
	role     models.HostRole
	hostname string
	ip       string
}

var _ = Describe("GetHostnameAndEffectiveRoleByIP", func() {

	hostRolesIpv4 := []hostNetProfile{{role: models.HostRoleMaster, hostname: "master-0", ip: "1.2.3.1"}, {role: models.HostRoleWorker, hostname: "worker-0", ip: "1.2.3.2"}}
	hostrolesIpv6 := []hostNetProfile{{role: models.HostRoleMaster, hostname: "master-1", ip: "1001:db8::11"}, {role: models.HostRoleWorker, hostname: "worker-1", ip: "1001:db8::12"}}
	clusterID := strfmt.UUID(uuid.New().String())
	infraEnvID := strfmt.UUID(uuid.New().String())

	Context("resolves hostname and role based on IP", func() {

		testCases := []struct {
			name             string
			hostRolesIpv4    []hostNetProfile
			hostRolesIpv6    []hostNetProfile
			targetIP         string
			expectedRole     models.HostRole
			expectedHostname string
			expectedError    error
		}{
			{
				name:             "resolves correctly when there are only IPv4 interfaces",
				hostRolesIpv4:    hostRolesIpv4,
				targetIP:         "1.2.3.1",
				expectedRole:     models.HostRoleMaster,
				expectedHostname: "master-0",
			},
			{
				name:             "resolves correctly when there are only IPv6 interfaces",
				hostRolesIpv6:    hostrolesIpv6,
				targetIP:         "1001:db8::12",
				expectedRole:     models.HostRoleWorker,
				expectedHostname: "worker-1",
			},
			{
				name:             "resolves correctly when there is a mix of IPv4 and IPv6 interfaces",
				hostRolesIpv4:    hostRolesIpv4,
				hostRolesIpv6:    hostrolesIpv6,
				targetIP:         "1001:db8::11",
				expectedRole:     models.HostRoleMaster,
				expectedHostname: "master-1",
			},
			{
				name:          "unable to resolve when there is a mix of IPv4 and IPv6 interfaces and no match is found",
				hostRolesIpv4: hostRolesIpv4,
				hostRolesIpv6: hostrolesIpv6,
				targetIP:      "1001:db8::30",
				expectedError: fmt.Errorf("host with IP %s not found in inventory", "1001:db8::30")},
		}

		for i := range testCases {
			test := testCases[i]
			It(test.name, func() {
				hosts := []*models.Host{}
				for _, v := range test.hostRolesIpv4 {
					netAddr := common.NetAddress{Hostname: v.hostname, IPv4Address: []string{fmt.Sprintf("%s/%d", v.ip, 24)}}
					h := hostutil.GenerateTestHostWithNetworkAddress(strfmt.UUID(uuid.New().String()), infraEnvID, clusterID, v.role, models.HostStatusKnown, netAddr)
					hosts = append(hosts, h)
				}
				for _, v := range test.hostRolesIpv6 {
					netAddr := common.NetAddress{Hostname: v.hostname, IPv6Address: []string{fmt.Sprintf("%s/%d", v.ip, 120)}}
					h := hostutil.GenerateTestHostWithNetworkAddress(strfmt.UUID(uuid.New().String()), infraEnvID, clusterID, v.role, models.HostStatusKnown, netAddr)
					hosts = append(hosts, h)
				}
				hostname, role, err := GetHostnameAndEffectiveRoleByIP(test.targetIP, hosts)
				if test.expectedError != nil {
					Expect(err).To(Equal(test.expectedError))
				} else {
					Expect(hostname).To(Equal(test.expectedHostname))
					Expect(role).To(Equal(test.expectedRole))
				}

			})
		}
	})

})

var _ = Describe("GetInterfaceByName", func() {
	interface1Name := "eth1"
	interface2Name := "eth2"
	interface1 := models.Interface{Name: interface1Name}
	interface2 := models.Interface{Name: interface2Name}

	Context("resolves interface from inventory based on name", func() {

		testCases := []struct {
			name              string
			interfaceName     string
			inventory         *models.Inventory
			expectedInterface *models.Interface
			expectedError     error
		}{
			{
				name:              "resolves interface correctly by name",
				interfaceName:     interface1Name,
				inventory:         &models.Inventory{Interfaces: []*models.Interface{&interface1, &interface2}},
				expectedInterface: &interface1,
			},
			{
				name:          "Returns error when interface list is empty",
				interfaceName: interface1Name,
				inventory:     &models.Inventory{Interfaces: []*models.Interface{}},
				expectedError: fmt.Errorf(noInterfaceForNameErrorTemplate, interface1Name),
			},
			{
				name:          "Returns error when interface list is nil",
				interfaceName: interface1Name,
				inventory:     nil,
				expectedError: fmt.Errorf(nilInventoryErrorTemplate),
			},
		}

		for i := range testCases {
			test := testCases[i]
			It(test.name, func() {
				resolvedInterface, err := GetInterfaceByName(test.interfaceName, test.inventory)
				if test.expectedError != nil {
					Expect(err).To(Equal(test.expectedError))
				} else {
					Expect(resolvedInterface).To(Equal(test.expectedInterface))
				}

			})
		}
	})

})

var _ = Describe("GetInterfaceByIp", func() {
	ipv4Cidr := "192.168.1.1"
	ipv6Cidr := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	interface1 := models.Interface{IPV4Addresses: []string{fmt.Sprintf("%s/%d", ipv4Cidr, 24)}}
	interface2 := models.Interface{IPV6Addresses: []string{fmt.Sprintf("%s/%d", ipv6Cidr, 120)}}

	Context("resolves interface from inventory based on ip", func() {

		testCases := []struct {
			name              string
			interfaceIp       string
			inventory         *models.Inventory
			expectedInterface *models.Interface
			expectedError     error
		}{
			{
				name:              "resolves interface correctly for ipv4",
				interfaceIp:       ipv4Cidr,
				inventory:         &models.Inventory{Interfaces: []*models.Interface{&interface1, &interface2}},
				expectedInterface: &interface1,
			},
			{
				name:              "resolves interface correctly for ipv6",
				interfaceIp:       ipv6Cidr,
				inventory:         &models.Inventory{Interfaces: []*models.Interface{&interface1, &interface2}},
				expectedInterface: &interface2,
			},
			{
				name:          "Returns error when inventory is nil",
				interfaceIp:   ipv4Cidr,
				inventory:     nil,
				expectedError: fmt.Errorf(nilInventoryErrorTemplate),
			},
		}

		for i := range testCases {
			test := testCases[i]
			It(test.name, func() {
				resolvedInterface, err := GetInterfaceByIp(test.interfaceIp, test.inventory)
				if test.expectedError != nil {
					Expect(err).To(Equal(test.expectedError))
				} else {
					Expect(resolvedInterface).To(Equal(test.expectedInterface))
				}

			})
		}
	})

})

var _ = Describe("GetHostByIP", func() {
	ipv4Cidr := "192.168.1.1"
	ipv6Cidr := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	interface1 := models.Interface{IPV4Addresses: []string{fmt.Sprintf("%s/%d", ipv4Cidr, 24)}}
	interface2 := models.Interface{IPV6Addresses: []string{fmt.Sprintf("%s/%d", ipv6Cidr, 120)}}

	inv1 := models.Inventory{Interfaces: []*models.Interface{&interface1}}
	inv2 := models.Inventory{Interfaces: []*models.Interface{&interface2}}

	marshalledInv1, err := common.MarshalInventory(&inv1)
	Expect(err).To(BeNil())
	marshalledInv2, err := common.MarshalInventory(&inv2)
	Expect(err).To(BeNil())

	host1 := models.Host{Inventory: marshalledInv1}
	host2 := models.Host{Inventory: marshalledInv2}

	Context("resolves a specific host from a list based on ip", func() {

		testCases := []struct {
			name          string
			interfaceIp   string
			hosts         []*models.Host
			expectedHost  *models.Host
			expectedError error
		}{
			{
				name:         "resolves host correctly for ipv4",
				interfaceIp:  ipv4Cidr,
				hosts:        []*models.Host{&host1, &host2},
				expectedHost: &host1,
			},
			{
				name:         "resolves host correctly for ipv6",
				interfaceIp:  ipv6Cidr,
				hosts:        []*models.Host{&host1, &host2},
				expectedHost: &host2,
			},
			{
				name:          "Returns error when hosts list is nil",
				interfaceIp:   ipv4Cidr,
				hosts:         nil,
				expectedError: fmt.Errorf(nilInventoryErrorTemplate),
			},
		}

		for i := range testCases {
			test := testCases[i]
			It(test.name, func() {
				resolvedHost, err := GetHostByIP(test.interfaceIp, test.hosts)
				if test.expectedError != nil {
					Expect(err).To(Equal(test.expectedError))
				} else {
					Expect(resolvedHost).To(Equal(test.expectedHost))
				}
			})
		}
	})

})
