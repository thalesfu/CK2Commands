package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥古斯塔瓦斯AugustavasBarony struct {
	feud.BaseBarony
}

var BAugustavas奥古斯塔瓦斯 feud.Barony = &奥古斯塔瓦斯AugustavasBarony{}

func init() {
    f := BAugustavas奥古斯塔瓦斯.(*奥古斯塔瓦斯AugustavasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "augustavas",
		TitleName: "奥古斯塔瓦斯",
		TitleCode: "b_augustavas",
	}
}
