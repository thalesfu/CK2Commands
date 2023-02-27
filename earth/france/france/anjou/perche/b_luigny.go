package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕伊尼LuignyBarony struct {
	feud.BaseBarony
}

var BLuigny吕伊尼 feud.Barony = &吕伊尼LuignyBarony{}

func init() {
    f := BLuigny吕伊尼.(*吕伊尼LuignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luigny",
		TitleName: "吕伊尼",
		TitleCode: "b_luigny",
	}
}
