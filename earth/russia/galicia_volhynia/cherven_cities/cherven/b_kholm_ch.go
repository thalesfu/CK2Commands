package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔姆Kholm_chBarony struct {
	feud.BaseBarony
}

var BKholm_ch霍尔姆 feud.Barony = &霍尔姆Kholm_chBarony{}

func init() {
    f := BKholm_ch霍尔姆.(*霍尔姆Kholm_chBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kholm_ch",
		TitleName: "霍尔姆",
		TitleCode: "b_kholm_ch",
	}
}
