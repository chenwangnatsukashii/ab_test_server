package constant

// 数据库常量
const (
	AppTable     = "application"
	DomainTable  = "domain"
	LayerTable   = "layer"
	ExpTable     = "experiment"
	ExpRelTable  = "exp_rel"
	ReportTable  = "report"
	PublishTable = "publish_config"

	// 应用-域-层-实验 状态常量
	DeleteState  = 0 // 删除
	NormalState  = 1 // 正常
	OnLineState  = 2 // 上线
	OffLineState = 3 // 下线
)
