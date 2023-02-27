package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡拉SkaraBarony struct {
	feud.BaseBarony
}

var BSkara斯卡拉 feud.Barony = &斯卡拉SkaraBarony{}

func init() {
    f := BSkara斯卡拉.(*斯卡拉SkaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skara",
		TitleName: "斯卡拉",
		TitleCode: "b_skara",
	}
}
