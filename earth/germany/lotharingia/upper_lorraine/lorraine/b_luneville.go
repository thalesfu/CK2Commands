package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕内维尔LunevilleBarony struct {
	feud.BaseBarony
}

var BLuneville吕内维尔 feud.Barony = &吕内维尔LunevilleBarony{}

func init() {
    f := BLuneville吕内维尔.(*吕内维尔LunevilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luneville",
		TitleName: "吕内维尔",
		TitleCode: "b_luneville",
	}
}
