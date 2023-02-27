package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马纳特ManatBarony struct {
	feud.BaseBarony
}

var BManat马纳特 feud.Barony = &马纳特ManatBarony{}

func init() {
    f := BManat马纳特.(*马纳特ManatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manat",
		TitleName: "马纳特",
		TitleCode: "b_manat",
	}
}
