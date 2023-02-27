package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆卡拉MukallaBarony struct {
	feud.BaseBarony
}

var BMukalla穆卡拉 feud.Barony = &穆卡拉MukallaBarony{}

func init() {
    f := BMukalla穆卡拉.(*穆卡拉MukallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mukalla",
		TitleName: "穆卡拉",
		TitleCode: "b_mukalla",
	}
}
