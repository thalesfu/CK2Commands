package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基奇卡斯KichkasBarony struct {
	feud.BaseBarony
}

var BKichkas基奇卡斯 feud.Barony = &基奇卡斯KichkasBarony{}

func init() {
    f := BKichkas基奇卡斯.(*基奇卡斯KichkasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kichkas",
		TitleName: "基奇卡斯",
		TitleCode: "b_kichkas",
	}
}
