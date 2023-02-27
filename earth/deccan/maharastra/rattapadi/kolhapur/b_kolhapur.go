package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘罗诃补罗KolhapurBarony struct {
	feud.BaseBarony
}

var BKolhapur拘罗诃补罗 feud.Barony = &拘罗诃补罗KolhapurBarony{}

func init() {
    f := BKolhapur拘罗诃补罗.(*拘罗诃补罗KolhapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolhapur",
		TitleName: "拘罗诃补罗",
		TitleCode: "b_kolhapur",
	}
}
