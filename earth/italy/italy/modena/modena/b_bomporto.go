package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邦波尔托BomportoBarony struct {
	feud.BaseBarony
}

var BBomporto邦波尔托 feud.Barony = &邦波尔托BomportoBarony{}

func init() {
    f := BBomporto邦波尔托.(*邦波尔托BomportoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bomporto",
		TitleName: "邦波尔托",
		TitleCode: "b_bomporto",
	}
}
