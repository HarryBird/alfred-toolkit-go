package Alfred

type Res struct {
	Items []Item `json:"items"`
}

func (r *Res) Set(its []Item) *Res {
	r.Items = its
	return r
}

func (r *Res) Append(it Item) *Res {
	r.Items = append(r.Items, it)
	return r
}
