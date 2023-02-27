package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉加内BalaganyBarony struct {
	feud.BaseBarony
}

var BBalagany巴拉加内 feud.Barony = &巴拉加内BalaganyBarony{}

func init() {
    f := BBalagany巴拉加内.(*巴拉加内BalaganyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balagany",
		TitleName: "巴拉加内",
		TitleCode: "b_balagany",
	}
}
