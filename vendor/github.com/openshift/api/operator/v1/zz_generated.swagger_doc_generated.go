package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_GenerationStatus = map[string]string{
	"":               "GenerationStatus keeps track of the generation for a given resource so that decisions about forced updates can be made.",
	"group":          "group is the group of the thing you're tracking",
	"resource":       "resource is the resource type of the thing you're tracking",
	"namespace":      "namespace is where the thing you're tracking is",
	"name":           "name is the name of the thing you're tracking",
	"lastGeneration": "lastGeneration is the last generation of the workload controller involved",
	"hash":           "hash is an optional field set for resources without generation that are content sensitive like secrets and configmaps",
}

func (GenerationStatus) SwaggerDoc() map[string]string {
	return map_GenerationStatus
}

var map_MyOperatorResource = map[string]string{
	"": "MyOperatorResource is an example operator configuration type",
}

func (MyOperatorResource) SwaggerDoc() map[string]string {
	return map_MyOperatorResource
}

var map_NodeStatus = map[string]string{
	"":                         "NodeStatus provides information about the current state of a particular node managed by this operator.",
	"nodeName":                 "nodeName is the name of the node",
	"currentRevision":          "currentRevision is the generation of the most recently successful deployment",
	"targetRevision":           "targetRevision is the generation of the deployment we're trying to apply",
	"lastFailedRevision":       "lastFailedRevision is the generation of the deployment we tried and failed to deploy.",
	"lastFailedRevisionErrors": "lastFailedRevisionErrors is a list of the errors during the failed deployment referenced in lastFailedRevision",
}

func (NodeStatus) SwaggerDoc() map[string]string {
	return map_NodeStatus
}

var map_OperatorCondition = map[string]string{
	"": "OperatorCondition is just the standard condition fields.",
}

func (OperatorCondition) SwaggerDoc() map[string]string {
	return map_OperatorCondition
}

var map_OperatorSpec = map[string]string{
	"":                           "OperatorSpec contains common fields operators need.  It is intended to be anonymous included inside of the Spec struct for your particular operator.",
	"managementState":            "managementState indicates whether and how the operator should manage the component",
	"logLevel":                   "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.",
	"operatorLogLevel":           "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.",
	"unsupportedConfigOverrides": "unsupportedConfigOverrides holds a sparse config that will override any previously set options.  It only needs to be the fields to override it will end up overlaying in the following order: 1. hardcoded defaults 2. observedConfig 3. unsupportedConfigOverrides",
	"observedConfig":             "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
}

func (OperatorSpec) SwaggerDoc() map[string]string {
	return map_OperatorSpec
}

var map_OperatorStatus = map[string]string{
	"observedGeneration": "observedGeneration is the last generation change you've dealt with",
	"conditions":         "conditions is a list of conditions and their status",
	"version":            "version is the level this availability applies to",
	"readyReplicas":      "readyReplicas indicates how many replicas are ready and at the desired state",
	"generations":        "generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.",
}

func (OperatorStatus) SwaggerDoc() map[string]string {
	return map_OperatorStatus
}

var map_StaticPodOperatorSpec = map[string]string{
	"":                        "StaticPodOperatorSpec is spec for controllers that manage static pods.",
	"forceRedeploymentReason": "forceRedeploymentReason can be used to force the redeployment of the operand by providing a unique string. This provides a mechanism to kick a previously failed deployment and provide a reason why you think it will work this time instead of failing again on the same config.",
	"failedRevisionLimit":     "failedRevisionLimit is the number of failed static pod installer revisions to keep on disk and in the api -1 = unlimited, 0 or unset = 5 (default)",
	"succeededRevisionLimit":  "succeededRevisionLimit is the number of successful static pod installer revisions to keep on disk and in the api -1 = unlimited, 0 or unset = 5 (default)",
}

func (StaticPodOperatorSpec) SwaggerDoc() map[string]string {
	return map_StaticPodOperatorSpec
}

