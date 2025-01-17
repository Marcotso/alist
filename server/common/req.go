package common

type PageReq struct {
	Page    int `json:"page_index" form:"page"`
	PerPage int `json:"page_size" form:"per_page"`
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func (p *PageReq) Validate() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PerPage < 1 {
		p.PerPage = MaxInt
	}
}
