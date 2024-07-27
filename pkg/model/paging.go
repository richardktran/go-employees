package model

type Paging struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func (p *Paging) Process() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}
