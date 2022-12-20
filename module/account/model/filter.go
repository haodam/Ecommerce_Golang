package accountmodel

type Filter struct {
	AccountId int   `json:"account_id,omitempty" form:"account_id"`
	Status    []int `json:"-"`
}
