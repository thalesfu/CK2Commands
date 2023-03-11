package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂科钦TykocinBarony struct {
	feud.BaseBarony
}

var BTykocin蒂科钦 feud.Barony = &蒂科钦TykocinBarony{}

func init() {
    f := BTykocin蒂科钦.(*蒂科钦TykocinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tykocin",
		TitleName: "蒂科钦",
		TitleCode: "b_tykocin",
	}
}
