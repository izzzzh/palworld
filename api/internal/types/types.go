// Code generated by goctl. DO NOT EDIT.
package types

type GetPalReq struct {
	ID int64 `path:"id"`
}

type Ability struct {
	Icon  int32  `json:"icon"`
	Name  string `json:"name"`
	Level int64  `json:"level"`
}

type Pal struct {
	ID               int64     `json:"id"`
	Number           string    `json:"number"`
	Name             string    `json:"name"`
	Icon             string    `json:"icon"`
	Description      string    `json:"description"`
	AttributeIds     []int32   `json:"attribute_ids"`
	HP               int32     `json:"hp"`
	Energy           int32     `json:"energy"`
	Defensively      int32     `json:"defensively"`
	Abilities        []Ability `json:"abilities"`
	Eat              int32     `json:"eat"`
	PassiveSkill     string    `json:"passive_skill"`
	PassiveSkillDesc string    `json:"passive_skill_desc"`
	ActiveSkills     []Skill   `json:"active_skills"`
}

type GetPalResp struct {
	Base
	Data Pal `json:"data"`
}

type ListPalReq struct {
	Name        string `json:"name,optional"`
	AttributeId int64  `json:"attribute_id,optional"`
	Page        int64  `json:"page,optional"`
}

type ListPal struct {
	ID           int64     `json:"id"`
	Number       string    `json:"number"`
	Name         string    `json:"name"`
	Icon         string    `json:"icon"`
	AttributeIds []int32   `json:"attribute_ids"`
	Abilities    []Ability `json:"abilities"`
}

type ListPalResp struct {
	Base
	Data []ListPal `json:"data"`
}

type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Skill struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	CT          int32  `json:"ct"`
	Power       int32  `json:"power"`
	AttributeID int32  `json:"attribute_id"`
}

type ListSkillReq struct {
	AttributeID int32 `json:"attribute_id,optional"`
}

type ListSkillResp struct {
	Base
	Data []Skill `json:"data"`
}

type ListGoodsreq struct {
	Name  string `json:"name,optional"`
	Types int    `json:"types,optional"`
}

type Material struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ListGoods struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Image       string     `json:"image"`
	Description string     `json:"description"`
	Quality     int32      `json:"quality"`
	Workload    int32      `json:"workload"`
	Materials   []Material `json:"materials"`
	Types       int32      `json:"types"`
}

type ListGoodsResp struct {
	Base
	Data []ListGoods `json:"data"`
}

type ListPalMateReq struct {
	ParentOne int64 `json:"parent_one,optional" form:"parent_one,optional"`
	ParentTwo int64 `json:"parent_two,optional" form:"parent_two,optional"`
	Result    int64 `json:"result,optional" form:"result,optional"`
	Page      int   `json:"page,optional" form:"page,optional"`
}

type MateInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type PalMate struct {
	ParentOne MateInfo `json:"parent_one"`
	ParentTwo MateInfo `json:"parent_two"`
	Result    MateInfo `json:"result"`
}

type ListPalMateResp struct {
	Base
	Data []*PalMate `json:"data"`
}

type GetTechnologyTreeResp struct {
	Base
	Data []TechnologyTree `json:"data"`
}

type Technology struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Ancient     bool   `json:"ancient"`
}

type TechnologyTree struct {
	Level int32        `json:"level"`
	Data  []Technology `json:"data"`
}
