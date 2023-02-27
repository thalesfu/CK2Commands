package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙塔尔奇诺MontalcinoBarony struct {
	feud.BaseBarony
}

var BMontalcino蒙塔尔奇诺 feud.Barony = &蒙塔尔奇诺MontalcinoBarony{}

func init() {
    f := BMontalcino蒙塔尔奇诺.(*蒙塔尔奇诺MontalcinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montalcino",
		TitleName: "蒙塔尔奇诺",
		TitleCode: "b_montalcino",
	}
}
