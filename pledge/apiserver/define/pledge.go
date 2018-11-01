package define

//押品概要信息
type PledgeGeneralInfo struct {
	ChannelSeq       string `json:"ChannelSeq"`       //渠道编号
	PledgeNoStorage  string `json:"PledgeNoStorage"`  //质押物编号
	PledgeName       string `json:"PledgeName"`       //质押物名称
	SocialCreditCode string `json:"SocialCreditCode"` //统一社会信用代码
	PledgeOwnerName  string `json:"PledgeOwnerName"`  //客户名称
	PledgeState     int    `json:"PledgeStatus"`     //质押物状态
	PledgeType       int    `json:"PledgeType"`       //质押物类型
}

//押品详细信息
type PledgeDetailInfo struct {
	CommodityVaretiesName string              `json:"CommodityVaretiesName"` //商品名称
	Specifications        string              `json:"Specifications"`        //规格
	Quantity              int                 `json:"Quantity"`              //数量
	CommodityCompany      string              `json:"CommodityCompany"`      //商品单位
	Manufacturer          string              `json:"Manufacturer"`          //生产厂商
	VoucherNo             string              `json:"VoucherNo"`             //凭证号
	VoucherName           string              `json:"VoucherName"`           //凭证名称
	CoreSocialCreditCode  string              `json:"CoreSocialCreditCode"`  //仓库企业统一社会信用代码
	CoreEnterpriseName    string              `json:"CoreEnterpriseName"`    //仓储企业名称
	DepotNo               string              `json:"DepotNo"`               //仓库编号
	DepotName             string              `json:"DepotName"`             //仓库名称
	IsBuyInsurance        int                 `json:"IsBuyInsurance"`        //是否购买保险
	WarehousingDate       string              `json:"WarehousingDate"`       //入库日期
	InventoryStatus       int                 `json:"InventoryStatus"`       //存货状态
	MonitorEquipmentIds   []string            `json:"MonitorEquipmentIds"`   //监控设备ID集
	RuleList              map[string][]string `json:"RuleList"`              //监控设备与规则ID的映射
}

//保险信息
type PledgeInsuranceInfo struct {
	ChannelSeq         string `json:"ChannelSeq"`         //渠道编号
	PledgeNoStorage    string `json:"PledgeNoStorage"`    //质押物编号
	PledgeName         string `json:"PledgeName"`         //质押物名称
	InsuranceCompany   string `json:"InsuranceCompany"`   //保险公司
	InsuranceNo        string `json:"InsuranceNo"`        //保单单号
	InsuranceType      int    `json:"InsuranceType"`      //保险类型
	InsuranceClass     int    `json:"InsuranceClass"`     //险种
	PolicyHolder       string `json:"PolicyHolder"`       //投保人
	Beneficiary        string `json:"Beneficiary"`        //受益人
	InsurancePremium   string `json:"InsurancePremium"`   //投保费
	InsuranceAmount    string `json:"InsuranceAmount"`    //投保金额
	InsuranceStartDate string `json:"InsuranceStartDate"` //保险起始日期
	InsuranceEndDate   string `json:"InsuranceEndDate"`   //保险结束日期
	InsuranceDate      string `json:"InsuranceDate"`      //出单日期
}

//公证信息
type PledgeNotarizationInfo struct {
	ChannelSeq       string `json:"ChannelSeq"`       //渠道编号
	NotarialDeedNo   string `json:"NotarialDeedNo"`   //公证书编号
	NotarialOffice   string `json:"NotarialOffice"`   //公证处
	RegistrationDate string `json:"RegistrationDate"` //登记日期
}

