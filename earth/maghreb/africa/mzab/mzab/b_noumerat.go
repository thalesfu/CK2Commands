package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努迈拉特NoumeratBarony struct {
	feud.BaseBarony
}

var BNoumerat努迈拉特 feud.Barony = &努迈拉特NoumeratBarony{}

func init() {
	f := BNoumerat努迈拉特.(*努迈拉特NoumeratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noumerat",
		TitleName: "努迈拉特",
		TitleCode: "b_noumerat",
	}
}
