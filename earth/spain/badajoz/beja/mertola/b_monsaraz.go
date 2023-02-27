package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙萨拉什MonsarazBarony struct {
	feud.BaseBarony
}

var BMonsaraz蒙萨拉什 feud.Barony = &蒙萨拉什MonsarazBarony{}

func init() {
    f := BMonsaraz蒙萨拉什.(*蒙萨拉什MonsarazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monsaraz",
		TitleName: "蒙萨拉什",
		TitleCode: "b_monsaraz",
	}
}