var map_StaticPodOperatorStatus = map[string]string{
	"":                              "StaticPodOperatorStatus is status for controllers that manage static pods.  There are different needs because individual node status must be tracked.",
	"latestAvailableRevision":       "latestAvailableRevision is the deploymentID of the most recent deployment",
	"latestAvailableRevisionReason": "latestAvailableRevisionReason describe the detailed reason for the most recent deployment",
	"nodeStatuses":                  "nodeStatuses track the deployment values and errors across individual nodes",
}

func (StaticPodOperatorStatus) SwaggerDoc() map[string]string {
	return map_StaticPodOperatorStatus
}

var map_Authentication = map[string]string{
	"": "Authentication provides information to configure an operator to manage authentication.",
}

func (Authentication) SwaggerDoc() map[string]string {
	return map_Authentication
}

var map_AuthenticationList = map[string]string{
	"": "AuthenticationList is a collection of items",
}

func (AuthenticationList) SwaggerDoc() map[string]string {
	return map_AuthenticationList
}

var map_ConsoleCustomization = map[string]string{
	"brand":                "brand is the default branding of the web console which can be overridden by providing the brand field.  There is a limited set of specific brand options. This field controls elements of the console such as the logo. Invalid value will prevent a console rollout.",
	"documentationBaseURL": "documentationBaseURL links to external documentation are shown in various sections of the web console.  Providing documentationBaseURL will override the default documentation URL. Invalid value will prevent a console rollout.",
	"customProductName":    "customProductName is the name that will be displayed in page titles, logo alt text, and the about dialog instead of the normal OpenShift product name.",
	"customLogoFile":       "customLogoFile replaces the default OpenShift logo in the masthead and about dialog. It is a reference to a ConfigMap in the openshift-config namespace. This can be created with a command like 'oc create configmap custom-logo --from-file=/path/to/file -n openshift-config'. Image size must be less than 1 MB due to constraints on the ConfigMap size. The ConfigMap key should include a file extension so that the console serves the file with the correct MIME type. Recommended logo specifications: Dimensions: Max height of 68px and max width of 200px SVG format preferred",
}

func (ConsoleCustomization) SwaggerDoc() map[string]string {
	return map_ConsoleCustomization
}

var map_ConsoleProviders = map[string]string{
	"statuspage": "statuspage contains ID for statuspage.io page that provides status info about.",
}

func (ConsoleProviders) SwaggerDoc() map[string]string {
	return map_ConsoleProviders
}

var map_ConsoleSpec = map[string]string{
	"customization": "customization is used to optionally provide a small set of customization options to the web console.",
	"providers":     "providers contains configuration for using specific service providers.",
}

func (ConsoleSpec) SwaggerDoc() map[string]string {
	return map_ConsoleSpec
}

var map_StatuspageProvider = map[string]string{
	"":       "StatuspageProvider provides identity for statuspage account.",
	"pageID": "pageID is the unique ID assigned by Statuspage for your page. This must be a public page.",
}

func (StatuspageProvider) SwaggerDoc() map[string]string {
	return map_StatuspageProvider
}

var map_DNS = map[string]string{
	"":       "DNS manages the CoreDNS component to provide a name resolution service for pods and services in the cluster.\n\nThis supports the DNS-based service discovery specification: https://github.com/kubernetes/dns/blob/master/docs/specification.md\n\nMore details: https://kubernetes.io/docs/tasks/administer-cluster/coredns",
	"spec":   "spec is the specification of the desired behavior of the DNS.",
	"status": "status is the most recently observed status of the DNS.",
}

func (DNS) SwaggerDoc() map[string]string {
	return map_DNS
}

var map_DNSList = map[string]string{
	"": "DNSList contains a list of DNS",
}

func (DNSList) SwaggerDoc() map[string]string {
	return map_DNSList
}

var map_DNSSpec = map[string]string{
	"": "DNSSpec is the specification of the desired behavior of the DNS.",
}

func (DNSSpec) SwaggerDoc() map[string]string {
	return map_DNSSpec
}

