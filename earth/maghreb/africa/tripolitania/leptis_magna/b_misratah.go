package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米苏拉塔MisratahBarony struct {
	feud.BaseBarony
}

var BMisratah米苏拉塔 feud.Barony = &米苏拉塔MisratahBarony{}

func init() {
    f := BMisratah米苏拉塔.(*米苏拉塔MisratahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "misratah",
		TitleName: "米苏拉塔",
		TitleCode: "b_misratah",
	}
}
