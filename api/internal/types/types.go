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

type CreatePalReq struct {
	Number       string    `json:"number"`
	Name         string    `json:"name"`
	Icon         string    `json:"icon"`
	Description  string    `json:"description"`
	AttributeIds []int32   `json:"attribute_ids"`
	HP           int32     `json:"hp"`
	Energy       int32     `json:"energy"`
	Defensively  int32     `json:"defensively"`
	Abilities    []Ability `json:"abilities"`
	Eat          int32     `json:"eat"`
	PassiveSkill string    `json:"passive_skill"`
}

type DeletePalReq struct {
	ID int64 `path:"id"`
}

type UpdatePalReq struct {
	ID           int64     `path:"id"`
	Number       string    `json:"number"`
	Name         string    `json:"name"`
	Icon         string    `json:"icon"`
	Description  string    `json:"description"`
	AttributeIds []int32   `json:"attribute_ids"`
	HP           int32     `json:"hp"`
	Energy       int32     `json:"energy"`
	Defensively  int32     `json:"defensively"`
	Abilities    []Ability `json:"abilities"`
	Eat          int32     `json:"eat"`
	PassiveSkill string    `json:"passive_skill"`
}

type CreatePalResp struct {
	ID int64 `json:"id"`
}

type UpdatePalResp struct {
	ID int64 `json:"id"`
}

type DeletePalResp struct {
	ID int64 `json:"id"`
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

type SkillInfo struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Desc        string     `json:"desc"`
	CT          int32      `json:"ct"`
	Power       int32      `json:"power"`
	AttributeID int32      `json:"attribute_id"`
	Pals        []SkillPal `json:"pals"`
}

type AddSkillReq struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	CT          int32  `form:"ct"`
	Power       int32  `form:"power"`
	AttributeID int32  `form:"attribute_id"`
}

type DeleteSkillReq struct {
	ID int32 `path:"id"`
}

type UpdateSkillReq struct {
	ID          int32  `path:"id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	CT          int32  `form:"ct"`
	Power       int32  `form:"power"`
	AttributeID int32  `form:"attribute_id"`
}

type GetSkillPalsReq struct {
	SkillID int32 `path:"id"`
}

type SkillPal struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ListGoodsreq struct {
	Name     string `json:"name,optional" form:"name,optional"`
	Types    string `json:"types,optional" form:"types,optional"`
	Page     int32  `json:"page,optional" form:"page,optional"`
	PageSize int32  `json:"page_size,optional" form:"page_size,optional"`
	Quality  int32  `json:"quality,optional" form:"quality,optional"`
}

type Material struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int32  `json:"count"`
	Image string `json:"image"`
}

type ListGoods struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Image       string     `json:"image"`
	Description string     `json:"description"`
	Quality     int32      `json:"quality"`
	Workload    int32      `json:"workload"`
	Materials   []Material `json:"materials"`
	Types       string     `json:"types"`
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

type Technology struct {
	ID          int64                `json:"id"`
	Name        string               `json:"name"`
	Icon        string               `json:"icon"`
	Description string               `json:"description"`
	Level       int32                `json:"level"`
	Cost        int32                `json:"cost"`
	Ancient     bool                 `json:"ancient"`
	Material    []TechnologyMaterial `json:"material"`
}

type TechnologyMaterial struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Count int32  `json:"count"`
}

type AddTechnologyMaterial struct {
	ID    int64 `json:"id"`
	Count int32 `json:"count"`
}

type TechnologyTree struct {
	Level int32        `json:"level"`
	Data  []Technology `json:"data"`
}

type AddTechnologyReq struct {
	Name        string                  `form:"name"`
	Icon        string                  `form:"icon"`
	Level       int32                   `form:"level"`
	Description string                  `form:"description"`
	Cost        int32                   `form:"cost"`
	Ancient     bool                    `form:"ancient"`
	Material    []AddTechnologyMaterial `form:"material"`
}

type GetTechnologyReq struct {
	ID int64 `path:"id"`
}

type GetTechnologyResp struct {
	Technology Technology `json:"technology"`
}

type Container struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Created int64  `json:"created"`
	State   string `json:"state"`
	Status  string `json:"status"`
	Health  string `json:"health,omitempty"`
}

type ListContainerResp struct {
	Base
	Data []Container `json:"data"`
}

type ContainerLogReq struct {
	ID string `json:"id,optional" form:"id,optional"`
}

type Event struct {
	Message string `json:"message"`
}

type Captcha struct {
	CaptchaId     string `json:"captcha_id"`
	CaptchaBase64 string `json:"captcha_base64"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResp struct {
	UserInfo *UserInfo `json:"user_info"`
	Token    *Token    `json:"token"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type RegisterReq struct {
	Username        string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
	CaptchaId       string `form:"captcha_id"`
	CaptchaCode     string `form:"captcha_code"`
}

type RegisterResp struct {
	UserInfo *UserInfo `json:"user_info"`
	Token    *Token    `json:"token"`
}

type GetUserReq struct {
	Id int64 `path:"id"`
}

type UserInfo struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type UserListReq struct {
	Page     int64 `form:"page"`
	PageSize int64 `form:"page_size"`
}

type UserListResp struct {
	List  []*UserInfo `json:"list"`
	Total int64       `json:"total"`
}

type UpdateUserReq struct {
	Id       int64  `path:"id"`
	Username string `form:"username"`
	Avatar   string `form:"avatar"`
}

type UpdateUserResp struct {
	Id int64 `json:"id"`
}

type DeleteUserReq struct {
	Id int64 `form:"id"`
}

type DeleteUserResp struct {
	Id int64 `path:"id"`
}

type AddUserReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Avatar   string `form:"avatar"`
	Role     int    `form:"role"`
}

type AddUserResp struct {
	Id int64 `json:"id"`
}

type Comment struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	Username   string    `json:"username"`
	Avatar     string    `json:"avatar"`
	Content    string    `json:"content"`
	CreateTime string    `json:"create_time"`
	ParentID   int64     `json:"parent_id"`
	RootID     int64     `json:"root_id"`
	Children   []Comment `json:"children"`
}

type AddCommentReq struct {
	UserID   int64  `form:"user_id"`
	Content  string `form:"content"`
	Category string `form:"category"`
	ObjectID int64  `form:"object_id"`
	RootID   int64  `form:"root_id"`
	ParentID int64  `form:"parent_id"`
}

type ListCommentReq struct {
	Category string `form:"category"`
	ObjectID int64  `form:"object_id"`
	Page     int32  `form:"page"`
	PageSize int32  `form:"page_size"`
}

type ListCommentResp struct {
	List  []Comment `json:"list"`
	Total int64     `json:"total"`
}
