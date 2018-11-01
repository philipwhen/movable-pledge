package define

type PledgeGeneralInfo struct {
	PledgeNo             string
	PledgeName           string
	PledgeOwnerNumber    string
	PledgeOwnerName      string
	PledgeType           string
	PledgeStatus         string
	PledgeEffectiveValue string
	Currency             string
	EnterpriseNo         string
	DisposalStatus       string
	CreatedBran          string
	CreatedUser          string
	CreatedDate          string
}
type PledgeDetailInfo struct {
	PledgeNo              string
	CommodityVaretiesName string
	CommodityBigClass     string
	CommodityInClass      string
	CommoditySmallClass   string
	Quantity              string
	CommodityCompany      string
	NewMarketPrice        string
	InitMarketPrice       string
	CurrentInvoicePrice   string
	WarehousingDate       string
	DepotNo               string
	DepotName             string
	IsBuyInsurance        string
	Manufacturer          string
	Specifications        string
	InventoryStatus       string
	VoucherNo             string
	VoucherName           string
	CoreSocialCreditCode  string
	CoreEnterpriseName    string
	MonitorEquipmentId    string
	Remark                string
}
type PledgeInsuranceInfo struct {
	PledgeNo           string
	PledgeName         string
	PledgeType         string
	InsuranceNumber    string
	InsuranceCompany   string
	InsuranceType      string
	TypesOfInsurance   string
	InsurPeople        string
	Beneficiary        string
	InsureFee          string
	InsureAmount       string
	InsuranceStartDate string
	InsuranceEndDate   string
	IssueDate          string
	Remark             string
}
type PledgeNotarizationInfo struct {
	PledgeNo           string
	NotarizationNumber string
	NotarialOffice     string
	RegisterDateStr    string
	Remark             string
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
type GoodsInfo struct {
	PledgeGeneralInfo      PledgeGeneralInfo
	PledgeDetailInfo       PledgeDetailInfo
	PledgeInsuranceInfo    []PledgeInsuranceInfo
	PledgeNotarizationInfo PledgeNotarizationInfo
	PledgePatrolDetailInfo []PledgePatrolDetailInfo
}
type QueryGoodsRequest struct {
	PledgeNo string
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
