package relationDB

import (
	"gitee.com/i-Things/share/stores"
	"github.com/i-Things/things/service/udsvr/internal/domain/scene"
	"time"
)

type UdSceneInfo struct {
	ID         int64             `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`     // id编号
	TenantCode stores.TenantCode `gorm:"column:tenant_code;type:VARCHAR(50);NOT NULL"`         // 租户编码
	ProjectID  stores.ProjectID  `gorm:"column:project_id;type:bigint;default:0;NOT NULL"`     // 项目ID(雪花ID)
	Tag        string            `gorm:"column:tag;type:VARCHAR(128);NOT NULL;default:normal"` //标签 admin: 管理员 normal: 普通
	HeadImg    string            `gorm:"column:head_img;type:VARCHAR(256);NOT NULL"`           // 头像
	Name       string            `gorm:"column:name;type:varchar(100);NOT NULL"`               // 名称
	Desc       string            `gorm:"column:desc;type:varchar(200);NOT NULL"`               // 描述
	//LastRunTime    time.Time         `gorm:"column:last_run_time;index;default:CURRENT_TIMESTAMP;NOT NULL"`
	Status         int64 `gorm:"column:status;type:BIGINT;default:1"` //状态
	UdSceneTrigger `gorm:"embedded;embeddedPrefix:trigger_"`
	UdSceneWhen    `gorm:"embedded;embeddedPrefix:when_"`
	UdSceneThen    `gorm:"embedded;embeddedPrefix:then_"`
	stores.SoftTime
}

func (m *UdSceneInfo) TableName() string {
	return "ud_scene_info"
}

type UdSceneTrigger struct {
	Type    string                  `gorm:"column:type;type:VARCHAR(25);NOT NULL"` //触发类型 device: 设备触发 timer: 定时触发 manual:手动触发
	Devices []*UdSceneTriggerDevice `gorm:"foreignKey:SceneID;references:ID"`
	Timers  []*UdSceneTriggerTimer  `gorm:"foreignKey:SceneID;references:ID"`
}

type UdSceneTriggerTimer struct {
	ID          int64        `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`       // id编号
	SceneID     int64        `gorm:"column:scene_id;index;type:bigint"`                      // 场景id编号
	ExecAt      int64        `gorm:"column:exec_at;index;type:bigint;NOT NULL"`              //执行时间 从0点加起来的秒数 如 1点就是 1*60*60
	ExecRepeat  int64        `gorm:"column:exec_repeat;index;type:bigint;default:0b1111111"` //重复 二进制周日到周六 11111111 这个参数只有定时触发才有
	LastRunTime time.Time    `gorm:"column:last_run_time;index;default:CURRENT_TIMESTAMP;NOT NULL"`
	SceneInfo   *UdSceneInfo `gorm:"foreignKey:ID;references:SceneID"`
	Status      int64        `gorm:"column:status;type:BIGINT;default:1"` //状态 同步场景联动的status
	stores.Time
}

func (m *UdSceneTriggerTimer) TableName() string {
	return "ud_scene_trigger_timer"
}

type UdSceneTriggerDevice struct {
	ID          int64                       `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`  // id编号
	SceneID     int64                       `gorm:"column:scene_id;index;type:bigint"`                 // 场景id编号
	ProductID   string                      `gorm:"column:product_id;index;type:VARCHAR(25);NOT NULL"` //产品id
	SelectType  scene.SelectType            `gorm:"column:select_type;type:VARCHAR(25);NOT NULL"`      //设备选择方式  all: 全部 fixed:指定的设备
	GroupID     int64                       `gorm:"column:group_id;index;type:bigint"`                 //group类型传GroupID
	DeviceName  string                      `gorm:"column:device_name;type:VARCHAR(255);"`             //选择的列表  选择的列表, fixed类型是设备名列表
	DeviceAlias string                      `gorm:"column:device_alias;type:VARCHAR(255);"`            //设备别名
	Type        string                      `gorm:"column:type;type:VARCHAR(25);NOT NULL"`             //触发类型  connected:上线 disConnected:下线 reportProperty:属性上报 reportEvent: 事件上报
	Schema      *UdSceneTriggerDeviceSchema `gorm:"embedded;embeddedPrefix:schema_"`                   //物模型类型的具体操作 reportProperty:属性上报 reportEvent: 事件上报
	SceneInfo   *UdSceneInfo                `gorm:"foreignKey:ID;references:SceneID"`
	stores.Time
}

