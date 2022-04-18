package model

//L3网络中是否存在可用VF网卡
type IsVfNicAvailableInL3NetworkRequest struct {
	ReqConfig
	L3NetworkUUID string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"` //	三层网络UUID
	HostUuid      string   `json:"hostUuid" bson:"hostUuid"`           //物理机UUID
	SystemTags    []string `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags      []string `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type IsVfNicAvailableInL3NetworkResponse struct {
	Error          ErrorCode `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	VfNicAvailable bool      `json:"vfNicAvailable" bson:"vfNicAvailable"`
}

//修改云主机网卡类型
type ChangeVmNicTypeRequest struct {
	ReqConfig
	VmNicUuid       string                `json:"vmNicUuid" bson:"vmNicUuid"` //云主机网卡UUID
	ChangeVmNicType ChangeVmNicTypeParams `json:"changeVmNicType" bson:"changeVmNicType"`
	SystemTags      []string              `json:"systemTags,omitempty" bson:"systemTags,omitempty"`
	UserTags        []string              `json:"userTags,omitempty" bson:"userTags,omitempty"`
}

type ChangeVmNicTypeParams struct {
	VmNicType string `json:"vmNicType" bson:"vmNicType"` //云主机网卡类型:VNIC
}

type ChangeVmNicTypeResponse struct {
	Error     ErrorCode      `json:"error,omitempty" bson:"error,omitempty"` //错误信息
	Inventory VmNicInventory `json:"inventory" bson:"inventory"`
}

type VmNicInventory struct {
	UUID           string   `json:"uuid" bson:"uuid"`                     //资源的UUID，唯一标示该资源
	VMInstanceUUID string   `json:"vmInstanceUuid" bson:"vmInstanceUuid"` //	云主机UUID
	L3NetworkUUID  string   `json:"l3NetworkUuid" bson:"l3NetworkUuid"`   //	三层网络UUID
	IP             string   `json:"ip" bson:"ip"`                         //ip地址
	Mac            string   `json:"mac" bson:"mac"`                       //mac地址
	HypervisorType string   `json:"hypervisorType" bson:"hypervisorType"` //虚拟化类型
	Netmask        string   `json:"netmask" bson:"netmask"`               //子网掩码
	Gateway        string   `json:"gateway" bson:"gateway"`               //网关
	MetaData       string   `json:"metaData" bson:"metaData"`             //内部使用的保留域，元数据
	IpVersion      int      `json:"ipVersion" bson:"ipVersion"`           //IP地址版本
	DeviceID       int      `json:"deviceId" bson:"deviceId"`             //设备ID 标识网卡在客户操作系统（guest operating system）以太网设备中顺序的整形数字。例如， 0通常代表eth0，1通常代表eth1。
	Type           string   `json:"type" bson:"type"`                     //
	CreateDate     string   `json:"createDate" bson:"createDate"`         //创建时间
	LastOpDate     string   `json:"lastOpDate" bson:"lastOpDate"`         //最后一次修改时间
	UsedIps        []UsedIp `json:"usedIps" bson:"usedIps"`
}