//巡库基本信息
type MonitorGeneralInfo struct {
	ChannelSeq           string `json:"ChannelSeq"`           //渠道编号
	PledgeOwnerNumber    string `json:"PledgeOwnerNumber"`    //统一社会信用代码
	PledgeOwnerName      string `json:"PledgeOwnerName"`      //客户名称
	CoreSocialCreditCode string `json:"CoreSocialCreditCode"` //仓库企业统一社会信用代码
	CoreEnterpriseName   string `json:"CoreEnterpriseName"`   //仓储企业名称
	DepotNo              string `json:"DepotNo"`              //仓库编号
	DepotName            string `json:"DepotName"`            //仓库名称
	PledgeNoStorage      string `json:"PledgeNoStorage"`      //质押物编号
	PledgeName           string `json:"PledgeName"`           //质押物名称
	Manufacturer         string `json:"Manufacturer"`         //生产厂商
	Specifications       string `json:"Specifications"`       //规格
	UnitPrice            int    `json:"UnitPrice"`            //单价
	Quantity             int    `json:"Quantity"`             //数量
	CommodityCompany     string `json:"CommodityCompany"`     //商品单位
	PledgeState          int    `json:"PledgeState"`          //质押物状态
	PledgeStartDate      string `json:"PledgeStartDate"`      //质押起始日
	PledgeEndDate        string `json:"PledgeEndDate"`        //质押到期日
	CheckDate            string `json:"CheckDate"`            //巡库时间
	CheckMode            int    `json:"CheckMode"`            //巡库方式
	MonitorEquipmentId   string `json:"MonitorEquipmentId"`   //监控设备ID
	CheckDuration        string `json:"CheckDuration"`        //巡库时长
	CheckStaff           string `json:"CheckStaff"`           //巡库人员
	CheckInstruction     string `json:"CheckInstruction"`     //巡库说明
}

//巡库结果
type MonitorCheckResult struct {
	RecordPledgeArrival             int `json:"RecordPledgeArrival"`             //记录未来货权质押到货情况
	CheckPledgeSupervise            int `json:"CheckPledgeSupervise"`            //核定库存质押监管情况
	SecuritySystemFormulation       int `json:"SecuritySystemFormulation"`       //是否制定安全保卫制度
	OperationPostSetting            int `json:"OperationPostSetting"`            //是否设定验货、入库、出库等内部操作岗位
	GoodsProceduresComplement       int `json:"GoodsProceduresComplement"`       //以往客户货物进出手续是否齐全
	AdministratorQuality            int `json:"AdministratorQuality"`            //管理人员综合素质
	OperatorExperience              int `json:"OperatorExperience"`              //操作人员行业经验
	SupervisionUnitCoordination     int `json:"SupervisionUnitCoordination"`     //监管单位对我行核库配合
	SuitGoodsSafekeeping            int `json:"SuitGoodsSafekeeping"`            //仓储环境是否适合质押货物保管
	HaveBasicStorageFacilities      int `json:"HaveBasicStorageFacilities"`      //仓库是否具备基本仓储设施
	SecuritySystemImplementation    int `json:"SecuritySystemImplementation"`    //安全保卫体系是否落实
	ThirdPartCoordination           int `json:"ThirdPartCoordination"`           //第三方或出质人对监管的配合
	ClearPosition                   int `json:"ClearPosition"`                   //质押物保管是否有明确的仓位
	PledgeFlagSuspension            int `json:"PledgeFlagSuspension"`            //质押物是否悬挂质押标识
	PledgorConsistency              int `json:"PledgorConsistency"`              //质押物运输单据、入库单与监管单位账务记载出质人是否相符
	ListConsistency                 int `json:"ListConsistency"`                 //质押物运输单据、入库单、监管单位账务与质押物清单是否相符
	TurnoverNormal                  int `json:"TurnoverNormal"`                  //存货周转情况是否正常
	DisplaySafetyRegulations        int `json:"DisplaySafetyRegulations"`        //质押物摆放是否符合安全规范
	TakeProtectionMeasures          int `json:"TakeProtectionMeasures"`          //如需要，质押物是否采取防护措施
	GoodsMatch                      int `json:"GoodsMatch"`                      //与购销合同（单证）约定货物是否相签
	TimeReceiverDestinationMatch    int `json:"TimeReceiverDestinationMatch"`    //到货时间、收货人、到货港（站）与购销合同约定是否相符
	SupervisionUnitStation          int `json:"SupervisionUnitStation"`          //监管单位是否进驻
	GoodsEntryExitRecordsComplement int `json:"GoodsEntryExitRecordsComplement"` //货物进出库记录是否齐全
	AuditSupervisionImplementation  int `json:"AuditSupervisionImplementation"`  //适时审核监督是否落实
	ValueBalanceMeetRequirement     int `json:"ValueBalanceMeetRequirement"`     //库存货物价值余额是否符合我行要求
	PledgeGoodsMatch                int `json:"PledgeGoodsMatch"`                //在库质押货物与质押物清单是否相符
}

