package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾伊拜克AibakBarony struct {
	feud.BaseBarony
}

var BAibak艾伊拜克 feud.Barony = &艾伊拜克AibakBarony{}

func init() {
	f := BAibak艾伊拜克.(*艾伊拜克AibakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aibak",
		TitleName: "艾伊拜克",
		TitleCode: "b_aibak",
	}
}
