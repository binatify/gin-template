package apiret

import "encoding/json"

type Ret struct {
	RequestID string      `json:"requestId"`
	Result    interface{} `json:"result"`
}

type List struct {
	Items  interface{} `json:"items"`
	Marker string      `json:"marker,omitempty"`
	Total  int         `json:"total,omitempty"`
}

func NewRet(r interface{}, reqID string) *Ret {
	return &Ret{
		RequestID: reqID,
		Result:    r,
	}
}

func (r *Ret) BindData(bd interface{}) error {
	buf, err := json.Marshal(r.Result)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, &bd)
}

func (r *Ret) BindListData(bd interface{}) error {
	buf, err := json.Marshal(r.Result)
	if err != nil {
		return err
	}

	var list List
	if err = json.Unmarshal(buf, &list); err != nil {
		return err
	}

	dat, err := json.Marshal(list.Items)
	if err != nil {
		return err
	}

	return json.Unmarshal(dat, &bd)
}

func NewList(data interface{}, markers ...string) *List {
	var marker string

	if len(markers) > 0 {
		marker = markers[0]
	}

	return &List{
		Items:  data,
		Marker: marker,
	}
}

func NewWithTotal(data interface{}, totals ...int) *List {
	total := -1
	if len(totals) > 0 {
		total = totals[0]
	}

	return &List{
		Items: data,
		Total: total,
	}
}