var map_DNSStatus = map[string]string{
	"":              "DNSStatus defines the observed status of the DNS.",
	"clusterIP":     "clusterIP is the service IP through which this DNS is made available.\n\nIn the case of the default DNS, this will be a well known IP that is used as the default nameserver for pods that are using the default ClusterFirst DNS policy.\n\nIn general, this IP can be specified in a pod's spec.dnsConfig.nameservers list or used explicitly when performing name resolution from within the cluster. Example: dig foo.com @<service IP>\n\nMore info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
	"clusterDomain": "clusterDomain is the local cluster DNS domain suffix for DNS services. This will be a subdomain as defined in RFC 1034, section 3.5: https://tools.ietf.org/html/rfc1034#section-3.5 Example: \"cluster.local\"\n\nMore info: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service",
	"conditions":    "conditions provide information about the state of the DNS on the cluster.\n\nThese are the supported DNS conditions:\n\n  * Available\n  - True if the following conditions are met:\n    * DNS controller daemonset is available.\n  - False if any of those conditions are unsatisfied.",
}

func (DNSStatus) SwaggerDoc() map[string]string {
	return map_DNSStatus
}

var map_Etcd = map[string]string{
	"": "Etcd provides information to configure an operator to manage kube-apiserver.",
}

func (Etcd) SwaggerDoc() map[string]string {
	return map_Etcd
}

var map_EtcdList = map[string]string{
	"":         "KubeAPISOperatorConfigList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (EtcdList) SwaggerDoc() map[string]string {
	return map_EtcdList
}

var map_EndpointPublishingStrategy = map[string]string{
	"":             "EndpointPublishingStrategy is a way to publish the endpoints of an IngressController, and represents the type and any additional configuration for a specific type.",
	"type":         "type is the publishing strategy to use. Valid values are:\n\n* LoadBalancerService\n\nPublishes the ingress controller using a Kubernetes LoadBalancer Service.\n\nIn this configuration, the ingress controller deployment uses container networking. A LoadBalancer Service is created to publish the deployment.\n\nSee: https://kubernetes.io/docs/concepts/services-networking/#loadbalancer\n\nIf domain is set, a wildcard DNS record will be managed to point at the LoadBalancer Service's external name. DNS records are managed only in DNS zones defined by dns.config.openshift.io/cluster .spec.publicZone and .spec.privateZone.\n\nWildcard DNS management is currently supported only on the AWS platform.\n\n* HostNetwork\n\nPublishes the ingress controller on node ports where the ingress controller is deployed.\n\nIn this configuration, the ingress controller deployment uses host networking, bound to node ports 80 and 443. The user is responsible for configuring an external load balancer to publish the ingress controller via the node ports.\n\n* Private\n\nDoes not publish the ingress controller.\n\nIn this configuration, the ingress controller deployment uses container networking, and is not explicitly published. The user must manually publish the ingress controller.",
	"loadBalancer": "loadBalancer holds parameters for the load balancer. Present only if type is LoadBalancerService.",
	"hostNetwork":  "hostNetwork holds parameters for the HostNetwork endpoint publishing strategy. Present only if type is HostNetwork.",
	"private":      "private holds parameters for the Private endpoint publishing strategy. Present only if type is Private.",
}

func (EndpointPublishingStrategy) SwaggerDoc() map[string]string {
	return map_EndpointPublishingStrategy
}

var map_HostNetworkStrategy = map[string]string{
	"": "HostNetworkStrategy holds parameters for the HostNetwork endpoint publishing strategy.",
}

func (HostNetworkStrategy) SwaggerDoc() map[string]string {
	return map_HostNetworkStrategy
}

var map_IngressController = map[string]string{
	"":       "IngressController describes a managed ingress controller for the cluster. The controller can service OpenShift Route and Kubernetes Ingress resources.\n\nWhen an IngressController is created, a new ingress controller deployment is created to allow external traffic to reach the services that expose Ingress or Route resources. Updating this resource may lead to disruption for public facing network connections as a new ingress controller revision may be rolled out.\n\nhttps://kubernetes.io/docs/concepts/services-networking/ingress-controllers\n\nWhenever possible, sensible defaults for the platform are used. See each field for more details.",
	"spec":   "spec is the specification of the desired behavior of the IngressController.",
	"status": "status is the most recently observed status of the IngressController.",
}

