package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那赐罗跋NasirabadBarony struct {
	feud.BaseBarony
}

var BNasirabad那赐罗跋 feud.Barony = &那赐罗跋NasirabadBarony{}

func init() {
    f := BNasirabad那赐罗跋.(*那赐罗跋NasirabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nasirabad",
		TitleName: "那赐罗跋",
		TitleCode: "b_nasirabad",
	}
}
