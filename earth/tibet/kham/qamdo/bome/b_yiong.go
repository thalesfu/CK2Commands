package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 易翁YiongBarony struct {
	feud.BaseBarony
}

var BYiong易翁 feud.Barony = &易翁YiongBarony{}

func init() {
	f := BYiong易翁.(*易翁YiongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yiong",
		TitleName: "易翁",
		TitleCode: "b_yiong",
	}
}