func (IngressController) SwaggerDoc() map[string]string {
	return map_IngressController
}

var map_IngressControllerList = map[string]string{
	"": "IngressControllerList contains a list of IngressControllers.",
}

func (IngressControllerList) SwaggerDoc() map[string]string {
	return map_IngressControllerList
}

var map_IngressControllerSpec = map[string]string{
	"":                           "IngressControllerSpec is the specification of the desired behavior of the IngressController.",
	"domain":                     "domain is a DNS name serviced by the ingress controller and is used to configure multiple features:\n\n* For the LoadBalancerService endpoint publishing strategy, domain is\n  used to configure DNS records. See endpointPublishingStrategy.\n\n* When using a generated default certificate, the certificate will be valid\n  for domain and its subdomains. See defaultCertificate.\n\n* The value is published to individual Route statuses so that end-users\n  know where to target external DNS records.\n\ndomain must be unique among all IngressControllers, and cannot be updated.\n\nIf empty, defaults to ingress.config.openshift.io/cluster .spec.domain.",
	"replicas":                   "replicas is the desired number of ingress controller replicas. If unset, defaults to 2.",
	"endpointPublishingStrategy": "endpointPublishingStrategy is used to publish the ingress controller endpoints to other networks, enable load balancer integrations, etc.\n\nIf unset, the default is based on infrastructure.config.openshift.io/cluster .status.platform:\n\n  AWS:      LoadBalancerService\n  Azure:    LoadBalancerService\n  GCP:      LoadBalancerService\n  Libvirt:  HostNetwork\n\nAny other platform types (including None) default to HostNetwork.\n\nendpointPublishingStrategy cannot be updated.",
	"defaultCertificate":         "defaultCertificate is a reference to a secret containing the default certificate served by the ingress controller. When Routes don't specify their own certificate, defaultCertificate is used.\n\nThe secret must contain the following keys and data:\n\n  tls.crt: certificate file contents\n  tls.key: key file contents\n\nIf unset, a wildcard certificate is automatically generated and used. The certificate is valid for the ingress controller domain (and subdomains) and the generated certificate's CA will be automatically integrated with the cluster's trust store.\n\nThe in-use certificate (whether generated or user-specified) will be automatically integrated with OpenShift's built-in OAuth server.",
	"namespaceSelector":          "namespaceSelector is used to filter the set of namespaces serviced by the ingress controller. This is useful for implementing shards.\n\nIf unset, the default is no filtering.",
	"routeSelector":              "routeSelector is used to filter the set of Routes serviced by the ingress controller. This is useful for implementing shards.\n\nIf unset, the default is no filtering.",
	"nodePlacement":              "nodePlacement enables explicit control over the scheduling of the ingress controller.\n\nIf unset, defaults are used. See NodePlacement for more details.",
}

func (IngressControllerSpec) SwaggerDoc() map[string]string {
	return map_IngressControllerSpec
}

var map_IngressControllerStatus = map[string]string{
	"":                           "IngressControllerStatus defines the observed status of the IngressController.",
	"availableReplicas":          "availableReplicas is number of observed available replicas according to the ingress controller deployment.",
	"selector":                   "selector is a label selector, in string format, for ingress controller pods corresponding to the IngressController. The number of matching pods should equal the value of availableReplicas.",
	"domain":                     "domain is the actual domain in use.",
	"endpointPublishingStrategy": "endpointPublishingStrategy is the actual strategy in use.",
	"conditions":                 "conditions is a list of conditions and their status.\n\nAvailable means the ingress controller deployment is available and servicing route and ingress resources (i.e, .status.availableReplicas equals .spec.replicas)\n\nThere are additional conditions which indicate the status of other ingress controller features and capabilities.\n\n  * LoadBalancerManaged\n  - True if the following conditions are met:\n    * The endpoint publishing strategy requires a service load balancer.\n  - False if any of those conditions are unsatisfied.\n\n  * LoadBalancerReady\n  - True if the following conditions are met:\n    * A load balancer is managed.\n    * The load balancer is ready.\n  - False if any of those conditions are unsatisfied.\n\n  * DNSManaged\n  - True if the following conditions are met:\n    * The endpoint publishing strategy and platform support DNS.\n    * The ingress controller domain is set.\n    * dns.config.openshift.io/cluster configures DNS zones.\n  - False if any of those conditions are unsatisfied.\n\n  * DNSReady\n  - True if the following conditions are met:\n    * DNS is managed.\n    * DNS records have been successfully created.\n  - False if any of those conditions are unsatisfied.",
}

