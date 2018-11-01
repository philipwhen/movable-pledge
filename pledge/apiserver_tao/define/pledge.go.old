package define

type PledgeGeneralInfo struct {
	ChannelSeq       string
	PledgeNoStorage  string
	PledgeName       string
	SocialCreditCode string
	PledgeOwnerName  string
	PledgeState      string
	PledgeType       string
}
type PledgeDetailInfo struct {
	CommodityVaretiesName string
	Specifications        string
	Quantity              string
	CommodityCompany      string
	Manufacturer          string
	VoucherNo             string
	VoucherName           string
	CoreSocialCreditCode  string
	CoreEnterpriseName    string
	DepotNo               string
	DepotName             string
	IsBuyInsurance        string
	WarehousingDate       string
	InventoryStatus       string
	MonitorEquipmentIds   []string
	RuleList              map[string][]string
}
type PledgeInsuranceInfo struct {
	ChannelSeq         string
	PledgeNoStorage    string
	PledgeName         string
	InsuranceCompany   string
	InsuranceNo        string
	InsuranceType      string
	InsuranceClass     string
	PolicyHolder       string
	Beneficiary        string
	InsurancePremium   string
	InsuranceAmount    string
	InsuranceStartDate string
	InsuranceEndDate   string
	InsuranceDate      string
}
type PledgeNotarizationInfo struct {
	ChannelSeq       string
	NotarialDeedNo   string
	NotarialOffice   string
	RegistrationDate string
}
type PledgePatrolDetailInfo struct {
	ID                      string
	Pledge_patrol_no        string
	Social_credit_code      string
	Enterprise_name         string
	Core_social_credit_code string
	Core_enterprise_name    string
	Ctra_no                 string
	Depot_no                string
	Depot_name              string
	Realtive_pledge_no      string
	Pledge_name             string
	Status                  string
	Manufacturer            string
	Specifications          string
	Commodity_company       string
	Quantity                string
	NEW_MARKET_PRICE        string
	Loan_time               string
	Loan_maturity           string
	Monitor_equipment_id    string
	Patrol_duration         string
	Patrol_type             string
	Patrol_person           string
	Patrol_directions       string
	Input_user              string
	Input_org               string
	Input_time              string
	Update_user             string
	Update_org              string
	Update_time             string
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
type GoodsInfo struct {
	PledgeGeneralInfo      PledgeGeneralInfo
	PledgeDetailInfo       PledgeDetailInfo
	PledgeInsuranceInfo    []PledgeInsuranceInfo
	PledgeNotarizationInfo PledgeNotarizationInfo
	PledgePatrolDetailInfo []PledgePatrolDetailInfo
	PledgeWarningInfo      []PledgeWarningInfo
}
type QueryGoodsRequest struct {
	PledgeNoStorage string
}
type StatusSyncRequest struct {
	ChannelSeq      string
	PledgeNoStorage string
	PledgeName      string
	PledgeState     string
}
type QueryGoodsResponse struct {
	ResponseCode string `json:"ResponseCode"` // 返回码
	ResponseMsg  string `json:"ResponseMsg"`  //返回信息
	GoodsInfo    GoodsInfo
}

type UploadMaterialsResponse struct {
	ResponseCode string `json:"ResponseCode"` // 返回码
	ResponseMsg  string `json:"ResponseMsg"`  //返回信息
}
