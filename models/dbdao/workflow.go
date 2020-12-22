package db

//WorkFlow 流程主表
type WorkFlow struct {
	Id              string         `gorm:"primary_key;type:varchar(50);"`
	WorkFlowName    string         `gorm:"column:work_flow_name;not null;comment:'名称'" json:"work_flow_name"`
	WorkFlowExplain string         `gorm:"column:work_flow_explain;not null;comment:'说明'"json:"work_flow_explain"`
	WorkFlowNodes   []WorkFlowNode `gorm:"foreignKey:foreign_id"  json:"work_flow_nodes"`
	WorkFlowLinks   []WorkFlowLink `gorm:"foreignKey:foreign_id"  json:"work_flow_links"`
	WorkFlowLogs    []WorkFlowLog  `gorm:"foreignKey:foreign_id"  json:"work_flow_logs"`
	Model
}

//WorkFlowNode 流程节点表
type WorkFlowNode struct {
	Id         string `gorm:"primary_key;type:varchar(50);"`
	NodeName   string `gorm:"column:node_name;not null;comment:'名称'" json:"node_name"`
	WorkFlowId string `gorm:"index;column:work_flow_id;comment:'流程ID'" json:"work_flow_id"`
	Model
}

//WorkFlowLink 流程线路
type WorkFlowLink struct {
	Id string `gorm:"primary_key;type:varchar(50);"`

	LinkName         string         `gorm:"column:link_name;not null;comment:'名称'" json:"link_name"`
	WorkFlowNodeForm string         `gorm:"index;column:work_flow_node_form;comment:'开始节点'"  json:"work_flow_node_form"`
	WorkFlowNodeTo   string         `gorm:"index;column:work_flow_node_to;comment:'下一节点'" json:"work_flow_node_to"`
	WorkFlowFuncs    []WorkFlowFunc `gorm:"many2many:work_flow_work_flow_func" json:"work_flow_funcs"`
	Model
}

//WorkFlowFunc 审批日志
type WorkFlowFunc struct {
	Id         string `gorm:"primary_key;type:varchar(50);" `
	FuncName   string `gorm:"column:func_name;not null;comment:'名称'" json:"func_name"`
	WorkFlowId string `gorm:"index;column:work_flow_id;comment:'流程ID'" json:"work_flow_id"`
	Model
}

type WorkFlowExample struct {
	Id         string `gorm:"primary_key;type:varchar(50);" `
	WorkFlowId string `gorm:"index;column:work_flow_id;comment:'流程ID'" json:"work_flow_id"`
	Name       string `gorm:"column:name,comment:'实例名称'" json:"name"`
	UserId     string `gorm:"column:user_id" json:"user_id"`
	Model
}

//WorkFlowLog 审批日志
type WorkFlowLog struct {
	Id                string `gorm:"primary_key;type:varchar(50);" `
	WorkFlowExampleId string `gorm:"index;column:work_flow_example_id;comment:'流程实例ID'"  json:"work_flow_example_id"`
	UserId            string `gorm:"column:user_id;comment:'审批人'" json:"user_id"`
	Start             string `gorm:"column:start;comment:'审批状态'" json:"start"` //读取的是sys_value表的code 不做直接关联
	Model
}
