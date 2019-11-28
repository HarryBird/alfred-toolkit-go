package Alfred

import (
	"strconv"
)

var uid int64

func newUID() int64 {
	uid++
	return uid
}

type Item struct {
	Uid          string `json:"uid"`
	Arg          string `json:"arg"`
	Type         string `json:"type"`
	Valid        bool   `json:"valid"`
	AutoComplete string `json:"autocomplete"`
	Title        string `json:"title"`
	SubTitle     string `json:"subtitle"`
	Icon         Icon   `json:"icon"`
}

func NewDefaultItem() Item {
	return Item{
		Uid:          strconv.FormatInt(newUID(), 10),
		Arg:          "Default Arg",
		Valid:        true,
		Type:         "default",
		AutoComplete: "AutoComplete",
		Title:        "Default Title",
		SubTitle:     "Default SubTitle",
		Icon:         NewDefaultIcon(),
	}
}

func NewErrorItem(err error) Item {
	return NewItem(
		"We had a error", err.Error(), "", "", "", "", false, NewErrorIcon(),
	)
}

func NewErrorTitleItem(title, subTitle string) Item {
	return NewItem(
		title, subTitle, "", "", "", "", false, NewErrorIcon(),
	)
}

func NewNoResultItem() Item {
	return NewItem(
		"No Result", "", "", "", "", "", false, NewFailIcon(),
	)
}

func NewCommonItem(title, subTitle, arg string) Item {
	return NewItem(title, subTitle, arg, arg, "", "", true, NewSuccIcon())
}

func NewItem(title, subTitle, arg, auto, uid, ty string, valid bool, icon Icon) Item {
	if uid == "" {
		uid = strconv.FormatInt(newUID(), 10)
	}

	if ty == "" {
		ty = "default"
	}

	return Item{
		Uid:          uid,
		Arg:          arg,
		Valid:        valid,
		Type:         ty,
		AutoComplete: auto,
		Title:        title,
		SubTitle:     subTitle,
		Icon:         icon,
	}
}
