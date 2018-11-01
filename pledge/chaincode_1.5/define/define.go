package define

type InvokeRequest struct {
	Key   string `json:"key"`   //存储数据的key
	Value string `json:"value"` //存储数据的value
}

type InvokeResponse struct {
	ResStatus ResponseStatus `json:"responseStatus"`
	Payload   interface{}    `json:"payload"`
}

type QueryResponse struct {
	ResStatus ResponseStatus `json:"responseStatus"`
	Page      Page           `json:"page"`
	Payload   interface{}    `json:"payload"`
}

type Page struct {
	CurrentPage  uint `json:"currentPage"`  //当前页码
	PageSize     uint `json:"pageSize"`     //每个页面显示个数
	TotalRecords uint `json:"totalRecords"` //总记录数
}

type ResponseStatus struct {
	StatusCode int    `json:"statusCode"` //错误码0:成功1:失败
	StatusMsg  string `json:"statusMsg"`  //错误信息
}

type NormalArgs struct {
	Args []string `json:"Args"`
}

type OperateRequest struct {
	MonitorEquipmentId []string   `json:"monitorEquipmentId"` //设备ID
	Rule_id            [][]string `json:"rule_id"`            //规则ID
	IsEnabled          [][]int    `json:"isEnabled"`          //是否开启
	Custom             string     `json:"custom"`             //自定义
}

type Custom struct {
	Appkey     string
	Time       string
	Token      string
	OpUserUuid string
}