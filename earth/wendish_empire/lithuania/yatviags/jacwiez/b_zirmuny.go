package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日尔穆内ZirmunyBarony struct {
	feud.BaseBarony
}

var BZirmuny日尔穆内 feud.Barony = &日尔穆内ZirmunyBarony{}

func init() {
    f := BZirmuny日尔穆内.(*日尔穆内ZirmunyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zirmuny",
		TitleName: "日尔穆内",
		TitleCode: "b_zirmuny",
	}
}
