package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉比亚格RathbegBarony struct {
	feud.BaseBarony
}

var BRathbeg拉比亚格 feud.Barony = &拉比亚格RathbegBarony{}

func init() {
    f := BRathbeg拉比亚格.(*拉比亚格RathbegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rathbeg",
		TitleName: "拉比亚格",
		TitleCode: "b_rathbeg",
	}
}
