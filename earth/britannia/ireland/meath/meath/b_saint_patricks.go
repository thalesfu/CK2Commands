package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣帕特里克Saint_patricksBarony struct {
	feud.BaseBarony
}

var BSaint_patricks圣帕特里克 feud.Barony = &圣帕特里克Saint_patricksBarony{}

func init() {
    f := BSaint_patricks圣帕特里克.(*圣帕特里克Saint_patricksBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saint_patricks",
		TitleName: "圣帕特里克",
		TitleCode: "b_saint_patricks",
	}
}
