package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布阿BuaBarony struct {
	feud.BaseBarony
}

var BBua布阿 feud.Barony = &布阿BuaBarony{}

func init() {
    f := BBua布阿.(*布阿BuaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bua",
		TitleName: "布阿",
		TitleCode: "b_bua",
	}
}
