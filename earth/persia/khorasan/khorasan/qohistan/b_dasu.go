package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达苏DasuBarony struct {
	feud.BaseBarony
}

var BDasu达苏 feud.Barony = &达苏DasuBarony{}

func init() {
	f := BDasu达苏.(*达苏DasuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dasu",
		TitleName: "达苏",
		TitleCode: "b_dasu",
	}
}
