package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗洛拉FloraBarony struct {
	feud.BaseBarony
}

var BFlora弗洛拉 feud.Barony = &弗洛拉FloraBarony{}

func init() {
	f := BFlora弗洛拉.(*弗洛拉FloraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flora",
		TitleName: "弗洛拉",
		TitleCode: "b_flora",
	}
}
