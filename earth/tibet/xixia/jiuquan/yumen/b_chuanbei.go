package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 川北ChuanbeiBarony struct {
	feud.BaseBarony
}

var BChuanbei川北 feud.Barony = &川北ChuanbeiBarony{}

func init() {
    f := BChuanbei川北.(*川北ChuanbeiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chuanbei",
		TitleName: "川北",
		TitleCode: "b_chuanbei",
	}
}
