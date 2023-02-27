package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞拉SaleBarony struct {
	feud.BaseBarony
}

var BSale塞拉 feud.Barony = &塞拉SaleBarony{}

func init() {
    f := BSale塞拉.(*塞拉SaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sale",
		TitleName: "塞拉",
		TitleCode: "b_sale",
	}
}
