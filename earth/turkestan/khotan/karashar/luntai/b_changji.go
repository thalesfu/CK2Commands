package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昌吉ChangjiBarony struct {
	feud.BaseBarony
}

var BChangji昌吉 feud.Barony = &昌吉ChangjiBarony{}

func init() {
	f := BChangji昌吉.(*昌吉ChangjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "changji",
		TitleName: "昌吉",
		TitleCode: "b_changji",
	}
}