func (IngressControllerStatus) SwaggerDoc() map[string]string {
	return map_IngressControllerStatus
}

var map_LoadBalancerStrategy = map[string]string{
	"":      "LoadBalancerStrategy holds parameters for a load balancer.",
	"scope": "scope indicates the scope at which the load balancer is exposed. Possible values are \"External\" and \"Internal\".  The default is \"External\".",
}

func (LoadBalancerStrategy) SwaggerDoc() map[string]string {
	return map_LoadBalancerStrategy
}

var map_NodePlacement = map[string]string{
	"":             "NodePlacement describes node scheduling configuration for an ingress controller.",
	"nodeSelector": "nodeSelector is the node selector applied to ingress controller deployments.\n\nIf unset, the default is:\n\n  beta.kubernetes.io/os: linux\n  node-role.kubernetes.io/worker: ''\n\nIf set, the specified selector is used and replaces the default.",
	"tolerations":  "tolerations is a list of tolerations applied to ingress controller deployments.\n\nThe default is an empty list.\n\nSee https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/",
}

func (NodePlacement) SwaggerDoc() map[string]string {
	return map_NodePlacement
}

var map_PrivateStrategy = map[string]string{
	"": "PrivateStrategy holds parameters for the Private endpoint publishing strategy.",
}

func (PrivateStrategy) SwaggerDoc() map[string]string {
	return map_PrivateStrategy
}

var map_KubeAPIServer = map[string]string{
	"": "KubeAPIServer provides information to configure an operator to manage kube-apiserver.",
}

func (KubeAPIServer) SwaggerDoc() map[string]string {
	return map_KubeAPIServer
}

var map_KubeAPIServerList = map[string]string{
	"":         "KubeAPIServerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (KubeAPIServerList) SwaggerDoc() map[string]string {
	return map_KubeAPIServerList
}

var map_KubeControllerManager = map[string]string{
	"": "KubeControllerManager provides information to configure an operator to manage kube-controller-manager.",
}

func (KubeControllerManager) SwaggerDoc() map[string]string {
	return map_KubeControllerManager
}

var map_KubeControllerManagerList = map[string]string{
	"":         "KubeControllerManagerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (KubeControllerManagerList) SwaggerDoc() map[string]string {
	return map_KubeControllerManagerList
}

var map_AdditionalNetworkDefinition = map[string]string{
	"":             "AdditionalNetworkDefinition configures an extra network that is available but not created by default. Instead, pods must request them by name. type must be specified, along with exactly one \"Config\" that matches the type.",
	"type":         "type is the type of network The only supported value is NetworkTypeRaw",
	"name":         "name is the name of the network. This will be populated in the resulting CRD This must be unique.",
	"namespace":    "namespace is the namespace of the network. This will be populated in the resulting CRD If not given the network will be created in the default namespace.",
	"rawCNIConfig": "rawCNIConfig is the raw CNI configuration json to create in the NetworkAttachmentDefinition CRD",
}

func (AdditionalNetworkDefinition) SwaggerDoc() map[string]string {
	return map_AdditionalNetworkDefinition
}

var map_ClusterNetworkEntry = map[string]string{
	"": "ClusterNetworkEntry is a subnet from which to allocate PodIPs. A network of size HostPrefix (in CIDR notation) will be allocated when nodes join the cluster. Not all network providers support multiple ClusterNetworks",
}

func (ClusterNetworkEntry) SwaggerDoc() map[string]string {
	return map_ClusterNetworkEntry
}

