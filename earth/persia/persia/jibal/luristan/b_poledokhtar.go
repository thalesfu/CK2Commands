package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔多赫塔尔PoledokhtarBarony struct {
	feud.BaseBarony
}

var BPoledokhtar波尔多赫塔尔 feud.Barony = &波尔多赫塔尔PoledokhtarBarony{}

func init() {
    f := BPoledokhtar波尔多赫塔尔.(*波尔多赫塔尔PoledokhtarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poledokhtar",
		TitleName: "波尔多赫塔尔",
		TitleCode: "b_poledokhtar",
	}
}
