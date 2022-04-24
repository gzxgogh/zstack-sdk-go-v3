package model

//创建负载均衡器
type CreateLoadBalancerRequest struct {
	ReqConfig
	Params     CreateLoadBalancerParams `json:"params" bson:"params"`
	SystemTags []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateLoadBalancerParams struct {
	VipUuid      string `json:"vipUuid" bson:"vipUuid"`
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除负载均衡器
type DeleteLoadBalancerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLoadBalancerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询负载均衡器
type QueryLoadBalancerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLoadBalancerResponse struct {
	Inventories []LoadBalancerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode               `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//刷新负载均衡器
type RefreshLoadBalancerRequest struct {
	ReqConfig
	UUID                string                    `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	RefreshLoadBalancer RefreshLoadBalancerParams `json:"updatePortForwardingRule" bson:"updatePortForwardingRule"`
	SystemTags          []string                  `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags            []string                  `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RefreshLoadBalancerParams struct {
}

type RefreshLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建负载均衡监听器
type CreateLoadBalancerListenerRequest struct {
	ReqConfig
	LoadBalancerUuid string                           `json:"loadBalancerUuid" bson:"loadBalancerUuid"` //负载均衡器UUID
	Params           CreateLoadBalancerListenerParams `json:"params" bson:"params"`
	SystemTags       []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags         []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateLoadBalancerListenerParams struct {
	Name                string   `json:"name" bson:"name"`                                     //资源名称
	Description         string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	InstancePort        int      `json:"instancePort,omitempty" bson:"instancePort,omitempty"` //云主机端口
	LoadBalancerPort    int      `json:"loadBalancerPort" bson:"loadBalancerPort"`             //负载均衡器端口
	Protocol            string   `json:"protocol,omitempty" bson:"protocol,omitempty"`         //协议:tcp,HTTP,https,dcp
	CertificateUuid     string   `json:"certificateUuid,omitempty" bson:"certificateUuid,omitempty"`
	ResourceUuid        string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"`               //资源UUID。若指定，镜像会使用该字段值作为UUID。
	HealthCheckProtocol string   `json:"healthCheckProtocol,omitempty" bson:"healthCheckProtocol,omitempty"` //健康检查协议
	HealthCheckMethod   string   `json:"healthCheckMethod,omitempty" bson:"healthCheckMethod,omitempty"`     //健康检查方法
	HealthCheckURI      string   `json:"healthCheckURI,omitempty" bson:"healthCheckURI,omitempty"`           //健康检查的URI
	HealthCheckHttpCode string   `json:"healthCheckHttpCode,omitempty" bson:"healthCheckHttpCode,omitempty"` //健康检查期望的返回码
	AclStatus           string   `json:"aclStatus,omitempty" bson:"aclStatus,omitempty"`                     //访问控制策略状态
	AclUuids            []string `json:"aclUuids,omitempty" bson:"aclUuids,omitempty"`                       //访问控制策略组
	AclType             string   `json:"aclType,omitempty" bson:"aclType,omitempty"`                         //访问控制策略类型:white,black
	TagUuids            []string `json:"tagUuids ,omitempty" bson:"tagUuids ,omitempty"`                     //源CIDR; 端口转发规则只作用于源CIDR的流量; 如果忽略不设置, 会默认设置为to 0.0.0.0/0
}

type CreateLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除负载均衡监听器
type DeleteLoadBalancerListenerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteLoadBalancerListenerResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询负载均衡监听器
type QueryLoadBalancerListenerRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryLoadBalancerListenerResponse struct {
	Inventories []LoadBalancerListenerInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新负载均衡监听器
type UpdateLoadBalancerListenerRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateLoadBalancerListener UpdateLoadBalancerListenerParams `json:"updateLoadBalancerListener" bson:"updateLoadBalancerListener"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLoadBalancerListenerParams struct {
	Name        string `json:"name" bson:"name"`                                   //资源名称
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
}

type UpdateLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//修改负载均衡监听器参数
type ChangeLoadBalancerListenerRequest struct {
	ReqConfig
	UUID                       string                           `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	UpdateLoadBalancerListener ChangeLoadBalancerListenerParams `json:"updateLoadBalancerListener" bson:"updateLoadBalancerListener"`
	SystemTags                 []string                         `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags                   []string                         `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeLoadBalancerListenerParams struct {
	ConnectionIdleTimeout int    `json:"connectionIdleTimeout,omitempty" bson:"connectionIdleTimeout,omitempty"`
	MaxConnection         int    `json:"maxConnection,omitempty" bson:"maxConnection,omitempty"`
	BalancerAlgorithm     string `json:"balancer_algorithm,omitempty" bson:"balancer_algorithm,omitempty"` //roundrobin,leastconn,source
	HealthCheckTarget     string `json:"healthCheckTarget,omitempty" bson:"healthCheckTarget,omitempty"`
	HealthyThreshold      int    `json:"healthyThreshold,omitempty" bson:"healthyThreshold,omitempty"`
	UnhealthyThreshold    int    `json:"unhealthyThreshold,omitempty" bson:"unhealthyThreshold,omitempty"`
	HealthCheckInterval   int    `json:"healthCheckInterval,omitempty" bson:"healthCheckInterval,omitempty"`
	HealthCheckProtocol   string `json:"healthCheckProtocol,omitempty" bson:"healthCheckProtocol,omitempty"` //健康检查协议:tcp,udp,HTTP
	HealthCheckMethod     string `json:"healthCheckMethod,omitempty" bson:"healthCheckMethod,omitempty"`     //健康检查方法:GET,HEAD
	HealthCheckURI        string `json:"healthCheckURI,omitempty" bson:"healthCheckURI,omitempty"`           //健康检查的URI
	HealthCheckHttpCode   string `json:"healthCheckHttpCode,omitempty" bson:"healthCheckHttpCode,omitempty"` //健康检查期望的返回码
	AclStatus             string `json:"aclStatus,omitempty" bson:"aclStatus,omitempty"`                     //访问控制策略状态
}

type ChangeLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取云主机网卡
type GetCandidateVmNicsForLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string   `json:"listenerUuid" bson:"listenerUuid"`                 //负载均衡监听器UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateVmNicsForLoadBalancerResponse struct {
	Inventories []VmNic   `json:"inventories" bson:"inventories"`
	Error       ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加云主机网卡到负载均衡器
type AddVmNicToLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string                       `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Params       AddVmNicToLoadBalancerParams `json:"params" bson:"params"`
	SystemTags   []string                     `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                     `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddVmNicToLoadBalancerParams struct {
	VmNicUuids []string `json:"vmNicUuids" bson:"vmNicUuids"`
}

type AddVmNicToLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从负载均衡器移除云主机网卡
type RemoveVmNicFromLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	VmNicUuid    string   `json:"vmNicUuid" bson:"vmNicUuid"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveVmNicFromLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新负载均衡器
type UpdateLoadBalancerRequest struct {
	ReqConfig
	UUID               string                   `json:"uuid" bson:"uuid"`
	UpdateLoadBalancer UpdateLoadBalancerParams `json:"updateLoadBalancer" bson:"updateLoadBalancer"`
	SystemTags         []string                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags           []string                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateLoadBalancerParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type UpdateLoadBalancerResponse struct {
	Inventory LoadBalancerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode             `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建证书
type CreateCertificateRequest struct {
	ReqConfig
	Params     CreateCertificateParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateCertificateParams struct {
	Certificate  string `json:"certificate" bson:"certificate"`
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type CreateCertificateResponse struct {
	Inventory CertificateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除证书
type DeleteCertificateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteCertificateResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询证书
type QueryCertificateRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryCertificateResponse struct {
	Inventories []CertificateInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode              `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加证书到负载均衡
type AddCertificateToLoadBalancerListenerRequest struct {
	ReqConfig
	ListenerUuid string                                     `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Params       AddCertificateToLoadBalancerListenerParams `json:"params" bson:"params"`
	SystemTags   []string                                   `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                                   `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddCertificateToLoadBalancerListenerParams struct {
	CertificateUuid string `json:"certificateUuid" bson:"certificateUuid"`
}

type AddCertificateToLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//从负载均衡移除证书
type RemoveCertificateFromLoadBalancerListenerRequest struct {
	ReqConfig
	ListenerUuid    string   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	CertificateUuid string   `json:"certificateUuid" bson:"certificateUuid"`
	SystemTags      []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags        []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveCertificateFromLoadBalancerListenerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//更新证书信息
type UpdateCertificateRequest struct {
	ReqConfig
	UUID       string                  `json:"uuid" bson:"uuid"` //
	Params     UpdateCertificateParams `json:"params" bson:"params"`
	SystemTags []string                `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type UpdateCertificateParams struct {
	Name         string `json:"name" bson:"name"`                                     //资源名称
	Description  string `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
}

type UpdateCertificateResponse struct {
	Inventory CertificateInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//创建访问控制策略组
type CreateAccessControlListRequest struct {
	ReqConfig
	Params     CreateAccessControlListParams `json:"params" bson:"params"`
	SystemTags []string                      `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                      `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type CreateAccessControlListParams struct {
	Name         string   `json:"name" bson:"name"`                                   //资源名称
	Description  string   `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	IpVersion    string   `json:"ipVersion,omitempty" bson:"ipVersion,omitempty"`
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type CreateAccessControlListResponse struct {
	Inventory AccessControlListInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                  `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除访问控制策略组
type DeleteAccessControlListRequest struct {
	ReqConfig
	UUID       string   `json:"uuid" bson:"uuid"`                                 //资源的UUID，唯一标示该资源
	DeleteMode string   `json:"deleteMode,omitempty" bson:"deleteMode,omitempty"` //删除模式(Permissive 或者 Enforcing, 默认 Permissive)
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type DeleteAccessControlListResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//查询访问控制策略组
type QueryAccessControlListRequest struct {
	ReqConfig
	UUID       string   `json:"uuid,omitempty" bson:"uuid,omitempty"`             //资源的UUID，唯一标示该资源
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type QueryAccessControlListResponse struct {
	Inventories []AccessControlListInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode                    `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//向访问控制策略组添加IP组
type AddAccessControlListEntryRequest struct {
	ReqConfig
	AclUuid    string                          `json:"aclUuid" bson:"aclUuid"` //资源的UUID，唯一标示该资源
	Params     AddAccessControlListEntryParams `json:"params" bson:"params"`
	SystemTags []string                        `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string                        `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddAccessControlListEntryParams struct {
	Entries      string   `json:"entries" bson:"entries"`
	Description  string   `json:"description,omitempty" bson:"description,omitempty"`   //详细描述
	ResourceUuid string   `json:"resourceUuid,omitempty" bson:"resourceUuid,omitempty"` //资源UUID。若指定，镜像会使用该字段值作为UUID。
	TagUuids     []string `json:"tagUuids,omitempty" bson:"tagUuids,omitempty"`
}

type AddAccessControlListEntryResponse struct {
	Inventory AccessControlListEntryInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                       `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除访问控制策略的IP组
type RemoveAccessControlListEntryRequest struct {
	ReqConfig
	AclUuid    string   `json:"aclUuid" bson:"aclUuid"` //资源的UUID，唯一标示该资源
	UUID       string   `json:"uuid" bson:"uuid"`
	SystemTags []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags   []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveAccessControlListEntryResponse struct {
	Error ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//添加监听器的访问控制策略
type AddAccessControlListToLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string                                   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Params       AddAccessControlListToLoadBalancerParams `json:"params" bson:"params"`
	SystemTags   []string                                 `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string                                 `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type AddAccessControlListToLoadBalancerParams struct {
	AclUuids []string `json:"aclUuids" bson:"aclUuids"`
	AclType  string   `json:"aclType" bson:"aclType"` //详细描述
}

type AddAccessControlListToLoadBalancerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//删除监听器访问控制策略
type RemoveAccessControlListFromLoadBalancerRequest struct {
	ReqConfig
	AclUuids     []string `json:"aclUuids" bson:"aclUuids"`
	ListenerUuid string   `json:"listenerUuid" bson:"listenerUuid"`                 //负载均衡监听器UUID
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type RemoveAccessControlListFromLoadBalancerResponse struct {
	Inventory LoadBalancerListenerInventory `json:"inventory" bson:"inventory"`
	Error     ErrorCode                     `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

//获取监听器可加载L3网络
type GetCandidateL3NetworksForLoadBalancerRequest struct {
	ReqConfig
	ListenerUuid string   `json:"listenerUuid" bson:"listenerUuid"` //负载均衡监听器UUID
	Limit        int      `json:"limit" bson:"limit"`
	Start        int      `json:"start" bson:"start"`
	SystemTags   []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"` //云主机系统标签
	UserTags     []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type GetCandidateL3NetworksForLoadBalancerResponse struct {
	Inventories []L3NetworkInventory `json:"inventories" bson:"inventories"`
	Error       ErrorCode            `json:"error,omitempty" bson:"error,omitempty"` //错误信息
}

type LoadBalancerInventory struct {
	UUID        string                          `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	VipUuid     string                          `json:"vipUuid" bson:"vipUuid"`
	Name        string                          `json:"name" bson:"name"`                                   //资源名称
	Description string                          `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	State       string                          `json:"state" bson:"state"`
	CreateDate  string                          `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate  string                          `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	Listeners   []LoadBalancerListenerInventory `json:"listeners" bson:"listeners"`
}

type LoadBalancerListenerInventory struct {
	UUID             string            `json:"uuid" bson:"uuid"`                                   //资源的UUID，唯一标示该资源
	Name             string            `json:"name" bson:"name"`                                   //资源名称
	Description      string            `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	LoadBalancerUuid string            `json:"loadBalancerUuid" bson:"loadBalancerUuid"`           //负载均衡器UUID
	InstancePort     string            `json:"instancePort" bson:"instancePort"`
	LoadBalancerPort string            `json:"loadBalancerPort" bson:"loadBalancerPort"`
	Protocol         string            `json:"protocol" bson:"protocol"`
	CreateDate       string            `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate       string            `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
	VmNicRefs        []VmNicRefs       `json:"vmNicRefs" bson:"vmNicRefs"`
	AclRefs          []AclRefs         `json:"aclRefs" bson:"aclRefs"`
	CertificateRefs  []CertificateRefs `json:"certificateRefs" bson:"certificateRefs"`
}

type VmNicRefs struct {
	Id           int64  `json:"id" bson:"id"` //资源的UUID，唯一标示该资源
	ListenerUuid string `json:"listenerUuid" bson:"listenerUuid"`
	VmNicUuid    string `json:"vmNicUuid" bson:"vmNicUuid"` //云主机网卡UUID
	Status       string `json:"status" bson:"status"`
	CreateDate   string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type AclRefs struct {
	Id           int64  `json:"id" bson:"id"`                     //资源的UUID，唯一标示该资源
	ListenerUuid string `json:"listenerUuid" bson:"listenerUuid"` //监听器唯一标识
	AclUuid      string `json:"aclUuid" bson:"aclUuid"`           //访问策略组唯一标识
	Type         string `json:"type" bson:"type"`                 //访问策略类型
	CreateDate   string `json:"createDate" bson:"createDate"`     //创建时间
	LastOpDate   string `json:"lastOpDate" bson:"lastOpDate"`     //最后一次修改时间
}

type CertificateInventory struct {
	UUID        string            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string            `json:"name" bson:"name"` //资源名称
	Certificate string            `json:"certificate" bson:"certificate"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string            `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string            `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	Listeners   []CertificateRefs `json:"listeners" bson:"listeners"`
}

type CertificateRefs struct {
	Id              int64  `json:"id" bson:"id"` //资源的UUID，唯一标示该资源
	ListenerUuid    string `json:"listenerUuid" bson:"listenerUuid"`
	CertificateUuid string `json:"certificateUuid" bson:"certificateUuid"`
	CreateDate      string `json:"createDate" bson:"createDate"` //创建时间
	LastOpDate      string `json:"lastOpDate" bson:"lastOpDate"` //最后一次修改时间
}

type AccessControlListInventory struct {
	UUID        string                            `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	Name        string                            `json:"name" bson:"name"` //资源名称
	IpVersion   string                            `json:"ipVersion" bson:"ipVersion"`
	Description string                            `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string                            `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string                            `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
	Entries     []AccessControlListEntryInventory `json:"entries" bson:"entries"`
}

type AccessControlListEntryInventory struct {
	UUID        string `json:"uuid" bson:"uuid"` //资源的UUID，唯一标示该资源
	AclUuid     string `json:"aclUuid" bson:"aclUuid"`
	IpEntries   string `json:"ipEntries" bson:"ipEntries"`
	Description string `json:"description,omitempty" bson:"description,omitempty"` //详细描述
	CreateDate  string `json:"createDate" bson:"createDate"`                       //创建时间
	LastOpDate  string `json:"lastOpDate" bson:"lastOpDate"`                       //最后一次修改时间
}