var map_DefaultNetworkDefinition = map[string]string{
	"":                    "DefaultNetworkDefinition represents a single network plugin's configuration. type must be specified, along with exactly one \"Config\" that matches the type.",
	"type":                "type is the type of network All NetworkTypes are supported except for NetworkTypeRaw",
	"openshiftSDNConfig":  "openShiftSDNConfig configures the openshift-sdn plugin",
	"ovnKubernetesConfig": "oVNKubernetesConfig configures the ovn-kubernetes plugin. This is currently not implemented.",
	"kuryrConfig":         "KuryrConfig configures the kuryr plugin",
}

func (DefaultNetworkDefinition) SwaggerDoc() map[string]string {
	return map_DefaultNetworkDefinition
}

var map_KuryrConfig = map[string]string{
	"":                     "KuryrConfig configures the Kuryr-Kubernetes SDN",
	"daemonProbesPort":     "The port kuryr-daemon will listen for readiness and liveness requests.",
	"controllerProbesPort": "The port kuryr-controller will listen for readiness and liveness requests.",
}

func (KuryrConfig) SwaggerDoc() map[string]string {
	return map_KuryrConfig
}

var map_Network = map[string]string{
	"": "Network describes the cluster's desired network configuration. It is consumed by the cluster-network-operator.",
}

func (Network) SwaggerDoc() map[string]string {
	return map_Network
}

var map_NetworkList = map[string]string{
	"": "NetworkList contains a list of Network configurations",
}

func (NetworkList) SwaggerDoc() map[string]string {
	return map_NetworkList
}

var map_NetworkSpec = map[string]string{
	"":                    "NetworkSpec is the top-level network configuration object.",
	"clusterNetwork":      "clusterNetwork is the IP address pool to use for pod IPs. Some network providers, e.g. OpenShift SDN, support multiple ClusterNetworks. Others only support one. This is equivalent to the cluster-cidr.",
	"serviceNetwork":      "serviceNetwork is the ip address pool to use for Service IPs Currently, all existing network providers only support a single value here, but this is an array to allow for growth.",
	"defaultNetwork":      "defaultNetwork is the \"default\" network that all pods will receive",
	"additionalNetworks":  "additionalNetworks is a list of extra networks to make available to pods when multiple networks are enabled.",
	"disableMultiNetwork": "disableMultiNetwork specifies whether or not multiple pod network support should be disabled. If unset, this property defaults to 'false' and multiple network support is enabled.",
	"deployKubeProxy":     "deployKubeProxy specifies whether or not a standalone kube-proxy should be deployed by the operator. Some network providers include kube-proxy or similar functionality. If unset, the plugin will attempt to select the correct value, which is false when OpenShift SDN and ovn-kubernetes are used and true otherwise.",
	"kubeProxyConfig":     "kubeProxyConfig lets us configure desired proxy configuration. If not specified, sensible defaults will be chosen by OpenShift directly. Not consumed by all network providers - currently only openshift-sdn.",
}

func (NetworkSpec) SwaggerDoc() map[string]string {
	return map_NetworkSpec
}

var map_NetworkStatus = map[string]string{
	"": "NetworkStatus is currently unused. Instead, status is reported in the Network.config.openshift.io object.",
}

func (NetworkStatus) SwaggerDoc() map[string]string {
	return map_NetworkStatus
}

var map_OVNKubernetesConfig = map[string]string{
	"":    "ovnKubernetesConfig is the proposed configuration parameters for networks using the ovn-kubernetes network project",
	"mtu": "mtu is the MTU to use for the tunnel interface. This must be 100 bytes smaller than the uplink mtu. Default is 1400",
}

func (OVNKubernetesConfig) SwaggerDoc() map[string]string {
	return map_OVNKubernetesConfig
}

var map_OpenShiftSDNConfig = map[string]string{
	"":                       "OpenShiftSDNConfig configures the three openshift-sdn plugins",
	"mode":                   "mode is one of \"Multitenant\", \"Subnet\", or \"NetworkPolicy\"",
	"vxlanPort":              "vxlanPort is the port to use for all vxlan packets. The default is 4789.",
	"mtu":                    "mtu is the mtu to use for the tunnel interface. Defaults to 1450 if unset. This must be 50 bytes smaller than the machine's uplink.",
	"useExternalOpenvswitch": "useExternalOpenvswitch tells the operator not to install openvswitch, because it will be provided separately. If set, you must provide it yourself.",
}

