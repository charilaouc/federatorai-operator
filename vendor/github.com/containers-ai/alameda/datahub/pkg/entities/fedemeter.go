package entities

import (
	"time"
)

type FedemeterCalculationInstance struct {
	DatahubEntity   `scope:"fedemeter" category:"calculation" type:"price"`
	Measurement     *Measurement `name:"calculation_price_instance" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time            *time.Time   `json:"time"            required:"false" column:"tag"`
	NodeName        string       `json:"nodename"        required:"true"  column:"tag"`
	ClusterName     string       `json:"clustername"     required:"true"  column:"tag"`
	Provider        string       `json:"provider"        required:"true"  column:"tag"`
	Region          string       `json:"region"          required:"true"  column:"tag"`
	Unit            string       `json:"unit"            required:"true"  column:"tag"`
	Granularity     string       `json:"granularity"     required:"true"  column:"tag"`
	CPU             float64      `json:"cpu"             required:"false" column:"field"`
	Memory          float64      `json:"memory"          required:"false" column:"field"`
	Description     string       `json:"description"     required:"false" column:"field"`
	DisplayName     string       `json:"displayname"     required:"false" column:"field"`
	InstanceNum     int32        `json:"instancenum"     required:"false" column:"field"`
	InstanceType    string       `json:"instancetype"    required:"false" column:"field"`
	NodeType        string       `json:"nodetype"        required:"false" column:"field"`
	OperatingSystem string       `json:"operatingsystem" required:"false" column:"field"`
	Period          int32        `json:"period"          required:"false" column:"field"`
	PreinstalledSW  string       `json:"preinstalledsw"  required:"false" column:"field"`
	StartTime       int64        `json:"starttime"       required:"false" column:"field"`
	Cost            float64      `json:"cost"            required:"false" column:"field"`
	TotalCost       float64      `json:"totalcost"       required:"false" column:"field"`
}

type FedemeterCalculationStorage struct {
	DatahubEntity `scope:"fedemeter" category:"calculation" type:"price"`
	Measurement   *Measurement `name:"calculation_price_storage" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time   `json:"time"        required:"false" column:"tag"`
	NodeName      string       `json:"nodename"    required:"true"  column:"tag"`
	ClusterName   string       `json:"clustername" required:"true"  column:"tag"`
	Provider      string       `json:"provider"    required:"true"  column:"tag"`
	Unit          string       `json:"unit"        required:"true"  column:"tag"`
	Granularity   string       `json:"granularity" required:"true"  column:"tag"`
	Description   string       `json:"description" required:"false" column:"field"`
	DisplayName   string       `json:"displayname" required:"false" column:"field"`
	Period        int32        `json:"period"      required:"false" column:"field"`
	StorageNum    int32        `json:"storagenum"  required:"false" column:"field"`
	StorageSize   int64        `json:"storagesize" required:"false" column:"field"`
	VolumeType    string       `json:"volumetype"  required:"false" column:"field"`
	StartTime     int64        `json:"starttime"   required:"false" column:"field"`
	Cost          float64      `json:"cost"        required:"false" column:"field"`
}

