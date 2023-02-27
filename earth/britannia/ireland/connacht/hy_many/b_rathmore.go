package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉沃尔RathmoreBarony struct {
	feud.BaseBarony
}

var BRathmore拉沃尔 feud.Barony = &拉沃尔RathmoreBarony{}

func init() {
    f := BRathmore拉沃尔.(*拉沃尔RathmoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rathmore",
		TitleName: "拉沃尔",
		TitleCode: "b_rathmore",
	}
}