func (OpenShiftSDNConfig) SwaggerDoc() map[string]string {
	return map_OpenShiftSDNConfig
}

var map_ProxyConfig = map[string]string{
	"":                   "ProxyConfig defines the configuration knobs for kubeproxy All of these are optional and have sensible defaults",
	"iptablesSyncPeriod": "The period that iptables rules are refreshed. Default: 30s",
	"bindAddress":        "The address to \"bind\" on Defaults to 0.0.0.0",
	"proxyArguments":     "Any additional arguments to pass to the kubeproxy process",
}

func (ProxyConfig) SwaggerDoc() map[string]string {
	return map_ProxyConfig
}

var map_OpenShiftAPIServer = map[string]string{
	"": "OpenShiftAPIServer provides information to configure an operator to manage openshift-apiserver.",
}

func (OpenShiftAPIServer) SwaggerDoc() map[string]string {
	return map_OpenShiftAPIServer
}

var map_OpenShiftAPIServerList = map[string]string{
	"":         "OpenShiftAPIServerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (OpenShiftAPIServerList) SwaggerDoc() map[string]string {
	return map_OpenShiftAPIServerList
}

var map_OpenShiftControllerManager = map[string]string{
	"": "OpenShiftControllerManager provides information to configure an operator to manage openshift-controller-manager.",
}

func (OpenShiftControllerManager) SwaggerDoc() map[string]string {
	return map_OpenShiftControllerManager
}

var map_OpenShiftControllerManagerList = map[string]string{
	"":         "OpenShiftControllerManagerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (OpenShiftControllerManagerList) SwaggerDoc() map[string]string {
	return map_OpenShiftControllerManagerList
}

var map_KubeScheduler = map[string]string{
	"": "KubeScheduler provides information to configure an operator to manage scheduler.",
}

func (KubeScheduler) SwaggerDoc() map[string]string {
	return map_KubeScheduler
}

var map_KubeSchedulerList = map[string]string{
	"":         "KubeSchedulerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (KubeSchedulerList) SwaggerDoc() map[string]string {
	return map_KubeSchedulerList
}

var map_ServiceCA = map[string]string{
	"":       "ServiceCA provides information to configure an operator to manage the service cert controllers",
	"spec":   "spec holds user settable values for configuration",
	"status": "status holds observed values from the cluster. They may not be overridden.",
}

func (ServiceCA) SwaggerDoc() map[string]string {
	return map_ServiceCA
}

var map_ServiceCAList = map[string]string{
	"":         "ServiceCAList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (ServiceCAList) SwaggerDoc() map[string]string {
	return map_ServiceCAList
}

var map_ServiceCatalogAPIServer = map[string]string{
	"": "ServiceCatalogAPIServer provides information to configure an operator to manage Service Catalog API Server",
}

func (ServiceCatalogAPIServer) SwaggerDoc() map[string]string {
	return map_ServiceCatalogAPIServer
}

var map_ServiceCatalogAPIServerList = map[string]string{
	"":         "ServiceCatalogAPIServerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (ServiceCatalogAPIServerList) SwaggerDoc() map[string]string {
	return map_ServiceCatalogAPIServerList
}

var map_ServiceCatalogControllerManager = map[string]string{
	"": "ServiceCatalogControllerManager provides information to configure an operator to manage Service Catalog Controller Manager",
}

func (ServiceCatalogControllerManager) SwaggerDoc() map[string]string {
	return map_ServiceCatalogControllerManager
}

var map_ServiceCatalogControllerManagerList = map[string]string{
	"":         "ServiceCatalogControllerManagerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (ServiceCatalogControllerManagerList) SwaggerDoc() map[string]string {
	return map_ServiceCatalogControllerManagerList
}

// AUTO-GENERATED FUNCTIONS END HERE