type FedemeterRecommendationJERI struct {
	DatahubEntity     `scope:"fedemeter" category:"recommendation" type:"jeri"`
	Measurement       *Measurement `name:"recommendation_jeri" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time              *time.Time   `json:"time"                required:"false" column:"tag"`
	ClusterName       string       `json:"clustername"         required:"true"  column:"tag"`
	Country           string       `json:"country"             required:"true"  column:"tag"`
	InstanceType      string       `json:"instancetype"        required:"true"  column:"tag"`
	Provider          string       `json:"provider"            required:"true"  column:"tag"`
	Rank              string       `json:"rank"                required:"true"  column:"tag"`
	Region            string       `json:"region"              required:"true"  column:"tag"`
	ReservedInstances string       `json:"reservedinstances"   required:"true"  column:"tag"`
	Granularity       string       `json:"granularity"         required:"true"  column:"tag"`
	MasterNum         int32        `json:"master_num"          required:"false" column:"field"`
	WorkerNum         int32        `json:"worker_num"          required:"false" column:"field"`
	MasterStorageSize float64      `json:"master_storage_size" required:"false" column:"field"`
	WorkerStorageSize float64      `json:"worker_storage_size" required:"false" column:"field"`
	AccCost           float64      `json:"acc_cost"            required:"false" column:"field"`
	DisplayName       string       `json:"display_name"        required:"false" column:"field"`
	ResourceName      string       `json:"resource_name"       required:"false" column:"field"`
	StartTime         int64        `json:"start_time"          required:"false" column:"field"`
	OndemandNum       int32        `json:"ondemand_num"        required:"false" column:"field"`
	MasterRiNum       int32        `json:"master_ri_num"       required:"false" column:"field"`
	WorkerRiNum       int32        `json:"worker_ri_num"       required:"false" column:"field"`
	MasterSpotNum     int32        `json:"master_spot_num"     required:"false" column:"field"`
	WorkerSpotNum     int32        `json:"worker_spot_num"     required:"false" column:"field"`
	MasterOndemandNum int32        `json:"master_ondemand_num" required:"false" column:"field"`
	WorkerOndemandNum int32        `json:"worker_ondemand_num" required:"false" column:"field"`
	Cost              float64      `json:"cost"                required:"false" column:"field"`
	TotalCost         float64      `json:"total_cost"          required:"false" column:"field"`
}

type FedemeterResourceHistoryCostApp struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_history" type:"cost"`
	Measurement    *Measurement `name:"resource_history_cost_app" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"           required:"false" column:"tag"`
	AppName        string       `json:"appname"        required:"true"  column:"tag"`
	NamespaceName  string       `json:"namespacename"  required:"true"  column:"tag"`
	ClusterName    string       `json:"clustername"    required:"true"  column:"tag"`
	Provider       string       `json:"provider"       required:"true"  column:"tag"`
	Type           string       `json:"type"           required:"true"  column:"tag"`
	Granularity    string       `json:"granularity"    required:"true"  column:"tag"`
	CreateTime     int64        `json:"createtime"     required:"false" column:"field"`
	StartTime      int64        `json:"starttime"      required:"false" column:"field"`
	CostPercentage float64      `json:"costpercentage" required:"false" column:"field"`
	WorkloadCost   float64      `json:"workloadcost"   required:"false" column:"field"`
}

type FedemeterResourceHistoryCostNamespace struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_history" type:"cost"`
	Measurement    *Measurement `name:"resource_history_cost_namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"           required:"false" column:"tag"`
	NamespaceName  string       `json:"namespacename"  required:"true"  column:"tag"`
	ClusterName    string       `json:"clustername"    required:"true"  column:"tag"`
	Provider       string       `json:"provider"       required:"true"  column:"tag"`
	Type           string       `json:"type"           required:"true"  column:"tag"`
	Granularity    string       `json:"granularity"    required:"true"  column:"tag"`
	CreateTime     int64        `json:"createtime"     required:"false" column:"field"`
	StartTime      int64        `json:"starttime"      required:"false" column:"field"`
	CostPercentage float64      `json:"costpercentage" required:"false" column:"field"`
	WorkloadCost   float64      `json:"workloadcost"   required:"false" column:"field"`
}

type FedemeterResourcePredictionCostApp struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_prediction" type:"cost"`
	Measurement    *Measurement `name:"resource_prediction_cost_app" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"           required:"false" column:"tag"`
	AppName        string       `json:"appname"        required:"true"  column:"tag"`
	NamespaceName  string       `json:"namespacename"  required:"true"  column:"tag"`
	ClusterName    string       `json:"clustername"    required:"true"  column:"tag"`
	Provider       string       `json:"provider"       required:"true"  column:"tag"`
	Type           string       `json:"type"           required:"true"  column:"tag"`
	Granularity    string       `json:"granularity"    required:"true"  column:"tag"`
	CreateTime     int64        `json:"createtime"     required:"false" column:"field"`
	StartTime      int64        `json:"starttime"      required:"false" column:"field"`
	CostPercentage float64      `json:"costpercentage" required:"false" column:"field"`
	WorkloadCost   float64      `json:"workloadcost"   required:"false" column:"field"`
}

type FedemeterResourcePredictionCostNamespace struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_prediction" type:"cost"`
	Measurement    *Measurement `name:"resource_prediction_cost_namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time   `json:"time"           required:"false" column:"tag"`
	NamespaceName  string       `json:"namespacename"  required:"true"  column:"tag"`
	ClusterName    string       `json:"clustername"    required:"true"  column:"tag"`
	Provider       string       `json:"provider"       required:"true"  column:"tag"`
	Type           string       `json:"type"           required:"true"  column:"tag"`
	Granularity    string       `json:"granularity"    required:"true"  column:"tag"`
	CreateTime     int64        `json:"createtime"     required:"false" column:"field"`
	StartTime      int64        `json:"starttime"      required:"false" column:"field"`
	CostPercentage float64      `json:"costpercentage" required:"false" column:"field"`
	WorkloadCost   float64      `json:"workloadcost"   required:"false" column:"field"`
}
