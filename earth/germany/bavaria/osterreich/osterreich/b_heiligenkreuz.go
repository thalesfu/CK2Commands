package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海利根克罗伊茨HeiligenkreuzBarony struct {
	feud.BaseBarony
}

var BHeiligenkreuz海利根克罗伊茨 feud.Barony = &海利根克罗伊茨HeiligenkreuzBarony{}

func init() {
	f := BHeiligenkreuz海利根克罗伊茨.(*海利根克罗伊茨HeiligenkreuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heiligenkreuz",
		TitleName: "海利根克罗伊茨",
		TitleCode: "b_heiligenkreuz",
	}
}
