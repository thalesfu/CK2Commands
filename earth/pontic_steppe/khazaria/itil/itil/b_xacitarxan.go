package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈吉塔尔汉XacitarxanBarony struct {
	feud.BaseBarony
}

var BXacitarxan哈吉塔尔汉 feud.Barony = &哈吉塔尔汉XacitarxanBarony{}

func init() {
    f := BXacitarxan哈吉塔尔汉.(*哈吉塔尔汉XacitarxanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xacitarxan",
		TitleName: "哈吉塔尔汉",
		TitleCode: "b_xacitarxan",
	}
}
