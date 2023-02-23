package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约代尔YoredaleBarony struct {
	feud.BaseBarony
}

var BYoredale约代尔 feud.Barony = &约代尔YoredaleBarony{}

func init() {
	f := BYoredale约代尔.(*约代尔YoredaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yoredale",
		TitleName: "约代尔",
		TitleCode: "b_yoredale",
	}
}
