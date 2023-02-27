package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦耶多KayattarBarony struct {
	feud.BaseBarony
}

var BKayattar迦耶多 feud.Barony = &迦耶多KayattarBarony{}

func init() {
    f := BKayattar迦耶多.(*迦耶多KayattarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kayattar",
		TitleName: "迦耶多",
		TitleCode: "b_kayattar",
	}
}