func (m *UdSceneTriggerDevice) TableName() string {
	return "ud_scene_trigger_device"
}

type UdSceneTriggerDeviceSchema struct {
	DataID   string   `gorm:"column:data_id;type:VARCHAR(255);"`        //选择为属性或事件时需要填该字段 属性的id及事件的id aa.bb.cc
	DataName string   `gorm:"column:data_name;type:VARCHAR(255);"`      //对应的物模型定义,只读
	TermType string   `gorm:"column:term_type;type:VARCHAR(255);"`      //动态条件类型  eq: 相等  not:不相等  btw:在xx之间  gt: 大于  gte:大于等于 lt:小于  lte:小于等于   in:在xx值之间
	Values   []string `gorm:"column:values;type:json;serializer:json;"` //比较条件列表
}

type UdSceneWhen struct {
	ValidRanges   scene.WhenRanges `gorm:"column:validRanges;type:json;serializer:json"`
	InvalidRanges scene.WhenRanges `gorm:"column:invalidRanges;type:json;serializer:json"`
	Conditions    scene.Conditions `gorm:"column:conditions;type:json;serializer:json"`
}

type UdSceneThen struct {
	Actions []*UdSceneThenAction `gorm:"foreignKey:SceneID;references:ID"`
}

type UdSceneThenAction struct {
	ID          int64                   `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"` // id编号
	TenantCode  stores.TenantCode       `gorm:"column:tenant_code;type:VARCHAR(50);NOT NULL"`     // 租户编码
	SceneID     int64                   `gorm:"column:scene_id;index;type:bigint"`                // 场景id编号
	ExecuteType scene.ActionExecuteType `gorm:"column:execute_type;type:VARCHAR(25);NOT NULL"`
	Delay       int64                   `gorm:"column:delay;type:bigint"`
	Device      *UdSceneActionDevice    `gorm:"embedded;embeddedPrefix:device_"`
}

func (m *UdSceneThenAction) TableName() string {
	return "ud_scene_then_action"
}

type UdSceneActionDevice struct {
	ProjectID   stores.ProjectID       `gorm:"column:project_id;type:bigint;default:2;NOT NULL"`  // 项目ID(雪花ID)
	AreaID      stores.AreaID          `gorm:"column:area_id;type:bigint;default:2;NOT NULL"`     // 项目区域ID(雪花ID)
	ProductID   string                 `gorm:"column:product_id;index;type:VARCHAR(25);NOT NULL"` //产品id
	SelectType  scene.SelectType       `gorm:"column:select_type;type:VARCHAR(25);NOT NULL"`      //设备选择方式
	DeviceName  string                 `gorm:"column:device_name;type:VARCHAR(255);"`             //选择的列表  选择的列表, fixed类型是设备名列表
	DeviceAlias string                 `gorm:"column:device_alias;type:VARCHAR(255);"`            //设备别名
	GroupID     int64                  `gorm:"column:group_id;index;type:bigint"`                 //group类型传GroupID
	Type        scene.ActionDeviceType `gorm:"column:type;type:VARCHAR(25);NOT NULL"`
	DataID      string                 `gorm:"column:data_id;index;type:VARCHAR(100);NOT NULL"`
	Value       string                 `gorm:"column:value;index;type:VARCHAR(500);NOT NULL"`
	DataName    string                 `gorm:"column:data_name;type:VARCHAR(500);NOT NULL"` //对应的物模型定义,只读
}