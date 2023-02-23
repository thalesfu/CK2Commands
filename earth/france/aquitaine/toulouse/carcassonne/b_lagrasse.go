package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉格拉斯LagrasseBarony struct {
	feud.BaseBarony
}

var BLagrasse拉格拉斯 feud.Barony = &拉格拉斯LagrasseBarony{}

func init() {
	f := BLagrasse拉格拉斯.(*拉格拉斯LagrasseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lagrasse",
		TitleName: "拉格拉斯",
		TitleCode: "b_lagrasse",
	}
}
