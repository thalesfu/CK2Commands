package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那乌NauBarony struct {
	feud.BaseBarony
}

var BNau那乌 feud.Barony = &那乌NauBarony{}

func init() {
	f := BNau那乌.(*那乌NauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nau",
		TitleName: "那乌",
		TitleCode: "b_nau",
	}
}
