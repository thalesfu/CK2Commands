package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉钦斯克KalachinskBarony struct {
	feud.BaseBarony
}

var BKalachinsk卡拉钦斯克 feud.Barony = &卡拉钦斯克KalachinskBarony{}

func init() {
	f := BKalachinsk卡拉钦斯克.(*卡拉钦斯克KalachinskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalachinsk",
		TitleName: "卡拉钦斯克",
		TitleCode: "b_kalachinsk",
	}
}