//巡库信息
type PledgePatrolDetailInfo struct {
	MonitorGeneralInfo MonitorGeneralInfo `json:"MonitorGeneralInfo"` //巡库基本信息
	MonitorCheckResult MonitorCheckResult `json:"MonitorCheckResult"` //巡库结果
}

type PledgeWarningInfo struct {
	ChannelSeq         string
	PledgeNoStorage    string
	PledgeName         string
	SocialCreditCode   string
	CoreEnterpriseName string
	WarningMsg         string
}
type PledgeWarningMsg struct {
	MonitorEquipmentId string
	RuleId             string
	ChannelSeq         string
	WarningContents    string
	WarningTime        string
	WarningType        string
	UserDefined        []string
}

//货物信息
type GoodsInfo struct {
	PledgeGeneralInfo      PledgeGeneralInfo        `json:"PledgeGeneralInfo"`      //押品概要信息
	PledgeDetailInfo       PledgeDetailInfo         `json:"PledgeDetailInfo"`       //押品详细信息
	PledgeInsuranceInfo    []PledgeInsuranceInfo    `json:"PledgeInsuranceInfo"`    //保险信息
	PledgeNotarizationInfo PledgeNotarizationInfo   `json:"PledgeNotarizationInfo"` //公证信息
	PledgePatrolDetailInfo []PledgePatrolDetailInfo `json:"PledgePatrolDetailInfo"` //巡库信息
	PledgeWarningInfo      []PledgeWarningInfo
}

//保险公证信息
type PledgeInsuranceNotarizationInfo struct {
	PledgeInsuranceInfo    []PledgeInsuranceInfo  `json:"PledgeInsuranceInfo"`    //保险信息
	PledgeNotarizationInfo PledgeNotarizationInfo `json:"PledgeNotarizationInfo"` //公证信息
}

//预警周期
type AlertPeriodSetting struct {
	ChannelSeq  string `json:"ChannelSeq"`  //渠道编号
	AlertPeriod string `json:"AlertPeriod"` //预警周期
}

type StatusSyncRequest struct {
	ChannelSeq      string
	PledgeNoStorage string
	PledgeName      string
	PledgeState     int
}

//查询押品信息
type QueryGoodsRequest struct {
	PledgeNoStorage string `json:"PledgeNoStorage"` //质押物编号
}
type QueryGoodsResponse struct {
	ResponseCode string `json:"ResponseCode"` // 返回码
	ResponseMsg  string `json:"ResponseMsg"`  //返回信息
	GoodsInfo    GoodsInfo
}

//查询预警周期
type QueryPeriodRequest struct {
	ChannelSeq string `json:"ChannelSeq"` //渠道编号
}
type QueryPeriodResponse struct {
	ResponseCode string `json:"ResponseCode"` // 返回码
	ResponseMsg  string `json:"ResponseMsg"`  //返回信息
	AlertPeriod  string `json:"AlertPeriod"`  //预警周期
}

//查询监控设备监控的押品
type QueryMonitorPledgeRequest struct {
	MonitorEquipmentId string `json:"MonitorEquipmentId"` //监控设备ID
}
type QueryMonitorPledgeResponse struct {
	ResponseCode        string   `json:"ResponseCode"`        // 返回码
	ResponseMsg         string   `json:"ResponseMsg"`         //返回信息
	PledgeNoStorageList []string `json:"PledgeNoStorageList"` //押品编号列表
}

type PledgeResponse struct {
	ResponseCode string `json:"ResponseCode"` //返回码
	ResponseMsg  string `json:"ResponseMsg"`  //返回信息
}
