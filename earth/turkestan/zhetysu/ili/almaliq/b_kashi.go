package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀什KashiBarony struct {
	feud.BaseBarony
}

var BKashi喀什 feud.Barony = &喀什KashiBarony{}

func init() {
    f := BKashi喀什.(*喀什KashiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kashi",
		TitleName: "喀什",
		TitleCode: "b_kashi",
	}
}
