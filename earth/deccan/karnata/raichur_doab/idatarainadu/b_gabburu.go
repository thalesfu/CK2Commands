package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加布罗GabburuBarony struct {
	feud.BaseBarony
}

var BGabburu加布罗 feud.Barony = &加布罗GabburuBarony{}

func init() {
    f := BGabburu加布罗.(*加布罗GabburuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabburu",
		TitleName: "加布罗",
		TitleCode: "b_gabburu",
	}
}
