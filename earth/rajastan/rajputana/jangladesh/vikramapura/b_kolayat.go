package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘罗耶多KolayatBarony struct {
	feud.BaseBarony
}

var BKolayat拘罗耶多 feud.Barony = &拘罗耶多KolayatBarony{}

func init() {
    f := BKolayat拘罗耶多.(*拘罗耶多KolayatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolayat",
		TitleName: "拘罗耶多",
		TitleCode: "b_kolayat",
	}
}
