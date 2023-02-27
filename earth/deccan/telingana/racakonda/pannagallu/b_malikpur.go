package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩力补罗MalikpurBarony struct {
	feud.BaseBarony
}

var BMalikpur摩力补罗 feud.Barony = &摩力补罗MalikpurBarony{}

func init() {
    f := BMalikpur摩力补罗.(*摩力补罗MalikpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malikpur",
		TitleName: "摩力补罗",
		TitleCode: "b_malikpur",
	}
}
