package types

type Meta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}

type CreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type UpdateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type ReprioritizeRequest struct {
	NewPriority int `json:"newPriority" validate:"required,gt=0"`
}

type GoodResponse struct {
	Id          int    `json:"id"`
	ProjectId   int    `json:"projectId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Removed     bool   `json:"removed"`
	CreatedAt   string `json:"createdAt"`
}

type RemoveResponse struct {
	Id        int  `json:"id"`
	ProjectId int  `json:"campaignId"`
	Removed   bool `json:"removed"`
}

type LogRemoved struct {
	Id        int  `json:"id"`
	ProjectId int  `json:"projectId"`
	Removed   bool `json:"removed"`
}

type ListResponse struct {
	Meta  *Meta           `json:"meta"`
	Goods []*GoodResponse `json:"goods"`
}

type Priority struct {
	Id          int `json:"id"`
	PriorityNum int `json:"priority"`
}

type PrioritiesResponse struct {
	Priorities []*Priority `json:"priorities"`
}
