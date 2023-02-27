package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基律纳KirunaBarony struct {
	feud.BaseBarony
}

var BKiruna基律纳 feud.Barony = &基律纳KirunaBarony{}

func init() {
    f := BKiruna基律纳.(*基律纳KirunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiruna",
		TitleName: "基律纳",
		TitleCode: "b_kiruna",
	}
}
